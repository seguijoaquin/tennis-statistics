package broker

import (
	"log"

	"github.com/streadway/amqp"
)

// MessageBroker represents a queue of messages
type MessageBroker struct {
	amqpChan     *amqp.Channel
	amqpConn     *amqp.Connection
	messageQueue *amqp.Queue
}

// GetBroker creates a new queue if it does not exists and retrieves it
func GetBroker(name string) (*MessageBroker, error) {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	if err := ch.ExchangeDeclare(
		name,    // name args[1]
		"topic", // type
		true,    // durable
		false,   // auto-deleted
		false,   // internal
		false,   // no-wait
		nil,     // arguments
	); err != nil {
		return nil, err
	}

	return &MessageBroker{
		amqpChan: ch,
		amqpConn: conn,
	}, nil
}

// Post to Broker
func (b *MessageBroker) Post(exchange string, routingKey string, body []byte) error {
	return b.amqpChan.Publish(
		exchange,   // exchange
		routingKey, // routing key
		false,      // mandatory
		false,      // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
}

// DeclareUnnamedMessageQueue returns an unnamed queue
func (b *MessageBroker) DeclareUnnamedMessageQueue() error {
	q, err := b.amqpChan.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	if err != nil {
		return err
	}

	b.messageQueue = &q

	return nil
}

// BindTopicsToQueue binds a list of topics to MessageBroker Queue
func (b *MessageBroker) BindTopicsToQueue(exchangeName string, topics []string) error {

	for _, s := range topics {
		log.Printf("Binding queue %s to exchange %s with routing key %s",
			b.messageQueue.Name, exchangeName, s)
		err := b.amqpChan.QueueBind(
			b.messageQueue.Name, // queue name
			s,                   // routing key
			exchangeName,        // exchange
			false,
			nil)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetMessages registers a new unnamed consumer and returns a channel to consume new Messages
func (b *MessageBroker) GetMessages() (<-chan amqp.Delivery, error) {
	return b.amqpChan.Consume(
		b.messageQueue.Name, // queue
		"",                  // consumer
		true,                // auto ack
		false,               // exclusive
		false,               // no local
		false,               // no wait
		nil,                 // args
	)
}

// Close to free resources
func (b *MessageBroker) Close() {
	b.amqpChan.Close()
	b.amqpConn.Close()
}
