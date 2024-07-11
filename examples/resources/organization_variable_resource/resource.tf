resource "terrakube_organization_variable" "sample1" {
  organization_id = data.terrakube_organization.org.id
  key             = "sample-env-var"
  value           = "sample-var2"
  description     = "sample env var2213241243"
  category        = "ENV"
  sensitive       = false
  hcl             = false
}

resource "terrakube_organization_variable" "sample2" {
  organization_id = data.terrakube_organization.org.id
  key             = "sample-terra-var"
  value           = "sample-terraform234"
  description     = "sample env var222212341234"
  category        = "TERRAFORM"
  sensitive       = false
  hcl             = false
}