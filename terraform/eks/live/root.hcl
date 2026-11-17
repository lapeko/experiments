locals {
  bucket_name = "terragrunt-eks-project-state"
  region      = "eu-central-1"
}

remote_state {
  backend = "s3"

  config = {
    bucket       = local.bucket_name
    key          = "${path_relative_to_include()}/terraform.tfstate"
    region       = local.region
    encrypt      = true
    use_lockfile = true
  }
}
