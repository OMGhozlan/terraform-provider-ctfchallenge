terraform {
  required_providers {
    ctfchallenge = {
      source = "github.com/omghozlan/ctfchallenge"
    }
  }
}

provider "ctfchallenge" {
  player_name = "your-name-here"
}

# Get a hint if you need it
data "ctfchallenge_hint" "expr_hint" {
  challenge_id = "expression_expert"
  level        = 0  # 0, 1, or 2
}

output "hint" {
  value = data.ctfchallenge_hint.expr_hint.hint
}

# --- YOUR SOLUTION HERE ---
# Compute: base64(sha256("terraform" + "expressions" + "rock"))

locals {
  combined = "terraformexpressionsrock"
  hashed   = sha256(local.combined)
  encoded  = base64encode(local.hashed)
}

# --- END SOLUTION ---

resource "ctfchallenge_flag_validator" "expressions" {
  challenge_id = "expression_expert"
  flag         = "flag{3xpr3ss10ns_unl0ck3d}"

  proof_of_work = {
    computed_value = local.encoded
  }
}

output "result" {
  value = ctfchallenge_flag_validator.expressions.message
}