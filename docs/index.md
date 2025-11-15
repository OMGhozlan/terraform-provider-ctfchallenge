---
page_title: "CTF Challenge Provider"
subcategory: ""
description: |-
  A Capture The Flag (CTF) game implemented as a Terraform provider. Learn Terraform concepts by solving interactive challenges.
---

# CTF Challenge Provider

The CTF Challenge provider is an educational tool that gamifies learning Terraform. It provides a comprehensive series of challenges that teach Terraform concepts through hands-on practice, from basic resource management to advanced validation patterns.

## Features

- 🎯 **30+ Progressive Challenges** - From beginner to advanced Terraform concepts
- 💡 **Multi-Category Learning** - Meta-arguments, validation, modules, and more
- 🏆 **Point System** - Earn up to 5,700+ points total
- 🔐 **Flag Capture** - Complete challenges to reveal flags as rewards
- 📚 **Educational** - Learn dependencies, expressions, modules, state management, pre/postconditions, and lifecycle rules
- 🔗 **Structure-Based Validation** - Validator inspects actual Terraform resource structures, lifecycle configurations, and meta-arguments

## Example Usage

```terraform
terraform {
  required_providers {
    ctfchallenge = {
      source  = "omghozlan/ctfchallenge"
      version = "~> 1.0"
    }
  }
}

provider "ctfchallenge" {
  player_name = "your-name"
}

# List all available challenges
data "ctfchallenge_list" "all" {}

output "available_challenges" {
  value = data.ctfchallenge_list.all.challenges
}

# Solve a validation challenge with actual resource structure
resource "ctfchallenge_validated_resource" "my_solution" {
  name           = "validated-resource"
  required_value = "substantial-value"
  
  lifecycle {
    precondition {
      condition     = length(var.required_value) >= 5
      error_message = "required_value must be at least 5 characters for security purposes"
    }
    
    postcondition {
      condition     = self.validated == true
      error_message = "Resource '${self.name}' failed validation checks. State: ${self.state}"
    }
    
    postcondition {
      condition     = self.quality_score >= 50
      error_message = "Quality score ${self.quality_score} is below minimum threshold of 50"
    }
  }
}

# Submit the resource structure for validation
resource "ctfchallenge_flag_validator" "challenge" {
  challenge_id = "precondition_guardian"
  
  resource_proof {
    resource_type = "ctfchallenge_validated_resource"
    resource_name = "my_solution"
    
    attributes = {
      name      = ctfchallenge_validated_resource.my_solution.name
      validated = tostring(ctfchallenge_validated_resource.my_solution.validated)
      solved    = tostring(ctfchallenge_validated_resource.my_solution.solved)
    }
    
    # The validator inspects this structure
    lifecycle_config = jsonencode({
      preconditions = [
        {
          condition     = "length(var.required_value) >= 5"
          error_message = "required_value must be at least 5 characters for security purposes"
        }
      ]
      postconditions = [
        {
          condition     = "self.validated == true"
          error_message = "Resource failed validation checks"
        },
        {
          condition     = "self.quality_score >= 50"
          error_message = "Quality score below minimum threshold"
        }
      ]
    })
  }
}

# Capture the flag!
output "flag" {
  value     = ctfchallenge_flag_validator.challenge.flag
  sensitive = true
}

output "validation_details" {
  value = ctfchallenge_flag_validator.challenge.validation_details
}
```

## Challenge Categories

### Fundamentals (450 points)
- **Terraform Basics** (100 points) - Resource dependencies
- **State Secrets** (200 points) - State management
- **Data Source Detective** (150 points) - Data source queries

### Meta-Arguments (1,575 points)
- **Count Master** (150 points) - Master the count meta-argument
- **For Each Wizard** (200 points) - Use for_each effectively
- **Dependency Chain** (175 points) - Explicit depends_on usage
- **Lifecycle Expert** (225 points) - Lifecycle rules mastery
- **Meta Grandmaster** (300 points) - Combine all meta-arguments
- **Dynamic Blocks** (180 points) - Dynamic block generation
- **Locals + Count Combo** (160 points) - Combine locals with count
- **Conditional Resources** (140 points) - Conditional creation patterns

### Validation (2,025 points)
- **Precondition Guardian** (150 points) - Input validation with preconditions
- **Postcondition Validator** (175 points) - Output validation with postconditions
- **Condition Master** (200 points) - Combined pre/postconditions
- **Data Validator** (160 points) - Data source validation
- **Output Contract** (180 points) - Module contract enforcement
- **Validation Chain** (250 points) - Interconnected validations
- **Module Contract** (300 points) - Comprehensive module design
- **Self-Reference Master** (190 points) - Advanced self usage
- **Conditional Validation** (220 points) - Complex boolean logic
- **Error Message Designer** (140 points) - Helpful error messages

### Advanced (1,150 points)
- **Expression Expert** (350 points) - Functions and expressions
- **Module Master** (400 points) - Module composition
- **Cryptographic Compute** (500 points) - Cryptographic functions

**Total Points Available:** 5,200+

## Structure-Based Validation

The provider now validates challenges by inspecting actual Terraform resource structures:

### Resource Structure Validation

