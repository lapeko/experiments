resource "aws_s3_bucket" "eks_project_tf_state" {
  bucket = "terragrunt-eks-project-state"
}
