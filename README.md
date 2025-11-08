# Terraform Provider CTF Challenge 🎯

[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org)
[![Terraform](https://img.shields.io/badge/terraform-1.0+-purple.svg)](https://www.terraform.io)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

A Capture The Flag (CTF) game implemented as a Terraform provider. Learn Terraform concepts by solving interactive challenges and capturing flags!

## 🎮 What is This?

This Terraform provider turns learning Infrastructure as Code into an interactive game. Instead of reading documentation, you'll solve progressively challenging puzzles that teach you Terraform concepts hands-on. Each challenge you complete reveals a flag as your reward!

**Perfect for:**
- 🎓 Learning Terraform interactively
- 🏢 Team training and workshops
- 🎯 CTF competitions and hackathons
- 💪 Practicing Terraform skills
- 🎉 Having fun with Infrastructure as Code!

## ✨ Features

- **8 Progressive Challenges** - From beginner to advanced (2,250 total points)
- **Real CTF Experience** - Complete challenges to capture flags
- **Hint System** - Get help when stuck (with point penalties)
- **Comprehensive Documentation** - Guides, examples, and walkthroughs
- **Educational** - Learn by doing, not just reading
- **Bonus Puzzles** - XOR puzzle and more for extra flags
- **No Cloud Required** - All challenges run locally

## 🏆 Challenges

| Challenge | Difficulty | Points | What You'll Learn |
|-----------|-----------|---------|-------------------|
| Terraform Basics | Beginner | 100 | Resource dependencies & execution order |
| State Secrets | Beginner | 200 | State management concepts |
| Data Source Detective | Beginner | 150 | Querying and filtering data |
| For-Each Wizard | Intermediate | 250 | Resource iteration with for_each |
| Dynamic Blocks | Intermediate | 300 | Dynamic block generation |
| Expression Expert | Intermediate | 350 | Functions and expressions |
| Module Master | Advanced | 400 | Module composition |
| Cryptographic Compute | Advanced | 500 | Cryptographic functions |

**Total: 2,250 points** + bonus flags!

## 🚀 Quick Start

### Installation

#### Using Terraform Registry

Add to your `terraform` block:

```terraform
terraform {
  required_providers {
    ctfchallenge = {
      source  = "omghozlan/ctfchallenge"
      version = "~> 1.0"
    }
  }
}
```

#### Local Development

```bash
# Clone the repository
git clone https://github.com/omghozlan/terraform-provider-ctfchallenge.git
cd terraform-provider-ctfchallenge

# Build the provider
go build -o terraform-provider-ctfchallenge

# Move to Terraform plugins directory (Linux/macOS example)
mkdir -p ~/.terraform.d/plugins/github.com/omghozlan/ctfchallenge/1.0.0/linux_amd64/
mv terraform-provider-ctfchallenge ~/.terraform.d/plugins/github.com/omghozlan/ctfchallenge/1.0.0/linux_amd64/
```

### Your First Challenge

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
  player_name = "your-name-here"  # Choose your hacker alias!
}

# List all available challenges
data "ctfchallenge_list" "all" {}

output "challenges" {
  value = data.ctfchallenge_list.all.challenges
}
```

Initialize and explore:

```bash
terraform init
terraform apply
```

### Solve Your First Challenge

```terraform
# Get challenge information
data "ctfchallenge_challenge_info" "basics" {
  challenge_id = "terraform_basics"
}

# Create your solution (3 dependent resources)
resource "null_resource" "first" {}
resource "null_resource" "second" { depends_on = [null_resource.first] }
resource "null_resource" "third" { depends_on = [null_resource.second] }

# Submit and capture the flag!
resource "ctfchallenge_flag_validator" "basics" {
  challenge_id = "terraform_basics"
  
  proof_of_work = {
    dependencies = "${null_resource.first.id},${null_resource.second.id},${null_resource.third.id}"
  }
}

# View your captured flag
output "flag" {
  value     = ctfchallenge_flag_validator.basics.flag
  sensitive = true
}

output "points" {
  value = ctfchallenge_flag_validator.basics.points
}
```

Apply and capture your first flag:

```bash
terraform apply
terraform output -raw flag
# Output: flag{t3rr4f0rm_d3p3nd3nc13s}
```

🎉 **Congratulations! You've captured your first flag!**

## 📚 Documentation

### Resources

- **[ctfchallenge_flag_validator](docs/resources/flag_validator.md)** - Submit challenge solutions and capture flags
- **[ctfchallenge_puzzle_box](docs/resources/puzzle_box.md)** - Solve logic puzzles for bonus flags

### Data Sources

- **[ctfchallenge_list](docs/data-sources/list.md)** - List all available challenges
- **[ctfchallenge_challenge_info](docs/data-sources/challenge_info.md)** - Get detailed challenge information
- **[ctfchallenge_hint](docs/data-sources/hint.md)** - Request hints (costs points)

### Guides

- **[Getting Started](docs/guides/getting-started.md)** - Step-by-step tutorial
- **[Challenge Walkthrough](docs/guides/challenge-walkthrough.md)** - Complete solutions (spoilers!)
- **[Advanced Tips](docs/guides/advanced-challenges.md)** - Pro strategies and techniques

## 🎯 How It Works

Unlike traditional learning where you read then practice, this provider follows the **CTF paradigm**:

1. **Choose a challenge** - Read the description to understand the goal
2. **Write your solution** - Use Terraform to solve the puzzle
3. **Submit proof of work** - Validate your solution
4. **Capture the flag** - If successful, the flag is revealed as your reward!

Flags follow the format: `flag{some_text_here}`

## 💡 Example: Expression Expert Challenge

This challenge teaches you Terraform's built-in functions:

```terraform
# The challenge: compute base64encode(sha256("terraformexpressionsrock"))

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
```

```bash
terraform apply
terraform output -raw flag
# Output: flag{3xpr3ss10ns_unl0ck3d}
```

**350 points earned!** 🎉

## 🎲 Bonus: XOR Puzzle

Solve the XOR puzzle for a bonus flag:

```terraform
# Find 5 numbers whose XOR equals 0
resource "ctfchallenge_puzzle_box" "xor" {
  inputs = {
    input_1 = "15"
    input_2 = "23"
    input_3 = "42"
    input_4 = "37"
    input_5 = "11"  # Calculated: 15 XOR 23 XOR 42 XOR 37 = 11
  }
}

output "bonus_flag" {
  value     = ctfchallenge_puzzle_box.xor.secret_output
  sensitive = true
}
```

## 🛠️ Development

### Prerequisites

- [Go](https://golang.org/doc/install) 1.21+
- [Terraform](https://www.terraform.io/downloads.html) 1.0+

### Building from Source

```bash
# Clone the repository
git clone https://github.com/omghozlan/terraform-provider-ctfchallenge.git
cd terraform-provider-ctfchallenge

# Download dependencies
go mod download

# Build
go build -o terraform-provider-ctfchallenge

# Run tests
go test -v ./...
```

### Project Structure

```
terraform-provider-ctfchallenge/
├── challenges/           # Challenge definitions and validators
│   └── validator.go
├── provider/             # Terraform provider implementation
│   ├── provider.go
│   ├── resource_flag_validator.go
│   ├── resource_puzzle_box.go
│   ├── data_source_hint.go
│   ├── data_source_list.go
│   └── data_source_challenge_info.go
├── docs/                 # Documentation
│   ├── index.md
│   ├── resources/
│   ├── data-sources/
│   └── guides/
├── examples/             # Example configurations
├── main.go              # Provider entry point
├── go.mod
└── README.md
```

### Running Tests

```bash
# Run all tests
go test -v ./...

# Run with coverage
go test -v -cover ./...

# Run specific test
go test -v -run TestValidateExpressions ./challenges
```

### Adding a New Challenge

1. Add challenge definition to `challenges/validator.go`:

```go
"my_challenge": {
    ID:          "my_challenge",
    Name:        "My Challenge",
    Description: "Learn something cool",
    Points:      300,
    Flag:        "flag{my_fl4g}",
    Difficulty:  "intermediate",
    Category:    "my-category",
    Validator:   validateMyChallenge,
},
```

2. Implement the validator function:

```go
func validateMyChallenge(input map[string]interface{}) (bool, string, error) {
    correctFlag := "flag{my_fl4g}"
    
    // Your validation logic here
    if /* challenge solved */ {
        return true, correctFlag, nil
    }
    
    return false, "", fmt.Errorf("challenge not solved")
}
```

3. Add hints to `GetHint()` function

4. Update documentation

5. Add example to `examples/`

## 🧪 Testing the Provider

Use the included examples to test functionality:

```bash
cd examples/challenge1_basics
terraform init
terraform apply

cd ../challenge2_expressions
terraform init
terraform apply
```

## 📖 Tips for Players

1. **Use `terraform console`** - Test expressions interactively
   ```bash
   $ terraform console
   > sha256("test")
   > base64encode(sha256("test"))
   ```

2. **Request hints strategically** - They cost 10-30 points but can save time

3. **Read error messages** - They often contain helpful debugging info

4. **Start with beginner challenges** - Build foundational knowledge

5. **Track your flags** - Keep a `flags.txt` file with your captures
   ```bash
   echo "terraform_basics: $(terraform output -raw flag)" >> flags.txt
   ```

6. **Use version control** - Commit your progress
   ```bash
   git add -A
   git commit -m "Solved Expression Expert - 350 points!"
   ```

## 🎓 What You'll Learn

By completing all challenges, you'll master:

- ✅ Resource dependencies and ordering (`depends_on`)
- ✅ Terraform state management
- ✅ Data sources and filtering
- ✅ The `for_each` meta-argument
- ✅ Dynamic block generation
- ✅ Terraform expressions and functions
- ✅ Module creation and composition
- ✅ Cryptographic functions (sha256, md5, base64encode)
- ✅ String manipulation and templating
- ✅ Best practices and patterns

## 🏅 Scoring

- **Total Points Available:** 2,250
- **Hint Penalties:** 10-30 points per hint
- **Bonus Puzzles:** Extra flags available!

Track your score:

```terraform
locals {
  completed = {
    terraform_basics    = 100
    state_secrets       = 200
    expression_expert   = 350
    # ... add as you complete
  }
  
  hints_used = {
    expression_expert_l0 = 10
    # ... track hint costs
  }
  
  total_earned = sum(values(local.completed))
  total_hints  = sum(values(local.hints_used))
  net_score    = local.total_earned - local.total_hints
}

output "scoreboard" {
  value = {
    points_earned   = local.total_earned
    hints_used      = local.total_hints
    net_score       = local.net_score
    completion_pct  = "${(local.total_earned / 2250) * 100}%"
  }
}
```

## 🤝 Contributing

Contributions are welcome! Here's how you can help:

- 🐛 Report bugs or issues
- 💡 Suggest new challenges
- 📝 Improve documentation
- ✨ Add features
- 🎯 Create additional puzzles

### Contribution Guidelines

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-challenge`)
3. Make your changes
4. Add tests if applicable
5. Update documentation
6. Commit your changes (`git commit -m 'Add amazing challenge'`)
7. Push to the branch (`git push origin feature/amazing-challenge`)
8. Open a Pull Request

