resource "terrakube_team" "team" {
  name             = "TERRAKUBE_SUPER_ADMIN"
  organization_id  = data.terrakube_organization.org.id
  manage_state     = false
  manage_workspace = false
  manage_module    = false
  manage_provider  = true
  manage_vcs       = true
  manage_template  = true
}
