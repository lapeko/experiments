variable "region" {}
variable "env" {}
variable "vpc_id" {}
variable "public_subnets" {
  type = map(object({
    cidr = string
    az = string
  }))
}
variable "private_subnets" {
  type = map(object({
    cidr = string
    az = string
  }))
}
variable "cluster_name" {}
