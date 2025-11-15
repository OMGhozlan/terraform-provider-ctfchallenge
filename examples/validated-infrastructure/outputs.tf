output "infrastructure_id" {
  description = "Primary infrastructure ID"
  value       = ctfchallenge_validated_resource.primary.computed_id

  precondition {
    condition     = ctfchallenge_validated_resource.primary.validated == true
    error_message = "Cannot expose infrastructure_id - primary resource not validated."
  }
}

output "instance_ids" {
  description = "All instance IDs"
  value       = ctfchallenge_validated_resource.secondary[*].computed_id

  precondition {
    condition     = length(ctfchallenge_validated_resource.secondary) >= var.min_instances
    error_message = "Insufficient instances created - cannot expose IDs."
  }

  precondition {
    condition     = alltrue([for r in ctfchallenge_validated_resource.secondary : r.validated])
    error_message = "Not all instances passed validation - cannot expose IDs."
  }
}

output "deployment_status" {
  description = "Overall deployment status"
  value = {
    environment    = var.environment
    primary_state  = ctfchallenge_validated_resource.primary.state
    instance_count = length(ctfchallenge_validated_resource.secondary)
    all_validated  = alltrue([for r in ctfchallenge_validated_resource.secondary : r.validated])
  }

  precondition {
    condition     = ctfchallenge_validated_resource.primary.quality_score >= 50
    error_message = <<-EOT
      Cannot expose deployment status - quality threshold not met.
      Current score: ${ctfchallenge_validated_resource.primary.quality_score}
      Required: 50+
    EOT
  }
}