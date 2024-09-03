resource "terrakube_webhook" "webhook" {
  organization_id = data.terrakube_organization.org.id
  path            = ["/terraform/.*.tf"]
  branch          = ["feat", "fix"]
  template_id     = data.terrakube_template.template.id
  workspace_id    = data.terrakube_workspace_vcs.workspace.id
}
