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

func (h *Handler) RegisterDevice(c *gin.Context) {
	var req RegisterDeviceRequest
	err := c.ShouldBindBodyWithJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

}

func (h *Handler) GetDeviceInfo(c *gin.Context) {

}

func (h *Handler) ListUserDevices(c *gin.Context) {

}

func (h *Handler) DeleteDevice(c *gin.Context) {

}