### Code of Conduct

- Be respectful and inclusive
- Provide constructive feedback
- Help others learn
- Keep it fun and educational!

## 📜 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- HashiCorp for Terraform and the excellent provider SDK
- The CTF community for inspiration
- All contributors and players!

## 📞 Support

- 📚 [Documentation](docs/)
- 🐛 [Issue Tracker](https://github.com/omghozlan/terraform-provider-ctfchallenge/issues)
- 💬 [Discussions](https://github.com/omghozlan/terraform-provider-ctfchallenge/discussions)

## 🎉 Hall of Fame

Completed all challenges? Submit a PR adding yourself to the Hall of Fame!

```markdown
| Player | Score | Date | Notes |
|--------|-------|------|-------|
| @yourname | 2,250 | 2025-01-15 | Perfect score! |
```

## 🚀 Roadmap

Future ideas:

- [ ] More challenges (cloud-specific scenarios)
- [ ] Leaderboard API integration
- [ ] Team mode
- [ ] Timed challenges
- [ ] Achievement badges
- [ ] Multi-language support
- [ ] Video walkthroughs

## 🎮 Ready to Play?

```bash
# Start your CTF journey now!
mkdir terraform-ctf-journey
cd terraform-ctf-journey
terraform init
# ... and let the flag hunting begin! 🏴‍☠️
```

---

**Made with ❤️ for the Terraform community**

*"The best way to learn is by doing. The most fun way to do is by playing."*

🎯 **Happy Flag Hunting!** 🏴‍☠️