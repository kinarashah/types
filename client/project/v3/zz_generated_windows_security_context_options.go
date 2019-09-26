package client

const (
	WindowsSecurityContextOptionsType                        = "windowsSecurityContextOptions"
	WindowsSecurityContextOptionsFieldGMSACredentialSpec     = "gmsaCredentialSpec"
	WindowsSecurityContextOptionsFieldGMSACredentialSpecName = "gmsaCredentialSpecName"
)

type WindowsSecurityContextOptions struct {
	GMSACredentialSpec     string `json:"gmsaCredentialSpec,omitempty" yaml:"gmsaCredentialSpec,omitempty"`
	GMSACredentialSpecName string `json:"gmsaCredentialSpecName,omitempty" yaml:"gmsaCredentialSpecName,omitempty"`
}
