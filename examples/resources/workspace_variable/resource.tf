resource "terrakube_workspace_variable" "sample1" {
  organization_id = data.terrakube_organization.org.id
  workspace_id    = terrakube_workspace_cli.sample1.id
  key             = "sample-env-var"
  value           = "sample-value2222"
  description     = "sample env var"
  category        = "ENV"
  sensitive       = false
  hcl             = false
}

resource "terrakube_workspace_variable" "sample2" {
  organization_id = data.terrakube_organization.org.id
  workspace_id    = terrakube_workspace_cli.sample1.id
  key             = "sample-terra-var"
  value           = "sample-TERRAFORM"
  description     = "sample env var"
  category        = "TERRAFORM"
  sensitive       = false
  hcl             = false
}