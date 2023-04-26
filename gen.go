package main

import (
	"fmt"
	"strings"
	"log"

	"io/ioutil"
	"path/filepath"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/golang/protobuf/proto"
	"github.com/springstar/protogen/pb"
)

type ProtoGen struct {
	files []string
	mds map[int32]*desc.MessageDescriptor

}

func newProtoGen() *ProtoGen {
	return &ProtoGen{
		mds: make(map[int32]*desc.MessageDescriptor),
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
					fmt.Println(md)
				}
			}
		}
	}


}



