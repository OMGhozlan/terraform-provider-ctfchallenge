---
page_title: "ctfchallenge_puzzle_box Resource - ctfchallenge"
subcategory: ""
description: |-
  Solve logic puzzles to discover bonus flags.
---

# ctfchallenge_puzzle_box (Resource)

The `puzzle_box` resource presents logic puzzles that can be solved for bonus flags. Currently features an XOR puzzle where you must find 5 numbers whose XOR (exclusive OR) equals zero.

## Example Usage

```terraform
# XOR Puzzle: Find 5 numbers whose XOR equals 0
resource "ctfchallenge_puzzle_box" "xor_puzzle" {
  inputs = {
    input_1 = "15"
    input_2 = "23"
    input_3 = "42"
    input_4 = "37"
    input_5 = "11"  # 15 XOR 23 XOR 42 XOR 37 XOR 11 = 0
  }
}

output "puzzle_solved" {
  value = ctfchallenge_puzzle_box.xor_puzzle.solved
}

output "puzzle_message" {
  value = ctfchallenge_puzzle_box.xor_puzzle.message
}

output "bonus_flag" {
  value     = ctfchallenge_puzzle_box.xor_puzzle.secret_output
  sensitive = true
}
```

## Example with Variables

```terraform
variable "puzzle_inputs" {
  type = map(string)
  default = {
    input_1 = "15"
    input_2 = "23"
    input_3 = "42"
    input_4 = "37"
    input_5 = "11"
  }
  description = "XOR puzzle inputs"
}

resource "ctfchallenge_puzzle_box" "dynamic_puzzle" {
  inputs = var.puzzle_inputs
}

output "result" {
  value = {
    solved  = ctfchallenge_puzzle_box.dynamic_puzzle.solved
    message = ctfchallenge_puzzle_box.dynamic_puzzle.message
  }
}
```

## Schema

### Required

- `inputs` (Map of String) The puzzle inputs. For the XOR puzzle, provide `input_1` through `input_5`. All values must be strings representing integers.

### Read-Only

- `id` (String) The unique identifier for this puzzle attempt.
- `solved` (Boolean) Whether the puzzle was successfully solved.
- `message` (String) A message describing the puzzle result or hint.
- `secret_output` (String, Sensitive) The secret flag revealed when the puzzle is solved. Empty if unsolved.

## XOR Puzzle Rules

The XOR puzzle requires you to provide 5 numbers (`input_1` through `input_5`) whose XOR (exclusive OR) equals zero.

### XOR Basics

```
Truth Table:
0 XOR 0 = 0
0 XOR 1 = 1
1 XOR 0 = 1
1 XOR 1 = 1

Properties:
• a XOR a = 0 (self-cancellation)
• a XOR 0 = a (identity)
• a XOR b = b XOR a (commutative)
• (a XOR b) XOR c = a XOR (b XOR c) (associative)
```

### Example Solutions

#### Solution 1: All Zeros (Easiest)

```terraform
resource "ctfchallenge_puzzle_box" "simple" {
  inputs = {
    input_1 = "0"
    input_2 = "0"
    input_3 = "0"
    input_4 = "0"
    input_5 = "0"
  }
}
```

### Result: `0 XOR 0 XOR 0 XOR 0 XOR 0 = 0` ✓

#### Solution 2: Pairs that Cancel

```terraform
resource "ctfchallenge_puzzle_box" "pairs" {
  inputs = {
    input_1 = "42"
    input_2 = "42"  # Cancels with input_1
    input_3 = "7"
    input_4 = "7"   # Cancels with input_3
    input_5 = "0"   # Identity element
  }
}
```

### Result: `42 XOR 42 XOR 7 XOR 7 XOR 0 = 0 XOR 0 XOR 0 = 0` ✓

#### Solution 3: Calculated Fifth Number (Advanced)

Choose any 4 numbers, then calculate the 5th:

**Using Python:**
```python
a, b, c, d = 15, 23, 42, 37
e = a ^ b ^ c ^ d
print(f"The 5th number is: {e}")
# Output: 11
```

**Using JavaScript:**
```javascript
let [a, b, c, d] = [15, 23, 42, 37];
let e = a ^ b ^ c ^ d;
console.log(`The 5th number is: ${e}`);
// Output: 11
```

