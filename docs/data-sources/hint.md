---
page_title: "ctfchallenge_hint Data Source - ctfchallenge"
subcategory: ""
description: |-
  Retrieves hints for challenges at the cost of points.
---

# ctfchallenge_hint (Data Source)

The `hint` data source provides hints for challenges. Each hint level costs points (10, 20, or 30 points depending on the level).

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
# Progressive hints
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
    hint_1 = data.ctfchallenge_hint.hint_level_0.hint
    hint_2 = data.ctfchallenge_hint.hint_level_1.hint
    hint_3 = data.ctfchallenge_hint.hint_level_2.hint
  }
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

## Example Hint Progression

For the "Expression Expert" challenge:

- **Level 0**: "Look at Terraform's hash and encoding functions"
- **Level 1**: "Combine sha256() and base64encode() functions"
- **Level 2**: "Concatenate the strings: 'terraform' + 'expressions' + 'rock', hash with sha256, then base64encode"