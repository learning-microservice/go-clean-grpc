locals {
  app_name = "${var.project_name}-${var.environment}"
}

terraform {
  required_version = ">= 1.14.9, < 2.0.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 6.0"
    }
  }

  # 【重要】Stateファイルの保存先
  # 最初の1回目（S3バケット作成前）は Makefile 内で -backend=false を指定して実行します
  #backend "s3" {}
}

provider "aws" {
  region = var.aws_region

  default_tags {
    tags = {
      Project   = var.project_name
      Env       = var.environment
      ManagedBy = "Terraform"
    }
  }
}

data "aws_caller_identity" "current" {}

module "ecr_app" {
  source                = "../modules/ecr"
  repository_name       = local.app_name
  push_access_role_arns = [aws_iam_role.github_actions.arn]
}
