package telemetry

import "github.com/gin-gonic/gin"

type UseCase interface {
}

type Handler struct {
	uc UseCase
}

func NewDeviceControlHandler(r *gin.RouterGroup, uc UseCase) {
	h := Handler{uc: uc}

	r.POST("/last_telemetry", h.GetLastTelemetry)

}

// GetLastTelemetry godoc
// @Summary Gets last telemetry for device
// @Schemes
// @Description get last telemetry for device
// @Tags example
// @Accept json
// @Produce json
// @Param requestBody body GetLastTelemetryRequest true "Gets last telemetry request"
// @Param Authorization header string true "Bearer"
// @Success 200 {string} ok
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /telemetry [post]
func (h *Handler) GetLastTelemetry(c *gin.Context) {
}
