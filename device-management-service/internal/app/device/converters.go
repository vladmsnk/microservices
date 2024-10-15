package device

import "device-management-service/internal/entities"

func toDeviceEntity(request RegisterDeviceRequest) entities.Device {
	return entities.Device{
		Name:        request.DeviceName,
		Home:        entities.Home{Name: request.HomeName},
		User:        entities.User{Name: request.Username},
		Description: request.DeviceDescription,
		Type:        entities.DeviceTypeFromString(request.DeviceType),
		Status:      entities.DeviceStatusUnspecified,
	}
}