**Using Go:**
```go
package main
import "fmt"

func main() {
    a, b, c, d := 15, 23, 42, 37
    e := a ^ b ^ c ^ d
    fmt.Printf("The 5th number is: %d\n", e)
    // Output: 11
}
```

**In Terraform:**
```terraform
resource "ctfchallenge_puzzle_box" "calculated" {
  inputs = {
    input_1 = "15"
    input_2 = "23"
    input_3 = "42"
    input_4 = "37"
    input_5 = "11"  # Calculated: 15 XOR 23 XOR 42 XOR 37 = 11
  }
}
```

### Result: `15 XOR 23 XOR 42 XOR 37 XOR 11 = 0` ✓

## Validation Results

### Successful Solution

```hcl
solved        = true
message       = "Puzzle solved! XOR of all inputs equals zero."
secret_output = "flag{xor_puzzl3_s0lv3d}"
```

### Failed Solution

```hcl
solved        = false
message       = "XOR result: 23 (must be 0)"
secret_output = ""
```

### Invalid Input

```hcl
solved        = false
message       = "Provide exactly 5 numbers (input_1 through input_5)"
secret_output = ""
```

## Viewing the Bonus Flag

After solving the puzzle:

```bash
terraform output -raw bonus_flag
```

Output:
```
flag{xor_puzzl3_s0lv3d}
```

## Complete Example

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

# Solve the XOR puzzle
resource "ctfchallenge_puzzle_box" "xor" {
  inputs = {
    input_1 = "15"
    input_2 = "23"
    input_3 = "42"
    input_4 = "37"
    input_5 = "11"
  }
}

# Display results
output "puzzle_status" {
  value = {
    solved  = ctfchallenge_puzzle_box.xor.solved
    message = ctfchallenge_puzzle_box.xor.message
  }
}

output "bonus_flag" {
  value     = ctfchallenge_puzzle_box.xor.secret_output
  sensitive = true
}
```

Apply and capture the flag:

```bash
terraform apply
terraform output -raw bonus_flag
```

## XOR Calculator Helper

Create a helper script to find valid combinations:

**xor_solver.py:**
```python
#!/usr/bin/env python3
import random

def find_xor_solution():
    # Generate 4 random numbers
    nums = [random.randint(1, 100) for _ in range(4)]
    
    # Calculate the 5th
    fifth = nums[0] ^ nums[1] ^ nums[2] ^ nums[3]
    
    # Verify
    result = nums[0] ^ nums[1] ^ nums[2] ^ nums[3] ^ fifth
    
    print(f"input_1 = \"{nums[0]}\"")
    print(f"input_2 = \"{nums[1]}\"")
    print(f"input_3 = \"{nums[2]}\"")
    print(f"input_4 = \"{nums[3]}\"")
    print(f"input_5 = \"{fifth}\"")
    print(f"\nVerification: {result} (should be 0)")

if __name__ == "__main__":
    find_xor_solution()
```

Run it:
```bash
chmod +x xor_solver.py
./xor_solver.py
```

## Tips

1. **Start simple** - All zeros is a valid solution!
2. **Use pairs** - Any number XORed with itself equals 0
3. **Calculate the 5th** - Pick 4 numbers, compute the 5th
4. **Test externally** - Use Python/JavaScript to verify before applying
5. **All values must be strings** - Numbers must be quoted

## Common Errors

### Missing Input

```
Error: Provide exactly 5 numbers (input_1 through input_5)
```

**Solution:** Ensure all 5 inputs are provided.

### Wrong Data Type

```
Error: Incorrect value type
```

**Solution:** All inputs must be strings, not numbers:
```terraform
# Wrong
inputs = { input_1 = 42 }

# Correct
inputs = { input_1 = "42" }
```

### Non-Zero Result

```
Puzzle not solved: XOR result: 15 (must be 0)
```

**Solution:** Recalculate your numbers to ensure XOR equals 0.

## Import

Puzzle box resources are ephemeral and cannot be imported.

## See Also

- [Advanced Challenges Guide](../guides/advanced-challenges.md)
- [Challenge Walkthrough](../guides/challenge-walkthrough.md)