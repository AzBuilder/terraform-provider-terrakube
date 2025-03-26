resource "terrakube_ssh" "ssh" {
  name            = "github-key"
  organization_id = data.terrakube_organization.org.id
  description     = "ssh key to get modules from github"
  private_key     = file("~/.ssh/id_rsa")
  ssh_type        = "rsa"
}
