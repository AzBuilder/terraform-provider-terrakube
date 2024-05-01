resource "terrakube_workspace_variable" "sample1" {
  organization_id = data.terrakube_organization.org.id
  workspace_id    = data.terrakube_workspace.workspace.id
  key             = "sample-env-var"
  value           = "sample-value"
  description     = "sample env var"
  category        = "ENV"
  sensitive       = false
  hcl             = false
}