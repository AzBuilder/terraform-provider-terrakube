resource "terrakube_collection_item" "sample1" {
  organization_id = data.terrakube_organization.org.id
  collection_id   = terrakube_collection.sample1.id
  key             = "sample-env-var"
  value           = "sample-value2222"
  description     = "sample env var"
  category        = "ENV"
  sensitive       = false
  hcl             = false
}

resource "terrakube_collection_item" "sample2" {
  organization_id = data.terrakube_organization.org.id
  collection_id   = terrakube_collection.sample1.id
  key             = "sample-terra-var"
  value           = "sample-TERRAFORM"
  description     = "sample env var"
  category        = "TERRAFORM"
  sensitive       = false
  hcl             = false
}