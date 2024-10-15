package telemetry

import "github.com/gin-gonic/gin"

type UseCase interface {
}

type Handler struct {
	uc UseCase
}

func NewTelemetryHandler(r *gin.RouterGroup, uc UseCase) {
	h := Handler{uc: uc}

	r.GET("/last_telemetry", h.GetLastTelemetry)
	r.GET("/telemetry", h.ListTelemetry)

}

// GetLastTelemetry godoc
// @Summary Gets last telemetry for device
// @Schemes
// @Description get last telemetry for device
// @Tags example
// @Accept json
// @Produce json
// @Param device_name path string true "Device name"
// @Success 200 {string} ok
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /last_telemetry/{device_name} [get]
func (h *Handler) GetLastTelemetry(c *gin.Context) {
}

// ListTelemetry godoc
// @Summary List telemetry for device
// @Schemes
// @Description list telemetry for device
// @Tags example
// @Accept json
// @Produce json
// @Param requestBody body ListTelemetryRequest true "List telemetry request"
// @Success 200 {string} ok
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /telemetry/{device_name} [get]
func (h *Handler) ListTelemetry(c *gin.Context) {

}
