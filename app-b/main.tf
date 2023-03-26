terraform {
  backend "s3" {
    bucket = "ntriller-terraform"
    key    = "app-b"
    region = "eu-central-1"
    dynamodb_table = "terraform-locks"
  }
}

resource "aws_iam_role" "test_role_b" {
  name = "test_role_b"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      },
    ]
  })
}

resource "time_sleep" "wait" {
  create_duration = "3s"
}
