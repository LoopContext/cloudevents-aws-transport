package cloudeventsaws

import (
	"github.com/jakubknejzlik/cloudevents-aws-transport/sns"
	"github.com/jakubknejzlik/cloudevents-aws-transport/sqs"
)

var (
	// SQS
	NewSQSTransport = sqs.New

	// SNS
	NewSNSTransport = sns.New
	WithPort        = sns.WithPort
	WithPath        = sns.WithPath
)
