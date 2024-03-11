resource "google_service_account" "backend_executor" {
  project      = var.project_id
  account_id   = var.run.executor.id
  display_name = "The Service Account for Backend executor in home"
}

resource "google_project_iam_member" "backend_executor" {
  for_each = var.run.executor.roles
  project  = var.project_id
  role     = each.value
  member   = "serviceAccount:${google_service_account.backend_executor.email}"
}
