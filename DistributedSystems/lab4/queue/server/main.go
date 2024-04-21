package main

import (
	"context"
	"log"
	"time"

	"github.com/google/uuid"
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
	_, err = stream.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:      "PRODUCT_USAGE_1",
		Retention: jetstream.WorkQueuePolicy,
		Subjects:  []string{"PRODUCT_USAGE_1.*.*"},
		Storage:   jetstream.FileStorage,
		Replicas:  1,
		MaxAge:    time.Second * 10,
	})
	if err != nil {
		panic(err)
	}

	counter := atomic.NewInt32(0)
	workers := 1000
	msgCount := 100_000
	for i := 0; i < workers; i++ {
		go func() {
			for i := 0; i < msgCount; i++ {
				counter.Inc()
				_, err := stream.Publish(ctx, "PRODUCT_USAGE_1.subscription."+uuid.NewString(), []byte(uuid.NewString()))
				if err != nil {
					panic(err)
				}
				if counter.Load()%1000 == 0 {
					log.Println("Published subscription", counter.Load())
				}
			}
		}()
	}
	time.Sleep(time.Second * 50)
}
