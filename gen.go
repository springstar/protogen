package main

import (
	"fmt"
	"strings"
	"log"
	"os"
	"io/ioutil"
	"path/filepath"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/golang/protobuf/proto"
	"github.com/springstar/protogen/pb"
	"github.com/gobuffalo/plush"
)

type ProtoGen struct {
	files []string
	mds map[int32]*desc.MessageDescriptor
	id2names map[int32]string

}

func newProtoGen() *ProtoGen {
	return &ProtoGen{
		mds: make(map[int32]*desc.MessageDescriptor),
		id2names: make(map[int32]string),
	}
}

func(g *ProtoGen) parse(path string) {
	files, err := ioutil.ReadDir(path)
	if (err != nil) {
		log.Fatal(err)
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		
		ext := filepath.Ext(f.Name())
		if strings.Compare(ext, ".proto") != 0 {
			continue
		}

		p := filepath.Join(path, f.Name())
		
		g.files = append(g.files, p)
	}

	var parser protoparse.Parser
	for _, f := range g.files {
		fds, err := parser.ParseFiles(f)
		if err != nil {
			fmt.Println(err)
			return
		}
		
		for _, fd := range fds {
			mds := fd.GetMessageTypes()
			for _, md := range mds {
				fmt.Println(md.GetName())
				options := md.GetOptions()
				mid, _ := proto.GetExtension(options, pb.E_Msgid)
				if (mid != nil) {
					g.mds[ *mid.(*int32)] = md
					g.id2names[*mid.(*int32)] = md.GetName()
				}
			}
		}
	}

}

func (g *ProtoGen) generate() {
	content, err := ioutil.ReadFile("template/funcdecl.tpl")
	if (err != nil) {
		log.Fatal(err)
	}

	template := string(content[:])

	ctx := plush.NewContext()
	var protos []string
	for _, p := range g.id2names {
		protos = append(protos, p)
	}

	ctx.Set("package", "package main")
	ctx.Set("imports", func() []string {
		return []string{
			"github.com/springstar/protogen/pb",
			"github.com/jhump/protoreflect/desc",
			"github.com/jhump/protoreflect/dynamic",
		}
	})

	ctx.Set("names", protos)
	ctx.Set("decl", func(p string) string {
		var sb strings.Builder
		sb.WriteString("func parse")
		sb.WriteString(p)
		sb.WriteString("(id int32, bytes []byte)")
		sb.WriteString(" *")
		sb.WriteString(p)
		return sb.String()
		
		
	})

	ctx.Set("lbrack", func() string {
		return " {"
	})

	ctx.Set("rbrack", func() string {
		return "}"
	})

	s, err := plush.Render(template, ctx)
	if (err != nil) {
		log.Fatal(err)
	}

	write(s)

	

}

func write(s string) {
    file, err := os.Create("serializer.go")
    if err != nil {
        return
    }
    defer file.Close()

    file.WriteString(s)
}





