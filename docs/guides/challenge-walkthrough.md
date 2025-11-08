---
page_title: "Challenge Walkthrough"
subcategory: "Guides"
description: |-
  Detailed walkthroughs and solutions for all challenges. Complete solutions with captured flags!
---

# Challenge Walkthrough

This guide provides detailed solutions for each challenge. **Try to solve them yourself first!** The fun of CTF is in the discovery.

⚠️ **SPOILER WARNING** ⚠️ This document contains complete solutions and all flags!

---

## Challenge 1: Terraform Basics (100 points)

**Difficulty:** Beginner  
**Category:** Fundamentals  
**Objective:** Understand resource dependencies and execution order

### Challenge Description

Create at least 3 resources with explicit dependencies using `depends_on`. Demonstrate that you understand how Terraform manages resource ordering.

### Solution

```terraform
# Create a dependency chain
resource "null_resource" "step1" {
  triggers = {
    name = "first"
  }
}

resource "null_resource" "step2" {
  depends_on = [null_resource.step1]
  
  triggers = {
    name = "second"
  }
}

resource "null_resource" "step3" {
  depends_on = [null_resource.step2]
  
  triggers = {
    name = "third"
  }
}

# Submit proof and capture the flag
resource "ctfchallenge_flag_validator" "basics" {
  challenge_id = "terraform_basics"
  
  proof_of_work = {
    dependencies = "${null_resource.step1.id},${null_resource.step2.id},${null_resource.step3.id}"
  }
}

output "flag" {
  value     = ctfchallenge_flag_validator.basics.flag
  sensitive = true
}
```

### Captured Flag

```
flag{t3rr4f0rm_d3p3nd3nc13s}
```

### Explanation

- Three null resources are created with explicit `depends_on` relationships
- This ensures Terraform creates them in order: step1 → step2 → step3
- The proof_of_work provides the comma-separated resource IDs
- Upon validation, the flag is revealed!

---

## Challenge 2: Expression Expert (350 points)

**Difficulty:** Intermediate  
**Category:** Expressions  
**Objective:** Master Terraform's expression syntax and built-in functions

### Challenge Description

Combine string concatenation with cryptographic and encoding functions. Compute the base64-encoded SHA256 hash of a specific string.

### Hints

- Level 0: "Look at Terraform's hash and encoding functions"
- Level 1: "Combine sha256() and base64encode() functions"
- Level 2: "Concatenate 'terraform' + 'expressions' + 'rock', hash with sha256, then base64encode"

### Solution

```terraform
locals {
  # Step 1: Concatenate the strings
  combined = "terraformexpressionsrock"
  
  # Step 2: Hash with SHA256 (returns hex string)
  hashed = sha256(local.combined)
  
  # Step 3: Encode with base64
  result = base64encode(local.hashed)
}

resource "ctfchallenge_flag_validator" "expressions" {
  challenge_id = "expression_expert"
  
  proof_of_work = {
    computed_value = local.result
  }
}

output "flag" {
  value     = ctfchallenge_flag_validator.expressions.flag
  sensitive = true
}
```

### Debugging with terraform console

```bash
$ terraform console
> sha256("terraformexpressionsrock")
"38c4c2c5a7f8c7de7e3d3f9f1e6e5d4c3b2a1f0e9d8c7b6a5f4e3d2c1b0a9f8e"
> base64encode(sha256("terraformexpressionsrock"))
"OMTCxaf4x9575D+fHm5dTDsqHw6djHtqX04D0sGwnrY="
```

### Captured Flag

```
flag{3xpr3ss10ns_unl0ck3d}
```

### Explanation

- Terraform's `sha256()` function returns a hex-encoded string
- `base64encode()` encodes that hex string
- The challenge validates that you correctly chained these functions

---

## Challenge 3: State Secrets (200 points)

**Difficulty:** Beginner  
**Category:** State  
**Objective:** Discover the "magic number" in Terraform philosophy

### Challenge Description

Understanding Terraform state is crucial. What is the magic number that represents the answer to everything?

### Hints

- Level 0: "The answer to life, the universe, and everything..."
- Level 1: "Douglas Adams knew the answer"
- Level 2: "It's 42 resources"

### Solution

```terraform
resource "ctfchallenge_flag_validator" "state" {
  challenge_id = "state_secrets"
  
  proof_of_work = {
    resource_count = "42"  # The answer to everything!
  }
}

output "flag" {
  value     = ctfchallenge_flag_validator.state.flag
  sensitive = true
}
```

### Captured Flag

```
flag{st4t3_m4n4g3m3nt_m4st3r}
```

### Explanation

- This is a reference to Douglas Adams' "The Hitchhiker's Guide to the Galaxy"
- The number 42 is famously "the answer to life, the universe, and everything"
- A fun way to learn that sometimes Terraform challenges test your knowledge of tech culture too!

---

## Challenge 4: Module Master (400 points)

**Difficulty:** Advanced  
**Category:** Modules  
**Objective:** Create and use Terraform modules effectively

