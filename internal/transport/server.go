package transport

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type Server struct {
	AuditServer *AuditServer
	MqUri       string
}

func NewServer(MqUri string, auditServer *AuditServer) *Server {
	return &Server{
		AuditServer: auditServer,
		MqUri:       MqUri,
	}
}

func (s *Server) StartListen() {
	conn, err := amqp.Dial(s.MqUri)
	if err != nil {
		log.Fatal("failed to connect to rabbitmq")
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("failed to open a channel")
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"logs", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	if err != nil {
		log.Fatal("failed to declare a queue")
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatal("failed to register a consumer")
	}

	ctx := context.Background()
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			err := s.AuditServer.Log(ctx, d.Body)
			if err != nil {
				log.Printf("failed to log message: %s", err)
			}
		}
	}()

	log.Printf(" [*] Server started. To exit press CTRL+C")
	<-forever
}
