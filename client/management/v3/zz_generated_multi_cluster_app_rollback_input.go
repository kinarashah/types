package client

const (
	MultiClusterAppRollbackInputType            = "multiClusterAppRollbackInput"
	MultiClusterAppRollbackInputFieldBatchSize  = "batchSize"
	MultiClusterAppRollbackInputFieldRevisionID = "revisionId"
)

type MultiClusterAppRollbackInput struct {
	BatchSize  int64  `json:"batchSize,omitempty" yaml:"batchSize,omitempty"`
	RevisionID string `json:"revisionId,omitempty" yaml:"revisionId,omitempty"`
}
