package main

import (
	"context"
	"log"
	// "fmt"
	// "time"

	cloudevents "github.com/cloudevents/sdk-go"
	cloudeventsaws "github.com/loopcontext/cloudevents-aws-transport"
)

// Receive ..
func Receive(event cloudevents.Event) {
	log.Printf("? %v", event)
	// do something with event.Context and event.Data (via event.DataAs(foo)
}

func main() {
	t, err := cloudeventsaws.NewSQSTransport("https://sqs.us-east-2.amazonaws.com/381430815459/loopcontext-test")
	// t, err := cloudeventsaws.NewSNSTransport("arn:aws:sqs:us-east-2:381430815459:loopcontext-test", cloudeventsaws.WithPort(8081))
	// t, err := cloudeventsaws.NewEventBridgeTransport("arn:aws:events:eu-central-1:458470902217:event-bus/test")
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
	// 	event.SetType("com.loopcontext.readme.sent")
	// 	event.SetSource("test")
	// 	event.SetTime(time.Now())
	// 	event.SetData(map[string]string{
	// 		"message": fmt.Sprintf("hello world %d", i),
	// 	})
	// 	log.Println(c.Send(context.Background(), event))
	// }

	log.Fatal(c.StartReceiver(context.Background(), Receive))
}
