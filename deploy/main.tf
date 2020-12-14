terraform {
  backend "s3" {
    bucket = "paqot-api-tfstate"
    key = "paqot-api.tfstate"
    region = "us-east-1"
    encrypt = true
    dynamodb_table = "paqot-api-tf-state-lock"
  }
}

provider "aws" {
  region = "us-east-1"
  version = "~> 3.20.0"
}