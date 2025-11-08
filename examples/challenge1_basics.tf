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

# --- YOUR SOLUTION HERE ---
# Create at least 3 dependent resources

resource "null_resource" "first" {}

resource "null_resource" "second" {
  depends_on = [null_resource.first]
}

resource "null_resource" "third" {
  depends_on = [null_resource.second]
}

# --- END SOLUTION ---

# Submit your solution
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

output "points_earned" {
  value = ctfchallenge_flag_validator.basics.points
}