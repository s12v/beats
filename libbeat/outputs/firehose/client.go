package firehose

import (
	"time"
	"github.com/elastic/beats/libbeat/outputs"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/elastic/beats/libbeat/publisher"
	"fmt"
	"github.com/elastic/beats/libbeat/logp"
)

type Client struct {
	client *firehose.Firehose

	timeout time.Duration

	stats *outputs.Stats
}

func NewClient(config *FirehoseConfig) (*Client, error) {

	// TODO: reuse session
	s := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(config.Region),
	}))

	// Create a Firehose client from just a session.
	svc := firehose.New(s)

	client := &Client{
		client: svc,
	}

	return client, nil
}

func (client *Client) Close() error {
	return nil
}

func (client *Client) Connect() error {
	return nil
}

func (client *Client) Publish(batch publisher.Batch) error {
	events := batch.Events()
	for i := range events {
		event := &events[i].Content
		logp.Info(fmt.Sprintf("event %d, time: %s", i, event.Timestamp))
	}

	return nil

}
