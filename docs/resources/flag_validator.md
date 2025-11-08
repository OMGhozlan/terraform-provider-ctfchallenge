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

## Example with Expressions Challenge

```terraform
locals {
  combined = "terraformexpressionsrock"
  hashed   = sha256(local.combined)
  encoded  = base64encode(local.hashed)
}

resource "ctfchallenge_flag_validator" "expressions" {
  challenge_id = "expression_expert"

  proof_of_work = {
    computed_value = local.encoded
  }
}

output "flag" {
  value     = ctfchallenge_flag_validator.expressions.flag
  sensitive = true
}

output "points" {
  value = ctfchallenge_flag_validator.expressions.points
}
```

## Example with Multiple Challenges

```terraform
# Challenge 1
resource "ctfchallenge_flag_validator" "basics" {
  challenge_id = "terraform_basics"
  proof_of_work = {
    dependencies = "${null_resource.first.id},${null_resource.second.id},${null_resource.third.id}"
  }
}

# Challenge 2
resource "ctfchallenge_flag_validator" "state" {
  challenge_id = "state_secrets"
  proof_of_work = {
    resource_count = "42"
  }
}

# Track all captured flags
output "all_flags" {
  value = {
    basics = ctfchallenge_flag_validator.basics.flag
    state  = ctfchallenge_flag_validator.state.flag
  }
  sensitive = true
}

output "total_points" {
  value = (
    ctfchallenge_flag_validator.basics.points +
    ctfchallenge_flag_validator.state.points
  )
}
```

## Viewing Your Captured Flag

After successful validation, view your flag with:

```bash
terraform output -raw captured_flag
```

Or for multiple flags:

```bash
terraform output -json all_flags | jq -r
```

## Schema

### Required

- `challenge_id` (String) The ID of the challenge to validate. Must be one of the available challenge IDs.
- `proof_of_work` (Map of String) Evidence that you completed the challenge requirements. The specific keys required depend on the challenge. All values must be strings.

### Read-Only

- `id` (String) The unique identifier for this validation attempt.
- `validated` (Boolean) Whether the challenge was successfully validated.
- `message` (String) A message describing the validation result.
- `flag` (String, Sensitive) **The flag revealed upon successful completion.** This is your reward! Empty if validation failed.
- `points` (Number) Points awarded for successfully completing this challenge. Zero if validation failed.
- `timestamp` (String) RFC3339 timestamp when the challenge was completed. Empty if validation failed.

## Proof of Work Requirements by Challenge

### terraform_basics (100 points)

Create at least 3 resources with dependencies.

```hcl
proof_of_work = {
  dependencies = "id1,id2,id3"  # Comma-separated list of resource IDs
}
```

**Example:**
```terraform
resource "null_resource" "first" {}
resource "null_resource" "second" { depends_on = [null_resource.first] }
resource "null_resource" "third" { depends_on = [null_resource.second] }

resource "ctfchallenge_flag_validator" "basics" {
  challenge_id = "terraform_basics"
  
  proof_of_work = {
    dependencies = "${null_resource.first.id},${null_resource.second.id},${null_resource.third.id}"
  }
}
```

### expression_expert (350 points)

Compute: `base64encode(sha256("terraform" + "expressions" + "rock"))`

```hcl
proof_of_work = {
  computed_value = "base64(sha256(...))"  # Result of expression computation
}
```

**Example:**
```terraform
locals {
  result = base64encode(sha256("terraformexpressionsrock"))
}

resource "ctfchallenge_flag_validator" "expressions" {
  challenge_id = "expression_expert"
  
  proof_of_work = {
    computed_value = local.result
  }
}
```

### state_secrets (200 points)

Find the magic number (hint: Douglas Adams).

```hcl
proof_of_work = {
  resource_count = "42"  # The answer to everything
}
```

**Example:**
```terraform
resource "ctfchallenge_flag_validator" "state" {
  challenge_id = "state_secrets"
  
  proof_of_work = {
    resource_count = "42"
  }
}
```

### module_master (400 points)

Create and use modules with outputs.

