package entities

type DeviceStatus int64

const (
	DeviceStatusUnspecified DeviceStatus = 0
	DeviceStatusOn          DeviceStatus = 1
	DeviceStatusOff         DeviceStatus = 2
)

func DeviceStatusFromString(deviceStatus string) DeviceStatus {
	switch deviceStatus {
	case DeviceStatusOn.String():
		return DeviceStatusOn
	case DeviceStatusOff.String():
		return DeviceStatusOff
	default:
		return DeviceStatusUnspecified
	}
}

func (s DeviceStatus) String() string {
	switch s {
	case DeviceStatusOn:
		return "on"
	case DeviceStatusOff:
		return "off"
	default:
		return ""
	}
}

type DeviceType int64

const (
	DeviceTypeUnspecified DeviceType = 0
	DeviceTypeLight       DeviceType = 1
	DeviceTypeTemperature DeviceType = 2
)

func (t DeviceType) String() string {
	switch t {
	case DeviceTypeLight:
		return "light"
	case DeviceTypeTemperature:
		return "temperature"
	default:
		return ""
	}
}

func DeviceTypeFromString(deviceType string) DeviceType {
	switch deviceType {
	case "light":
		return DeviceTypeLight
	case "temperature":
		return DeviceTypeTemperature
	default:
		return DeviceTypeUnspecified
	}
}

type Device struct {
	Name        string
	Description string
	Status      DeviceStatus
	Type        DeviceType
	User        User
	Home        Home
}
