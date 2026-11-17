provider "aws" {
  region = var.aws_region
}

terraform {
  required_version = "~>1.17.2"

  required_providers {
    aws = {
      version = "~>6.73.0"
    }
  }
}
