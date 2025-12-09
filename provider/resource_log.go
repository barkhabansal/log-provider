package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"strings"
)

type logResource struct {
}

func NewLogResource() resource.Resource {
	return &logResource{}
}

type logModel struct {
	ID      types.String `tfsdk:"id"`
	Message types.String `tfsdk:"message"`
	Level   types.String `tfsdk:"level"`
}

func (l logResource) Metadata(_ context.Context, _ resource.MetadataRequest, response *resource.MetadataResponse) {
	response.TypeName = "log-provider_log"
}

func (l logResource) Schema(_ context.Context, _ resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Description: "The ID of the log entry.",
				Computed:    true,
			},
			"message": schema.StringAttribute{
				Description: "The log message.",
				Required:    true,
			},
			"level": schema.StringAttribute{
				Description: "The log level (e.g., INFO (default), WARN, ERROR).",
				Optional:    true,
			},
		},
	}
}

func (l logResource) Create(ctx context.Context, request resource.CreateRequest, response *resource.CreateResponse) {
	var plan logModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)

	if response.Diagnostics.HasError() {
		return
	}

	level := strings.ToUpper(plan.Level.ValueString())
	msg := plan.Message.ValueString()

	fields := map[string]any{"level": level, "message": msg, "operation": "create"}

	switch level {
	case "ERROR":
		tflog.Error(ctx, msg, fields)
	case "WARN":
		tflog.Warn(ctx, msg, fields)
	default:
		tflog.Info(ctx, msg, fields)
	}

	plan.ID = types.StringValue(fmt.Sprintf("log-%s", msg))
	response.State.Set(ctx, &plan)

}

func (l logResource) Read(ctx context.Context, request resource.ReadRequest, response *resource.ReadResponse) {
}

func (l logResource) Update(ctx context.Context, request resource.UpdateRequest, response *resource.UpdateResponse) {
	var plan logModel
	response.Diagnostics.Append(request.Plan.Get(ctx, &plan)...)

	if response.Diagnostics.HasError() {
		return
	}

	level := strings.ToUpper(plan.Level.ValueString())
	msg := plan.Message.ValueString()

	fields := map[string]any{"level": level, "message": msg, "operation": "update"}

	switch level {
	case "ERROR":
		tflog.Error(ctx, msg, fields)
	case "WARN":
		tflog.Warn(ctx, msg, fields)
	default:
		tflog.Info(ctx, msg, fields)
	}

	response.State.Set(ctx, &plan)
}

func (l logResource) Delete(ctx context.Context, request resource.DeleteRequest, response *resource.DeleteResponse) {
	var state logModel
	response.Diagnostics.Append(request.State.Get(ctx, &state)...)

	if response.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, fmt.Sprintf("Deleting %s", state.ID))
}
