package main

// WATERMILL_PUBSUB_TYPE=rabbitmq NATS_CLUSTER_ID=test-cluster NATS_URL=nats://localhost:4223 AMQP_URL=amqp://guest:guest@localhost:5672/ go run main.go
// WATERMILL_PUBSUB_TYPE=nats NATS_CLUSTER_ID=test-cluster NATS_URL=nats://localhost:4223 go run main.go
import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/components/metrics"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/sean830314/model-data-collector/internal"
)

var (
	metricsAddr         = ":8081"
	watermillPubsubType = os.Getenv("WATERMILL_PUBSUB_TYPE")
	natsClusterID       = os.Getenv("NATS_CLUSTER_ID")
	natsURL             = os.Getenv("NATS_URL")
	// You probably want to ship your own implementation of `watermill.LoggerAdapter`.
	logger     = watermill.NewStdLogger(false, false) // debug=false, trace=false
	dlaSubject = "dla_subject"
	piiSubject = "pii_subject"
)

func main() {
	router, err := message.NewRouter(message.RouterConfig{}, logger)
	if err != nil {
		log.Fatal(err)
	}

	promRegistry, closeMetricsServer := metrics.CreateRegistryAndServeHTTP(metricsAddr)
	defer closeMetricsServer()

	metricsBuilder := metrics.NewPrometheusMetricsBuilder(promRegistry, "demo", "hello")
	metricsBuilder.AddPrometheusRouterMetrics(router)

	// SignalsHandler gracefully shutdowns Router when receiving SIGTERM
	router.AddPlugin(plugin.SignalsHandler)

	// Router level middleware are executed for every message sent to the router
	router.AddMiddleware(
		// CorrelationID will copy the correlation id from the incoming message's metadata to the produced messages
		middleware.CorrelationID,
		// Timeout makes the handler cancel the incoming message's context after a specified time
		middleware.Timeout(time.Second*10),
		// Throttle provides a middleware that limits the amount of messages processed per unit of time
		middleware.NewThrottle(10, time.Second).Middleware,
		// After MaxRetries, the message is Nacked and it's up to the PubSub to resend it
		middleware.Retry{
			MaxRetries: 5,
			Logger:     logger,
		}.Middleware,

		// Recoverer handles panics from handlers
		middleware.Recoverer,
	)

	var publisher message.Publisher
	var subscriber message.Subscriber
	if watermillPubsubType == "nats" {
		publisher, err = internal.NewNATSPublisher(logger, natsClusterID, natsURL)
		if err != nil {
			log.Fatal(err)
		}
		subscriber, err = internal.NewNATSSubscriber(logger, natsClusterID, watermill.NewShortUUID(), natsURL)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Run Nats Streaming mode")
	} else {
		pubsub := internal.NewGoChannel(logger)
		publisher = pubsub
		subscriber = pubsub
		fmt.Println("Run Go Channel mode")
	}

	router.AddNoPublisherHandler(
		"pdf_download_handler",
		dlaSubject,
		subscriber,
		internal.PdfDownloadHandler{}.HandlerWithoutPublish,
	)

	router.AddNoPublisherHandler(
		"text_download_handler",
		piiSubject,
		subscriber,
		internal.TextDownloadHandler{}.HandlerWithoutPublish,
	)

	router.AddNoPublisherHandler(
		"bot_transfer_handler",
		piiSubject,
		subscriber,
		internal.BioTransferHandler{}.HandlerWithoutPublish,
	)

	// Producing some incoming messages in background
	go publishMessages(piiSubject, publisher, 10)
	go publishMessages(dlaSubject, publisher, 7)
	// Run the router
	ctx := context.Background()
	if err := router.Run(ctx); err != nil {
		log.Fatal(err)
	}
}

func publishMessages(topic string, publisher message.Publisher, delay time.Duration) {
	for {
		msg := message.NewMessage(watermill.NewUUID(), []byte(fmt.Sprintf("topic: %s, time: %s\n", topic, time.Now().String())))
		middleware.SetCorrelationID(watermill.NewUUID(), msg)

		fmt.Printf("\n\n\nSending message %s, correlation id: %s\n", msg.UUID, middleware.MessageCorrelationID(msg))
		if err := publisher.Publish(topic, msg); err != nil {
			log.Fatal(err)
		}
		time.Sleep(delay * time.Second)
	}
}
