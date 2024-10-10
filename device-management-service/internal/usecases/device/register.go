package device

import (
	"context"
	"device-management-service/internal/entities"
	"device-management-service/internal/storage"
	"errors"
	"fmt"
)

func (i *Implementation) RegisterDevice(ctx context.Context, device entities.Device) error {
	_, err := i.userRepo.GetUserByUsername(ctx, device.User.Name)
	if err != nil {
		return fmt.Errorf("i.userRepo.GetUserByUsername: %w", err)
	}

	_, err = i.homeRepo.GetHomeByName(ctx, device.Home.Name)
	if err != nil {
		return fmt.Errorf("i.homeRepo.GetHomeByName: %w", err)
	}

	_, err = i.storage.GetDeviceByName(ctx, device.Name)
	if err != nil {
		if !errors.Is(err, storage.ErrNotFound) {
			return fmt.Errorf("i.storage.GetDeviceByName: %w", err)
		}
	} else {
		return ErrDeviceAlreadyExists
	}

	err = i.storage.CreateDevice(ctx, device)
	if err != nil {
		return fmt.Errorf("i.storage.CreateDevice: %w", err)
	}

	return nil
}
