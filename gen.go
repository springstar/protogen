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

var typeMap map[string]string = map[string]string{
	"TYPE_STRING": "string",
	"TYPE_INT32": "int32",
	"TYPE_INT": "int",
	"TYPE_INT64": "int64",
	"TYPE_MESSAGE": "msg",
	"TYPE_ENUM": "enum",
	"TYPE_BOOL": "bool",
}

type Field struct {
	ptype string
	typ string
	name string
	repeated bool
}

type ProtoGen struct {
	files []string
	mds map[int32]*desc.MessageDescriptor
	fids map[string][]*desc.FieldDescriptor
	id2names map[int32]string

}

func newProtoGen() *ProtoGen {
	return &ProtoGen{
		mds: make(map[int32]*desc.MessageDescriptor),
		id2names: make(map[int32]string),
		fids: make(map[string][]*desc.FieldDescriptor),
	}
}

func readTemplate(tpl string) string {
	content, err := ioutil.ReadFile(tpl)
	if (err != nil) {
		log.Fatal(err)
	}

	template := string(content[:])
	return template
}

func writeFile(s string, f string) {
    file, err := os.Create(f)
    if err != nil {
        return
    }
    defer file.Close()

    file.WriteString(s)
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
				options := md.GetOptions()
				mid, _ := proto.GetExtension(options, pb.E_Msgid)
				if (mid != nil) {
					g.mds[ *mid.(*int32)] = md
					g.id2names[*mid.(*int32)] = md.GetName()
					var fids []*desc.FieldDescriptor
					fids = append(fids, md.GetFields()...)
					g.fids[md.GetName()] = fids
				}
			}
		}
	}


	template := readTemplate("template/msgid.tpl")
	ctx := plush.NewContext()

	ctx.Set("id2names", g.id2names)

	s, err := plush.Render(template, ctx)
	if (err != nil) {
		log.Fatal(err)
	}

	writeFile(s, "msg/msgid.go")

}

func (g *ProtoGen) generate() {
	template := readTemplate("template/funcdecl.tpl")
	ctx := plush.NewContext()
	var protos []string
	for _, p := range g.id2names {
		protos = append(protos, p)
	}

	ctx.Set("package", "package msg")
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

	var tm map[string][]Field = make(map[string][]Field)
	
	for name, fields := range g.fids {
		var fids []Field
		for _, field := range fields {
			t := typeMap[field.GetType().String()]

			var fld Field
			
			if field.GetMessageType() != nil {
				fld.typ = field.GetMessageType().GetName()
			} else if field.GetEnumType() != nil {
				fld.typ = field.GetEnumType().GetName()
			} else {
				fld.typ = t
			}

			fld.ptype = t

			fld.name = field.GetName()
			// fmt.Printf("%s, %s, %s\n", fld.typ, fld.ptype, fld.name)

			if field.IsRepeated() {
				fld.repeated = true
			}

			fids = append(fids, fld)
		}

		tm[name] = fids
	}


	ctx.Set("fields", func(name string) []string {
		var fns []string
		fields := g.fids[name]
		for _, f := range fields {
			fns = append(fns, f.GetName())
		}
		
		return fns
	})

	ctx.Set("params", func(msg string) string {					
		fields := tm[msg]
		var sb strings.Builder
		c := len(fields)
		var i int = 0
		for _, field := range fields {
			k := field.typ
			v := field.name
			pt := field.ptype
			repeated := field.repeated

			sb.WriteString(v)
			sb.WriteString(" ")

			if pt == "msg" || pt == "enum" {
				var sbb strings.Builder
				if pt == "msg" {
					if repeated {
						sbb.WriteString("[]*pb.")						
					} else {
						sbb.WriteString("*pb.")						
					}
				} else {
					sbb.WriteString("pb.")
				}

				// cname := strings.Title(k)
				sbb.WriteString(k)
	
				sb.WriteString(sbb.String())

			} else {
				if repeated {
					sb.WriteString("[]")
				}

				sb.WriteString(k)
			}

			if i < c-1 {
				sb.WriteString(", ")
			}
			i++
		}

		return sb.String()
	})

	s, err := plush.Render(template, ctx)
	if (err != nil) {
		log.Fatal(err)
	}
	
	writeFile(s, "msg/serializer.go")	

}






