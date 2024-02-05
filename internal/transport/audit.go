package transport

import (
	"context"
	"encoding/json"
	"errors"
	audit "github.com/ninja-way/mq-audit-log/pkg/models"
)

type AuditService interface {
	Insert(ctx context.Context, logItem audit.LogItem) error
}

type AuditServer struct {
	service AuditService
}

func NewAuditServer(service AuditService) *AuditServer {
	return &AuditServer{
		service: service,
	}
}

func (h *AuditServer) Log(ctx context.Context, logMsg []byte) error {
	logItem := audit.LogItem{}
	err := json.Unmarshal(logMsg, &logItem)
	if err != nil {
		return errors.New("bad log message")
	}

	return h.service.Insert(ctx, logItem)
}
