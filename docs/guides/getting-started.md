---
page_title: "Getting Started with CTF Challenge Provider"
subcategory: "Guides"
description: |-
  A step-by-step guide to solving your first challenge.
---

# Getting Started

This guide walks you through setting up the CTF Challenge provider and solving your first challenge.

## Step 1: Configure the Provider

Create a new directory and a `main.tf` file:

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
  player_name = "your-name-here"
}
```

Initialize Terraform:

```bash
terraform init
```

## Step 2: Explore Available Challenges

List all challenges to see what's available:

```terraform
data "ctfchallenge_list" "all" {}

output "challenges" {
  value = data.ctfchallenge_list.all.challenges
}
```

Apply to see the output:

```bash
terraform apply
```

## Step 3: Get Challenge Information

Let's start with the "Terraform Basics" challenge:

```terraform
data "ctfchallenge_challenge_info" "basics" {
  challenge_id = "terraform_basics"
}

output "challenge_info" {
  value = {
    name        = data.ctfchallenge_challenge_info.basics.name
    description = data.ctfchallenge_challenge_info.basics.description
    points      = data.ctfchallenge_challenge_info.basics.points
  }
}
```

## Step 4: Request a Hint (Optional)

If you need help:

```terraform
data "ctfchallenge_hint" "basics_hint" {
  challenge_id = "terraform_basics"
  level        = 0  # Start with level 0
}

output "hint" {
  value = data.ctfchallenge_hint.basics_hint.hint
}
```

## Step 5: Solve the Challenge

For the "Terraform Basics" challenge, you need to create at least 3 dependent resources:

```terraform
# Create a chain of dependent resources
resource "null_resource" "first" {
  triggers = {
    timestamp = timestamp()
  }
}

resource "null_resource" "second" {
  depends_on = [null_resource.first]
  
  triggers = {
    timestamp = timestamp()
  }
}

resource "null_resource" "third" {
  depends_on = [null_resource.second]
  
  triggers = {
    timestamp = timestamp()
  }
}
```

## Step 6: Submit Your Solution

Now validate your solution with the flag:

```terraform
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

output "points" {
  value = ctfchallenge_flag_validator.basics.points
}
```

## Step 7: Apply and See Results

```bash
terraform apply
```

If successful, you'll see:

```
Outputs:

points = 100
result = "🎉 Congratulations! You solved 'Terraform Basics' and earned 100 points!"
```

## Next Steps

- Try the **Expression Expert** challenge to learn about Terraform functions
- Solve the **State Secrets** challenge to understand state management
- Work through all 8 challenges to earn the full 2,250 points!

## Tips

1. **Read the error messages** - They often contain hints about what's missing
2. **Use hints strategically** - They cost points but can save you time
3. **Experiment** - Terraform is declarative, so it's safe to try different approaches
4. **Check the documentation** - Each resource and data source has detailed examples

## Complete Example

Here's a complete working example for the first challenge:

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
  player_name = "alice"
}

# Your solution
resource "null_resource" "first" {}
resource "null_resource" "second" { depends_on = [null_resource.first] }
resource "null_resource" "third" { depends_on = [null_resource.second] }

# Validation
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
```

Happy learning! 🚀