variable "project_id" {
  description = "The ID of Google Cloud Platform."
  type        = string
}

variable "location" {
  description = "The location of the resource."
  type        = string
}

variable "run" {
  description = "Settings for the Cloud Run in Backend."
  type = object({
    name               = string
    image              = string
    max_instance_count = optional(number, 1)
    min_instance_count = optional(number, 0)
    cpu                = optional(string, "1")
    memory             = optional(string, "128Mi")
    executor = object({
      id = string
      roles = optional(
        set(string), [
          "roles/cloudsql.client",
          "roles/cloudtrace.agent",
        ]
      )
    })
    env = object({
      HOSTNAME   = string
      DB_USER    = string
      DB_PWD     = string
      DB_NAME    = string
      DB_TCPHOST = string
      DB_PORT    = number
    })
  })
}

variable "db" {
  description = "Settings for the Cloud SQL."
  type = object({
    instance_name = string
    database_name = optional(string, "home")
    version       = optional(string, "db-f1-micro")
    charset       = optional(string, "utf8mb4")
    collation     = optional(string, "utf8mb4_unicode_ci")
  })
}
