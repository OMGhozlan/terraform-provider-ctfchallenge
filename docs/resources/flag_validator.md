---
page_title: "ctfchallenge_flag_validator Resource - ctfchallenge"
subcategory: ""
description: |-
  Validates challenge solutions and reveals flags upon successful completion.
---

# ctfchallenge_flag_validator (Resource)

The `flag_validator` resource is used to submit your solution to a challenge. You provide proof that you completed the required work, and if validation succeeds, **the flag is revealed as your reward**.

This follows the standard CTF paradigm where flags are captured as proof of completion, not provided as input.

## Example Usage

```terraform
# Create dependent resources (your solution)
resource "null_resource" "first" {}

resource "null_resource" "second" {
  depends_on = [null_resource.first]
}

resource "null_resource" "third" {
  depends_on = [null_resource.second]
}

# Submit your solution
resource "ctfchallenge_flag_validator" "basics" {
  challenge_id = "terraform_basics"

  proof_of_work = {
    dependencies = "${null_resource.first.id},${null_resource.second.id},${null_resource.third.id}"
  }
}

# The flag is revealed upon successful validation!
output "captured_flag" {
  value     = ctfchallenge_flag_validator.basics.flag
  sensitive = true
}

output "result" {
  value = ctfchallenge_flag_validator.basics.message
}

output "points_earned" {
  value = ctfchallenge_flag_validator.basics.points
}
```

## Structure-Based Validation

For validation challenges, submit the actual structure of your resources:

### Precondition Challenge with Resource Structure

```terraform
variable "min_length" {
  type    = number
  default = 5
}

resource "ctfchallenge_validated_resource" "guarded" {
  name           = "precondition-demo"
  required_value = var.min_length >= 5 ? "valid-value" : ""
  
  lifecycle {
    precondition {
      condition     = var.min_length >= 5
      error_message = "min_length must be at least 5 characters for security compliance"
    }
    
    precondition {
      condition     = var.min_length <= 100
      error_message = "min_length cannot exceed 100 characters to prevent buffer issues"
    }
  }
}

resource "ctfchallenge_flag_validator" "precondition_challenge" {
  challenge_id = "precondition_guardian"
  
  resource_proof {
    resource_type = "ctfchallenge_validated_resource"
    resource_name = "guarded"
    
    attributes = {
      name      = ctfchallenge_validated_resource.guarded.name
      validated = tostring(ctfchallenge_validated_resource.guarded.validated)
    }
    
    # Submit your lifecycle configuration
    lifecycle_config = jsonencode({
      preconditions = [
        {
          condition     = "var.min_length >= 5"
          error_message = "min_length must be at least 5 characters for security compliance"
        },
        {
          condition     = "var.min_length <= 100"
          error_message = "min_length cannot exceed 100 characters to prevent buffer issues"
        }
      ]
    })
  }
}

output "flag" {
  value     = ctfchallenge_flag_validator.precondition_challenge.flag
  sensitive = true
}

output "details" {
  value = ctfchallenge_flag_validator.precondition_challenge.validation_details
}
```

### Postcondition Challenge with Self References

```terraform
resource "ctfchallenge_validated_resource" "post_validated" {
  name           = "postcondition-demo"
  required_value = "important-data"
  
  lifecycle {
    postcondition {
      condition     = self.validated == true
      error_message = "Resource '${self.name}' failed validation. Current state: ${self.state}"
    }
    
    postcondition {
      condition     = self.quality_score >= 50
      error_message = "Quality score ${self.quality_score} below minimum 50"
    }
    
    postcondition {
      condition     = self.solved == true
      error_message = "Resource not in solved state"
    }
  }
}

resource "ctfchallenge_flag_validator" "postcondition_challenge" {
  challenge_id = "postcondition_validator"
  
  resource_proof {
    resource_type = "ctfchallenge_validated_resource"
    resource_name = "post_validated"
    
    attributes = {
      name          = ctfchallenge_validated_resource.post_validated.name
      validated     = tostring(ctfchallenge_validated_resource.post_validated.validated)
      solved        = tostring(ctfchallenge_validated_resource.post_validated.solved)
      quality_score = tostring(ctfchallenge_validated_resource.post_validated.quality_score)
    }
    
    lifecycle_config = jsonencode({
      postconditions = [
        {
          condition     = "self.validated == true"
          error_message = "Resource failed validation. Current state: ${self.state}"
        },
        {
          condition     = "self.quality_score >= 50"
          error_message = "Quality score ${self.quality_score} below minimum 50"
        },
        {
          condition     = "self.solved == true"
          error_message = "Resource not in solved state"
        }
      ]
    })
  }
}
```

### Data Source Validation

```terraform
data "ctfchallenge_challenge_info" "validated_data" {
  challenge_id = "data_validator"
  
  lifecycle {
    postcondition {
      condition     = self.points > 0
      error_message = "Challenge '${self.name}' has invalid points: ${self.points}"
    }
    
    postcondition {
      condition     = contains(["beginner", "intermediate", "advanced"], self.difficulty)
      error_message = "Invalid difficulty level: ${self.difficulty}"
    }
  }
}

resource "ctfchallenge_flag_validator" "data_challenge" {
  challenge_id = "data_validator"
  
  data_source_proof {
    data_source_type = "ctfchallenge_challenge_info"
    data_source_name = "validated_data"
    
    attributes = {
      name       = data.ctfchallenge_challenge_info.validated_data.name
      points     = tostring(data.ctfchallenge_challenge_info.validated_data.points)
      difficulty = data.ctfchallenge_challenge_info.validated_data.difficulty
    }
    
    lifecycle_config = jsonencode({
      postconditions = [
        {
          condition     = "self.points > 0"
          error_message = "Challenge has invalid points"
        },
        {
          condition     = "contains([\"beginner\", \"intermediate\", \"advanced\"], self.difficulty)"
          error_message = "Invalid difficulty level"
        }
      ]
    })
  }
}
```

