package schema

type ReadyStatus string

const (
	ReadyStatusReady    ReadyStatus = "ready"
	ReadyStatusNotReady ReadyStatus = "not_ready"
)

type DependencyStatus string

const (
	DependencyStatusReady   DependencyStatus = "ready"
	DependencyStatusPending DependencyStatus = "pending"
	DependencyStatusFailed  DependencyStatus = "failed"
)

type DependencyName string

const (
	DependencyNameDatabase DependencyName = "database"
)

type ReadyDependency struct {
	Name   DependencyName   `json:"name"`
	Status DependencyStatus `json:"status"`
}

type ReadyResponse struct {
	Status       ReadyStatus        `json:"status"`
	Dependencies []*ReadyDependency `json:"dependencies"`
}
