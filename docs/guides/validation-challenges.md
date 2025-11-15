---
page_title: "Validation Challenges Guide"
subcategory: "Guides"
description: |-
  Complete guide to Terraform precondition and postcondition challenges.
---

# Validation Challenges Guide

This guide covers all validation-related challenges focusing on Terraform's precondition and postcondition features.

## Overview

Terraform 1.2+ introduced custom validation conditions:

- **Preconditions** - Validate before resource creation/modification
- **Postconditions** - Validate after resource creation/read
- **Output Preconditions** - Validate before exposing outputs
- **Check Blocks** - Continuous validation

## Challenge List

| Challenge | Points | Difficulty | Focus |
|-----------|--------|----------|-------|
| Precondition Guardian | 150 | Intermediate | Input validation |
| Postcondition Validator | 175 | Intermediate | Output validation |
| Condition Master | 200 | Intermediate | Combined conditions |
| Data Validator | 160 | Intermediate | Data source validation |
| Output Contract | 180 | Intermediate | Module contracts |
| Validation Chain | 250 | Advanced | Interconnected validation |
| Module Contract | 300 | Advanced | Module design |
| Self-Reference Master | 190 | Intermediate | Advanced self usage |
| Conditional Validation | 220 | Advanced | Complex logic |
| Error Message Designer | 140 | Beginner | Helpful errors |

**Total:** 1,965 points

## Precondition Guardian (150 points)

### Objective
Use preconditions to validate inputs before resource creation.

### Solution

```terraform
variable "instance_size" {
  type = string
}

resource "ctfchallenge_validated_resource" "guarded" {
  name           = "precondition-example"
  required_value = var.instance_size
  
  lifecycle {
    precondition {
      condition     = var.instance_size != ""
      error_message = "instance_size cannot be empty. Provide: small, medium, or large."
    }
    
    precondition {
      condition     = contains(["small", "medium", "large"], var.instance_size)
      error_message = "Invalid instance_size '${var.instance_size}'. Must be: small, medium, or large."
    }
  }
}

resource "ctfchallenge_flag_validator" "precondition_guardian" {
  challenge_id = "precondition_guardian"
  
  proof_of_work = {
    uses_precondition    = "true"
    condition_expression = "var.instance_size != '' && contains([...], var.instance_size)"
    checks_input         = "true"
    error_message        = "instance_size cannot be empty. Provide: small, medium, or large."
    in_lifecycle_block   = "true"
    validates            = "variable"
  }
}
```

