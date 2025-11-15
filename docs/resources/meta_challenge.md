page_title: "ctfchallenge_meta_challenge Resource - ctfchallenge"
subcategory: ""
description: |-
  Specialized resource for meta-argument focused challenges.
---

# ctfchallenge_meta_challenge (Resource)

The `meta_challenge` resource is specifically designed for challenges focused on Terraform's meta-arguments like count, for_each, depends_on, and lifecycle.

## Example Usage

### Count Challenge

```terraform
resource "ctfchallenge_puzzle_box" "counted" {
  count = 3

  inputs = {
    index = count.index
    name  = "item-${count.index}"
  }
}

resource "ctfchallenge_meta_challenge" "count_demo" {
  challenge_type = "count"
  
  configuration = {
    count_value    = "3"
    resource_ids   = join(",", ctfchallenge_puzzle_box.counted[*].id)
    uses_count_index = "true"
  }
}
```

### For_each Challenge

```terraform
locals {
  items = toset(["alpha", "beta", "gamma"])
}

resource "ctfchallenge_puzzle_box" "foreach_items" {
  for_each = local.items
  
  inputs = {
    key = each.key
  }
}

resource "ctfchallenge_meta_challenge" "foreach_demo" {
  challenge_type = "for_each"
  
  configuration = {
    foreach_type = "set"
    items        = join(",", local.items)
    uses_each    = "true"
  }
}
```

### Lifecycle Challenge

```terraform
resource "ctfchallenge_puzzle_box" "lifecycle_demo" {
  inputs = {
    value = "important"
  }
  
  lifecycle {
    create_before_destroy = true
    ignore_changes       = [inputs]
  }
}

resource "ctfchallenge_meta_challenge" "lifecycle_demo" {
  challenge_type = "lifecycle"
  
  configuration = {
    uses_create_before_destroy = "true"
    ignore_changes            = "inputs"
    lifecycle_rules_count     = "2"
  }
}
```

## Schema

### Required

- `challenge_type` (String) Type of meta-argument challenge. Valid values: `count`, `for_each`, `depends_on`, `lifecycle`.
- `configuration` (Map of String) Your challenge solution configuration. Required keys depend on challenge_type.

### Optional

- `metadata` (List of Object) Additional metadata about your solution.
  - `meta_arguments_used` (List of String) List of meta-arguments used.
  - `resource_count` (Number) Number of resources created.
  - `complexity_score` (Number) Self-assessed complexity (1-10).
  - `notes` (String) Notes about your implementation.
- `hints_used` (Number) Number of hints used for this challenge. Defaults to 0.

### Read-Only

- `id` (String) The unique identifier for this challenge attempt.
- `validation_result` (String) Result of the meta-argument validation.
- `success` (Boolean) Whether the challenge was completed successfully.

## Configuration Requirements by Type

### count

```hcl
configuration = {
  count_value       = "3"                    # Number of resources
  resource_ids      = "id1,id2,id3"         # Created resource IDs
  uses_count_index  = "true"                # Uses count.index
}
```

### for_each

```hcl
configuration = {
  foreach_type  = "set"  # or "map"
  items         = "a,b,c"
  uses_each     = "true"
}
```

### depends_on

```hcl
configuration = {
  dependency_chain_length = "3"
  resource_chain         = "first,second,third"
  uses_depends_on        = "true"
}
```

### lifecycle

```hcl
configuration = {
  uses_create_before_destroy = "true"
  ignore_changes            = "attribute_name"
  lifecycle_rules_count     = "2"
}
```

## See Also

- [Meta-Argument Challenges Guide](../guides/meta-arguments.md)
- [Flag Validator Resource](flag_validator.md)