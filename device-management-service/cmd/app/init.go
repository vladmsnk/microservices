package main

import (
	"device-management-service/cmd/docs"
	"device-management-service/internal/app/control"
	"device-management-service/internal/app/device"
	"device-management-service/internal/app/user"
	"github.com/gin-gonic/gin"
)

func setUpRouter(r *gin.Engine) {

	docs.SwaggerInfo.BasePath = "/api/v1"

	deviceHandlers := r.Group("/api/v1/device")
	device.NewDeviceHandler(deviceHandlers, nil)

	controlHandlers := r.Group("/api/v1/control")
	control.NewDeviceControlHandler(controlHandlers, nil)

	userHandlers := r.Group("/api/v1/user")
	user.NewHandler(userHandlers)
}