### Key Concepts
- Preconditions run BEFORE resource operations
- Cannot reference `self` (resource doesn't exist yet)
- Validate variables, data sources, other resources
- Provide clear error messages

## Postcondition Validator (175 points)

### Objective
Use postconditions with `self` to validate resource attributes after creation.

### Solution

```terraform
resource "ctfchallenge_validated_resource" "postcondition_demo" {
  name           = "postcondition-example"
  required_value = "test-value"
  
  lifecycle {
    postcondition {
      condition     = self.solved == true
      error_message = "Resource '${self.name}' failed to reach solved state. Current: ${self.state}"
    }
    
    postcondition {
      condition     = self.quality_score >= 50
      error_message = "Quality score ${self.quality_score} below minimum 50."
    }
    
    postcondition {
      condition     = self.computed_id != ""
      error_message = "Failed to generate valid computed_id."
    }
  }
}

resource "ctfchallenge_flag_validator" "postcondition_validator" {
  challenge_id = "postcondition_validator"
  
  proof_of_work = {
    uses_postcondition       = "true"
    uses_self                = "true"
    validated_attribute      = "self.solved, self.quality_score, self.computed_id"
    validates_after_creation = "true"
    error_message            = "Resource failed to reach solved state"
    in_lifecycle_block       = "true"
  }
}
```

### Key Concepts
- Postconditions run AFTER resource operations
- Must use `self` to reference the resource
- Validate computed attributes
- Ensure resource is in expected state

## Condition Master (200 points)

### Objective
Combine preconditions and postconditions in a single resource.

### Solution

```terraform
variable "min_quality_score" {
  type    = number
  default = 50
}

resource "ctfchallenge_validated_resource" "combined" {
  name           = "combined-conditions"
  required_value = "substantial-value"
  
  lifecycle {
    # Precondition: Validate BEFORE creation
    precondition {
      condition     = var.min_quality_score >= 0 && var.min_quality_score <= 100
      error_message = "min_quality_score must be 0-100, got ${var.min_quality_score}"
    }
    
    # Postcondition: Validate AFTER creation
    postcondition {
      condition     = self.quality_score >= var.min_quality_score
      error_message = "Score ${self.quality_score} < required ${var.min_quality_score}"
    }
    
    postcondition {
      condition     = self.validated == true
      error_message = "Resource failed validation checks"
    }
  }
}

resource "ctfchallenge_flag_validator" "condition_master" {
  challenge_id = "condition_master"
  
  proof_of_work = {
    uses_precondition       = "true"
    uses_postcondition      = "true"
    postcondition_uses_self = "true"
    precondition_uses_self  = "false"
    precondition_validates  = "input variables"
    postcondition_validates = "resource state and attributes"
    precondition_error      = "min_quality_score must be 0-100"
    postcondition_error     = "Score below required minimum"
  }
}
```

### Key Concepts
- Can use both pre and post conditions together
- Preconditions validate inputs
- Postconditions validate outputs
- Different error messages for different failures

## Data Validator (160 points)

### Objective
Use postconditions to validate data source outputs.

### Solution

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
    
    postcondition {
      condition     = self.name != "" && self.description != ""
      error_message = "Challenge data incomplete. Name: '${self.name}'"
    }
  }
}

resource "ctfchallenge_flag_validator" "data_validator" {
  challenge_id = "data_validator"
  
  data_source_proof {
    data_source_type = "ctfchallenge_challenge_info"
    data_source_id   = data.ctfchallenge_challenge_info.validated_data.id
    attributes = {
      name       = data.ctfchallenge_challenge_info.validated_data.name
      points     = tostring(data.ctfchallenge_challenge_info.validated_data.points)
      difficulty = data.ctfchallenge_challenge_info.validated_data.difficulty
    }
  }
  
  proof_of_work = {
    uses_data_source         = "true"
    uses_postcondition       = "true"
    uses_self                = "true"
    validated_data_attribute = "points, difficulty, name"
    validation_purpose       = "Ensure challenge data is complete and valid"
  }
}
```

### Key Concepts
- Data sources typically use postconditions
- Validate fetched data meets expectations
- Use `self` to reference data attributes
- Prevent invalid data from propagating

## Output Contract (180 points)

### Objective
Use preconditions in output blocks to enforce module contracts.

### Solution

```terraform
output "validated_resource_id" {
  description = "ID of the validated resource"
  value       = ctfchallenge_validated_resource.combined.computed_id
  
  precondition {
    condition     = ctfchallenge_validated_resource.combined.validated == true
    error_message = "Cannot output resource ID - validation failed. Check validation rules."
  }
  
  precondition {
    condition     = ctfchallenge_validated_resource.combined.state == "active"
    error_message = "Cannot output - not in active state. Current: ${ctfchallenge_validated_resource.combined.state}"
  }
}

resource "ctfchallenge_flag_validator" "output_contract" {
  challenge_id = "output_contract"
  
  proof_of_work = {
    uses_output_block       = "true"
    uses_precondition       = "true"
    validates_before_output = "resource validation state"
    enforces_module_contract = "true"
    consumer_friendly_error  = "Cannot output resource ID - validation failed"
  }
}
```

### Key Concepts
- Output preconditions validate before exposing values
- Enforce module contracts
- Provide helpful error messages to consumers
- Prevent outputting invalid/incomplete data

## Validation Chain (250 points)

### Objective
Create interconnected resources with validation dependencies.

### Solution

```terraform
resource "ctfchallenge_validated_resource" "chain_base" {
  name           = "chain-base"
  required_value = "foundation"
  
  lifecycle {
    precondition {
      condition     = length(self.name) >= 5
      error_message = "Base name must be >= 5 characters"
    }
    
    postcondition {
      condition     = self.quality_score >= 100
      error_message = "Base quality ${self.quality_score} too low"
    }
  }
}

