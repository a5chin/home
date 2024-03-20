terraform {
  backend "gcs" {
    bucket = "home-tfstate"
    prefix = "frontend"
  }
}

data "terraform_remote_state" "backend" {
  backend   = "gcs"
  workspace = terraform.workspace
  config = {
    bucket = "home-tfstate"
    prefix = "backend"
  }
}
