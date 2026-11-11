variable "vpc_id" {}

variable "subnet_cidr" {
  default = "10.0.0.0/18"
}

variable "env" {
  default = "dev"
}
