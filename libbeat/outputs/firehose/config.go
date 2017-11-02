package firehose

import (
	"time"
)

type FirehoseConfig struct {
	Region     string             `config:"region"`
	BatchSize  int                `config:"batch_size"`
	MaxRetries int                `config:"max_retries"`
	Timeout    time.Duration      `config:"timeout"`
	Backoff    Backoff            `config:"backoff"`
}

type Backoff struct {
	Init time.Duration
	Max  time.Duration
}

const (
	defaultBatchSize = 50
)

var (
	defaultConfig = FirehoseConfig{
		Region:     "us-east-1",
		Timeout:    90 * time.Second,
		MaxRetries: 3,
		Backoff: Backoff{
			Init: 1 * time.Second,
			Max:  60 * time.Second,
		},
	}
)

func (c *FirehoseConfig) Validate() error {
	return nil
}