resource "ctfchallenge_validated_resource" "chain_middle" {
  name           = "chain-middle"
  required_value = "middleware"
  
  depends_on = [ctfchallenge_validated_resource.chain_base]
  
  lifecycle {
    precondition {
      condition     = ctfchallenge_validated_resource.chain_base.validated == true
      error_message = "Cannot create - base not validated"
    }
    
    postcondition {
      condition     = self.quality_score >= 100
      error_message = "Middle quality insufficient"
    }
  }
}

resource "ctfchallenge_validated_resource" "chain_final" {
  name           = "chain-final"
  required_value = "completion"
  
  depends_on = [
    ctfchallenge_validated_resource.chain_base,
    ctfchallenge_validated_resource.chain_middle
  ]
  
  lifecycle {
    precondition {
      condition = (
        ctfchallenge_validated_resource.chain_base.validated == true &&
        ctfchallenge_validated_resource.chain_middle.validated == true
      )
      error_message = "Cannot create - previous chain not validated"
    }
    
    postcondition {
      condition     = self.state == "active"
      error_message = "Failed to reach active state"
    }
  }
}

resource "ctfchallenge_flag_validator" "validation_chain" {
  challenge_id = "validation_chain"
  
  proof_of_work = {
    resource_count            = "3"
    total_conditions          = "7"
    conditions_interconnected = "true"
    validation_flow           = "base -> middle -> final with cascading validation"
    uses_depends_on           = "true"
  }
}
```

## Best Practices

### Writing Good Error Messages

```terraform
# ❌ Bad - Not helpful
error_message = "Invalid value"

# ✅ Good - Specific and actionable
error_message = <<-EOT
  Invalid instance_size: '${var.instance_size}'
  
  Allowed values: small, medium, large
  
  Example: instance_size = "medium"
EOT
```

### Validation Timing

```terraform
# Precondition - Checks BEFORE operation
lifecycle {
  precondition {
    condition     = var.count >= 1
    error_message = "count must be >= 1"
  }
}

# Postcondition - Checks AFTER operation
lifecycle {
  postcondition {
    condition     = self.id != ""
    error_message = "Failed to create resource"
  }
}
```

### Complex Boolean Logic

```terraform
lifecycle {
  precondition {
    condition = (
      var.environment == "prod" 
        ? var.monitoring == true && var.backup == true
        : true
    )
    error_message = "Production requires monitoring AND backup"
  }
}
```

## Common Patterns

### Environment-Specific Validation

```terraform
lifecycle {
  precondition {
    condition = (
      var.environment != "prod" || 
      (var.redundancy >= 3 && var.monitoring == true)
    )
    error_message = "Production requires redundancy >= 3 and monitoring"
  }
}
```

### Range Validation

```terraform
lifecycle {
  precondition {
    condition     = var.port >= 1024 && var.port <= 65535
    error_message = "Port must be 1024-65535, got ${var.port}"
  }
}
```

### List/Set Validation

```terraform
lifecycle {
  precondition {
    condition     = length(var.allowed_ips) > 0
    error_message = "Must specify at least one allowed IP"
  }
  
  precondition {
    condition = alltrue([
      for ip in var.allowed_ips :
      can(regex("^\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}$", ip))
    ])
    error_message = "All IPs must be valid IPv4 addresses"
  }
}
```

## See Also

- [Terraform Custom Conditions](https://www.terraform.io/language/expressions/custom-conditions)
- [Validated Resource](../resources/validated_resource.md)
- [Flag Validator Resource](../resources/flag_validator.md)