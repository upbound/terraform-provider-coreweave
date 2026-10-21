// Package xpprovider exports the provider constructor for Crossplane
// providers built with Upjet, which link this provider into their own binary
// and drive it in-process instead of through the Terraform CLI.
package xpprovider

import (
	"context"

	fwprovider "github.com/hashicorp/terraform-plugin-framework/provider"

	"github.com/coreweave/terraform-provider-coreweave/internal/provider"
)

// GetProvider returns a new instance of the CoreWeave terraform-plugin-framework
// provider. The version is reported in the User-Agent of every API call.
func GetProvider(_ context.Context, version string) (fwprovider.Provider, error) {
	return provider.New(version)(), nil
}
