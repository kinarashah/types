package client

const (
	DeploymentRollbackInputType          = "deploymentRollbackInput"
	DeploymentRollbackInputFieldRevision = "revision"
)

type DeploymentRollbackInput struct {
	Revision string `json:"revision,omitempty" yaml:"revision,omitempty"`
}
