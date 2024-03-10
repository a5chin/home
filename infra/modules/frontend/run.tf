resource "google_cloud_run_v2_service" "main" {
  name     = var.run.name
  location = var.location

  template {
    scaling {
      max_instance_count = var.run.max_instance_count
      min_instance_count = var.run.min_instance_count
    }

    service_account = google_service_account.frontend_executor.email

    containers {
      image = var.run.image

      resources {
        cpu_idle = true
        limits = {
          cpu    = var.run.cpu
          memory = var.run.memory
        }
      }
    }
  }
}
