package main

func main() {
	g := newProtoGen()
	g.parse("msg/protocol")
	g.generate()
}