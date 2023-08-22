terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}
provider "aws" {
  region =var.region
}

resource "aws_instance" "web" {
  ami           = "ami-08a52ddb321b32a8c"
  instance_type = var.instance_type

  tags = {
    Name = var.instance_name
  }
}