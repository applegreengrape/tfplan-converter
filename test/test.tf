resource "aws_s3_bucket" "b2" {
  bucket = "tf-plan-converter-test"
  acl    = "private"

  tags = {
    Name        = "My bucket"
    Environment = "Dev"
  }
}