resource "aws_subnet" "subnet_public" {
  for_each = var.public_subnets

  vpc_id                  = var.vpc_id
  cidr_block              = each.value.cidr
  availability_zone       = each.value.az
  map_public_ip_on_launch = true

  tags = {
    Name                       = "${var.env}-public-${each.key}"
    kubernetes.io / role / elb = "1"
  }
}

resource "aws_subnet" "subnet_private" {
  for_each = var.private_subnets

  vpc_id            = var.vpc_id
  cidr_block        = each.value.cidr
  availability_zone = each.value.az

  tags = {
    Name                                = "${var.env}-private-${each.key}"
    kubernetes.io / role / internal-elb = "1"
    karpenter.sh / discovery            = "${var.env}-${var.cluster_name}"
  }
}