```terraform
resource "ctfchallenge_flag_validator" "challenge" {
  challenge_id = "postcondition_validator"
  
  resource_proof {
    resource_type = "ctfchallenge_validated_resource"
    resource_name = "my_solution"
    
    attributes = {
      validated = "true"
      solved    = "true"
    }
    
    # Lifecycle structure is validated
    lifecycle_config = jsonencode({
      postconditions = [
        {
          condition     = "self.validated == true"
          error_message = "Validation failed"
        }
      ]
    })
  }
}
```

The validator checks:
- ✅ Presence of lifecycle blocks
- ✅ Correct use of preconditions (no `self` reference)
- ✅ Correct use of postconditions (requires `self` reference)
- ✅ Error message quality and length
- ✅ Complexity of validation logic
- ✅ Proper use of meta-arguments

### Data Source Structure Validation

```terraform
data "ctfchallenge_challenge_info" "validated" {
  challenge_id = "data_validator"
  
  lifecycle {
    postcondition {
      condition     = self.points > 0
      error_message = "Challenge has invalid points: ${self.points}"
    }
  }
}

resource "ctfchallenge_flag_validator" "data_challenge" {
  challenge_id = "data_validator"
  
  data_source_proof {
    data_source_type = "ctfchallenge_challenge_info"
    data_source_name = "validated"
    
    attributes = {
      points = tostring(data.ctfchallenge_challenge_info.validated.points)
    }
    
    lifecycle_config = jsonencode({
      postconditions = [
        {
          condition     = "self.points > 0"
          error_message = "Challenge has invalid points"
        }
      ]
    })
  }
}
```

### Module Structure Validation

```terraform
resource "ctfchallenge_flag_validator" "module_challenge" {
  challenge_id = "module_contract"
  
  module_proof {
    module_name = "my_module"
    
    input_validations = jsonencode([
      {
        type          = "precondition"
        condition     = "var.instance_count >= 1"
        error_message = "instance_count must be at least 1"
        target        = "var.instance_count"
      },
      {
        type          = "precondition"
        condition     = "contains([\"dev\", \"prod\"], var.environment)"
        error_message = "environment must be dev or prod"
        target        = "var.environment"
      }
    ])
    
    output_validations = jsonencode([
      {
        type          = "postcondition"
        condition     = "length(module.my_module.resource_ids) > 0"
        error_message = "Module must create at least one resource"
        target        = "output.resource_ids"
      }
    ])
    
    resources_count = 5
  }
}
```

## Detailed Validation Feedback

When validation fails, you receive detailed feedback:

```
Error: Challenge validation failed

Checking precondition 1...
  ✗ Precondition uses 'self' incorrectly
  Preconditions should NOT use 'self' - the resource doesn't exist yet
  Hint: Use var.* or local.* to validate inputs before creation

Validation details:
  1. Found resource 'my_solution' with lifecycle block
  2. Precondition has empty condition expression
  3. Error message must be at least 10 characters and descriptive
```

## Schema

### Optional

- `player_name` (String) Your player name for the CTF. Can also be set via the `TF_CTF_PLAYER` environment variable. Defaults to `"anonymous"`.
- `api_endpoint` (String) Optional API endpoint for score tracking. Can also be set via the `TF_CTF_API` environment variable.

## Getting Started

See the [Getting Started Guide](guides/getting-started.md) for a step-by-step walkthrough of your first challenge.

## Resources

- [ctfchallenge_flag_validator](resources/flag_validator.md) - Validate challenge solutions and capture flags
- [ctfchallenge_puzzle_box](resources/puzzle_box.md) - Solve logic puzzles for bonus flags
- [ctfchallenge_meta_challenge](resources/meta_challenge.md) - Meta-argument focused challenges
- [ctfchallenge_validated_resource](resources/validated_resource.md) - Resource with validation support

## Data Sources

- [ctfchallenge_hint](data-sources/hint.md) - Get hints for challenges
- [ctfchallenge_list](data-sources/list.md) - List all available challenges
- [ctfchallenge_challenge_info](data-sources/challenge_info.md) - Get detailed challenge information
- [ctfchallenge_validation_helper](data-sources/validation_helper.md) - Validation assistance

## Learning Paths

### Beginner Path (490 points)
1. Terraform Basics
2. State Secrets
3. Data Source Detective
4. Conditional Resources
5. Error Message Designer

### Intermediate Path (1,915 points)
1. Count Master
2. For Each Wizard
3. Locals + Count Combo
4. Dynamic Blocks
5. Precondition Guardian
6. Postcondition Validator
7. Condition Master
8. Data Validator
9. Output Contract
10. Self-Reference Master

### Advanced Path (2,895 points)
1. Expression Expert
2. Module Master
3. Dependency Chain
4. Lifecycle Expert
5. Meta Grandmaster
6. Validation Chain
7. Module Contract
8. Conditional Validation
9. Cryptographic Compute

## How CTF Challenges Work

In traditional CTF (Capture The Flag) competitions, you complete a challenge and receive a flag as proof of completion. This provider follows that paradigm:

1. **Read the challenge description** - Understand what you need to accomplish
2. **Build your solution** - Write Terraform code that meets the requirements
3. **Submit resource structure** - The validator inspects your actual Terraform configurations
4. **Capture the flag** - If successful, the flag is revealed as your reward!

The flag format is: `flag{some_text_here}`

## Viewing Captured Flags

Flags are marked as sensitive outputs. To view them:

```bash
terraform output -raw flag
```

Good luck, and happy flag hunting! 🚀