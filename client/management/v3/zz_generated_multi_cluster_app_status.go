package client

const (
	MultiClusterAppStatusType            = "multiClusterAppStatus"
	MultiClusterAppStatusFieldConditions = "conditions"
)

type MultiClusterAppStatus struct {
	Conditions []AppCondition `json:"conditions,omitempty" yaml:"conditions,omitempty"`
}
