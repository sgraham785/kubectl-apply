package handler

import (
	"github.com/canary-health/kubectl-apply/config"
)

// Handler struct returned by NewHandler
type Handler struct {
	// Kms *kmsClient
	// Cmd  CmdHandler
	// Kube *kubeConfigHandler
	// Dec  *decoderHandler
	Aws *awsSessHandler
}

// NewHandler is a reusable handler
func NewHandler(c *config.Config) *Handler {
	// fmt.Printf("Config in Handler = %q", c.Kubeconfig)
	return &Handler{
		// Kube: newKubeConfigHandler(c.Kubeconfig),
		// Dec:  newDecoderHandler(c.File),
		Aws: newAwsSessionHandler(),
	}
}