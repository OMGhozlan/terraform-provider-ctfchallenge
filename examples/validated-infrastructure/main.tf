resource "ctfchallenge_validated_resource" "primary" {
  name           = "primary-${var.environment}"
  required_value = var.environment

  lifecycle {
    precondition {
      condition     = var.environment == "prod" ? var.enable_monitoring == true : true
      error_message = "Production environment MUST have monitoring enabled for compliance."
    }

    precondition {
      condition     = var.environment == "prod" ? var.min_instances >= 3 : true
      error_message = "Production requires minimum 3 instances for high availability."
    }

    postcondition {
      condition     = self.validated == true
      error_message = "Primary resource validation failed - check configuration."
    }

    postcondition {
      condition     = self.quality_score >= (var.environment == "prod" ? 90 : 50)
      error_message = "${var.environment} environment quality requirements not met."
    }
  }
}

resource "ctfchallenge_validated_resource" "secondary" {
  count = var.min_instances

  name           = "secondary-${var.environment}-${count.index}"
  required_value = "instance-${count.index}"

  depends_on = [ctfchallenge_validated_resource.primary]

  lifecycle {
    precondition {
      condition     = ctfchallenge_validated_resource.primary.validated == true
      error_message = "Cannot create secondary resources - primary not validated."
    }

    postcondition {
      condition     = self.state == "active"
      error_message = "Secondary instance ${count.index} failed to activate."
    }
  }
}