.DEFAULT_GOAL := help

# Setting environment variables
export TF_VAR_project_name  ?= go-clean-grpc
export TF_VAR_environment   ?= dev
export TF_VAR_aws_region    ?= ap-northeast-1
export TF_VAR_github_repo   ?= learning-microservice/go-clean-grpc
export TF_VAR_github_branch ?= main
export TF_VAR_github_environment ?= development

ifndef AWS_PROFILE
$(error "環境変数 AWS_PROFILE が設定されていません. export AWS_PROFILE=xxx を設定して下さい.")
endif

.PHONY: help
help: ## Display this help screen
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

# ~~~ Setup Actions ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

.PHONY: setup
setup: ## Setup AWS resources and sync to GitHub environment variables and secrets
	@$(MAKE) .terraform.apply
	@$(MAKE) .sync-github-vars

.PHONY: destroy
destroy: ## Destroy AWS resources and github environment variables and secrets
	@$(MAKE) .delete-github-vars
	@$(MAKE) .terraform.destroy

.terraform.apply:
	@echo "Applying Terraform..."
	@terraform init
	@terraform apply -auto-approve

.terraform.destroy:
	@echo "Destroying Terraform..."
	@terraform init -upgrade -reconfigure
	@terraform destroy -auto-approve

.sync-github-vars:
	@echo "Ensuring GitHub environment '$(TF_VAR_github_environment)' exists..."
	@gh api -X PUT /repos/$(TF_VAR_github_repo)/environments/$(TF_VAR_github_environment) > /dev/null 2>&1 || true
	@echo "Syncing GitHub environment variables and secrets..."
	@terraform output -json github_actions_environment_secrets | jq -r 'to_entries[] | "\(.key)=\(.value)"' | while IFS='=' read -r key value; do \
		gh secret set "$$key" --body "$$value" --env $(TF_VAR_github_environment); \
	done
	@echo "Syncing GitHub Variables..."
	@terraform output -json github_actions_environment_variables | jq -r 'to_entries[] | "\(.key)=\(.value)"' | while IFS='=' read -r key value; do \
		gh variable set "$$key" --body "$$value" --env $(TF_VAR_github_environment); \
	done

.delete-github-vars:
	@echo "Deleting GitHub Secrets..."
	@terraform output -json github_actions_environment_secrets | jq -r 'to_entries[] | .key' | while read -r key; do \
		gh secret delete "$$key" --env $(TF_VAR_github_environment) || true; \
	done
	@echo "Deleting GitHub Variables..."
	@terraform output -json github_actions_environment_variables | jq -r 'to_entries[] | .key' | while read -r key; do \
		gh variable delete "$$key" --env $(TF_VAR_github_environment) || true; \
	done
