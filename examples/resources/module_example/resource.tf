resource "terrakube_module" "module" {
  name            = "vpc"
  organization_id = data.terrakube_organization.org.id
  description     = "cloudposse module"
  provider_name   = "aws"
  source          = "https://github.com/terraform-aws-modules/terraform-aws-vpc.git"
}