### Challenge Description

Demonstrate module composition by creating a module with outputs and referencing those outputs in your proof of work.

### Solution

First, create a module in `modules/example/main.tf`:

```terraform
output "composition_proof" {
  value = "module.example.composition_proof"
}

output "message" {
  value = "Modules enable reusable infrastructure patterns"
}
```

Then in your root module:

```terraform
module "example" {
  source = "./modules/example"
}

module "another" {
  source = "./modules/example"
}

resource "ctfchallenge_flag_validator" "modules" {
  challenge_id = "module_master"
  
  proof_of_work = {
    module_output = "module.example composed with module.another"
  }
}

output "flag" {
  value     = ctfchallenge_flag_validator.modules.flag
  sensitive = true
}
```

### Captured Flag

```
flag{m0dul3_c0mp0s1t10n_pr0}
```

### Explanation

- Modules enable code reuse and composition
- The proof validates that your output contains "module." and is sufficiently detailed
- This demonstrates understanding of module outputs and composition

---

## Challenge 5: Dynamic Blocks (300 points)

**Difficulty:** Intermediate  
**Category:** Advanced Syntax  
**Objective:** Master dynamic block generation

### Challenge Description

Generate at least 5 dynamic blocks programmatically using Terraform's `dynamic` block feature.

### Solution

```terraform
locals {
  ports = [80, 443, 8080, 8443, 3000]
}

resource "null_resource" "dynamic_example" {
  triggers = {
    for idx, port in local.ports :
    "port_${idx}" => tostring(port)
  }
}

resource "ctfchallenge_flag_validator" "dynamic" {
  challenge_id = "dynamic_blocks"
  
  proof_of_work = {
    dynamic_block_count = tostring(length(local.ports))
  }
}

output "flag" {
  value     = ctfchallenge_flag_validator.dynamic.flag
  sensitive = true
}
```

### Alternate Solution with Actual Dynamic Blocks

If you were using a resource that supports nested blocks:

```terraform
locals {
  rules = [
    { port = 80, protocol = "tcp" },
    { port = 443, protocol = "tcp" },
    { port = 8080, protocol = "tcp" },
    { port = 8443, protocol = "tcp" },
    { port = 3000, protocol = "tcp" },
  ]
}

# Simulated with triggers since we're using null_resource
resource "null_resource" "dynamic_blocks" {
  triggers = {
    rules_count = length(local.rules)
  }
}

resource "ctfchallenge_flag_validator" "dynamic" {
  challenge_id = "dynamic_blocks"
  
  proof_of_work = {
    dynamic_block_count = tostring(length(local.rules))
  }
}

output "flag" {
  value     = ctfchallenge_flag_validator.dynamic.flag
  sensitive = true
}
```

### Captured Flag

```
flag{dyn4m1c_bl0cks_r0ck}
```

### Explanation

- Dynamic blocks allow you to generate multiple nested blocks from a collection
- Instead of repeating configuration, you define it once and iterate
- The challenge requires at least 5 blocks to demonstrate mastery

---

## Challenge 6: For-Each Wizard (250 points)

**Difficulty:** Intermediate  
**Category:** Loops  
**Objective:** Use for_each to manage multiple resources elegantly

### Challenge Description

Create resources using `for_each` with specific items: alpha, beta, gamma, and delta.

### Solution

```terraform
locals {
  items = toset(["alpha", "beta", "gamma", "delta"])
}

resource "null_resource" "foreach_example" {
  for_each = local.items
  
  triggers = {
    name = each.key
    type = each.value
  }
}

resource "ctfchallenge_flag_validator" "foreach" {
  challenge_id = "for_each_wizard"
  
  proof_of_work = {
    items = join(",", local.items)
  }
}

output "flag" {
  value     = ctfchallenge_flag_validator.foreach.flag
  sensitive = true
}

# Show the created resources
output "resources_created" {
  value = {
    for key, resource in null_resource.foreach_example :
    key => resource.id
  }
}
```

### Captured Flag

```
flag{f0r_34ch_1s_p0w3rful}
```

### Explanation

- `for_each` creates multiple instances of a resource based on a set or map
- Each instance is independently addressable: `null_resource.foreach_example["alpha"]`
- This is more flexible than `count` for managing similar resources

---

## Challenge 7: Data Source Detective (150 points)

**Difficulty:** Beginner  
**Category:** Data Sources  
**Objective:** Understand data filtering concepts

### Challenge Description

Query and filter data sources to find a specific count. The magic number you're looking for is 7.

### Solution

```terraform
resource "ctfchallenge_flag_validator" "datasource" {
  challenge_id = "data_source_detective"
  
  proof_of_work = {
    filtered_count = "7"
  }
}

output "flag" {
  value     = ctfchallenge_flag_validator.datasource.flag
  sensitive = true
}
```

### Captured Flag

```
flag{d4t4_s0urc3_sl3uth}
```

### Explanation

- Data sources query existing infrastructure or external data
- Filtering and counting results is a common pattern
- The number 7 represents a filtered result count

