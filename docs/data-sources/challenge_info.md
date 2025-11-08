---
page_title: "ctfchallenge_challenge_info Data Source - ctfchallenge"
subcategory: ""
description: |-
  Retrieves detailed information about a specific challenge.
---

# ctfchallenge_challenge_info (Data Source)

The `challenge_info` data source provides detailed information about a specific challenge without requesting hints or attempting validation.

## Example Usage

```terraform
data "ctfchallenge_challenge_info" "basics" {
  challenge_id = "terraform_basics"
}

output "challenge_details" {
  value = {
    name        = data.ctfchallenge_challenge_info.basics.name
    description = data.ctfchallenge_challenge_info.basics.description
    points      = data.ctfchallenge_challenge_info.basics.points
    difficulty  = data.ctfchallenge_challenge_info.basics.difficulty
    category    = data.ctfchallenge_challenge_info.basics.category
  }
}
```

## Example with Multiple Challenges

```terraform
locals {
  challenge_ids = [
    "terraform_basics",
    "expression_expert",
    "state_secrets"
  ]
}

data "ctfchallenge_challenge_info" "challenges" {
  for_each = toset(local.challenge_ids)
  
  challenge_id = each.key
}

output "challenge_overview" {
  value = {
    for id, info in data.ctfchallenge_challenge_info.challenges :
    id => {
      name       = info.name
      points     = info.points
      difficulty = info.difficulty
    }
  }
}
```

## Schema

### Required

- `challenge_id` (String) The ID of the challenge to get information about.

### Read-Only

- `id` (String) The challenge ID (same as input).
- `name` (String) The human-readable name of the challenge.
- `description` (String) A detailed description of what the challenge teaches.
- `points` (Number) Points awarded for completing this challenge.
- `difficulty` (String) The difficulty level (`beginner`, `intermediate`, or `advanced`).
- `category` (String) The category this challenge belongs to.

## Valid Challenge IDs

- `terraform_basics`
- `expression_expert`
- `state_secrets`
- `module_master`
- `dynamic_blocks`
- `for_each_wizard`
- `data_source_detective`
- `cryptographic_compute`