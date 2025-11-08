terraform {
  required_providers {
    ctfchallenge = {
      source = "github.com/omghozlan/ctfchallenge"
    }
  }
}

provider "ctfchallenge" {}

# List all challenges
data "ctfchallenge_list" "all" {}

output "all_challenges" {
  value = data.ctfchallenge_list.all.challenges
}

output "total_points_available" {
  value = data.ctfchallenge_list.all.total_points
}

# List only beginner challenges
data "ctfchallenge_list" "beginner" {
  difficulty = "beginner"
}

output "beginner_challenges" {
  value = data.ctfchallenge_list.beginner.challenges
}