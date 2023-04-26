package main

import (
	"github.com/springstar/protogen/pb"
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/dynamic"
)

func TestGen(t *testing.T) {
	g := newProtoGen()
	g.parse("msg/protocol")

	for k, v := range g.mds {
		fmt.Println(k)
		switch k {
		case 111:
			testCSLoing(t, v)
		case 101:	
		}
	}
}

func testCSLoing(t *testing.T, m *desc.MessageDescriptor) {
	packet := getLoginPacket()
	login := dynamicMsg(packet, &pb.CSLogin{}, m).(*pb.CSLogin)

	assert.Equal(t, "test", login.GetAccount())
	assert.Equal(t, "123456", login.GetPassword())
	assert.Equal(t, "sjdflkas12", login.GetToken())
	assert.Equal(t, int32(1), login.GetServerId())
	assert.Equal(t, int32(1035), login.GetVersion())


}

func TestDispatch(t *testing.T) {
	g := newProtoGen()
	g.parse("msg/protocol")

	msgid := 111
	md := g.mds[int32(msgid)]
	assert.NotNil(t, md)


}

func getLoginPacket() []byte {
	msg := &pb.CSLogin{
		Account: "test",
		Password: "123456",
		Token: "sjdflkas12",
		ServerId: 1,
		Version: 1035,
	}

	packet, _ := proto.Marshal(msg)
	return packet
}

func dynamicMsg(packet []byte, msg proto.Message, md *desc.MessageDescriptor) proto.Message {
	dmsg := dynamic.NewMessage(md)
	dmsg.Unmarshal(packet)
	dmsg.ConvertTo(msg)
	
	return msg
}

// parseCSLogin(bytes []byte) {

//}