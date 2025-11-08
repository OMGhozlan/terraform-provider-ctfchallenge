---
page_title: "ctfchallenge_puzzle_box Resource - ctfchallenge"
subcategory: ""
description: |-
  Solve logic puzzles to discover bonus flags.
---

# ctfchallenge_puzzle_box (Resource)

The `puzzle_box` resource presents logic puzzles that can be solved for bonus flags. Each puzzle has specific rules that must be satisfied.

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

output "secret_flag" {
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
}

resource "ctfchallenge_puzzle_box" "dynamic_puzzle" {
  inputs = var.puzzle_inputs
}
```

## Schema

### Required

- `inputs` (Map of String) The puzzle inputs. The specific keys required depend on the puzzle type. All values must be strings.

### Read-Only

- `id` (String) The unique identifier for this puzzle attempt.
- `solved` (Boolean) Whether the puzzle was successfully solved.
- `message` (String) A message describing the puzzle result or hint.
- `secret_output` (String, Sensitive) The secret flag revealed when the puzzle is solved.

## XOR Puzzle Rules

The XOR puzzle requires you to provide 5 numbers (`input_1` through `input_5`) whose XOR (exclusive OR) equals zero.

### XOR Basics
- `a XOR a = 0`
- `a XOR 0 = a`
- `a XOR b XOR c = (a XOR b) XOR c`

### Example Solutions

**Solution 1: All zeros**
```hcl
inputs = {
  input_1 = "0"
  input_2 = "0"
  input_3 = "0"
  input_4 = "0"
  input_5 = "0"
}
```

**Solution 2: Pairs that cancel**
```hcl
inputs = {
  input_1 = "42"
  input_2 = "42"
  input_3 = "7"
  input_4 = "7"
  input_5 = "0"
}
```

**Solution 3: Working backward**
```hcl
# If you want specific first 4 numbers, calculate the 5th:
# input_5 = input_1 XOR input_2 XOR input_3 XOR input_4
inputs = {
  input_1 = "15"
  input_2 = "23"
  input_3 = "42"
  input_4 = "37"
  input_5 = "11"  # Calculated: 15 XOR 23 XOR 42 XOR 37
}
```

## Import

Puzzle box resources are ephemeral and cannot be imported.