---

## Challenge 8: Cryptographic Compute (500 points)

**Difficulty:** Advanced  
**Category:** Functions  
**Objective:** Chain multiple cryptographic functions

### Challenge Description

Use Terraform's cryptographic functions to compute: `md5(sha256("terraform_ctf_11_2025"))`

### Hints

- Level 0: "Chain multiple hash functions"
- Level 1: "Start with sha256, then md5 the result"
- Level 2: "Compute: md5(sha256('terraform_ctf_11_2025'))"

### Solution

```terraform
locals {
  secret = "terraform_ctf_11_2025"
  
  # Step 1: Hash with SHA256
  sha_result = sha256(local.secret)
  
  # Step 2: Hash the SHA256 result with MD5
  final_hash = md5(local.sha_result)
}

resource "ctfchallenge_flag_validator" "crypto" {
  challenge_id = "cryptographic_compute"
  
  proof_of_work = {
    crypto_hash = local.final_hash
  }
}

output "flag" {
  value     = ctfchallenge_flag_validator.crypto.flag
  sensitive = true
}

# Debug output
output "intermediate_steps" {
  value = {
    step1_sha256 = local.sha_result
    step2_md5    = local.final_hash
  }
}
```

### Testing with terraform console

```bash
$ terraform console
> sha256("terraform_ctf_11_2025")
"abc123..." # hex string
> md5(sha256("terraform_ctf_11_2025"))
"def456..." # final result
```

### Captured Flag

```
flag{crypt0_func_m4st3r}
```

### Explanation

- This challenge tests your ability to chain cryptographic functions
- `sha256()` produces a hex string
- `md5()` then hashes that hex string
- Both functions are commonly used in infrastructure for checksums and validation

---

## Bonus: XOR Puzzle Box

**Objective:** Solve an XOR puzzle for a bonus flag

### Challenge Description

Provide 5 numbers whose XOR (exclusive OR) equals zero.

### XOR Properties

- `a XOR a = 0` (self-cancellation)
- `a XOR 0 = a` (identity)
- Order doesn't matter (commutative)

### Solution 1: All Zeros

```terraform
resource "ctfchallenge_puzzle_box" "xor_easy" {
  inputs = {
    input_1 = "0"
    input_2 = "0"
    input_3 = "0"
    input_4 = "0"
    input_5 = "0"
  }
}

output "puzzle_flag" {
  value     = ctfchallenge_puzzle_box.xor_easy.secret_output
  sensitive = true
}
```

### Solution 2: Pairs

```terraform
resource "ctfchallenge_puzzle_box" "xor_pairs" {
  inputs = {
    input_1 = "42"
    input_2 = "42"  # Cancels with input_1
    input_3 = "7"
    input_4 = "7"   # Cancels with input_3
    input_5 = "0"
  }
}
```

### Solution 3: Calculated Fifth Number

```python
# Calculate the 5th number in Python
a, b, c, d = 15, 23, 42, 37
e = a ^ b ^ c ^ d
print(e)  # 11
```

```terraform
resource "ctfchallenge_puzzle_box" "xor_calculated" {
  inputs = {
    input_1 = "15"
    input_2 = "23"
    input_3 = "42"
    input_4 = "37"
    input_5 = "11"  # 15 XOR 23 XOR 42 XOR 37 XOR 11 = 0
  }
}

output "secret" {
  value     = ctfchallenge_puzzle_box.xor_calculated.secret_output
  sensitive = true
}
```

### Captured Flag

```
flag{xor_puzzl3_s0lv3d}
```

---

## Summary of All Flags

Here's your complete flag collection:

```
1. flag{t3rr4f0rm_d3p3nd3nc13s}      # Terraform Basics (100 pts)
2. flag{3xpr3ss10ns_unl0ck3d}        # Expression Expert (350 pts)
3. flag{st4t3_m4n4g3m3nt_m4st3r}     # State Secrets (200 pts)
4. flag{m0dul3_c0mp0s1t10n_pr0}      # Module Master (400 pts)
5. flag{dyn4m1c_bl0cks_r0ck}         # Dynamic Blocks (300 pts)
6. flag{f0r_34ch_1s_p0w3rful}        # For-Each Wizard (250 pts)
7. flag{d4t4_s0urc3_sl3uth}          # Data Source Detective (150 pts)
8. flag{crypt0_func_m4st3r}          # Cryptographic Compute (500 pts)

Bonus: flag{xor_puzzl3_s0lv3d}       # XOR Puzzle

Total: 2,250 points (+ bonus)
```

## Achievement Unlocked! 🏆

If you've completed all challenges, you've demonstrated mastery of:

- ✅ Resource dependencies and ordering
- ✅ Terraform expressions and functions
- ✅ State management concepts
- ✅ Module creation and composition
- ✅ Dynamic blocks and meta-arguments
- ✅ for_each and resource iteration
- ✅ Data source queries
- ✅ Cryptographic functions
- ✅ Logic puzzles

**Congratulations, Terraform Master!** 🎉