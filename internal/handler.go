package internal

import (
	"fmt"

	"github.com/ThreeDotsLabs/watermill/message"
)

type PdfDownloadHandler struct{}

type TextDownloadHandler struct{}

type BioTransferHandler struct{}

// HandlerWithoutPublish receives and handles subscribe message without publishing it again
func (h PdfDownloadHandler) HandlerWithoutPublish(msg *message.Message) error {
	fmt.Printf(
		"\n> PdfDownloadHandler received message: %s\n> %s\n> metadata: %v\n",
		msg.UUID, string(msg.Payload), msg.Metadata,
	)
	return nil
}

// HandlerWithoutPublish receives and handles subscribe message without publishing it again
func (h TextDownloadHandler) HandlerWithoutPublish(msg *message.Message) error {
	fmt.Printf(
		"\n> TextDownloadHandler received message: %s\n> %s\n> metadata: %v\n",
		msg.UUID, string(msg.Payload), msg.Metadata,
	)
	return nil
}

// HandlerWithoutPublish receives and handles subscribe message without publishing it again
func (h BioTransferHandler) HandlerWithoutPublish(msg *message.Message) error {
	fmt.Printf(
		"\n> BioTransferHandler received message: %s\n> %s\n> metadata: %v\n",
		msg.UUID, string(msg.Payload), msg.Metadata,
	)
	return nil
}
