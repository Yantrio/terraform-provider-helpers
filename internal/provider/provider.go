// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"
	"net/http"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Ensure HelpersProvider satisfies various provider interfaces.
var _ provider.Provider = &HelpersProvider{}
var _ provider.ProviderWithFunctions = &HelpersProvider{}

// HelpersProvider defines the provider implementation.
type HelpersProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// HelpersProviderModel describes the provider data model.
type HelpersProviderModel struct {
}

func (p *HelpersProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "helpers"
	resp.Version = p.version
}

func (p *HelpersProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (p *HelpersProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data HelpersProviderModel

	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	client := http.DefaultClient
	resp.DataSourceData = client
	resp.ResourceData = client
}

func (p *HelpersProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *HelpersProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *HelpersProvider) Functions(ctx context.Context) []func() function.Function {
	return []func() function.Function{
		NewEchoFunction,
	}
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &HelpersProvider{
			version: version,
		}
	}
}
