resource "terrakube_workspace_access" "workspace_access" {
  name            = "my_terrakube_team"
  organization_id = "my_organization_id"
  workspace_id    = "my_workspace_id"

  manage_job       = true
  manage_state     = false
  manage_workspace = false
}