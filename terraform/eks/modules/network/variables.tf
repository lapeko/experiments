variable "region" {
  type    = string
  default = "eu-central-1"
}

variable "env" {
  type    = string
  default = "dev"
}

variable "cluster_name" {
  type    = string
  default = "cluster"
}

variable "cidr" {
  type    = string
  default = "10.0.0.0/16"
}

variable "public_subnets" {
  type = list(object({
    cidr = string
    az   = string
  }))

  default = [
    {
      cidr = "10.0.0.0/19"
      az   = "eu-central-1a"
    },
    {
      cidr = "10.0.32.0/19"
      az   = "eu-central-1b"
    }
  ]
}

variable "private_subnets" {
  type = list(object({
    cidr = string
    az   = string
  }))

  default = [
    {
      cidr = "10.0.64.0/19"
      az   = "eu-central-1a"
    },
    {
      cidr = "10.0.96.0/19"
      az   = "eu-central-1b"
    }
  ]
}
