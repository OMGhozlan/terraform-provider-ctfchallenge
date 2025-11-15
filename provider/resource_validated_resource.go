package provider

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceValidatedResource() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceValidatedResourceCreate,
		ReadContext:   resourceValidatedResourceRead,
		UpdateContext: resourceValidatedResourceUpdate,
		DeleteContext: resourceValidatedResourceDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Resource name",
			},
			"required_value": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A required value that will be validated",
			},
			"optional_value": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "An optional value",
			},
			"validation_rules": {
				Type:        schema.TypeList,
				Optional:    true,
				MaxItems:    1,
				Description: "Validation rules for this resource",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"has_precondition": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Whether preconditions are defined",
						},
						"has_postcondition": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Whether postconditions are defined",
						},
						"validates_input": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Whether input is validated",
						},
						"validates_output": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Whether output is validated",
						},
					},
				},
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Resource state after creation",
			},
			"validated": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether validation passed",
			},
			"validation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "When validation occurred",
			},
			// These represent attributes that would be validated with 'self'
			"computed_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Computed identifier",
			},
			"solved": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether resource is in solved state",
			},
			"quality_score": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Quality score of the resource",
			},
		},
	}
}

func resourceValidatedResourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	name := d.Get("name").(string)
	requiredValue := d.Get("required_value").(string)

	// Simulate validation
	validated := len(requiredValue) > 0

	// Set computed values that can be referenced with 'self'
	computedID := fmt.Sprintf("validated-%s-%d", name, time.Now().Unix())
	d.Set("computed_id", computedID)
	d.Set("state", "active")
	d.Set("validated", validated)
	d.Set("validation_timestamp", time.Now().UTC().Format(time.RFC3339))
	d.Set("solved", validated)
	d.Set("quality_score", len(requiredValue)*10)

	if validated {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Resource Validated Successfully",
			Detail:   fmt.Sprintf("Resource '%s' passed all validation checks", name),
		})
	}

	d.SetId(computedID)
	return diags
}

func resourceValidatedResourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func resourceValidatedResourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return resourceValidatedResourceCreate(ctx, d, m)
}

func resourceValidatedResourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	d.SetId("")
	return nil
}
