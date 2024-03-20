module "frontend" {
  source     = "../../../modules/frontend"
  project_id = var.project_id
  location   = var.location
  run        = var.run
  backend    = data.terraform_remote_state.backend.outputs
}
