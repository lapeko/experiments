resource "aws_vpc" "vpc_main" {
  cidr_block = var.cidr

  enable_dns_hostnames = true
  enable_dns_support   = true

  tags = {
    Name = "${var.env}-main"
  }
}

resource "aws_subnet" "subnet_public" {
  for_each = var.public_subnets

  vpc_id                  = aws_vpc.vpc_main.id
  cidr_block              = each.value.cidr
  availability_zone       = each.value.az
  map_public_ip_on_launch = true

  tags = {
    Name                     = "${var.env}-public-${each.key}"
    "kubernetes.io/role/elb" = "1"
  }
}

resource "aws_subnet" "subnet_private" {
  for_each = var.private_subnets

  vpc_id            = aws_vpc.vpc_main.id
  cidr_block        = each.value.cidr
  availability_zone = each.value.az

  tags = {
    Name                              = "${var.env}-private-${each.key}"
    "kubernetes.io/role/internal-elb" = "1"
    "karpenter.sh/discovery"          = "${var.env}-${var.cluster_name}"
  }
}

resource "aws_internet_gateway" "igw_main" {
  vpc_id = aws_vpc.vpc_main.id

  tags = {
    Name = "${var.env}-igw"
  }
}

resource "aws_eip" "eip_nat" {
  domain = "vpc"

  tags = {
    Name = "${var.env}-nat-eip"
  }

  depends_on = [aws_internet_gateway.igw_main]
}

resource "aws_nat_gateway" "nat_gateway" {
  allocation_id = aws_eip.eip_nat.id
  subnet_id     = values(aws_subnet.subnet_public)[0].id

  tags = {
    Name = "${var.env}-nat-gateway"
  }
}

resource "aws_route_table" "public" {
  vpc_id = aws_vpc.vpc_main.id

  route {
    cidr_block = "0.0.0.0/0"
    gateway_id = aws_internet_gateway.igw_main.id
  }

  tags = {
    Name = "${var.env}-public-rt"
  }
}

resource "aws_route_table_association" "public" {
  for_each       = aws_subnet.subnet_public
  subnet_id      = each.value.id
  route_table_id = aws_route_table.public.id
}

resource "aws_route_table" "private" {
  vpc_id = aws_vpc.vpc_main.id

  route {
    cidr_block     = "0.0.0.0/0"
    nat_gateway_id = aws_nat_gateway.nat_gateway.id
  }

  tags = {
    Name = "${var.env}-private-rt"
  }
}

resource "aws_route_table_association" "private" {
  for_each       = aws_subnet.subnet_private
  subnet_id      = each.value.id
  route_table_id = aws_route_table.private.id
}
