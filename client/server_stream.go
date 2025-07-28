package main

import (
	"context"
	"io"
	"log"

	pb "github.com/surajgoraicse/go-grpc-basic/proto"
)

func callSayHelloServerStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	log.Println("Streaming started")
	stream, err := client.SayHelloServerStreaming(context.Background(), names)
	if err != nil {
		log.Fatalln("could not send names ", err)
	}
	// printing the stream
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil{
			log.Fatalln("Error while streaming ", err)
		} 
		log.Println(message)

	}
	log.Println("Streaming finished")

}
