---
page_title: "ctfchallenge_list Data Source - ctfchallenge"
subcategory: ""
description: |-
  Lists all available challenges with optional filtering.
---

# ctfchallenge_list (Data Source)

The `list` data source retrieves all available challenges, optionally filtered by difficulty or category. Use this to get an overview of all flags you can capture!

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

output "beginner_total_points" {
  value = data.ctfchallenge_list.beginner.total_points
}
```

## Example Creating a Challenge Dashboard

```terraform
data "ctfchallenge_list" "all" {}

locals {
  challenges_by_difficulty = {
    for challenge in data.ctfchallenge_list.all.challenges :
    challenge.difficulty => challenge...
  }
  
  challenges_by_category = {
    for challenge in data.ctfchallenge_list.all.challenges :
    challenge.category => challenge...
  }
}

output "challenge_dashboard" {
  value = {
    total_challenges = length(data.ctfchallenge_list.all.challenges)
    total_points     = data.ctfchallenge_list.all.total_points
    
    by_difficulty = {
      beginner     = length([for c in data.ctfchallenge_list.all.challenges : c if c.difficulty == "beginner"])
      intermediate = length([for c in data.ctfchallenge_list.all.challenges : c if c.difficulty == "intermediate"])
      advanced     = length([for c in data.ctfchallenge_list.all.challenges : c if c.difficulty == "advanced"])
    }
    
    beginner_points     = sum([for c in data.ctfchallenge_list.all.challenges : c.points if c.difficulty == "beginner"])
    intermediate_points = sum([for c in data.ctfchallenge_list.all.challenges : c.points if c.difficulty == "intermediate"])
    advanced_points     = sum([for c in data.ctfchallenge_list.all.challenges : c.points if c.difficulty == "advanced"])
  }
}
```

## Example - Recommended Challenge Path

```terraform
data "ctfchallenge_list" "beginner" {
  difficulty = "beginner"
}

data "ctfchallenge_list" "intermediate" {
  difficulty = "intermediate"
}

data "ctfchallenge_list" "advanced" {
  difficulty = "advanced"
}

output "recommended_path" {
  value = {
    step_1 = {
      description = "Start here! Complete all beginner challenges"
      challenges  = [for c in data.ctfchallenge_list.beginner.challenges : c.name]
      points      = data.ctfchallenge_list.beginner.total_points
    }
    step_2 = {
      description = "Level up with intermediate challenges"
      challenges  = [for c in data.ctfchallenge_list.intermediate.challenges : c.name]
      points      = data.ctfchallenge_list.intermediate.total_points
    }
    step_3 = {
      description = "Master Terraform with advanced challenges"
      challenges  = [for c in data.ctfchallenge_list.advanced.challenges : c.name]
      points      = data.ctfchallenge_list.advanced.total_points
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

## Valid Difficulty Values

- `beginner` - New to Terraform or the concept
- `intermediate` - Familiar with basics, ready for more complex scenarios
- `advanced` - Expert-level challenges requiring deep knowledge

## Valid Category Values

- **fundamentals** - Basic Terraform concepts and resource management
- **expressions** - Terraform expression syntax and functions
- **state** - State management and understanding
- **modules** - Module creation and composition
- **advanced-syntax** - Dynamic blocks and meta-arguments
- **loops** - for_each and count
- **data-sources** - Querying and filtering data
- **functions** - Built-in function usage

## All Challenges Summary

| Difficulty | Count | Points |
|-----------|--------|--------|
| Beginner | 3 | 450 |
| Intermediate | 3 | 900 |
| Advanced | 2 | 900 |
| **Total** | **8** | **2,250** |

## Tips

1. **Start with beginners** - Build a solid foundation
2. **Filter by category** - Focus on one skill area at a time
3. **Track your progress** - Keep note of which flags you've captured
4. **Mix difficulties** - If stuck on a hard one, try an easier challenge to build momentum