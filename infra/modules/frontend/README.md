<!-- BEGIN_TF_DOCS -->
## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_google"></a> [google](#provider\_google) | n/a |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [google_cloud_run_v2_service.main](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/cloud_run_v2_service) | resource |
| [google_cloud_run_v2_service_iam_member.frontend_invoker](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/cloud_run_v2_service_iam_member) | resource |
| [google_project_iam_member.frontend_executor](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/project_iam_member) | resource |
| [google_service_account.frontend_executor](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/service_account) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_location"></a> [location](#input\_location) | The location of resources. | `string` | n/a | yes |
| <a name="input_project_id"></a> [project\_id](#input\_project\_id) | The ID of the Project. | `string` | n/a | yes |
| <a name="input_run"></a> [run](#input\_run) | The settings for the Cloud Run in Frontend | <pre>object({<br>    name               = optional(string, "home-frontend")<br>    image              = string<br>    max_instance_count = optional(number, 1)<br>    min_instance_count = optional(number, 0)<br>    cpu                = optional(string, "1")<br>    memory             = optional(string, "128Mi")<br>    executor = object({<br>      id = string<br>      roles = optional(<br>        set(string), [<br>          "roles/cloudtrace.agent",<br>        ]<br>      )<br>    })<br>  })</pre> | n/a | yes |

## Outputs

No outputs.
<!-- END_TF_DOCS -->