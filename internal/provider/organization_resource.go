package provider

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/google/jsonapi"
	"io"
	"net/http"
	"strings"
	"terraform-provider-terrakube/internal/client"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ resource.Resource = &OrganizationResource{}
var _ resource.ResourceWithImportState = &OrganizationResource{}

type OrganizationResource struct {
	client   *http.Client
	endpoint string
	token    string
}

type OrganizationResourceModel struct {
	ID            types.String `tfsdk:"id"`
	Name          types.String `tfsdk:"name"`
	Description   types.String `tfsdk:"description"`
	ExecutionMode types.String `tfsdk:"execution_mode"`
}

func NewOrganizationResource() resource.Resource {
	return &OrganizationResource{}
}

func (r *OrganizationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_organization"
}

func (r *OrganizationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Organization Id",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Organization name",
			},
			"description": schema.StringAttribute{
				Required:    true,
				Description: "Organization description",
			},
			"execution_mode": schema.StringAttribute{
				Required:    true,
				Description: "Select default execution mode for the organization (remote or local)",
			},
		},
	}
}

func (r *OrganizationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(*TerrakubeConnectionData)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Organization Resource Configure Type",
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

	tflog.Debug(ctx, "Configuring Organization resource", map[string]any{"success": true})
}

func (r *OrganizationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan OrganizationResourceModel

	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	bodyRequest := &client.OrganizationEntity{
		Name:          plan.Name.ValueString(),
		Description:   plan.Description.ValueString(),
		ExecutionMode: plan.ExecutionMode.ValueString(),
	}

	var out = new(bytes.Buffer)
	err := jsonapi.MarshalPayload(out, bodyRequest)

	if err != nil {
		resp.Diagnostics.AddError("Unable to marshal payload", fmt.Sprintf("Unable to marshal payload: %s", err))
		return
	}

	organizationRequest, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/api/v1/organization", r.endpoint), strings.NewReader(out.String()))
	organizationRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", r.token))
	organizationRequest.Header.Add("Content-Type", "application/vnd.api+json")
	if err != nil {
		resp.Diagnostics.AddError("Error creating organization resource request", fmt.Sprintf("Error creating organization resource request: %s", err))
		return
	}

	organizationResponse, err := r.client.Do(organizationRequest)
	if err != nil {
		resp.Diagnostics.AddError("Error executing organization resource request", fmt.Sprintf("Error executing organization resource request: %s", err))
		return
	}

	bodyResponse, err := io.ReadAll(organizationResponse.Body)
	if err != nil {
		tflog.Error(ctx, "Error reading organization resource response")
	}
	newOrganization := &client.OrganizationEntity{}

	err = jsonapi.UnmarshalPayload(strings.NewReader(string(bodyResponse)), newOrganization)

	if err != nil {
		resp.Diagnostics.AddError("Error unmarshal payload response", fmt.Sprintf("Error unmarshal payload response: %s", err))
		return
	}

	tflog.Info(ctx, "Body Response", map[string]any{"bodyResponse": string(bodyResponse)})

	plan.ID = types.StringValue(newOrganization.ID)
	plan.Name = types.StringValue(newOrganization.Name)
	plan.Description = types.StringValue(newOrganization.Description)
	plan.ExecutionMode = types.StringValue(newOrganization.ExecutionMode)

	tflog.Info(ctx, "Organization Resource Created", map[string]any{"success": true})

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *OrganizationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state OrganizationResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	organizationRequest, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/organization/%s", r.endpoint, state.ID.ValueString()), nil)
	organizationRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", r.token))
	organizationRequest.Header.Add("Content-Type", "application/vnd.api+json")
	if err != nil {
		resp.Diagnostics.AddError("Error creating organization resource request", fmt.Sprintf("Error creating organization resource request: %s", err))
		return
	}

	organizationResponse, err := r.client.Do(organizationRequest)
	if err != nil {
		resp.Diagnostics.AddError("Error executing organization resource request", fmt.Sprintf("Error executing organization resource request: %s", err))
		return
	}

	bodyResponse, err := io.ReadAll(organizationResponse.Body)
	if err != nil {
		tflog.Error(ctx, "Error reading organization resource response")
	}
	organization := &client.OrganizationEntity{}

	tflog.Info(ctx, "Body Response", map[string]any{"bodyResponse": string(bodyResponse)})
	err = jsonapi.UnmarshalPayload(strings.NewReader(string(bodyResponse)), organization)

	if err != nil {
		resp.Diagnostics.AddError("Error unmarshal payload response", fmt.Sprintf("Error unmarshal payload response: %s", err))
		return
	}

	tflog.Info(ctx, "Body Response", map[string]any{"bodyResponse": string(bodyResponse)})

	state.Description = types.StringValue(organization.Description)
	state.ExecutionMode = types.StringValue(organization.ExecutionMode)

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Organization Resource reading", map[string]any{"success": true})
}

