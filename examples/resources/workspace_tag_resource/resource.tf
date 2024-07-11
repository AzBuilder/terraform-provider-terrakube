resource "terrakube_workspace_tag" "example" {
  tag_id          = terrakube_tag.example.id
  workspace_id    = terrakube_workspace.example.id
  organization_id = terrakube_organization.example.id
}