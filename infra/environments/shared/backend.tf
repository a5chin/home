terraform {
  backend "gcs" {
    bucket = "home-tfstate"
    prefix = "shared"
  }
}
