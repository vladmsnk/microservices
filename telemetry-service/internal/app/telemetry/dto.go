package telemetry

type ListTelemetryRequest struct {
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
}
