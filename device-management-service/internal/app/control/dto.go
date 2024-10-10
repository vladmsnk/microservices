package control

type TurnOnDeviceRequest struct {
	DeviceName string `json:"device_name"`
	HomeName   string `json:"home_name"`
	Username   string `json:"username"`
}

type TurnOnDeviceResponse struct {
}

type TurnOffDeviceRequest struct {
	DeviceName string `json:"device_name"`
	HomeName   string `json:"home_name"`
	Username   string `json:"username"`
}

type TurnOffDeviceResponse struct {
}

type GetDeviceStatusRequest struct {
	DeviceName string `json:"device_name"`
	HomeName   string `json:"home_name"`
	Username   string `json:"username"`
}

type SendCommandToDeviceRequest struct {
	DeviceName string            `json:"device_name"`
	HomeName   string            `json:"home_name"`
	Username   string            `json:"username"`
	Command    string            `json:"command"`
	Attributes map[string]string `json:"attributes"`
}

type GetDeviceStateRequest struct {
	DeviceName string `json:"device_name"`
	HomeName   string `json:"home_name"`
	Username   string `json:"username"`
}
