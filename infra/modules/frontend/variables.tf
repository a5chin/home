variable "project_id" {
  description = "The ID of the Project."
  type        = string
}

variable "location" {
  description = "The location of resources."
  type        = string
}

variable "run" {
  description = "The settings for the Cloud Run in Frontend."
  type = object({
    name               = string
    image              = string
    max_instance_count = optional(number, 1)
    min_instance_count = optional(number, 0)
    cpu                = optional(string, "1000m")
    memory             = optional(string, "128Mi")
    executor = object({
      id = string
      roles = optional(
        set(string), [
          "roles/cloudtrace.agent",
        ]
      )
    })
  })
}

variable "backend" {
  description = "Outputs in the Backend."
}
