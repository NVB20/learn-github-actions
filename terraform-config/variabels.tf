# AWS 
variable "region" {
    description = "The region of the EC2 instances"
    type        = string
    default     = "us-east-1"
}
#EC2 variabels
variable "instance_name" {
 type        = string
 description = "Instance name for the EC2 instance"
 default     = "golang-instance"
}
variable "instance_type" {
 type        = string
 description = "Instance type for the EC2 instance"
 default     = "t2.micro"
 sensitive   = true
}