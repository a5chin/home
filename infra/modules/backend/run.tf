resource "google_cloud_run_v2_service" "backend" {
  name     = var.run.name
  location = var.location

  template {
    scaling {
      max_instance_count = var.run.max_instance_count
      min_instance_count = var.run.min_instance_count
    }

    vpc_access {
      connector = google_vpc_access_connector.main.id
      egress    = "PRIVATE_RANGES_ONLY"
    }

    service_account = google_service_account.backend_executor.email

    containers {
      image = var.run.image

      resources {
        cpu_idle = true
        limits = {
          cpu    = var.run.cpu
          memory = var.run.memory
        }
      }

      dynamic "env" {
        for_each = var.run.env
        content {
          name  = env.key
          value = env.value
        }
      }
    }
  }

  depends_on = [
    google_sql_database_instance.main
  ]
}


