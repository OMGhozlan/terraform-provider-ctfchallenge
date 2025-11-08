resource "null_resource" "example" {
  triggers = {
    name = var.name
    env  = var.environment
  }
}