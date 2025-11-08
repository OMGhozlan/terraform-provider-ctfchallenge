---
page_title: "Getting Started with CTF Challenge Provider"
subcategory: "Guides"
description: |-
  A step-by-step guide to solving your first challenge and capturing your first flag.
---

# Getting Started

This guide walks you through setting up the CTF Challenge provider and capturing your first flag!

## What is a CTF?

CTF (Capture The Flag) is a gamified learning approach where you solve challenges to earn flags. In this provider:

1. You read a challenge description
2. You write Terraform code to solve it
3. You submit proof of your work
4. **The flag is revealed as your reward!**

Flags look like this: `flag{some_text_here}`

## Step 1: Configure the Provider

Create a new directory for your CTF workspace:

```bash
mkdir terraform-ctf
cd terraform-ctf
```

Create a `main.tf` file:

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
  player_name = "your-name-here"  # Choose your player name!
}
```

Initialize Terraform:

```bash
terraform init
```

## Step 2: Explore Available Challenges

List all challenges to see what flags you can capture:

```terraform
data "ctfchallenge_list" "all" {}

output "challenges" {
  value = data.ctfchallenge_list.all.challenges
}

output "total_flags_available" {
  value = length(data.ctfchallenge_list.all.challenges)
}

output "total_points_available" {
  value = data.ctfchallenge_list.all.total_points
}
```

Apply to see the output:

```bash
terraform apply -auto-approve
```

You should see 8 challenges worth 2,250 total points!

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
    difficulty  = data.ctfchallenge_challenge_info.basics.difficulty
  }
}
```

Apply again:

```bash
terraform apply -auto-approve
```

Read the challenge description carefully!

## Step 4: Request a Hint (Optional)

If you need help, you can request hints. But remember, hints cost points!

```terraform
data "ctfchallenge_hint" "basics_hint" {
  challenge_id = "terraform_basics"
  level        = 0  # Start with level 0 (costs 10 points)
}

output "hint" {
  value = data.ctfchallenge_hint.basics_hint.hint
}

output "hint_cost" {
  value = data.ctfchallenge_hint.basics_hint.cost
}
```

## Step 5: Solve the Challenge

For the "Terraform Basics" challenge, you need to create at least 3 dependent resources. Here's the solution:

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

**Explanation:**
- We create 3 null_resource resources
- Each one depends on the previous one using `depends_on`
- This creates a dependency chain showing you understand resource ordering

## Step 6: Submit Your Solution and Capture the Flag!

Now for the exciting part - submitting your proof of work to capture the flag:

```terraform
resource "ctfchallenge_flag_validator" "basics" {
  challenge_id = "terraform_basics"

  proof_of_work = {
    dependencies = "${null_resource.first.id},${null_resource.second.id},${null_resource.third.id}"
  }
}

# The flag is revealed when you succeed!
output "flag" {
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

## Step 7: Apply and View Your Flag

Apply your solution:

```bash
terraform apply
```

If successful, you'll see:

```
Outputs:

points_earned = 100
result = "🎉 Congratulations! You solved 'Terraform Basics' and earned 100 points!"
```

Now capture your flag:

```bash
terraform output -raw flag
```

You should see:

```
flag{t3rr4f0rm_d3p3nd3nc13s}
```

🎉 **Congratulations! You've captured your first flag!**

## Step 8: Keep Your Flags

Create a `flags.txt` file to track your captures:

```bash
echo "terraform_basics: $(terraform output -raw flag)" >> flags.txt
```

## Complete Working Example

Here's a complete `main.tf` for your first challenge:

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

# Get challenge info
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

# Your solution - create 3 dependent resources
resource "null_resource" "first" {}
resource "null_resource" "second" { depends_on = [null_resource.first] }
resource "null_resource" "third" { depends_on = [null_resource.second] }

# Submit and capture the flag
resource "ctfchallenge_flag_validator" "basics" {
  challenge_id = "terraform_basics"
  
  proof_of_work = {
    dependencies = "${null_resource.first.id},${null_resource.second.id},${null_resource.third.id}"
  }
}

# Outputs
output "flag" {
  value     = ctfchallenge_flag_validator.basics.flag
  sensitive = true
}

output "result" {
  value = ctfchallenge_flag_validator.basics.message
}

output "points" {
  value = ctfchallenge_flag_validator.basics.points
}
```

## Next Challenges

Now that you've captured your first flag, try these next:

1. **State Secrets** (Beginner, 200 points) - Find the magic number
2. **Data Source Detective** (Beginner, 150 points) - Master data filtering
3. **Expression Expert** (Intermediate, 350 points) - Learn Terraform functions

## Tips for Success

1. **Read error messages carefully** - They often contain hints about what's missing
2. **Use hints strategically** - 10-30 points for a hint is worth it if you're stuck
3. **Experiment safely** - Terraform is declarative, so you can try different approaches
4. **Use `terraform console`** - Test expressions interactively
5. **Check the documentation** - Each resource has detailed examples
6. **The flag is the reward** - You never need to know the flag beforehand!

## Tracking Your Progress

Create a simple scoreboard:

```terraform
locals {
  completed_challenges = {
    terraform_basics = 100  # Add points as you complete challenges
  }
  
  total_earned = sum(values(local.completed_challenges))
  total_possible = 2250
  completion_percentage = (local.total_earned / local.total_possible) * 100
}

output "scoreboard" {
  value = {
    challenges_completed = length(local.completed_challenges)
    total_points         = local.total_earned
    completion           = "${local.completion_percentage}%"
  }
}
```

## What If I Get Stuck?

1. **Request a hint** - Start with level 0
2. **Read Terraform docs** - Check the [Terraform documentation](https://www.terraform.io/docs)
3. **Use terraform console** - Test expressions interactively
4. **Check the walkthrough** - See [Challenge Walkthrough](challenge-walkthrough.md) for solutions
5. **Take a break** - Sometimes stepping away helps!

## File Organization

As you progress, organize your challenges:

```
terraform-ctf/
├── main.tf                    # Provider configuration
├── challenge_1_basics.tf      # Terraform Basics
├── challenge_2_expressions.tf # Expression Expert
├── challenge_3_state.tf       # State Secrets
├── flags.txt                  # Your captured flags
└── README.md                  # Your notes
```

Happy flag hunting! 🚀

Remember: In CTF, the flag is your **trophy**, not your **ticket**!