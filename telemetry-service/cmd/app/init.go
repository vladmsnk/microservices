package main

import (
	"github.com/gin-gonic/gin"
	"telemetry-service/cmd/docs"
	"telemetry-service/internal/app/telemetry"
)

func setUpRouter(r *gin.Engine) {
	docs.SwaggerInfo.BasePath = "/api/v1"

	telemetryHandlers := r.Group("/api/v1/telemetry")
	telemetry.NewTelemetryHandler(telemetryHandlers, nil)
}
