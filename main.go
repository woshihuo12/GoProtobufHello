package main

import (
"github.com/woshihuo12/GoProtobufHello/test"
	"github.com/golang/protobuf/proto"
	"log"
)

func main() {
	test1 := &example.Test{
		Id:1,
		Label:"hell000o",
		Type:13,

	}

	data, err := proto.Marshal(test1)
	if err!=nil{
		log.Fatal("marshling error: ", err)
	}

	test2 := &example.Test{}

	err = proto.Unmarshal(data, test2)
	if err != nil {
		log.Fatal("unmarsh1 error", err)
	}

	log.Println(test2.Id)
	log.Println(test2.Label)
	log.Println(test2.Type)
}
