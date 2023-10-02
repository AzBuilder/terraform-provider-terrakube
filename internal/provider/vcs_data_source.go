package provider

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/google/jsonapi"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"io"
	"net/http"
	"reflect"
	"strings"
	"terraform-provider-terrakube/internal/client"
)

var (
	_ datasource.DataSource              = &VcsDataSource{}
	_ datasource.DataSourceWithConfigure = &VcsDataSource{}
)

type VcsDataSourceModel struct {
	ID             types.String `tfsdk:"id"`
	OrganizationId types.String `tfsdk:"organization_id"`
	Name           types.String `tfsdk:"name"`
	Description    types.String `tfsdk:"description"`
}

type VcsDataSource struct {
	client   *http.Client
	endpoint string
	token    string
}

func NewVcsDataSource() datasource.DataSource {
	return &VcsDataSource{}
}

func (d *VcsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, res *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	providerData, ok := req.ProviderData.(*TerrakubeConnectionData)
	if !ok {
		res.Diagnostics.AddError(
			"Unexpected Vcs Data Source Configure Type",
			fmt.Sprintf("Expected *TerrakubeConnectionData got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	customTransport := http.DefaultTransport.(*http.Transport).Clone()
	customTransport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	client := http.Client{Transport: customTransport}

	d.client = &client
	d.endpoint = providerData.Endpoint
	d.token = providerData.Token

	tflog.Info(ctx, "Creating Vcs datasource")
}

func (d *VcsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vcs"
}

func (d *VcsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Vcs Id",
			},
			"organization_id": schema.StringAttribute{
				Required:    true,
				Description: "Terrakube organization id",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Vcs Name",
			},
			"description": schema.StringAttribute{
				Computed:    true,
				Description: "Vcs description information",
			},
		},
	}
}

func (d *VcsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state VcsDataSourceModel

	req.Config.Get(ctx, &state)

	requestVcs, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/api/v1/organization/%s/vcs?filter[vcs]=name==%s", d.endpoint, state.OrganizationId.ValueString(), state.Name.ValueString()), nil)
	requestVcs.Header.Add("Authorization", fmt.Sprintf("Bearer %s", d.token))
	requestVcs.Header.Add("Content-Type", "application/vnd.api+json")
	if err != nil {
		tflog.Error(ctx, "Error creating vcs datasource request")
	}

	responseVcs, err := d.client.Do(requestVcs)
	if err != nil {
		tflog.Error(ctx, "Error executing vcs datasource request")
	}

	body, err := io.ReadAll(responseVcs.Body)
	if err != nil {
		tflog.Error(ctx, "Error reading vcs response")
	}

	var vcss []interface{}

	vcss, err = jsonapi.UnmarshalManyPayload(strings.NewReader(string(body)), reflect.TypeOf(new(client.VcsEntity)))

	for _, vcs := range vcss {
		data, _ := vcs.(*client.VcsEntity)
		state.ID = types.StringValue(data.ID)
		state.Description = types.StringValue(data.Description)
	}

	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
