terraform {
  required_providers {
    terrakube = {
      source = "registry.terraform.io/alfespa17/terrakube"
    }
  }
}

provider "terrakube" {
  endpoint             = "http://terrakube-api.minikube.net"
  token                = "12345"
  insecure_http_client = true
}