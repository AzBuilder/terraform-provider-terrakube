resource "terrakube_organization_tag" "example" {
  name            = "example"
  organization_id = terraform.workspace.id
}