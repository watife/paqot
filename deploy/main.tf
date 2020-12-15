terraform {
  backend "s3" {
    bucket         = "paqot-api-tfstate"
    key            = "paqot-api.tfstate"
    region         = "us-east-1"
    encrypt        = true
    dynamodb_table = "paqot-api-tf-state-lock"
  }

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.20.0"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}

locals {
  prefix = "${var.prefix}-${terraform.workspace}"
}