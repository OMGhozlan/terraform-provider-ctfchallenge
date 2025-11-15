package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceValidationHelper() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceValidationHelperRead,
		Schema: map[string]*schema.Schema{
			"resource_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "ID of the resource to validate",
			},
			"validation_type": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "basic",
				Description: "Type of validation (basic, strict, comprehensive)",
			},
			// Computed attributes for validation
			"is_valid": {
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Whether the resource passes validation",
			},
			"validation_score": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Validation score (0-100)",
			},
			"status": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Validation status",
			},
			"recommendations": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Validation recommendations",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func dataSourceValidationHelperRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	resourceID := d.Get("resource_id").(string)
	validationType := d.Get("validation_type").(string)

	// Simulate validation
	score := 85
	isValid := true
	status := "passed"

	if validationType == "strict" {
		score = 90
	} else if validationType == "comprehensive" {
		score = 95
	}

	recommendations := []interface{}{
		"Consider adding more validation rules",
		"Document your validation logic",
		"Add error recovery mechanisms",
	}

	d.Set("is_valid", isValid)
	d.Set("validation_score", score)
	d.Set("status", status)
	d.Set("recommendations", recommendations)
	d.SetId(fmt.Sprintf("validation-%s", resourceID))

	return diags
}