### Module Contract Validation

```terraform
resource "ctfchallenge_flag_validator" "module_challenge" {
  challenge_id = "module_contract"
  
  module_proof {
    module_name = "infrastructure"
    
    # Input validations (preconditions on variables)
    input_validations = jsonencode([
      {
        type          = "precondition"
        condition     = "var.instance_count >= 1 && var.instance_count <= 10"
        error_message = "instance_count must be between 1 and 10 for cost control"
        target        = "var.instance_count"
      },
      {
        type          = "precondition"
        condition     = "contains([\"dev\", \"staging\", \"prod\"], var.environment)"
        error_message = "environment must be one of: dev, staging, prod"
        target        = "var.environment"
      }
    ])
    
    # Output validations (postconditions)
    output_validations = jsonencode([
      {
        type          = "postcondition"
        condition     = "length(module.infrastructure.resource_ids) > 0"
        error_message = "Module must create at least one resource to be valid"
        target        = "output.resource_ids"
      },
      {
        type          = "postcondition"
        condition     = "module.infrastructure.deployment_status == \"healthy\""
        error_message = "Module deployment health check failed"
        target        = "output.deployment_status"
      }
    ])
    
    resources_count = 5
  }
}
```

### Meta-Argument Validation

```terraform
resource "ctfchallenge_validated_resource" "with_meta_args" {
  count = 3
  
  name           = "resource-${count.index}"
  required_value = "value-${count.index}"
  
  depends_on = [null_resource.prerequisite]
  
  lifecycle {
    create_before_destroy = true
    
    precondition {
      condition     = count.index >= 0
      error_message = "Invalid count index"
    }
  }
}

resource "ctfchallenge_flag_validator" "meta_challenge" {
  challenge_id = "meta_grandmaster"
  
  resource_proof {
    resource_type = "ctfchallenge_validated_resource"
    resource_name = "with_meta_args"
    
    attributes = {
      count_value = "3"
    }
    
    meta_arguments = {
      count      = "3"
      depends_on = "null_resource.prerequisite"
    }
    
    lifecycle_config = jsonencode({
      create_before_destroy = true
      preconditions = [
        {
          condition     = "count.index >= 0"
          error_message = "Invalid count index"
        }
      ]
    })
  }
}
```

## Schema

### Required

- `challenge_id` (String) The ID of the challenge to validate.

### Optional (Choose One)

- `proof_of_work` (Map of String) Manual proof for basic challenges. All values must be strings.

- `resource_proof` (List of Object) Proof from Terraform resources with structure validation.
  - `resource_type` (String) - Type of resource (e.g., "ctfchallenge_validated_resource")
  - `resource_name` (String) - Name/identifier of the resource
  - `attributes` (Map of String) - Resource attributes (all strings)
  - `lifecycle_config` (String) - **JSON-encoded lifecycle configuration** including preconditions/postconditions
  - `meta_arguments` (Map of String) - Meta-arguments used (count, for_each, depends_on, etc.)

- `data_source_proof` (List of Object) Proof from data sources.
  - `data_source_type` (String) - Type of data source
  - `data_source_name` (String) - Name/identifier
  - `attributes` (Map of String) - Data attributes (all strings)
  - `lifecycle_config` (String) - **JSON-encoded lifecycle configuration**

- `module_proof` (List of Object, MaxItems: 1) Proof from module configuration.
  - `module_name` (String) - Name of the module
  - `input_validations` (String) - **JSON-encoded array of input validation rules**
  - `output_validations` (String) - **JSON-encoded array of output validation rules**
  - `resources_count` (Number) - Number of resources in module

### Read-Only

- `id` (String) Unique identifier for this validation.
- `validated` (Boolean) Whether the challenge was successfully validated.
- `message` (String) Validation result message.
- `flag` (String, Sensitive) **The flag revealed upon success.**
- `points` (Number) Points awarded (0 if failed).
- `timestamp` (String) When completed (RFC3339).
- `proof_source` (String) Source of proof (manual, resources, data_sources, module).
- `validation_details` (List of String) **Detailed validation feedback** showing what passed/failed.

## Validation Details Output

When using structure-based validation, you get detailed feedback:

```terraform
output "validation_feedback" {
  value = ctfchallenge_flag_validator.challenge.validation_details
}
```

Example output for successful validation:

```
validation_details = [
  "✓ Found resource 'my_solution' with lifecycle block",
  "Checking precondition 1...",
  "  ✓ Has condition: var.min_length >= 5",
  "  ✓ Does not use 'self' (correct for precondition)",
  "  ✓ Has descriptive error message (52 chars)",
  "✓ Precondition challenge completed! Your resource properly validates inputs before creation."
]
```

Example output for failed validation:

```
validation_details = [
  "✓ Found resource 'my_solution' with lifecycle block",
  "Checking postcondition 1...",
  "  ✓ Has condition: self.validated == true",
  "  ✗ Postconditions must use 'self' to reference resource attributes",
  "  Example: self.solved == true or self.status == \"active\""
]
```

## Tips

1. **For validation challenges, use structure-based proof** - The validator inspects your lifecycle blocks
2. **JSON-encode lifecycle configurations** - Use `jsonencode()` to convert to string
3. **All attribute values must be strings** - Use `tostring()` for numbers/bools
4. **Check validation_details** - Detailed feedback shows exactly what passed/failed
5. **Test incrementally** - Start with one condition, then add more

## See Also

- [Validation Challenges Guide](../guides/validation-challenges.md)
- [Validated Resource](validated_resource.md)
- [Getting Started Guide](../guides/getting-started.md)