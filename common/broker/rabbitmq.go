package common

import (
	"fmt"

	"github.com/streadway/amqp"
)

// Connect to RabbitMQ
func ConnectMQ(cfg *RabbitMQ) (ch *amqp.Channel, err error) {
	conn, err := amqp.Dial("amqp://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err = conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open a channel: %v", err)
	}

	return ch, nil
}
