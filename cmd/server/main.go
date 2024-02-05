package main

import (
	"context"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/ninja-way/mq-audit-log/internal/config"
	"github.com/ninja-way/mq-audit-log/internal/repository"
	"github.com/ninja-way/mq-audit-log/internal/service"
	"github.com/ninja-way/mq-audit-log/internal/transport"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	db, err := pgx.Connect(ctx, cfg.DBConn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close(ctx)

	auditRepo := repository.NewAudit(db)
	auditService := service.NewAudit(auditRepo)
	auditSrv := transport.NewAuditServer(auditService)

	srv := transport.NewServer(cfg.MQ.URI, auditSrv)
	defer srv.CloseServerConnection()
	srv.StartListen()
}
