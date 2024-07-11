resource "terrakube_organization_template" "example" {
  name            = "example"
  organization_id = terrakube_organization.example.id
  description     = "Example organization template"
  version         = "1.0.0"
  content         = <<EOF
flow:
  - type: "terraformPlan"
    name: "Plan"
    step: 100
  EOF
}