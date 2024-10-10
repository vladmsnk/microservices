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
func (h *Handler) TurnOnDevice(c *gin.Context) {

}

func (h *Handler) TurnOffDevice(c *gin.Context) {

}

func (h *Handler) GetDeviceStatus(c *gin.Context) {

}

func (h *Handler) SendCommandToDevice(c *gin.Context) {

}

func (h *Handler) GetDeviceState(c *gin.Context) {

}
