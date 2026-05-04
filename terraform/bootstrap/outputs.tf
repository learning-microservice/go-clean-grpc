output "tfstate_bucket_name" {
  value = aws_s3_bucket.terraform_state.bucket
}

output "tfstate_lock_table_name" {
  value = aws_dynamodb_table.terraform_lock.name
}

output "github_actions_role_arn" {
  value     = aws_iam_role.github_actions.arn
  sensitive = true
}

output "ecr_repository_url" {
  value = module.ecr_app.repository_url
}

output "tfstate_backend_bootstrap_config" {
  value = local.backend_bootstrap_config
}

output "tfstate_backend_app_config" {
  value = local.backend_app_config
}

output "github_actions_environment_secrets" {
  value = {
    AWS_ROLE_ARN = aws_iam_role.github_actions.arn
  }
  sensitive = true
}

output "github_actions_environment_variables" {
  value = {
    APP_NAME           = local.app_name
    AWS_REGION         = var.aws_region
    ECR_REPOSITORY_URL = module.ecr_app.repository_url
    TF_CLI_ARGS_init   = local.backend_app_config
    TF_WORKING_DIR     = "terraform/envs/${var.environment}"
  }
}

locals {
  backend_bootstrap_config = join(" ", [
    "-backend-config=\"region=${var.aws_region}\"",
    "-backend-config=\"bucket=${aws_s3_bucket.terraform_state.bucket}\"",
    "-backend-config=\"key=bootstrap.tfstate\"",
    "-backend-config=\"dynamodb_table=${aws_dynamodb_table.terraform_lock.name}\"",
    "-backend-config=\"encrypt=true\"",
  ])
  backend_app_config = join(" ", [
    "-backend-config=\"region=${var.aws_region}\"",
    "-backend-config=\"bucket=${aws_s3_bucket.terraform_state.bucket}\"",
    "-backend-config=\"key=${local.app_name}.tfstate\"",
    "-backend-config=\"dynamodb_table=${aws_dynamodb_table.terraform_lock.name}\"",
    "-backend-config=\"encrypt=true\"",
  ])
}