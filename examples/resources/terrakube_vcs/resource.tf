resource "terrakube_vcs" "vcs" {
  organization_id = data.terrakube_organization.org.id
  name            = "Github"
  description     = "test github connection"
  vcs_type        = "GITHUB"
  client_id       = "Iv1.1b9b7b1b1b1b1b1b"
  client_secret   = "hellotest"
  endpoint        = "https://github.com"
  api_url         = "https://api.github.com"
}
