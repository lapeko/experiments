provider "aws" {
  region = var.region
}

module "vpc" {
  source = "../../../modules/vpc"

  vpc_cidr = var.vpn_cidr
  env = var.env
}
