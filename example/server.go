package main

import (
	"context"
	"log"

	cloudevents "github.com/cloudevents/sdk-go"
	cloudeventsaws "github.com/loopcontext/cloudevents-aws-transport"
)

func Receive(event cloudevents.Event) {
	log.Printf("? %v", event)
	// do something with event.Context and event.Data (via event.DataAs(foo)
}

func main() {
	t, err := cloudeventsaws.NewSQSTransport("https://sqs.eu-central-1.amazonaws.com/770719141847/rdk-dev-orm-events")
	// t, err := cloudeventsaws.NewSNSTransport("arn:aws:sns:eu-central-1:458470902217:test", cloudeventsaws.WithPort(8081))
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
	// 	event.SetType("com.cloudevents.readme.sent")
	// 	event.SetSource("test")
	// 	event.SetTime(time.Now())
	// 	event.SetData(map[string]string{
	// 		"message": fmt.Sprintf("hello world %d", i),
	// 	})
	// 	log.Println(c.Send(context.Background(), event))
	// }

	log.Fatal(c.StartReceiver(context.Background(), Receive))
}
