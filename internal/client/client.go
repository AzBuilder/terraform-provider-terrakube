package client

import (
	_ "github.com/google/jsonapi"
)

type OrganizationEntity struct {
	ID          string `jsonapi:"primary,organization"`
	Name        string `jsonapi:"attr,name"`
	Description string `jsonapi:"attr,description"`
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
