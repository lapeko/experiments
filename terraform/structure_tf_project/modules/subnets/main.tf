resource "aws_subnet" "private" {
  vpc_id = var.vpc_id
  cidr_block = var.subnet_cidr

  tags = {
    Name = "${var.env}-private"
  }
}
