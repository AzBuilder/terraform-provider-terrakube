terraform {
  required_providers {
    terrakube = {
      source = "registry.terraform.io/alfespa17/terrakube"
    }
  }
}

provider "terrakube" {
  endpoint = "http://terrakube-api.minikube.net"
  token    = "eyJhbGciOiJIUzI1NiJ9.eyJpc3MiOiJUZXJyYWt1YmUiLCJzdWIiOiJMZXN0ZXIgKFRva2VuKSIsImF1ZCI6IlRlcnJha3ViZSIsImp0aSI6IjI5ZmZhMzc5LTE1NWUtNDhlYS04MDJhLTMwZWUyZmRlMjQwOSIsImVtYWlsIjoiYWRtaW5AZXhhbXBsZS5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwibmFtZSI6Ikxlc3RlciAoVG9rZW4pIiwiZ3JvdXBzIjpbIlRFUlJBS1VCRV9BRE1JTiIsIlRFUlJBS1VCRV9ERVZFTE9QRVJTIl0sImlhdCI6MTY5NTc2NjUwMSwiZXhwIjoxNjk2MzcxMzAxfQ._cDRy4oH11OcWTRFiyxZhTENODwhBStmWwqrGj4yPRs"
}

data "terrakube_organization" "org" {
  name = "simple"
}

data "terrakube_vcs" "vcs" {
  name = "sample"
  organization_id = data.terrakube_organization.org.id
}

data "terrakube_ssh" "ssh" {
  name = "hola"
  organization_id = data.terrakube_organization.org.id
}

resource "terrakube_team" "team" {
  name             = "TERRAKUBE_SUPER_ADMIN"
  organization_id  = data.terrakube_organization.org.id
  manage_workspace = false
  manage_module    = false
  manage_provider  = true
  manage_vcs       = true
  manage_template  = true
}