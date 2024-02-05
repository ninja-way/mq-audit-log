package transport

import (
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

type Server struct {
	AuditServer *AuditServer

	conn *amqp.Connection
	ch   *amqp.Channel
	q    amqp.Queue
}

func NewServer(MqUri string, auditServer *AuditServer) *Server {
	conn, err := amqp.Dial(MqUri)
	if err != nil {
		log.Fatal("failed to connect to rabbitmq")
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("failed to open a channel")
	}

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

	return &Server{
		AuditServer: auditServer,
		conn:        conn,
		ch:          ch,
		q:           q,
	}
}

func (s *Server) CloseServerConnection() {
	s.ch.Close()
	s.conn.Close()
}

func (s *Server) StartListen() {
	msgs, err := s.ch.Consume(
		s.q.Name, // queue
		"",       // consumer
		true,     // auto-ack
		false,    // exclusive
		false,    // no-local
		false,    // no-wait
		nil,      // args
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
