data "terraform_remote_state" "vpc" {
  backend = "local"
  config = {
    path = "../vpc/terraform.tfstate"
  }
}

module "subnets" {
  source = "../../../modules/subnets"
  vpc_id = data.terraform_remote_state.vpc.outputs.vpc_id
  subnet_cidr = var.subnet_cidr
  env = var.env
}
