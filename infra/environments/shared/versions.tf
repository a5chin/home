terraform {
  required_version = "1.7.4"
  required_providers {
    google = {
      source  = "hashicorp/google"
      version = "5.19.0"
    }
    google-beta = {
      source  = "hashicorp/google-beta"
      version = "5.19.0"
    }
  }
}
