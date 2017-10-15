package memory

import (
	"app/pkg/domain"
	"context"
	"encoding/json"

	"github.com/vardius/golog"
	messagebus "github.com/vardius/message-bus"
)

type commandBus struct {
	messageBus messagebus.MessageBus
	logger     golog.Logger
}

func (bus *commandBus) Publish(command string, ctx context.Context, payload json.RawMessage, out chan<- error) {
	bus.logger.Debug(ctx, "[API CommandBus|Publish]: %s %q\n", command, payload)
	bus.messageBus.Publish(command, ctx, payload, out)
}
func (bus *commandBus) Subscribe(command string, fn domain.CommandHandler) error {
	bus.logger.Info(nil, "[API CommandBus|Subscribe]: %s\n", command)
	return bus.messageBus.Subscribe(command, fn)
}
func (bus *commandBus) Unsubscribe(command string, fn domain.CommandHandler) error {
	bus.logger.Info(nil, "[API CommandBus|Unsubscribe]: %s\n", command)
	return bus.messageBus.Unsubscribe(command, fn)
}

func NewCommandBus(log golog.Logger) domain.CommandBus {
	return &commandBus{messageBus, log}
}