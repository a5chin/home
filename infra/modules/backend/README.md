<!-- BEGIN_TF_DOCS -->
## Requirements

No requirements.

## Providers

| Name | Version |
|------|---------|
| <a name="provider_google"></a> [google](#provider\_google) | n/a |
| <a name="provider_google-beta"></a> [google-beta](#provider\_google-beta) | n/a |

## Modules

No modules.

## Resources

| Name | Type |
|------|------|
| [google-beta_google_vpc_access_connector.main](https://registry.terraform.io/providers/hashicorp/google-beta/latest/docs/resources/google_vpc_access_connector) | resource |
| [google_cloud_run_v2_service.backend](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/cloud_run_v2_service) | resource |
| [google_cloud_run_v2_service_iam_member.backend_invoker](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/cloud_run_v2_service_iam_member) | resource |
| [google_compute_global_address.private_ip_address](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_global_address) | resource |
| [google_compute_network.main](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_network) | resource |
| [google_compute_subnetwork.main](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/compute_subnetwork) | resource |
| [google_project_iam_member.backend_executor](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/project_iam_member) | resource |
| [google_service_account.backend_executor](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/service_account) | resource |
| [google_service_networking_connection.private_vpc_connection](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/service_networking_connection) | resource |
| [google_sql_database.main](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/sql_database) | resource |
| [google_sql_database_instance.main](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/sql_database_instance) | resource |
| [google_sql_user.main](https://registry.terraform.io/providers/hashicorp/google/latest/docs/resources/sql_user) | resource |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:--------:|
| <a name="input_db"></a> [db](#input\_db) | Settings for the Cloud SQL. | <pre>object({<br>    instance_name = string<br>    database_name = optional(string, "home")<br>    version       = optional(string, "db-f1-micro")<br>    charset       = optional(string, "utf8mb4")<br>    collation     = optional(string, "utf8mb4_unicode_ci")<br>  })</pre> | n/a | yes |
| <a name="input_frontend"></a> [frontend](#input\_frontend) | Outputs in the Frontend. | `any` | n/a | yes |
| <a name="input_location"></a> [location](#input\_location) | The location of the resource. | `string` | n/a | yes |
| <a name="input_project_id"></a> [project\_id](#input\_project\_id) | The ID of Google Cloud Platform. | `string` | n/a | yes |
| <a name="input_run"></a> [run](#input\_run) | Settings for the Cloud Run in Backend. | <pre>object({<br>    name               = string<br>    image              = string<br>    max_instance_count = optional(number, 1)<br>    min_instance_count = optional(number, 0)<br>    cpu                = optional(string, "1")<br>    memory             = optional(string, "128Mi")<br>    executor = object({<br>      id = string<br>      roles = optional(<br>        set(string), [<br>          "roles/cloudsql.client",<br>          "roles/cloudtrace.agent",<br>        ]<br>      )<br>    })<br>    env = object({<br>      HOSTNAME   = string<br>      DB_USER    = string<br>      DB_PWD     = string<br>      DB_NAME    = string<br>      DB_TCPHOST = string<br>      DB_PORT    = number<br>    })<br>  })</pre> | n/a | yes |

## Outputs

No outputs.
<!-- END_TF_DOCS -->