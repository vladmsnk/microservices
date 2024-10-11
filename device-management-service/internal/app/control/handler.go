package control

import "github.com/gin-gonic/gin"

type UseCase interface {
}

type Handler struct {
	uc UseCase
}

func NewDeviceControlHandler(r *gin.RouterGroup, uc UseCase) {
	h := Handler{uc: uc}

	r.POST("/turn_on", h.TurnOnDevice)
	r.POST("/turn_off", h.TurnOffDevice)
	r.GET("/status", h.GetDeviceStatus)
	r.POST("/command", h.SendCommandToDevice)
	r.GET("/state", h.GetDeviceState)
}

// TurnOnDevice godoc
// @Summary turn on device endpoint
// @Schemes
// @Description Turns on a device
// @Tags example
// @Accept json
// @Produce json
// @Param requestBody body TurnOnDeviceRequest true "Turn on device"
// @Param Authorization header string true "Bearer"
// @Success 200 {string} ok
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /control/turn_on [post]
func (h *Handler) TurnOnDevice(c *gin.Context) {

}

// TurnOffDevice godoc
// @Summary turn off device endpoint
// @Schemes
// @Description Turns off a device
// @Tags example
// @Accept json
// @Produce json
// @Param requestBody body TurnOffDeviceRequest true "Turn off device"
// @Param Authorization header string true "Bearer"
// @Success 200 {string} ok
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /control/turn_off [post]
func (h *Handler) TurnOffDevice(c *gin.Context) {

}

// GetDeviceStatus godoc
// @Summary gets device status
// @Schemes
// @Description Gets device status
// @Tags example
// @Accept json
// @Produce json
// @Param requestBody body GetDeviceStatusRequest true "Gets device status"
// @Param Authorization header string true "Bearer"
// @Success 200 {string} ok
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /control/status [get]
func (h *Handler) GetDeviceStatus(c *gin.Context) {

}

// SendCommandToDevice godoc
// @Summary sends command to device
// @Schemes
// @Description sends command to device
// @Tags example
// @Accept json
// @Produce json
// @Param requestBody body SendCommandToDeviceRequest true "Sends command to device"
// @Param Authorization header string true "Bearer"
// @Success 200 {string} ok
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /control/command [post]
func (h *Handler) SendCommandToDevice(c *gin.Context) {

}

// GetDeviceState godoc
// @Summary gets device state
// @Schemes
// @Description gets device state
// @Tags example
// @Accept json
// @Produce json
// @Param requestBody body GetDeviceStateRequest true "Gets device state"
// @Param Authorization header string true "Bearer"
// @Success 200 {string} ok
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /control/state [post]
func (h *Handler) GetDeviceState(c *gin.Context) {

}
