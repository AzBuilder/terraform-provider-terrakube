data "terrakube_organization" "org" {
  name = "simple"
}

data "terrakube_vcs" "vcs" {
  name            = "sample"
  organization_id = data.terrakube_organization.org.id
}