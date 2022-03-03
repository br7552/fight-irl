# Input variable definitions

variable "aws_region" {
  description = "AWS region for all resources."

  type    = string
  default = "us-east-1"
}

variable "mapkey" {
  description = "Google Maps API key"
  type        = string
  sensitive   = true
}
