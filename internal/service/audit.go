package service

import (
	"context"

	audit "github.com/ninja-way/mq-audit-log/pkg/models"
)

type Repository interface {
	Insert(ctx context.Context, item audit.LogItem) error
}

type Audit struct {
	repo Repository
}

func NewAudit(repo Repository) *Audit {
	return &Audit{
		repo: repo,
	}
}

func (s *Audit) Insert(ctx context.Context, item audit.LogItem) error {
	return s.repo.Insert(ctx, item)
}
