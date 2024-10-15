package device

import (
	"context"
	"device-management-service/internal/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UseCase interface {
	RegisterDevice(ctx context.Context, device entities.Device) error
	GetDeviceInfo(ctx context.Context, name string) (entities.Device, error)
}

type Handler struct {
	uc UseCase
}

func NewDeviceHandler(r *gin.RouterGroup, uc UseCase) {
	h := Handler{uc: uc}

	r.POST("/register", h.RegisterDevice)
	r.GET("/info", h.GetDeviceInfo)
	r.GET("/list", h.ListUserDevices)
	r.DELETE("/delete", h.DeleteDevice)
}

// RegisterDevice godoc
// @Summary register device endpoint
// @Schemes
// @Description Registers a new device in the system
// @Tags example
// @Accept json
// @Produce json
// @Param requestBody body RegisterDeviceRequest true "Register device"
// @Param Authorization header string true "Bearer"
// @Success 200 {string} ok
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /device/register [post]
func (h *Handler) RegisterDevice(c *gin.Context) {
	var req RegisterDeviceRequest
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate request

	// register device

	// return response

	//todo: implement register device
}

// GetDeviceInfo godoc
// @Summary get device info endpoint
// @Schemes
// @Description get device info
// @Tags example
// @Accept json
// @Produce json
// @Param requestBody body GetDeviceInfoRequest true "Get device info"
// @Param Authorization header string true "Bearer"
// @Success 200 {string} ok
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /device/info [post]
func (h *Handler) GetDeviceInfo(c *gin.Context) {
	// Get device name from request

	// Get device info

	// return response

	//todo: implement get device info
}

// ListUserDevices godoc
// @Summary list user devices endpoint
// @Schemes
// @Description Lists all devices for the authenticated user
// @Tags example
// @Accept json
// @Produce json
// @Param requestBody body ListUserDevicesRequest true "List user devices"
// @Param Authorization header string true "Bearer"
// @Success 200 {string} ok
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /device/list [post]
func (h *Handler) ListUserDevices(c *gin.Context) {
	//todo: implement list user devices
}

// DeleteDevice godoc
// @Summary delete device endpoint
// @Schemes
// @Description delete device
// @Tags example
// @Accept json
// @Produce json
// @Param requestBody body DeleteDeviceRequest true "Delete device"
// @Param Authorization header string true "Bearer"
// @Success 200 {string} ok
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /device/delete [post]
func (h *Handler) DeleteDevice(c *gin.Context) {
	//todo: implement delete device
}
