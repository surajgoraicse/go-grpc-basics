package main

import (
	"io"
	"log"

	pb "github.com/surajgoraicse/go-grpc-basic/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string
	// processing the stream
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Message: messages})
		}
		if err != nil {
			return err
		}
		log.Println("Got request with name : ", req.Name)
		messages = append(messages, "Hello "+req.Name)
	}
}