func (r *OrganizationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan OrganizationResourceModel
	var state OrganizationResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	bodyRequest := &client.OrganizationEntity{
		Description:   plan.Description.ValueString(),
		ExecutionMode: plan.ExecutionMode.ValueString(),
		Name:          plan.Name.ValueString(),
		ID:            state.ID.ValueString(),
	}

	var out = new(bytes.Buffer)
	err := jsonapi.MarshalPayload(out, bodyRequest)

	if err != nil {
		resp.Diagnostics.AddError("Unable to marshal payload", fmt.Sprintf("Unable to marshal payload: %s", err))
		return
	}

	organizationRequest, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/api/v1/organization/%s", r.endpoint, state.ID.ValueString()), strings.NewReader(out.String()))
	organizationRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", r.token))
	organizationRequest.Header.Add("Content-Type", "application/vnd.api+json")
	if err != nil {
		resp.Diagnostics.AddError("Error creating organization resource request", fmt.Sprintf("Error creating organization resource request: %s", err))
		return
	}

	organizationResponse, err := r.client.Do(organizationRequest)
	if err != nil {
		resp.Diagnostics.AddError("Error organization resource request", fmt.Sprintf("Error executing organization resource request: %s", err))
		return
	}

	bodyResponse, err := io.ReadAll(organizationResponse.Body)
	if err != nil {
		tflog.Error(ctx, "Error reading organization resource response")
	}

	tflog.Info(ctx, "Body Response", map[string]any{"success": string(bodyResponse)})

	organizationRequest, err = http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/organization/%s", r.endpoint, state.ID.ValueString()), nil)
	organizationRequest.Header.Add("Authorization", fmt.Sprintf("Bearer %s", r.token))
	organizationRequest.Header.Add("Content-Type", "application/vnd.api+json")
	if err != nil {
		resp.Diagnostics.AddError("Error creating organization resource request", fmt.Sprintf("Error creating organization resource request: %s", err))
		return
	}

	organizationResponse, err = r.client.Do(organizationRequest)
	if err != nil {
		resp.Diagnostics.AddError("Error executing organization resource request", fmt.Sprintf("Error executing organizationñ resource request: %s", err))
		return
	}

	bodyResponse, err = io.ReadAll(organizationResponse.Body)
	if err != nil {
		resp.Diagnostics.AddError("Error reading organization resource response body", fmt.Sprintf("Error reading organization resource response body: %s", err))
	}

	tflog.Info(ctx, "Body Response", map[string]any{"bodyResponse": string(bodyResponse)})

	organization := &client.OrganizationEntity{}
	err = jsonapi.UnmarshalPayload(strings.NewReader(string(bodyResponse)), organization)

	if err != nil {
		resp.Diagnostics.AddError("Error unmarshal payload response", fmt.Sprintf("Error unmarshal payload response: %s", err))
		return
	}

	plan.ID = types.StringValue(state.ID.ValueString())
	plan.Name = types.StringValue(organization.Name)
	plan.Description = types.StringValue(organization.Description)
	plan.ExecutionMode = types.StringValue(organization.ExecutionMode)

	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *OrganizationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data OrganizationResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	bodyRequest := &client.OrganizationEntity{
		Disabled: true,
		ID:       data.ID.ValueString(),
	}

	var out = new(bytes.Buffer)
	err := jsonapi.MarshalPayload(out, bodyRequest)

	if err != nil {
		resp.Diagnostics.AddError("Unable to marshal payload", fmt.Sprintf("Unable to marshal payload: %s", err))
		return
	}

	reqOrg, err := http.NewRequest(http.MethodPatch, fmt.Sprintf("%s/api/v1/organization/%s", r.endpoint, data.ID.ValueString()), strings.NewReader(out.String()))
	reqOrg.Header.Add("Authorization", fmt.Sprintf("Bearer %s", r.token))
	if err != nil {
		resp.Diagnostics.AddError("Error creating organization resource request", fmt.Sprintf("Error creating organization resource request: %s", err))
		return
	}

	_, err = r.client.Do(reqOrg)
	if err != nil {
		resp.Diagnostics.AddError("Error executing organization resource request", fmt.Sprintf("Error executing organization resource request: %s", err))
		return
	}
}

func (r *OrganizationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("id"), req, resp)
}
