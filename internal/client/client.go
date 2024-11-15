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

type OrganizationTemplateEntity struct {
	ID          string `jsonapi:"primary,template"`
	Name        string `jsonapi:"attr,name"`
	Description string `jsonapi:"attr,description"`
	Version     string `jsonapi:"attr,version"`
	Content     string `jsonapi:"attr,tcl"`
}

type OrganizationTagEntity struct {
	ID   string `jsonapi:"primary,tag"`
	Name string `jsonapi:"attr,name"`
}

type TeamEntity struct {
	ID              string `jsonapi:"primary,team"`
	Name            string `jsonapi:"attr,name"`
	ManageState     bool   `jsonapi:"attr,manageState"`
	ManageWorkspace bool   `jsonapi:"attr,manageWorkspace"`
	ManageModule    bool   `jsonapi:"attr,manageModule"`
	ManageProvider  bool   `jsonapi:"attr,manageProvider"`
	ManageVcs       bool   `jsonapi:"attr,manageVcs"`
	ManageTemplate  bool   `jsonapi:"attr,manageTemplate"`
}

type TeamTokenEntity struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Days        int32  `json:"days"`
	Hours       int32  `json:"hours"`
	Minutes     int32  `json:"minutes"`
	Group       string `json:"group"`
	Value       string `json:"token"`
}

type WorkspaceEntity struct {
	ID            string     `jsonapi:"primary,workspace"`
	Name          string     `jsonapi:"attr,name"`
	Description   string     `jsonapi:"attr,description"`
	Source        string     `jsonapi:"attr,source"`
	Branch        string     `jsonapi:"attr,branch"`
	Folder        string     `jsonapi:"attr,folder"`
	TemplateId    string     `jsonapi:"attr,defaultTemplate"`
	IaCType       string     `jsonapi:"attr,iacType"`
	IaCVersion    string     `jsonapi:"attr,terraformVersion"`
	ExecutionMode string     `jsonapi:"attr,executionMode"`
	Deleted       bool       `jsonapi:"attr,deleted"`
	Vcs           *VcsEntity `jsonapi:"relation,vcs,omitempty"`
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
	ID             string `jsonapi:"primary,vcs"`
	Name           string `jsonapi:"attr,name"`
	Description    string `jsonapi:"attr,description"`
	VcsType        string `jsonapi:"attr,vcsType"`
	ConnectionType string `jsonapi:"attr,connectionType"`
	ClientId       string `jsonapi:"attr,clientId"`
	ClientSecret   string `jsonapi:"attr,clientSecret"`
	PrivateKey     string `jsonapi:"attr,privateKey"`
	Endpoint       string `jsonapi:"attr,endpoint"`
	ApiUrl         string `jsonapi:"attr,apiUrl"`
	Status         string `jsonapi:"attr,status"`
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
	Folder      *string    `jsonapi:"attr,folder"`
	TagPrefix   *string    `jsonapi:"attr,tagPrefix"`
}

type CollectionEntity struct {
	ID          string `jsonapi:"primary,collection"`
	Name        string `jsonapi:"attr,name"`
	Description string `jsonapi:"attr,description"`
	Priority    int32  `jsonapi:"attr,priority"`
}

type CollectionItemEntity struct {
	ID          string `jsonapi:"primary,item"`
	Key         string `jsonapi:"attr,key"`
	Value       string `jsonapi:"attr,value"`
	Description string `jsonapi:"attr,description"`
	Category    string `jsonapi:"attr,category"`
	Sensitive   bool   `jsonapi:"attr,sensitive"`
	Hcl         bool   `jsonapi:"attr,hcl"`
}

type CollectionReferenceEntity struct {
	ID          string            `jsonapi:"primary,reference"`
	Description string            `jsonapi:"attr,description"`
	Workspace   *WorkspaceEntity  `jsonapi:"relation,workspace,omitempty"`
	Collection  *CollectionEntity `jsonapi:"relation,collection,omitempty"`
}

type WorkspaceWebhookEntity struct {
	ID           string `jsonapi:"primary,webhook"`
	Path         string `jsonapi:"attr,path"`
	Branch       string `jsonapi:"attr,branch"`
	TemplateId   string `jsonapi:"attr,templateId"`
	RemoteHookId string `jsonapi:"attr,remoteHookId"`
	Event        string `jsonapi:"attr,event"`
}

type WorkspaceScheduleEntity struct {
	ID         string `jsonapi:"primary,schedule"`
	Schedule   string `jsonapi:"attr,cron"`
	TemplateId string `jsonapi:"attr,templateReference"`
}
