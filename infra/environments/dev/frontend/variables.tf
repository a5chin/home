variable "project_id" {
  type        = string
  description = "The ID of the Project."
}

variable "location" {
  type        = string
  default     = "us-central1"
  description = "The location of resources."
}

variable "run" {
  type = object({
    image = string
    executor = object({
      id = string
    })
  })
  description = "The settings for the Cloud Run in Frontend"
}
