terraform {
  backend "gcs" {
    bucket = "home-tfstate"
    prefix = "frontend"
  }
}
