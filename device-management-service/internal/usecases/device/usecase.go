package device

import (
	"context"
	"fmt"

	"device-management-service/internal/entities"
)

var (
	ErrDeviceAlreadyExists = fmt.Errorf("device already exists")
	ErrDeviceNotFound      = fmt.Errorf("device not found")
)

type UseCase interface {
	RegisterDevice(ctx context.Context, device entities.Device) error
	GetDeviceInfo(ctx context.Context, name string) (entities.Device, error)
}

type Storage interface {
	CreateDevice(ctx context.Context, device entities.Device) error
	GetDeviceByName(ctx context.Context, name string) (entities.Device, error)
}

type UserRepository interface {
	GetUserByUsername(ctx context.Context, username string) (entities.User, error)
}

type HomeRepository interface {
	GetHomeByName(ctx context.Context, name string) (entities.Home, error)
}

type Implementation struct {
	storage  Storage
	userRepo UserRepository
	homeRepo HomeRepository
}

func New(storage Storage, userRepo UserRepository) *Implementation {
	return &Implementation{storage: storage, userRepo: userRepo}
}
