package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type logProvider struct {
	version string
}

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &logProvider{
			version: version,
		}
	}
}

func (l logProvider) Metadata(_ context.Context, _ provider.MetadataRequest, response *provider.MetadataResponse) {
	response.TypeName = "log-provider"
	response.Version = l.version
}

func (l logProvider) Schema(_ context.Context, _ provider.SchemaRequest, response *provider.SchemaResponse) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"platform": schema.StringAttribute{
				Description: "The platform where logs will be sent. Supported platforms are 'aws', 'private_cloud', and 'azure'.",
				Required:    true,
			},
			"region": schema.StringAttribute{
				Description: "The region where logs will be sent. e.g. 'stl'",
				Required:    true,
			},
			"environment": schema.StringAttribute{
				Description: "The environment where logs will be sent. Supported environments are 'non-prod' and 'prod'",
				Required:    true,
			},
		},
	}
}

func (l logProvider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {

}

func (l logProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

func (l logProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewLogResource,
	}
}
