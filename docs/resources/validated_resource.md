page_title: "ctfchallenge_validated_resource Resource - ctfchallenge"
subcategory: ""
description: |-
  A resource with built-in validation support for testing pre/postconditions.
---

# ctfchallenge_validated_resource (Resource)

The `validated_resource` provides a resource specifically designed for practicing Terraform's precondition and postcondition features. It exposes computed attributes that can be validated with `self` references.

## Example Usage

### Basic Usage

```terraform
resource "ctfchallenge_validated_resource" "example" {
  name           = "my-resource"
  required_value = "important-data"
}

output "resource_state" {
  value = {
    validated = ctfchallenge_validated_resource.example.validated
    solved    = ctfchallenge_validated_resource.example.solved
    score     = ctfchallenge_validated_resource.example.quality_score
  }
}
```

### With Preconditions

```terraform
variable "instance_size" {
  type = string
}

resource "ctfchallenge_validated_resource" "with_precondition" {
  name           = "validated"
  required_value = var.instance_size
  
  lifecycle {
    precondition {
      condition     = var.instance_size != ""
      error_message = "instance_size cannot be empty"
    }
    
    precondition {
      condition     = contains(["small", "medium", "large"], var.instance_size)
      error_message = "instance_size must be small, medium, or large"
    }
  }
}
```

### With Postconditions

```terraform
resource "ctfchallenge_validated_resource" "with_postcondition" {
  name           = "validated"
  required_value = "substantial-value"
  
  lifecycle {
    postcondition {
      condition     = self.validated == true
      error_message = "Resource failed validation: ${self.state}"
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
```

### Combined Pre/Postconditions

```terraform
variable "min_score" {
  type    = number
  default = 50
}

resource "ctfchallenge_validated_resource" "combined" {
  name           = "combined-validation"
  required_value = "good-value"
  
  validation_rules {
    has_precondition  = true
    has_postcondition = true
    validates_input   = true
    validates_output  = true
  }
  
  lifecycle {
    # Precondition: Validate BEFORE creation
    precondition {
      condition     = var.min_score >= 0 && var.min_score <= 100
      error_message = "min_score must be 0-100, got ${var.min_score}"
    }
    
    # Postcondition: Validate AFTER creation
    postcondition {
      condition     = self.quality_score >= var.min_score
      error_message = "Score ${self.quality_score} < required ${var.min_score}"
    }
    
    postcondition {
      condition     = self.validated == true
      error_message = "Resource validation failed"
    }
  }
}
```

## Schema

### Required

- `name` (String) Resource name.
- `required_value` (String) A required value that will be validated.

### Optional

- `optional_value` (String) An optional value.
- `validation_rules` (List of Object) Validation rules for this resource.
  - `has_precondition` (Boolean) Whether preconditions are defined.
  - `has_postcondition` (Boolean) Whether postconditions are defined.
  - `validates_input` (Boolean) Whether input is validated.
  - `validates_output` (Boolean) Whether output is validated.

### Read-Only

- `id` (String) The resource identifier.
- `state` (String) Resource state after creation.
- `validated` (Boolean) Whether validation passed.
- `validation_timestamp` (String) When validation occurred.
- `computed_id` (String) Computed identifier.
- `solved` (Boolean) Whether resource is in solved state.
- `quality_score` (Number) Quality score of the resource.

## Attributes for Validation

These computed attributes are designed to be validated in postconditions:

- `validated` - Boolean indicating validation success
- `solved` - Boolean indicating solved state
- `quality_score` - Numeric quality metric (0-1000+)
- `state` - String state value
- `computed_id` - String identifier

## Using with Flag Validator

```terraform
resource "ctfchallenge_validated_resource" "solution" {
  name           = "my-solution"
  required_value = "answer"
  
  lifecycle {
    postcondition {
      condition     = self.validated == true
      error_message = "Solution not validated"
    }
  }
}

resource "ctfchallenge_flag_validator" "challenge" {
  challenge_id = "postcondition_validator"
  
  resource_proof {
    resource_type = "ctfchallenge_validated_resource"
    resource_id   = ctfchallenge_validated_resource.solution.id
    attributes = {
      validated = tostring(ctfchallenge_validated_resource.solution.validated)
      solved    = tostring(ctfchallenge_validated_resource.solution.solved)
    }
  }
}
```

## See Also

- [Validation Challenges Guide](../guides/validation-challenges.md)
- [Flag Validator Resource](flag_validator.md)
- [Terraform Preconditions/Postconditions](https://www.terraform.io/language/expressions/custom-conditions)