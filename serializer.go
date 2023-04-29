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

func parseCSTest(id int32, bytes []byte) *pb.CSTest {
    msg :=  &pb.CSTest{}
    md := descriptors[id]
    dmsg := dynamic.NewMessage(md)
    dmsg.Unmarshal(bytes)
    dmsg.ConvertTo(msg)
    return msg
}




func serializeCSLogin(account string, password string, token string, serverId int32, version int32) *pb.CSLogin {
    msg := &pb.CSLogin{

    }
    return msg
}

func serializeAddress(state string, province string, city string, code int32, user *pb.User, sex pb.Sex) *pb.Address {
    msg := &pb.Address{
		User: user,
		Sex: sex,
    }
    return msg
}

func serializeCSTest(code string, money int64) *pb.CSTest {
    msg := &pb.CSTest{

    }
    return msg
}
