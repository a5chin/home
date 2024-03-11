variable "project_id" {
  description = "The ID of the Project."
  type        = string
}

variable "location" {
  description = "The location of resources."
  type        = string
  default     = "us-central1"
}

variable "run" {
  description = "The settings for the Cloud Run in Frontend."
  type = object({
    name  = optional(string, "home-backend")
    image = string
    executor = optional(
      object({
        id = string
      }), {
        id = "home-backend"
      }
    )
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
  })
}
