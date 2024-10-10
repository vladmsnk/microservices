package device

import (
	"context"
	"errors"
	"fmt"

	"device-management-service/internal/entities"
	"device-management-service/internal/storage"
)

func (i *Implementation) GetDeviceInfo(ctx context.Context, name string) (entities.Device, error) {
	device, err := i.storage.GetDeviceByName(ctx, name)
	if err != nil {
		if errors.Is(err, storage.ErrNotFound) {
			return entities.Device{}, ErrDeviceNotFound
		}
		return entities.Device{}, fmt.Errorf("i.storage.GetDeviceByName: %w", err)
	}

	return device, nil
}
