resource "google_artifact_registry_repository" "main" {
  location      = var.location
  repository_id = "docker-repo"
  format        = "DOCKER"
}
