resource "aws_dynamodb_table" "basic_dynamodb_table" {
  name           = "terraform-locks"
  billing_mode   = "PAY_PER_REQUEST"
  hash_key       = "LockID"

  attribute {
    name = "LockID"
    type = "S"
  }
}

resource "aws_s3_bucket" "state_bucket" {
  bucket = "ntriller-terraform"
}
