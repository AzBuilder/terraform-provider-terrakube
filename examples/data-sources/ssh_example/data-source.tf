data "terrakube_organization" "org" {
  name = "simple"
}

data "terrakube_ssh" "ssh" {
  name            = "sample"
  organization_id = data.terrakube_organization.org.id
}