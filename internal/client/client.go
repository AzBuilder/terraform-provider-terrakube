package client

import (
	_ "github.com/google/jsonapi"
)

type OrganizationEntity struct {
	ID            string `jsonapi:"primary,organization"`
	Name          string `jsonapi:"attr,name"`
	Description   string `jsonapi:"attr,description"`
	ExecutionMode string `jsonapi:"attr,executionMode"`
	Disabled      bool   `jsonapi:"attr,disabled"`
}

type OrganizationTagEntity struct {
	ID   string `jsonapi:"primary,tag"`
	Name string `jsonapi:"attr,name"`
}

type TeamEntity struct {
	ID              string `jsonapi:"primary,team"`
	Name            string `jsonapi:"attr,name"`
	ManageWorkspace bool   `jsonapi:"attr,manageWorkspace"`
	ManageModule    bool   `jsonapi:"attr,manageModule"`
	ManageProvider  bool   `jsonapi:"attr,manageProvider"`
	ManageVcs       bool   `jsonapi:"attr,manageVcs"`
	ManageTemplate  bool   `jsonapi:"attr,manageTemplate"`
}

type WorkspaceEntity struct {
	ID            string `jsonapi:"primary,workspace"`
	Name          string `jsonapi:"attr,name"`
	Description   string `jsonapi:"attr,description"`
	Source        string `jsonapi:"attr,source"`
	Branch        string `jsonapi:"attr,branch"`
	Folder        string `jsonapi:"attr,folder"`
	IaCType       string `jsonapi:"attr,iacType"`
	IaCVersion    string `jsonapi:"attr,terraformVersion"`
	ExecutionMode string `jsonapi:"attr,executionMode"`
	Deleted       bool   `jsonapi:"attr,deleted"`
}

type WorkspaceTagEntity struct {
	ID    string `jsonapi:"primary,workspacetag"`
	TagID string `jsonapi:"attr,tagId"`
}

type WorkspaceVariableEntity struct {
	ID          string `jsonapi:"primary,variable"`
	Key         string `jsonapi:"attr,key"`
	Value       string `jsonapi:"attr,value"`
	Description string `jsonapi:"attr,description"`
	Category    string `jsonapi:"attr,category"`
	Sensitive   bool   `jsonapi:"attr,sensitive"`
	Hcl         bool   `jsonapi:"attr,hcl"`
}

type OrganizationVariableEntity struct {
	ID          string `jsonapi:"primary,globalvar"`
	Key         string `jsonapi:"attr,key"`
	Value       string `jsonapi:"attr,value"`
	Description string `jsonapi:"attr,description"`
	Category    string `jsonapi:"attr,category"`
	Sensitive   *bool  `jsonapi:"attr,sensitive,omitempty"`
	Hcl         bool   `jsonapi:"attr,hcl"`
}

type VcsEntity struct {
	ID          string `jsonapi:"primary,vcs"`
	Name        string `jsonapi:"attr,name"`
	Description string `jsonapi:"attr,description"`
}

type SshEntity struct {
	ID          string `jsonapi:"primary,ssh"`
	Name        string `jsonapi:"attr,name"`
	Description string `jsonapi:"attr,description"`
}

type ModuleEntity struct {
	ID          string     `jsonapi:"primary,module"`
	Name        string     `jsonapi:"attr,name"`
	Description string     `jsonapi:"attr,description"`
	Provider    string     `jsonapi:"attr,provider"`
	Source      string     `jsonapi:"attr,source"`
	Vcs         *VcsEntity `jsonapi:"relation,vcs,omitempty"`
	Ssh         *SshEntity `jsonapi:"relation,ssh,omitempty"`
}
