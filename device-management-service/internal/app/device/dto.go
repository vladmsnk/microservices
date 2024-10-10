package device

type RegisterDeviceRequest struct {
	DeviceName        string `json:"device_name"`
	HomeName          string `json:"home_name"`
	Username          string `json:"username"`
	DeviceDescription string `json:"device_description"`
	DeviceType        string `json:"device_type"`
}

type RegisterDeviceResponse struct {
	DeviceName string `json:"device_name"`
}

type GetDeviceInfoRequest struct {
	DeviceName string `json:"device_name"`
	HomeName   string `json:"home_name"`
	Username   string `json:"username"`
}

type GetDeviceInfoResponse struct {
	DeviceName        string `json:"device_name"`
	HomeName          string `json:"home_name"`
	Username          string `json:"username"`
	DeviceDescription string `json:"device_description"`
	DeviceType        string `json:"device_type"`
}
