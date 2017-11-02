package firehose

import (
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/outputs"
	"github.com/elastic/beats/libbeat/beat"
)

func init() {
	outputs.RegisterType("firehose", New)
}

func New(
	beat beat.Info,
	stats *outputs.Stats,
	cfg *common.Config,
) (outputs.Group, error) {
	if !cfg.HasField("batch_size") {
		cfg.SetInt("batch_size", -1, defaultBatchSize)
	}

	config := defaultConfig
	if err := cfg.Unpack(&config); err != nil {
		return outputs.Fail(err)
	}

	var client outputs.NetworkClient
	client, err := NewClient(&config)
	if err != nil {
		return outputs.Fail(err)
	}

	client = outputs.WithBackoff(client, config.Backoff.Init, config.Backoff.Max)
	return outputs.Success(config.BatchSize, config.MaxRetries, client)
}
