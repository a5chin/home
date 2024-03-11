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
    name  = optional(string, "home-frontend")
    image = string
    executor = object({
      id = string
    })
  })
}
