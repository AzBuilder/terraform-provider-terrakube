---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "terrakube_workspace_vcs Resource - terraform-provider-terrakube"
subcategory: ""
description: |-
  
---

# terrakube_workspace_vcs (Resource)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `branch` (String) Workspace VCS branch
- `description` (String) Workspace VCS description
- `execution_mode` (String) Workspace VCS execution mode (remote or local)
- `folder` (String) Workspace VCS folder
- `iac_type` (String) Workspace VCS IaC type (Supported values terraform or tofu)
- `iac_version` (String) Workspace VCS VCS type
- `name` (String) Workspace VCS name
- `organization_id` (String) Terrakube organization id
- `repository` (String) Workspace VCS repository
- `template_id` (String) Default template ID for the workspace

### Optional

- `vcs_id` (String) VCS connection ID for private workspaces

### Read-Only

- `id` (String) Workspace CLI Id
