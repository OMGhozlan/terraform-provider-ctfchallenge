terraform {
  required_providers {
    ctfchallenge = {
      source = "github.com/omghozlan/ctfchallenge"
    }
  }
}

provider "ctfchallenge" {}

# Puzzle: Find 5 numbers whose XOR equals 0
resource "ctfchallenge_puzzle_box" "xor_puzzle" {
  inputs = {
    input_1 = "15"
    input_2 = "23"
    input_3 = "42"
    input_4 = "37"
    input_5 = "11" # 15 XOR 23 XOR 42 XOR 37 XOR 11 = 0
  }
}

output "puzzle_solved" {
  value = ctfchallenge_puzzle_box.xor_puzzle.solved
}

output "puzzle_message" {
  value = ctfchallenge_puzzle_box.xor_puzzle.message
}

output "secret" {
  value     = ctfchallenge_puzzle_box.xor_puzzle.secret_output
  sensitive = true
}