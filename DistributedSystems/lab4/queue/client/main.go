package main

import (
	"context"
	"log"
	"time"

	"github.com/dustin/go-humanize"
	"github.com/nats-io/nats.go"
	jetstream "github.com/nats-io/nats.go/jetstream"
	"go.uber.org/atomic"
)

func main() {
	ctx := context.Background()
	nc, err := nats.Connect("nats://host.docker.internal:4222")
	if err != nil {
		panic(err)
	}
	stream, err := jetstream.New(nc)
	if err != nil {
		panic(err)
	}
	counter := atomic.NewInt32(0)
	workers := 100
	for i := 0; i < workers; i++ {
		go func() {
			consumer, err := stream.CreateOrUpdateConsumer(ctx, "PRODUCT_USAGE_1", jetstream.ConsumerConfig{
				AckPolicy: jetstream.AckExplicitPolicy,
				Durable:   "TEST",
			})
			if err != nil {
				panic(err)
			}
			for {
				msg, err := consumer.FetchBytes(humanize.MiByte * 1)
				if err != nil || msg.Error() != nil {
					panic(coalesceErr(err, msg.Error()))
				}
				for m := range msg.Messages() {
					err := m.Ack()
					if err != nil {
						panic(err)
					}
					counter.Inc()
					if counter.Load()%1000 == 0 {
						log.Printf("Received %d messages", counter.Load())
					}
				}
			}
		}()
	}
	time.Sleep(100 * time.Second)
}

func coalesceErr(data ...error) error {
	for _, item := range data {
		if item != nil {
			return item
		}
	}
	return nil
}
