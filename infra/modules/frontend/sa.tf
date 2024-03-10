resource "google_service_account" "frontend_executor" {
  project      = var.project_id
  account_id   = var.run.executor.id
  display_name = "The Service Account for Frontend executor in home"
}

resource "google_project_iam_member" "frontend_executor" {
  for_each = var.run.executor.roles
  project  = var.project_id
  role     = each.value
  member   = "serviceAccount:${google_service_account.frontend_executor.email}"
}

resource "google_cloud_run_v2_service_iam_member" "frontend_invoker" {
  project  = google_cloud_run_v2_service.main.project
  location = google_cloud_run_v2_service.main.location
  name     = google_cloud_run_v2_service.main.name
  role     = "roles/run.invoker"
  member   = "allUsers"
}
