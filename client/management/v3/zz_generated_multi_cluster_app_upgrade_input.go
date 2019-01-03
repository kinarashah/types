package client

const (
	MultiClusterAppUpgradeInputType                     = "multiClusterAppUpgradeInput"
	MultiClusterAppUpgradeInputFieldAnswers             = "answers"
	MultiClusterAppUpgradeInputFieldBatchSize           = "batchSize"
	MultiClusterAppUpgradeInputFieldTemplateVersionName = "templateVersionName"
)

type MultiClusterAppUpgradeInput struct {
	Answers             map[string]string `json:"answers,omitempty" yaml:"answers,omitempty"`
	BatchSize           int64             `json:"batchSize,omitempty" yaml:"batchSize,omitempty"`
	TemplateVersionName string            `json:"templateVersionName,omitempty" yaml:"templateVersionName,omitempty"`
}
