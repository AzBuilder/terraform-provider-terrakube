resource "terrakube_workspace_cli" "sample1" {
  organization_id = data.terrakube_organization.org.id
  name            = "work-from-provider1"
  description     = "sample"
  execution_mode  = "remote"
  iac_type        = "terraform"
  iac_version     = "1.5.7"
}

resource "terrakube_workspace_cli" "sample2" {
  organization_id = data.terrakube_organization.org.id
  name            = "work-from-provider2"
  description     = "sample"
  execution_mode  = "local"
  iac_type        = "tofu"
  iac_version     = "1.7.0"
}
