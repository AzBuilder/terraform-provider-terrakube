package provider

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"github.com/google/jsonapi"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"io"
	"net/http"
	"strconv"
	"strings"
	"terraform-provider-terrakube/internal/client"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &WorkspaceVcsResource{}
var _ resource.ResourceWithImportState = &WorkspaceVcsResource{}

type WorkspaceVcsResource struct {
	client   *http.Client
	endpoint string
	token    string
}

type WorkspaceVcsResourceModel struct {
	ID             types.String `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	OrganizationId types.String `tfsdk:"organization_id"`
	Description    types.String `tfsdk:"description"`
	IaCType        types.String `tfsdk:"iac_type"`
	IaCVersion     types.String `tfsdk:"iac_version"`
	Repository     types.String `tfsdk:"repository"`
	Branch         types.String `tfsdk:"branch"`
	ExecutionMode  types.String `tfsdk:"execution_mode"`
}

func NewWorkspaceVcsResource() resource.Resource {
	return &WorkspaceVcsResource{}
}

func (r *WorkspaceVcsResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_workspace_vcs"
}

func (r *WorkspaceVcsResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Workspace CLI Id",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"organization_id": schema.StringAttribute{
				Required:    true,
				Description: "Terrakube organization id",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Workspace VCS name",
			},
			"description": schema.StringAttribute{
				Required:    true,
				Description: "Workspace VCS description",
			},
			"execution_mode": schema.StringAttribute{
				Required:    true,
				Description: "Workspace VCS execution mode (remote or local)",
			},
			"iac_type": schema.StringAttribute{
				Required:    true,
				Description: "Workspace VCS IaC type (Supported values terraform or tofu)",
			},
			"iac_version": schema.StringAttribute{
				Required:    true,
				Description: "Workspace VCS VCS type",
			},
			"repository": schema.StringAttribute{
				Required:    true,
				Description: "Workspace VCS repository",
			},
			"branch": schema.StringAttribute{
				Required:    true,
				Description: "Workspace VCS branch",
			},
		},
	}
}

func (r *WorkspaceVcsResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(*TerrakubeConnectionData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Team Resource Configure Type",
			fmt.Sprintf("Expected *TerrakubeConnectionData, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	if providerData.InsecureHttpClient {
		if custom, ok := http.DefaultTransport.(*http.Transport); ok {
			customTransport := custom.Clone()
			customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
			r.client = &http.Client{Transport: customTransport}
		} else {
			r.client = &http.Client{}
		}
	} else {
		r.client = &http.Client{}
	}

	r.endpoint = providerData.Endpoint
	r.token = providerData.Token

	tflog.Debug(ctx, "Configuring Workspace VCS resource", map[string]any{"success": true})
}

func (r *WorkspaceVcsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan WorkspaceVcsResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	bodyRequest := &client.WorkspaceEntity{
		Name:          plan.Name.ValueString(),
		Description:   plan.Description.ValueString(),
		Source:        plan.Repository.ValueString(),
		Branch:        plan.Branch.ValueString(),
		IaCType:       plan.IaCType.ValueString(),
		IaCVersion:    plan.IaCVersion.ValueString(),
		ExecutionMode: plan.ExecutionMode.ValueString(),
	}

	var out = new(bytes.Buffer)
	err := jsonapi.MarshalPayload(out, bodyRequest)

	if err != nil {
		resp.Diagnostics.AddError("Unable to marshal payload", fmt.Sprintf("Unable to marshal payload: %s", err))
		return
	}

	workspaceVcsRequest, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/organization/%s/workspace", r.endpoint, plan.OrganizationId.ValueString()), strings.NewReader(out.String()))
	workspaceVcsRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", r.token))
	workspaceVcsRequest.Header.Add("Content-Type", "application/vnd.api+json")
	if err != nil {
		resp.Diagnostics.AddError("Error creating workspace vcs resource request", fmt.Sprintf("Error creating workspace vcs resource request: %s", err))
		return
	}

	workspaceVcsResponse, err := r.client.Do(workspaceVcsRequest)
	if err != nil {
		resp.Diagnostics.AddError("Error executing workspace vcs resource request", fmt.Sprintf("Error executing workspace vcs resource request: %s", err))
		return
	}

	bodyResponse, err := io.ReadAll(workspaceVcsResponse.Body)
	if err != nil {
		tflog.Error(ctx, "Error reading workspace vcs resource response")
	}
	newWorkspaceVcs := &client.WorkspaceEntity{}

	fmt.Println(string(bodyResponse))
	err = jsonapi.UnmarshalPayload(strings.NewReader(string(bodyResponse)), newWorkspaceVcs)

	if err != nil {
		resp.Diagnostics.AddError("Error unmarshal payload response", fmt.Sprintf("Error unmarshal payload response: %s", err))
		return
	}

	tflog.Info(ctx, "Body Response", map[string]any{"bodyResponse": string(bodyResponse)})

	plan.ID = types.StringValue(newWorkspaceVcs.ID)
	plan.Name = types.StringValue(newWorkspaceVcs.Name)
	plan.Description = types.StringValue(newWorkspaceVcs.Description)
	plan.Repository = types.StringValue(newWorkspaceVcs.Source)
	plan.Branch = types.StringValue(newWorkspaceVcs.Branch)
	plan.IaCType = types.StringValue(newWorkspaceVcs.IaCType)
	plan.IaCVersion = types.StringValue(newWorkspaceVcs.IaCVersion)
	plan.ExecutionMode = types.StringValue(newWorkspaceVcs.ExecutionMode)

	tflog.Info(ctx, "Workspace VCS Resource Created", map[string]any{"success": true})

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *WorkspaceVcsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state WorkspaceVcsResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	workspaceRequest, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/organization/%s/workspace/%s", r.endpoint, state.OrganizationId.ValueString(), state.ID.ValueString()), nil)
	workspaceRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", r.token))
	workspaceRequest.Header.Add("Content-Type", "application/vnd.api+json")
	if err != nil {
		resp.Diagnostics.AddError("Error creating workspace vcs resource request", fmt.Sprintf("Error creating workspace cli resource request: %s", err))
		return
	}

	workspaceResponse, err := r.client.Do(workspaceRequest)
	if err != nil {
		resp.Diagnostics.AddError("Error executing workspace vcs resource request", fmt.Sprintf("Error executing workspace cli resource request: %s", err))
		return
	}

	bodyResponse, err := io.ReadAll(workspaceResponse.Body)
	if err != nil {
		tflog.Error(ctx, "Error reading workspace vcs resource response")
	}
	workspace := &client.WorkspaceEntity{}

	tflog.Info(ctx, "Body Response", map[string]any{"bodyResponse": string(bodyResponse)})
	err = jsonapi.UnmarshalPayload(strings.NewReader(string(bodyResponse)), workspace)

	if err != nil {
		resp.Diagnostics.AddError("Error unmarshal payload response", fmt.Sprintf("Error unmarshal payload response: %s", err))
		return
	}

	tflog.Info(ctx, "Body Response", map[string]any{"bodyResponse": string(bodyResponse)})

	state.Name = types.StringValue(workspace.Name)
	state.Description = types.StringValue(workspace.Description)
	state.ExecutionMode = types.StringValue(workspace.ExecutionMode)
	state.Repository = types.StringValue(workspace.Source)
	state.Branch = types.StringValue(workspace.Branch)
	state.IaCType = types.StringValue(workspace.IaCType)
	state.IaCVersion = types.StringValue(workspace.IaCVersion)
	state.ID = types.StringValue(workspace.ID)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Workspace vcs Resource reading", map[string]any{"success": true})
}

func (r *WorkspaceVcsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan WorkspaceVcsResourceModel
	var state WorkspaceVcsResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	bodyRequest := &client.WorkspaceEntity{
		IaCVersion:    plan.IaCVersion.ValueString(),
		IaCType:       plan.IaCType.ValueString(),
		ExecutionMode: plan.ExecutionMode.ValueString(),
		Description:   plan.Description.ValueString(),
		Source:        plan.Repository.ValueString(),
		Branch:        plan.Repository.ValueString(),
		Name:          plan.Name.ValueString(),
		ID:            state.ID.ValueString(),
	}

	var out = new(bytes.Buffer)
	err := jsonapi.MarshalPayload(out, bodyRequest)

	if err != nil {
		resp.Diagnostics.AddError("Unable to marshal payload", fmt.Sprintf("Unable to marshal payload: %s", err))
		return
	}

	organizationRequest, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/api/v1/organization/%s/workspace/%s", r.endpoint, state.OrganizationId.ValueString(), state.ID.ValueString()), strings.NewReader(out.String()))
	organizationRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", r.token))
	organizationRequest.Header.Add("Content-Type", "application/vnd.api+json")
	if err != nil {
		resp.Diagnostics.AddError("Error creating workspace vcs resource request", fmt.Sprintf("Error creating workspace vcs resource request: %s", err))
		return
	}

	organizationResponse, err := r.client.Do(organizationRequest)
	if err != nil {
		resp.Diagnostics.AddError("Error executing workspace vcs resource request", fmt.Sprintf("Error executing workspace vcs resource request: %s", err))
		return
	}

	bodyResponse, err := io.ReadAll(organizationResponse.Body)
	if err != nil {
		tflog.Error(ctx, "Error reading workspace vcs resource response")
	}

	tflog.Info(ctx, "Body Response", map[string]any{"success": string(bodyResponse)})

	organizationRequest, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/organization/%s/workspace/%s", r.endpoint, state.OrganizationId.ValueString(), state.ID.ValueString()), nil)
	organizationRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", r.token))
	organizationRequest.Header.Add("Content-Type", "application/vnd.api+json")
	if err != nil {
		resp.Diagnostics.AddError("Error creating workspace vcs resource request", fmt.Sprintf("Error creating workspace vcs resource request: %s", err))
		return
	}

	organizationResponse, err = r.client.Do(organizationRequest)
	if err != nil {
		resp.Diagnostics.AddError("Error executing workspace vcs resource request", fmt.Sprintf("Error executing workspace vcs resource request: %s", err))
		return
	}

	bodyResponse, err = io.ReadAll(organizationResponse.Body)
	if err != nil {
		resp.Diagnostics.AddError("Error reading workspace vcs resource response body", fmt.Sprintf("Error reading workspace vcs resource response body: %s", err))
	}

	tflog.Info(ctx, "Body Response", map[string]any{"bodyResponse": string(bodyResponse)})

	workspace := &client.WorkspaceEntity{}
	err = jsonapi.UnmarshalPayload(strings.NewReader(string(bodyResponse)), workspace)

	if err != nil {
		resp.Diagnostics.AddError("Error unmarshal payload response", fmt.Sprintf("Error unmarshal payload response: %s", err))
		return
	}

	plan.ID = types.StringValue(state.ID.ValueString())
	plan.Name = types.StringValue(workspace.Name)
	plan.Description = types.StringValue(workspace.Description)
	plan.Repository = types.StringValue(workspace.Source)
	plan.Branch = types.StringValue(workspace.Branch)
	plan.IaCType = types.StringValue(workspace.IaCType)
	plan.IaCVersion = types.StringValue(workspace.IaCVersion)
	plan.ExecutionMode = types.StringValue(workspace.ExecutionMode)

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *WorkspaceVcsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data WorkspaceVcsResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	ll := len(chars)
	b := make([]byte, 4)

	if _, err := rand.Read(b); err != nil {
		resp.Diagnostics.AddError("Error generating random string to delete workspace", fmt.Sprintf("Error generating random string to delete workspace: %s", err))
		return
	}

	for i := 0; i < 4; i++ {
		b[i] = chars[int(b[i])%ll]
	}

	tflog.Info(ctx, "Send patch request to mark as deleted...")
	tflog.Info(ctx, fmt.Sprintf("%s_DEL_%s", data.Name.ValueString(), string(b)))

	bodyRequest := &client.WorkspaceEntity{
		ID:            data.ID.ValueString(),
		Name:          fmt.Sprintf("%s_DEL_%s", data.Name.ValueString(), string(b)), // FORCE A NAME CHANGE WITH THE SAME LOGIC THAT IN THE UI
		Description:   data.Description.ValueString(),
		Source:        data.Repository.ValueString(),
		Branch:        data.Branch.ValueString(),
		IaCType:       data.IaCType.ValueString(),
		IaCVersion:    data.IaCVersion.ValueString(),
		ExecutionMode: data.ExecutionMode.ValueString(),
		Deleted:       true,
	}

	var out = new(bytes.Buffer)
	err := jsonapi.MarshalPayload(out, bodyRequest)

	tflog.Info(ctx, "Request Body...")
	tflog.Info(ctx, out.String())

	if err != nil {
		resp.Diagnostics.AddError("Unable to marshal payload", fmt.Sprintf("Unable to marshal payload: %s", err))
		return
	}

	workspaceVcsRequest, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/api/v1/organization/%s/workspace/%s", r.endpoint, data.OrganizationId.ValueString(), data.ID.ValueString()), strings.NewReader(out.String()))
	workspaceVcsRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", r.token))
	workspaceVcsRequest.Header.Add("Content-Type", "application/vnd.api+json")
	if err != nil {
		resp.Diagnostics.AddError("Error creating vcs resource request", fmt.Sprintf("Error creating vcs resource request: %s", err))
		return
	}

	workspaceVcsResponse, err := r.client.Do(workspaceVcsRequest)
	if err != nil {
		resp.Diagnostics.AddError("Error executing vcs resource request", fmt.Sprintf("Error executing vcs resource request: %s", err))
		return
	}

	tflog.Info(ctx, "Delete response code: "+strconv.Itoa(workspaceVcsResponse.StatusCode))

}

func (r *WorkspaceVcsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
