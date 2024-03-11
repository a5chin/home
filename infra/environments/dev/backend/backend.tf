terraform {
  backend "gcs" {
    bucket = "home-tfstate"
    prefix = "backend"
  }
}
