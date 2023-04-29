package main

import  (
    "github.com/springstar/protogen/pb"
    "github.com/jhump/protoreflect/desc"
    "github.com/jhump/protoreflect/dynamic"   
)

var descriptors map[int32]*desc.MessageDescriptor = make(map[int32]*desc.MessageDescriptor)

func AddDescriptor(id int32, desc *desc.MessageDescriptor) {
    descriptors[id] = desc
}


func parseCSTest(id int32, bytes []byte) *pb.CSTest {
    msg :=  &pb.CSTest{}
    md := descriptors[id]
    dmsg := dynamic.NewMessage(md)
    dmsg.Unmarshal(bytes)
    dmsg.ConvertTo(msg)
    return msg
}

func parseCSLogin(id int32, bytes []byte) *pb.CSLogin {
    msg :=  &pb.CSLogin{}
    md := descriptors[id]
    dmsg := dynamic.NewMessage(md)
    dmsg.Unmarshal(bytes)
    dmsg.ConvertTo(msg)
    return msg
}

func parseAddress(id int32, bytes []byte) *pb.Address {
    msg :=  &pb.Address{}
    md := descriptors[id]
    dmsg := dynamic.NewMessage(md)
    dmsg.Unmarshal(bytes)
    dmsg.ConvertTo(msg)
    return msg
}




func serializeCSTest(sex int32, user  *pb.User, money int64, code string) *pb.CSTest {
    return nil
}

func serializeCSLogin(code string, sex int32, user  *pb.User, money int64) *pb.CSLogin {
    return nil
}

func serializeAddress(user  *pb.User, money int64, code string, sex int32) *pb.Address {
    return nil
}
