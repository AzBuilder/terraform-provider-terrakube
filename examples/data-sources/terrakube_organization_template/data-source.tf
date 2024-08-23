data "terrakube_organization" "org" {
  name = "simple"
}

data "terrakube_organization_template" "template" {
  name            = "sample"
  organization_id = data.terrakube_organization.org.id
}