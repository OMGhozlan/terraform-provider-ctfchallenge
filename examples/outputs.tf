output "resource_id" {
  value = null_resource.example.id
}

output "composition_proof" {
  value = "module.${var.name}.resource_id"
}