package storage

import (
	"context"
	"fmt"

	"device-management-service/internal/entities"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	ErrNotFound = fmt.Errorf("not found")
)

type Storage struct {
	db *pgxpool.Pool
}

func New(db *pgxpool.Pool) *Storage {
	return &Storage{db: db}
}

func (s *Storage) CreateDevice(ctx context.Context, device entities.Device) error {
	return nil
}
