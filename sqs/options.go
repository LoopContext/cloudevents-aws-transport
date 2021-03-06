package sqs

import "github.com/loopcontext/cloudevents-aws-transport/encoding"

// Option is the function signature required to be considered an nats.Option.
type Option func(*Transport) error

// WithEncoding sets the encoding for NATS transport.
func WithEncoding(encoding encoding.Encoding) Option {
	return func(t *Transport) error {
		t.Encoding = encoding
		return nil
	}
}
