---
page_title: "Challenge Walkthrough"
subcategory: "Guides"
description: |-
  Detailed walkthroughs and solutions for all challenges.
---

# Challenge Walkthrough

This guide provides detailed solutions for each challenge. Try to solve them yourself first!

## Challenge 1: Terraform Basics (100 points)

**Objective:** Understand resource dependencies

**Solution:**

```terraform
resource "null_resource" "step1" {}
resource "null_resource" "step2" { depends_on = [null_resource.step1] }
resource "null_resource" "step3" { depends_on = [null_resource.step2] }

resource "ctfchallenge_flag_validator" "basics" {
  challenge_id = "terraform_basics"
  flag         = "flag{t3rr4f0rm_d3p3nd3nc13s}"
  
  proof_of_work = {
    dependencies = "${null_resource.step1.id},${null_resource.step2.id},${null_resource.step3.id}"
  }
}
```

---

## Challenge 2: Expression Expert (350 points)

**Objective:** Master Terraform expressions and functions

**Hint:** Concatenate "terraform" + "expressions" + "rock", apply sha256, then base64encode

**Solution:**

```terraform
locals {
  # Concatenate the strings
  combined = "terraformexpressionsrock"
  
  # Hash with SHA256
  hashed = sha256(local.combined)
  
  # Encode with base64
  result = base64encode(local.hashed)
}

resource "ctfchallenge_flag_validator" "expressions" {
  challenge_id = "expression_expert"
  flag         = "flag{3xpr3ss10ns_unl0ck3d}"
  
  proof_of_work = {
    computed_value = local.result
  }
}
```

---

## Challenge 3: State Secrets (200 points)

**Objective:** Understand the "magic number" in Terraform

**Hint:** The answer to life, the universe, and everything

**Solution:**

```terraform
resource "ctfchallenge_flag_validator" "state" {
  challenge_id = "state_secrets"
  flag         = "flag{st4t3_m4n4g3m3nt_m4st3r}"
  
  proof_of_work = {
    resource_count = "42"  # Douglas Adams reference!
  }
}
```

---

## Challenge 4: Module Master (400 points)

**Objective:** Create and use Terraform modules

**Solution:**

First, create a module in `modules/example/main.tf`:

```terraform
output "message" {
  value = "Hello from module"
}
```

Then in your root module:

```terraform
module "example" {
  source = "./modules/example"
}

resource "ctfchallenge_flag_validator" "modules" {
  challenge_id = "module_master"
  flag         = "flag{m0dul3_c0mp0s1t10n_pr0}"
  
  proof_of_work = {
    module_output = module.example.message
  }
}
```

---

## Challenge 5: Dynamic Blocks (300 points)

**Objective:** Generate multiple blocks dynamically

**Solution:**

```terraform
locals {
  ports = [80, 443, 8080, 8443, 3000]
}

resource "null_resource" "dynamic_example" {
  count = 1
  
  triggers = {
    for idx, port in local.ports :
    "port_${idx}" => port
  }
}

resource "ctfchallenge_flag_validator" "dynamic" {
  challenge_id = "dynamic_blocks"
  flag         = "flag{dyn4m1c_bl0cks_r0ck}"
  
  proof_of_work = {
    dynamic_block_count = tostring(length(local.ports))
  }
}
```

---

## Challenge 6: For-Each Wizard (250 points)

**Objective:** Use for_each to create multiple resources

**Solution:**

```terraform
locals {
  items = toset(["alpha", "beta", "gamma", "delta"])
}

resource "null_resource" "foreach_example" {
  for_each = local.items
  
  triggers = {
    name = each.key
  }
}

resource "ctfchallenge_flag_validator" "foreach" {
  challenge_id = "for_each_wizard"
  flag         = "flag{f0r_34ch_1s_p0w3rful}"
  
  proof_of_work = {
    items = join(",", local.items)
  }
}
```

---

## Challenge 7: Data Source Detective (150 points)

**Objective:** Filter data effectively

**Solution:**

```terraform
resource "ctfchallenge_flag_validator" "datasource" {
  challenge_id = "data_source_detective"
  flag         = "flag{d4t4_s0urc3_sl3uth}"
  
  proof_of_work = {
    filtered_count = "7"
  }
}
```

---

## Challenge 8: Cryptographic Compute (500 points)

**Objective:** Chain cryptographic functions

**Hint:** Compute md5(sha256("terraform_ctf_2024"))

**Solution:**

```terraform
locals {
  secret = "terraform_ctf_2024"
  sha    = sha256(local.secret)
  result = md5(local.sha)
}

resource "ctfchallenge_flag_validator" "crypto" {
  challenge_id = "cryptographic_compute"
  flag         = "flag{crypt0_func_m4st3r}"
  
  proof_of_work = {
    crypto_hash = local.result
  }
}
```

---

## Bonus: XOR Puzzle

**Solution:**

```terraform
resource "ctfchallenge_puzzle_box" "xor" {
  inputs = {
    input_1 = "15"
    input_2 = "23"
    input_3 = "42"
    input_4 = "37"
    input_5 = "11"  # 15 XOR 23 XOR 42 XOR 37 XOR 11 = 0
  }
}

output "secret_flag" {
  value     = ctfchallenge_puzzle_box.xor.secret_output
  sensitive = true
}
```

To see the secret:
```bash
terraform output -raw secret_flag
```
