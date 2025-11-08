---
page_title: "ctfchallenge_list Data Source - ctfchallenge"
subcategory: ""
description: |-
  Lists all available challenges with optional filtering.
---

# ctfchallenge_list (Data Source)

The `list` data source retrieves all available challenges, optionally filtered by difficulty or category.

## Example Usage

```terraform
# List all challenges
data "ctfchallenge_list" "all" {}

output "all_challenges" {
  value = data.ctfchallenge_list.all.challenges
}

output "total_points" {
  value = data.ctfchallenge_list.all.total_points
}
```

## Example with Filtering

```terraform
# List only beginner challenges
data "ctfchallenge_list" "beginner" {
  difficulty = "beginner"
}

# List challenges in a specific category
data "ctfchallenge_list" "expressions" {
  category = "expressions"
}

# List advanced challenges
data "ctfchallenge_list" "advanced" {
  difficulty = "advanced"
}

output "beginner_challenges" {
  value = data.ctfchallenge_list.beginner.challenges
}
```

## Example Creating a Dashboard

```terraform
data "ctfchallenge_list" "all" {}

locals {
  challenges_by_difficulty = {
    for challenge in data.ctfchallenge_list.all.challenges :
    challenge.difficulty => challenge...
  }
}

output "challenge_dashboard" {
  value = {
    total_challenges = length(data.ctfchallenge_list.all.challenges)
    total_points     = data.ctfchallenge_list.all.total_points
    by_difficulty    = {
      beginner     = length([for c in data.ctfchallenge_list.all.challenges : c if c.difficulty == "beginner"])
      intermediate = length([for c in data.ctfchallenge_list.all.challenges : c if c.difficulty == "intermediate"])
      advanced     = length([for c in data.ctfchallenge_list.all.challenges : c if c.difficulty == "advanced"])
    }
  }
}
```

## Schema

### Optional

- `difficulty` (String) Filter challenges by difficulty level. Valid values: `beginner`, `intermediate`, `advanced`.
- `category` (String) Filter challenges by category. Valid values: `fundamentals`, `expressions`, `state`, `modules`, `advanced-syntax`, `loops`, `data-sources`, `functions`.

### Read-Only

- `id` (String) The unique identifier for this data source query.
- `challenges` (List of Object) The list of challenges matching the filter criteria.
  - `id` (String) The challenge ID.
  - `name` (String) The challenge name.
  - `description` (String) A brief description of the challenge.
  - `points` (Number) Points awarded for completing this challenge.
  - `difficulty` (String) The difficulty level.
  - `category` (String) The challenge category.
- `total_points` (Number) The total points available from all listed challenges.

## Categories

- **fundamentals** - Basic Terraform concepts and resource management
- **expressions** - Terraform expression syntax and functions
- **state** - State management and understanding
- **modules** - Module creation and composition
- **advanced-syntax** - Dynamic blocks and meta-arguments
- **loops** - for_each and count
- **data-sources** - Querying and filtering data
- **functions** - Built-in function usage