package main

import (
	"fmt"
	"log"

	"github.com/Rosalita/protobuf/pb"
	"google.golang.org/protobuf/proto"
)

func main() {

	// To generate a protobuf message:
	// 1. Instantiate the type from the protoc generated code.
	person := pb.Person{
		Name: "Rosie",
		Age:  123,
	}

	// 2. Encode the data with proto.Marshal
	out, err := proto.Marshal(&person)
	if err != nil {
		log.Fatal(err)
	}

	// The generated message is []byte
	fmt.Println(out) // [10 5 82 111 115 105 101 16 123]

	// To read a protobuf message:
	in := []byte{10, 5, 82, 111, 115, 105, 101, 16, 123}

	// 1. Create an empty result type
	var result pb.Person

	// 2. Use proto.Unmarshal to marshal the data into the result.
	if err := proto.Unmarshal(in, &result); err != nil {
		log.Fatal(err)
	}

	fmt.Println(result.String()) // name:"Rosie" age:123
}
