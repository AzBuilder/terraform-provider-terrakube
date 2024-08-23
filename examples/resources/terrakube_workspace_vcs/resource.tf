resource "terrakube_workspace_cli" "sample1" {
  organization_id = data.terrakube_organization.org.id
  name            = "work-from-provider1"
  description     = "sample"
  execution_mode  = "remote"
  repository      = "https://github.com/AzBuilder/terrakube-docker-compose.git"
  branch          = "main"
  folder          = "/"
  template_id     = terrakube_organization_template.example.id
  iac_type        = "terraform"
  iac_version     = "1.5.7"
}
