package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns the schema for the ctfchallenge provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"player_name": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TF_CTF_PLAYER", "anonymous"),
				Description: "Your player name for the CTF",
			},
			"api_endpoint": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("TF_CTF_API", ""),
				Description: "Optional API endpoint for score tracking",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ctfchallenge_flag_validator": resourceFlagValidator(),
			"ctfchallenge_puzzle_box":     resourcePuzzleBox(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"ctfchallenge_hint":           dataSourceHint(),
			"ctfchallenge_list":           dataSourceChallengeList(),
			"ctfchallenge_challenge_info": dataSourceChallengeInfo(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

// ProviderConfig holds the provider configuration.
type ProviderConfig struct {
	PlayerName  string
	APIEndpoint string
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics

	config := &ProviderConfig{
		PlayerName:  d.Get("player_name").(string),
		APIEndpoint: d.Get("api_endpoint").(string),
	}

	return config, diags
}