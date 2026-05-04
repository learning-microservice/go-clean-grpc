variable "aws_region" {
  type        = string
  description = "AWSリージョン"
}

variable "project_name" {
  type        = string
  description = "プロジェクト名"
}

variable "environment" {
  description = "環境名（dev, prdなど）"
  type        = string
}

variable "github_repo" {
  type        = string
  description = "ユーザー名/リポジトリ名 (例: learning-microservice/go-clean-grpc)"
}

variable "github_environment" {
  type        = string
  description = "Github Environment (例: development)"
}

variable "github_branch" {
  type        = string
  description = "Githubブランチ名 (例: main)"
}
