---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "terrakube_collection_item Resource - terrakube"
subcategory: ""
description: |-
  Create collection item that will be used by this workspace only.
---

# terrakube_collection_item (Resource)

Create collection item that will be used by this workspace only.

## Example Usage

```terraform
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
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `category` (String) Variable category (ENV or TERRAFORM). ENV variables are injected in workspace environment at runtime.
- `collection_id` (String) Terrakube collection id
- `description` (String) Variable description
- `hcl` (Boolean) Parse this field as HashiCorp Configuration Language (HCL). This allows you to interpolate values at runtime.
- `key` (String) Variable key
- `organization_id` (String) Terrakube organization id
- `sensitive` (Boolean) Sensitive variables are never shown in the UI or API. They may appear in Terraform logs if your configuration is designed to output them.
- `value` (String) Variable value

### Read-Only

- `id` (String) Collection Id

## Import

Import is supported using the following syntax:

```shell
# Organization Workspace Variable can be import with organization_id,collection_id,id
terraform import terrakube_workspace_variable.example 00000000-0000-0000-0000-000000000000,00000000-0000-0000-0000-000000000000,00000000-0000-0000-0000-000000000000
```