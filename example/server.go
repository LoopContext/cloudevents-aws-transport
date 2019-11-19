package main

import (
	"context"
	"log"

	cloudevents "github.com/cloudevents/sdk-go"
	"github.com/jakubknejzlik/cloudevents-aws-transport"
)

func Receive(event cloudevents.Event) {
	log.Printf("? %v", event)
	// do something with event.Context and event.Data (via event.DataAs(foo)
}

func main() {
	// t, err := cloudeventsaws.NewSQSTransport("https://sqs.eu-central-1.amazonaws.com/458470902217/sqs-queue-test")
	t, err := cloudeventsaws.NewSNSTransport("arn:aws:sns:eu-central-1:458470902217:test",cloudeventsaws.WithPort(8081))
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}
	c, err := cloudevents.NewClient(t)
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}

	// for i := 0; i < 1; i++ {
	// 	event := cloudevents.NewEvent()
	// 	event.SetID(fmt.Sprintf("test123 %d", i))
	// 	event.SetType("com.cloudevents.readme.sent")
	// 	event.SetSource("http://localhost:8080/")
	// 	event.SetData(fmt.Sprintf("hello world %d", i))
	// 	log.Println(c.Send(context.Background(), event))
	// }

	log.Fatal(c.StartReceiver(context.Background(), Receive))
}
