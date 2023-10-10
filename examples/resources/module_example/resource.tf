resource "terrakube_module" "module" {
  name            = "vpc"
  organization_id = data.terrakube_organization.org.id
  description     = "cloudposse module"
  provider_name   = "aws"
  source          = "https://github.com/terraform-aws-modules/terraform-aws-vpc.git"
}

resource "terrakube_module" "module" {
  name            = "vpc_private"
  organization_id = data.terrakube_ssh.ssh.organization_id
  description     = "cloudposse module"
  provider_name   = "aws"
  source          = "https://github.com/terraform-aws-modules/terraform-aws-vpc.git"
  vcs_id          = data.terrakube_vcs.vcs.id
}

resource "terrakube_module" "module" {
  name            = "vpc_private_ssh"
  organization_id = data.terrakube_ssh.ssh.organization_id
  description     = "cloudposse module"
  provider_name   = "aws"
  source          = "https://github.com/terraform-aws-modules/terraform-aws-vpc.git"
  ssh_id          = data.terrakube_ssh.ssh.id
}