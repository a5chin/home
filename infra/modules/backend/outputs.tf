output "base_url" {
  description = "The URL for the Backend Cloud Run."
  value = google_cloud_run_v2_service.backend.uri
}
