package client

const (
	ListMetaType                    = "listMeta"
	ListMetaFieldContinue           = "continue"
	ListMetaFieldRemainingItemCount = "remainingItemCount"
	ListMetaFieldResourceVersion    = "resourceVersion"
	ListMetaFieldSelfLink           = "selfLink"
)

type ListMeta struct {
	Continue           string `json:"continue,omitempty" yaml:"continue,omitempty"`
	RemainingItemCount *int64 `json:"remainingItemCount,omitempty" yaml:"remainingItemCount,omitempty"`
	ResourceVersion    string `json:"resourceVersion,omitempty" yaml:"resourceVersion,omitempty"`
	SelfLink           string `json:"selfLink,omitempty" yaml:"selfLink,omitempty"`
}
