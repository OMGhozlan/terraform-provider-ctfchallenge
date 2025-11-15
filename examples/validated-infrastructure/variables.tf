variable "environment" {
  type        = string
  description = "Deployment environment"

  validation {
    condition     = contains(["dev", "staging", "prod"], var.environment)
    error_message = "Environment must be dev, staging, or prod."
  }
}

variable "min_instances" {
  type        = number
  description = "Minimum number of instances"

  validation {
    condition     = var.min_instances >= 1 && var.min_instances <= 10
    error_message = "min_instances must be between 1 and 10."
  }
}

variable "enable_monitoring" {
  type        = bool
  description = "Enable monitoring"
  default     = true
}