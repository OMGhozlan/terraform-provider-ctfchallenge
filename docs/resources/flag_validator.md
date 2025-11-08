---
page_title: "ctfchallenge_flag_validator Resource - ctfchallenge"
subcategory: ""
description: |-
  Validates challenge solutions and awards points for correct flags.
---

# ctfchallenge_flag_validator (Resource)

The `flag_validator` resource is used to submit your solution to a challenge. You provide the challenge ID, the flag you discovered, and proof that you completed the required work.

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
  flag         = "flag{t3rr4f0rm_d3p3nd3nc13s}"

  proof_of_work = {
    dependencies = "${null_resource.first.id},${null_resource.second.id},${null_resource.third.id}"
  }
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
  flag         = "flag{3xpr3ss10ns_unl0ck3d}"

  proof_of_work = {
    computed_value = local.encoded
  }
}
```

## Schema

### Required

- `challenge_id` (String) The ID of the challenge to validate. Must be one of the available challenge IDs.
- `flag` (String, Sensitive) The flag you discovered for this challenge. Format: `flag{...}`

### Optional

- `proof_of_work` (Map of String) Evidence that you completed the challenge requirements. The specific keys required depend on the challenge. All values must be strings.

### Read-Only

- `id` (String) The unique identifier for this validation attempt.
- `validated` (Boolean) Whether the challenge was successfully validated.
- `message` (String) A message describing the validation result.
- `points` (Number) Points awarded for successfully completing this challenge.
- `timestamp` (String) RFC3339 timestamp when the challenge was completed.

## Proof of Work Requirements by Challenge

### terraform_basics
```hcl
proof_of_work = {
  dependencies = "id1,id2,id3"  # Comma-separated list of resource IDs
}
```

### expression_expert
```hcl
proof_of_work = {
  computed_value = "base64(sha256(...))"  # Result of expression computation
}
```

### state_secrets
```hcl
proof_of_work = {
  resource_count = "42"  # The magic number
}
```

### module_master
```hcl
proof_of_work = {
  module_output = "module.name.output"  # Module output reference
}
```

### dynamic_blocks
```hcl
proof_of_work = {
  dynamic_block_count = "5"  # Number of dynamic blocks generated
}
```

### for_each_wizard
```hcl
proof_of_work = {
  items = "alpha,beta,gamma,delta"  # Items used in for_each
}
```

### data_source_detective
```hcl
proof_of_work = {
  filtered_count = "7"  # Number of filtered items
}
```

### cryptographic_compute
```hcl
proof_of_work = {
  crypto_hash = "abc123..."  # Result of cryptographic computation
}
```

## Import

Flag validator resources are ephemeral and cannot be imported.