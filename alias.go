package cloudeventsaws

import (
	"github.com/loopcontext/cloudevents-aws-transport/eventbridge"
	"github.com/loopcontext/cloudevents-aws-transport/sns"
	"github.com/loopcontext/cloudevents-aws-transport/sqs"
)

var (
	// SQS
	NewSQSTransport = sqs.New

	// SNS
	NewSNSTransport = sns.New
	WithPort        = sns.WithPort
	WithPath        = sns.WithPath

	// EventBridge
	NewEventBridgeTransport = eventbridge.New
)
