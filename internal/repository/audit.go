package repository

import (
	"context"
	"github.com/jackc/pgx/v5"

	audit "github.com/ninja-way/mq-audit-log/pkg/models"
)

type Audit struct {
	db *pgx.Conn
}

func NewAudit(db *pgx.Conn) *Audit {
	return &Audit{
		db: db,
	}
}

func (r *Audit) Insert(ctx context.Context, item audit.LogItem) error {
	_, err := r.db.Exec(ctx, "INSERT INTO logs (entity, action, entity_id, timestamp) VALUES ($1, $2, $3, $4)",
		item.Entity, item.Action, item.EntityID, item.Timestamp)
	return err
}