```hcl
proof_of_work = {
  module_output = "module.name.output"  # Module output reference
}
```

**Example:**
```terraform
module "example" {
  source = "./modules/example"
}

resource "ctfchallenge_flag_validator" "modules" {
  challenge_id = "module_master"
  
  proof_of_work = {
    module_output = "module.example with composition"
  }
}
```

### dynamic_blocks (300 points)

Generate at least 5 dynamic blocks.

```hcl
proof_of_work = {
  dynamic_block_count = "5"  # Number of dynamic blocks generated
}
```

**Example:**
```terraform
locals {
  blocks = [1, 2, 3, 4, 5]
}

resource "ctfchallenge_flag_validator" "dynamic" {
  challenge_id = "dynamic_blocks"
  
  proof_of_work = {
    dynamic_block_count = tostring(length(local.blocks))
  }
}
```

### for_each_wizard (250 points)

Use for_each with specific items: alpha, beta, gamma, delta.

```hcl
proof_of_work = {
  items = "alpha,beta,gamma,delta"  # Items used in for_each
}
```

**Example:**
```terraform
locals {
  items = toset(["alpha", "beta", "gamma", "delta"])
}

resource "null_resource" "foreach" {
  for_each = local.items
  triggers = { name = each.key }
}

resource "ctfchallenge_flag_validator" "foreach" {
  challenge_id = "for_each_wizard"
  
  proof_of_work = {
    items = join(",", local.items)
  }
}
```

### data_source_detective (150 points)

Filter data to find the magic count: 7.

```hcl
proof_of_work = {
  filtered_count = "7"  # Number of filtered items
}
```

**Example:**
```terraform
resource "ctfchallenge_flag_validator" "datasource" {
  challenge_id = "data_source_detective"
  
  proof_of_work = {
    filtered_count = "7"
  }
}
```

### cryptographic_compute (500 points)

Compute: `md5(sha256("terraform_ctf_11_2025"))`

```hcl
proof_of_work = {
  crypto_hash = "abc123..."  # Result of cryptographic computation
}
```

**Example:**
```terraform
locals {
  secret = "terraform_ctf_11_2025"
  result = md5(sha256(local.secret))
}

resource "ctfchallenge_flag_validator" "crypto" {
  challenge_id = "cryptographic_compute"
  
  proof_of_work = {
    crypto_hash = local.result
  }
}
```

## Validation Results

### Successful Validation

```hcl
validated = true
message   = "🎉 Congratulations! You solved 'Challenge Name' and earned X points!"
flag      = "flag{...}"  # The captured flag!
points    = X
timestamp = "2025-01-15T10:30:00Z"
```

### Failed Validation

```hcl
validated = false
message   = "❌ Challenge failed: <error details>"
flag      = ""
points    = 0
timestamp = ""
```

## Common Errors

### Missing Required Keys

```
Error: Challenge validation failed
│ Challenge failed: provide 'dependencies' in proof_of_work
```

**Solution:** Ensure your proof_of_work includes all required keys for the challenge.

### Wrong Data Type

```
Error: Incorrect value type
│ Inappropriate value for attribute "proof_of_work": all values must be strings
```

**Solution:** Convert all values to strings using `tostring()` or string interpolation.

### Incorrect Value

```
Error: Challenge validation failed
│ Challenge failed: create at least 3 dependent resources (found 2)
```

**Solution:** Review the challenge requirements and adjust your solution.

## Tips

1. **Test your expressions** in `terraform console` before submitting
2. **All proof_of_work values must be strings** - use `tostring()` if needed
3. **Check for typos** in challenge_id and proof_of_work keys
4. **Read error messages carefully** - they often tell you exactly what's missing
5. **The flag is sensitive** - use `terraform output -raw flag` to view it

## Import

Flag validator resources are ephemeral and cannot be imported. Each validation creates a new resource instance.

## See Also

- [Getting Started Guide](../guides/getting-started.md)
- [Challenge Walkthrough](../guides/challenge-walkthrough.md)
- [All Challenges List](../data-sources/list.md)
