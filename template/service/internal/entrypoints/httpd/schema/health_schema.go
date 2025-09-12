package schema

type HealthStatus string

const (
	HealthStatusOK HealthStatus = "ok"
)

type HealthResponse struct {
	Status    HealthStatus `json:"status"`
	Timestamp int64        `json:"timestamp"`
	IpAddress string       `json:"ip_address"`
}
