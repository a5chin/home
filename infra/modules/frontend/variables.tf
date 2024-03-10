variable "project_id" {
  type        = string
  description = "The ID of the Project."
}

variable "location" {
  type        = string
  description = "The location of resources."
}

variable "run" {
  type = object({
    name  = optional(string, "home-frontend")
    image = string
    max_instance_count = optional(number, 1)
    min_instance_count = optional(number, 0)
    cpu                = optional(string, "1")
    memory             = optional(string, "128Mi")
    executor           = object({
      id = string
      roles = optional(
        set(string), [
          "roles/cloudtrace.agent",
        ]
      )
    })
  })
 description = "The settings for the Cloud Run in Frontend"
}
