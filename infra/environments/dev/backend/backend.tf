terraform {
  backend "gcs" {
    bucket = "home-tfstate"
    prefix = "backend"
  }
}

data "terraform_remote_state" "frontend" {
  backend   = "gcs"
  workspace = terraform.workspace
  config = {
    bucket = "home-tfstate"
    prefix = "frontend"
  }
}
