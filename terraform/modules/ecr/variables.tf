variable "repository_name" {
  description = "ECRリポジトリ名（例: go-clean-grpc-dev）"
  type        = string
}

variable "push_access_role_arns" {
  description = "ECRへのPush権限を付与するIAMロールのARNリスト"
  type        = list(string)
}
