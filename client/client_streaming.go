package main

import (
	"context"
	"log"
	"time"

	pb "github.com/surajgoraicse/go-grpc-basic/proto"
)

// sending stream of name to the server
func callSayHelloClientStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	log.Println("Client streaming started")

	stream ,err :=client.SayHelloClientStreaming(context.Background())
	
	if err != nil {
		log.Fatalln("count not send names from client side ", err)
	}
	// sending name one by one
	for _, name := range names.Names{
		req := &pb.HelloRequest{
			Name: name,
		}
		if err :=stream.Send(req); err != nil {
			log.Fatalln("Error while sending ", err)
		}
		log.Println("Sending the request with name : ", name)
		time.Sleep(1 * time.Second)
	}

	// server sends some response after the completion of the client streaming
	res, err := stream.CloseAndRecv()
	log.Println("Client streaming finished")
	if err != nil {
		log.Fatalln("Error while receiving ", err)
	}
	log.Println(res.Message)
}
