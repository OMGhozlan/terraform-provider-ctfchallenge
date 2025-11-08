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

## Example - Building a Progress Tracker

```terraform
# Get info for all challenges you're working on
data "ctfchallenge_challenge_info" "basics" {
  challenge_id = "terraform_basics"
}

data "ctfchallenge_challenge_info" "expressions" {
  challenge_id = "expression_expert"
}

data "ctfchallenge_challenge_info" "state" {
  challenge_id = "state_secrets"
}

locals {
  challenges = [
    data.ctfchallenge_challenge_info.basics,
    data.ctfchallenge_challenge_info.expressions,
    data.ctfchallenge_challenge_info.state,
  ]
  
  total_possible_points = sum([for c in local.challenges : c.points])
}

output "challenge_tracker" {
  value = {
    challenges      = [for c in local.challenges : c.name]
    total_points    = local.total_possible_points
    beginner_count  = length([for c in local.challenges : c if c.difficulty == "beginner"])
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

- `terraform_basics` - Beginner (100 points)
- `expression_expert` - Intermediate (350 points)
- `state_secrets` - Beginner (200 points)
- `module_master` - Advanced (400 points)
- `dynamic_blocks` - Intermediate (300 points)
- `for_each_wizard` - Intermediate (250 points)
- `data_source_detective` - Beginner (150 points)
- `cryptographic_compute` - Advanced (500 points)

## Categories

- **fundamentals** - Basic Terraform concepts and resource management
- **expressions** - Terraform expression syntax and functions
- **state** - State management and understanding
- **modules** - Module creation and composition
- **advanced-syntax** - Dynamic blocks and meta-arguments
- **loops** - for_each and count
- **data-sources** - Querying and filtering data
- **functions** - Built-in function usage
```

## `docs/data-sources/hint.md`

```markdown
---
page_title: "ctfchallenge_hint Data Source - ctfchallenge"
subcategory: ""
description: |-
  Retrieves hints for challenges at the cost of points.
---

# ctfchallenge_hint (Data Source)

The `hint` data source provides hints for challenges. Each hint level costs points (10, 20, or 30 points depending on the level). Use hints strategically to help you capture flags without getting completely stuck!

## Example Usage

```terraform
# Get the first hint (level 0)
data "ctfchallenge_hint" "basic_hint" {
  challenge_id = "terraform_basics"
  level        = 0
}

output "hint_text" {
  value = data.ctfchallenge_hint.basic_hint.hint
}

output "hint_cost" {
  value = data.ctfchallenge_hint.basic_hint.cost
}
```

## Example with Multiple Hint Levels

```terraform
# Progressive hints for the Expression Expert challenge
data "ctfchallenge_hint" "hint_level_0" {
  challenge_id = "expression_expert"
  level        = 0
}

data "ctfchallenge_hint" "hint_level_1" {
  challenge_id = "expression_expert"
  level        = 1
}

data "ctfchallenge_hint" "hint_level_2" {
  challenge_id = "expression_expert"
  level        = 2
}

output "all_hints" {
  value = {
    hint_1 = {
      text = data.ctfchallenge_hint.hint_level_0.hint
      cost = data.ctfchallenge_hint.hint_level_0.cost
    }
    hint_2 = {
      text = data.ctfchallenge_hint.hint_level_1.hint
      cost = data.ctfchallenge_hint.hint_level_1.cost
    }
    hint_3 = {
      text = data.ctfchallenge_hint.hint_level_2.hint
      cost = data.ctfchallenge_hint.hint_level_2.cost
    }
  }
}
```

## Example - Conditional Hint Usage

```terraform
variable "need_help" {
  type        = bool
  default     = false
  description = "Set to true if you want to see a hint"
}

data "ctfchallenge_hint" "conditional_hint" {
  count = var.need_help ? 1 : 0
  
  challenge_id = "cryptographic_compute"
  level        = 0
}

output "hint" {
  value = var.need_help ? data.ctfchallenge_hint.conditional_hint[0].hint : "No hint requested"
}
```

## Schema

### Required

- `challenge_id` (String) The ID of the challenge to get a hint for.

### Optional

- `level` (Number) The hint level (0-2). Higher levels provide more detailed hints but cost more points. Defaults to `0`.

### Read-Only

- `id` (String) The unique identifier for this hint.
- `hint` (String) The hint text.
- `cost` (Number) The point penalty for requesting this hint.

## Hint Levels

- **Level 0** (10 points): General direction or concept
- **Level 1** (20 points): More specific guidance
- **Level 2** (30 points): Near-complete solution

## Point Strategy

When deciding whether to use hints:

- **Challenge Points - Hint Costs = Net Points**
- Example: Expression Expert (350 points) - Level 2 hint (30 points) = 320 net points
- Sometimes it's better to use a hint and move forward than to get stuck!

## Example Hint Progression

For the "Expression Expert" challenge:

- **Level 0**: "Look at Terraform's hash and encoding functions"
- **Level 1**: "Combine sha256() and base64encode() functions"
- **Level 2**: "Concatenate the strings: 'terraform' + 'expressions' + 'rock', hash with sha256, then base64encode"

For the "Cryptographic Compute" challenge:

- **Level 0**: "Chain multiple hash functions"
- **Level 1**: "Start with sha256, then md5 the result"
- **Level 2**: "Compute: md5(sha256('terraform_ctf_11_2025'))"

## Tips

1. **Start with Level 0** - Often the gentle nudge is all you need
2. **Google is your friend** - Hints combined with Terraform documentation are powerful
3. **Use terraform console** - Test expressions before committing to your solution
4. **Track your point total** - Keep a running tally of earned points minus hint costs

## Warning

When you request a hint, Terraform will show a warning message indicating the point cost:

```
Warning: Hint requested (-10 points)

Hint level 0 for challenge 'terraform_basics'
```

This is normal behavior and helps you track your hint usage.