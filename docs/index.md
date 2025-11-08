---
page_title: "CTF Challenge Provider"
subcategory: ""
description: |-
  A Capture The Flag (CTF) game implemented as a Terraform provider. Learn Terraform concepts by solving interactive challenges.
---

# CTF Challenge Provider

The CTF Challenge provider is an educational tool that gamifies learning Terraform. It provides a series of challenges that teach Terraform concepts through hands-on practice.

## Features

- 🎯 **8 Progressive Challenges** - From beginner to advanced Terraform concepts
- 💡 **Hint System** - Get help when you're stuck (with point penalties)
- 🏆 **Point System** - Earn up to 2,250 points total
- 🔐 **Flag Capture** - Complete challenges to reveal flags as rewards
- 📚 **Educational** - Learn dependencies, expressions, modules, state management, and more

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

# Get challenge information
data "ctfchallenge_challenge_info" "basics" {
  challenge_id = "terraform_basics"
}

# Solve a challenge by providing proof of work
resource "null_resource" "first" {}
resource "null_resource" "second" { depends_on = [null_resource.first] }
resource "null_resource" "third" { depends_on = [null_resource.second] }

resource "ctfchallenge_flag_validator" "basics" {
  challenge_id = "terraform_basics"
  
  proof_of_work = {
    dependencies = "${null_resource.first.id},${null_resource.second.id},${null_resource.third.id}"
  }
}

# Capture the flag!
output "flag" {
  value     = ctfchallenge_flag_validator.basics.flag
  sensitive = true
}
```

## Available Challenges

| Challenge | Difficulty | Points | Category |
|-----------|-----------|---------|----------|
| [Terraform Basics](#terraform-basics) | Beginner | 100 | Fundamentals |
| [Expression Expert](#expression-expert) | Intermediate | 350 | Expressions |
| [State Secrets](#state-secrets) | Beginner | 200 | State |
| [Module Master](#module-master) | Advanced | 400 | Modules |
| [Dynamic Blocks](#dynamic-blocks) | Intermediate | 300 | Advanced Syntax |
| [For-Each Wizard](#for-each-wizard) | Intermediate | 250 | Loops |
| [Data Source Detective](#data-source-detective) | Beginner | 150 | Data Sources |
| [Cryptographic Compute](#cryptographic-compute) | Advanced | 500 | Functions |

**Total Points Available:** 2,250

## Challenge Descriptions

### Terraform Basics
Learn about resource dependencies and the `depends_on` argument. Create a chain of dependent resources to understand execution order. Complete the challenge to capture your first flag!

### Expression Expert
Master Terraform's expression syntax and built-in functions including string manipulation, hashing, and encoding. Solve cryptographic puzzles to reveal the flag.

### State Secrets
Understand Terraform state management and how Terraform tracks resources. Discover the "magic number" hidden in state concepts to unlock the flag.

### Module Master
Learn to create and use Terraform modules for code reuse and organization. Build composable infrastructure patterns and capture the flag through proper module composition.

### Dynamic Blocks
Master the `dynamic` block feature for generating multiple nested blocks programmatically from data structures. Generate enough blocks to earn your flag.

### For-Each Wizard
Use the `for_each` meta-argument to create multiple similar resources elegantly without repetition. Create the right resources to reveal your flag.

### Data Source Detective
Query and filter data sources effectively. Learn to extract and transform data for use in your infrastructure code. Filter correctly to capture the flag.

### Cryptographic Compute
Use Terraform's cryptographic and hashing functions to solve complex puzzles. Chain multiple functions together correctly to reveal the flag.

## Schema

### Optional

- `player_name` (String) Your player name for the CTF. Can also be set via the `TF_CTF_PLAYER` environment variable. Defaults to `"anonymous"`.
- `api_endpoint` (String) Optional API endpoint for score tracking. Can also be set via the `TF_CTF_API` environment variable.

## Getting Started

See the [Getting Started Guide](guides/getting-started.md) for a step-by-step walkthrough of your first challenge.

## Resources

- [ctfchallenge_flag_validator](resources/flag_validator.md) - Validate challenge solutions and capture flags
- [ctfchallenge_puzzle_box](resources/puzzle_box.md) - Solve logic puzzles for bonus flags

## Data Sources

- [ctfchallenge_hint](data-sources/hint.md) - Get hints for challenges
- [ctfchallenge_list](data-sources/list.md) - List all available challenges
- [ctfchallenge_challenge_info](data-sources/challenge_info.md) - Get detailed information about a specific challenge

## How CTF Challenges Work

In traditional CTF (Capture The Flag) competitions, you complete a challenge and receive a flag as proof of completion. This provider follows that paradigm:

1. **Read the challenge description** - Understand what you need to accomplish
2. **Build your solution** - Write Terraform code that meets the requirements
3. **Submit proof of work** - Validate your solution with the flag_validator resource
4. **Capture the flag** - If successful, the flag is revealed as your reward!

The flag format is: `flag{some_text_here}`

## Viewing Captured Flags

Flags are marked as sensitive outputs. To view them:

```bash
terraform output -raw flag
```

Good luck, and happy flag hunting! 🚀