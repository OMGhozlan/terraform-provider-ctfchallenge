package challenges

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

// Challenge defines the structure for each CTF challenge.
type Challenge struct {
	ID          string
	Name        string
	Description string
	Points      int
	Flag        string
	Difficulty  string
	Category    string
	Validator   func(input map[string]interface{}) (bool, string, error)
}

// Challenges holds all available challenges.
var Challenges = map[string]*Challenge{
	"terraform_basics": {
		ID:          "terraform_basics",
		Name:        "Terraform Basics",
		Description: "Understand resource dependencies and outputs",
		Points:      100,
		Flag:        "flag{t3rr4f0rm_d3p3nd3nc13s}",
		Difficulty:  "beginner",
		Category:    "fundamentals",
		Validator:   validateBasics,
	},
	"expression_expert": {
		ID:          "expression_expert",
		Name:        "Expression Expert",
		Description: "Master Terraform expressions and functions",
		Points:      350,
		Flag:        "flag{3xpr3ss10ns_unl0ck3d}",
		Difficulty:  "intermediate",
		Category:    "expressions",
		Validator:   validateExpressions,
	},
	"state_secrets": {
		ID:          "state_secrets",
		Name:        "State Secrets",
		Description: "Understand Terraform state management",
		Points:      200,
		Flag:        "flag{st4t3_m4n4g3m3nt_m4st3r}",
		Difficulty:  "beginner",
		Category:    "state",
		Validator:   validateState,
	},
	"module_master": {
		ID:          "module_master",
		Name:        "Module Master",
		Description: "Create and use Terraform modules effectively",
		Points:      400,
		Flag:        "flag{m0dul3_c0mp0s1t10n_pr0}",
		Difficulty:  "advanced",
		Category:    "modules",
		Validator:   validateModules,
	},
	"dynamic_blocks": {
		ID:          "dynamic_blocks",
		Name:        "Dynamic Blocks Challenge",
		Description: "Master dynamic block generation",
		Points:      300,
		Flag:        "flag{dyn4m1c_bl0cks_r0ck}",
		Difficulty:  "intermediate",
		Category:    "advanced-syntax",
		Validator:   validateDynamicBlocks,
	},
	"for_each_wizard": {
		ID:          "for_each_wizard",
		Name:        "For-Each Wizard",
		Description: "Use for_each to manage multiple resources elegantly",
		Points:      250,
		Flag:        "flag{f0r_34ch_1s_p0w3rful}",
		Difficulty:  "intermediate",
		Category:    "loops",
		Validator:   validateForEach,
	},
	"data_source_detective": {
		ID:          "data_source_detective",
		Name:        "Data Source Detective",
		Description: "Query and filter data sources effectively",
		Points:      150,
		Flag:        "flag{d4t4_s0urc3_sl3uth}",
		Difficulty:  "beginner",
		Category:    "data-sources",
		Validator:   validateDataSource,
	},
	"cryptographic_compute": {
		ID:          "cryptographic_compute",
		Name:        "Cryptographic Compute",
		Description: "Use Terraform's cryptographic functions",
		Points:      500,
		Flag:        "flag{crypt0_func_m4st3r}",
		Difficulty:  "advanced",
		Category:    "functions",
		Validator:   validateCrypto,
	},
}

// validateBasics checks for a solution to the "Terraform Basics" challenge.
// Players must create at least 3 dependent resources.
func validateBasics(input map[string]interface{}) (bool, string, error) {
	correctFlag := "flag{t3rr4f0rm_d3p3nd3nc13s}"

	if depsStr, ok := input["dependencies"].(string); ok {
		// Count comma-separated values or JSON array items
		deps := strings.Split(depsStr, ",")
		if len(deps) >= 3 {
			return true, correctFlag, nil
		}
		return false, "", fmt.Errorf("create at least 3 dependent resources (found %d)", len(deps))
	}
	return false, "", fmt.Errorf("provide 'dependencies' as a comma-separated string in proof_of_work")
}

// validateExpressions checks for the "Expression Expert" challenge.
// Players must compute: base64(sha256("terraform" + "expressions" + "rock"))
// Note: Terraform's sha256() returns hex, so we accept base64 of that hex string
func validateExpressions(input map[string]interface{}) (bool, string, error) {
	correctFlag := "flag{3xpr3ss10ns_unl0ck3d}"

	if result, ok := input["computed_value"].(string); ok {
		expected := "terraformexpressionsrock"

		// Terraform's sha256() returns hex string, not raw bytes
		// So we compute what Terraform would produce:
		hash := sha256.Sum256([]byte(expected))
		hexString := hex.EncodeToString(hash[:])                            // This matches Terraform's sha256() output
		expectedB64 := base64.StdEncoding.EncodeToString([]byte(hexString)) // base64 of hex string

		if result == expectedB64 {
			return true, correctFlag, nil
		}
		return false, "", fmt.Errorf("computed value doesn't match expected hash")
	}
	return false, "", fmt.Errorf("provide 'computed_value' in proof_of_work")
}

// validateState checks understanding of state management.
// Players must provide a resource count that matches a specific pattern.
func validateState(input map[string]interface{}) (bool, string, error) {
	correctFlag := "flag{st4t3_m4n4g3m3nt_m4st3r}"

	if countStr, ok := input["resource_count"].(string); ok {
		count, err := strconv.Atoi(countStr)
		if err != nil {
			return false, "", fmt.Errorf("resource_count must be a number")
		}
		// The magic number is 42 (a nod to the answer to everything)
		if count == 42 {
			return true, correctFlag, nil
		}
		return false, "", fmt.Errorf("the state reveals a different count is needed")
	}
	return false, "", fmt.Errorf("provide 'resource_count' in proof_of_work")
}

// validateModules checks if player created outputs from modules correctly.
func validateModules(input map[string]interface{}) (bool, string, error) {
	correctFlag := "flag{m0dul3_c0mp0s1t10n_pr0}"

	if moduleOutput, ok := input["module_output"].(string); ok {
		// Check if output contains evidence of module composition
		if strings.Contains(moduleOutput, "module.") && len(moduleOutput) > 20 {
			return true, correctFlag, nil
		}
		return false, "", fmt.Errorf("module output doesn't show proper composition")
	}
	return false, "", fmt.Errorf("provide 'module_output' in proof_of_work")
}

// validateDynamicBlocks checks if player used dynamic blocks correctly.
func validateDynamicBlocks(input map[string]interface{}) (bool, string, error) {
	correctFlag := "flag{dyn4m1c_bl0cks_r0ck}"

	if blockCount, ok := input["dynamic_block_count"].(string); ok {
		count, err := strconv.Atoi(blockCount)
		if err != nil {
			return false, "", fmt.Errorf("dynamic_block_count must be a number")
		}
		if count >= 5 {
			return true, correctFlag, nil
		}
		return false, "", fmt.Errorf("generate at least 5 dynamic blocks")
	}
	return false, "", fmt.Errorf("provide 'dynamic_block_count' in proof_of_work")
}

// validateForEach validates the for_each challenge.
// Players must create resources for a specific set of items.
func validateForEach(input map[string]interface{}) (bool, string, error) {
	correctFlag := "flag{f0r_34ch_1s_p0w3rful}"

	if items, ok := input["items"].(string); ok {
		// Expecting a JSON-like representation of the set/map
		requiredItems := []string{"alpha", "beta", "gamma", "delta"}
		matchCount := 0
		for _, item := range requiredItems {
			if strings.Contains(items, item) {
				matchCount++
			}
		}
		if matchCount == len(requiredItems) {
			return true, correctFlag, nil
		}
		return false, "", fmt.Errorf("create for_each resources with items: alpha, beta, gamma, delta")
	}
	return false, "", fmt.Errorf("provide 'items' in proof_of_work")
}

// validateDataSource checks data source filtering skills.
func validateDataSource(input map[string]interface{}) (bool, string, error) {
	correctFlag := "flag{d4t4_s0urc3_sl3uth}"

	if filteredCount, ok := input["filtered_count"].(string); ok {
		count, err := strconv.Atoi(filteredCount)
		if err != nil {
			return false, "", fmt.Errorf("filtered_count must be a number")
		}
		// Specific count from filtering a mock data source
		if count == 7 {
			return true, correctFlag, nil
		}
		return false, "", fmt.Errorf("incorrect filter result")
	}
	return false, "", fmt.Errorf("provide 'filtered_count' in proof_of_work")
}

// validateCrypto validates cryptographic function mastery.
// Players must provide: md5(sha256("terraform_ctf_2024"))
func validateCrypto(input map[string]interface{}) (bool, string, error) {
	correctFlag := "flag{crypt0_func_m4st3r}"

	if hash, ok := input["crypto_hash"].(string); ok {
		secret := "terraform_ctf_2024"
		sha := sha256.Sum256([]byte(secret))
		md5Hash := md5.Sum(sha[:])
		expected := hex.EncodeToString(md5Hash[:])

		if hash == expected {
			return true, correctFlag, nil
		}
		return false, "", fmt.Errorf("cryptographic hash doesn't match expected value")
	}
	return false, "", fmt.Errorf("provide 'crypto_hash' in proof_of_work")
}

// GetHint provides hints for a given challenge at different levels.
func GetHint(challengeID string, level int) string {
	hints := map[string][]string{
		"terraform_basics": {
			"Start by creating resources with depends_on",
			"You need exactly 3 resources in a dependency chain",
			"Pass the resource IDs as a comma-separated string in dependencies",
		},
		"expression_expert": {
			"Look at Terraform's hash and encoding functions",
			"Combine sha256() and base64encode() functions",
			"Concatenate the strings: 'terraform' + 'expressions' + 'rock', hash with sha256, then base64encode",
		},
		"state_secrets": {
			"The answer to life, the universe, and everything...",
			"Douglas Adams knew the answer",
			"It's 42 resources",
		},
		"module_master": {
			"Create a module with outputs",
			"Reference module outputs using module.<name>.<output>",
			"Your module_output should show the full module reference path",
		},
		"dynamic_blocks": {
			"Use the dynamic block with for_each",
			"Generate blocks from a list or map",
			"Create at least 5 dynamic blocks using count or for_each inside the dynamic block",
		},
		"for_each_wizard": {
			"Use for_each with a set or map",
			"The required items are Greek letters",
			"Create resources for: alpha, beta, gamma, delta",
		},
		"data_source_detective": {
			"Use a data source and count the results",
			"Filter or process the data source output",
			"The expected filtered count is 7",
		},
		"cryptographic_compute": {
			"Chain multiple hash functions",
			"Start with sha256, then md5 the result",
			"Compute: md5(sha256('terraform_ctf_2024'))",
		},
	}

	if challengeHints, exists := hints[challengeID]; exists {
		if level < len(challengeHints) {
			return challengeHints[level]
		}
		return "No more hints available for this challenge"
	}
	return "No hints available for this challenge"
}

// GetAllChallengeIDs returns a list of all challenge IDs.
func GetAllChallengeIDs() []string {
	ids := make([]string, 0, len(Challenges))
	for id := range Challenges {
		ids = append(ids, id)
	}
	return ids
}

// ValidatePuzzleInput validates input for the puzzle box resource.
func ValidatePuzzleInput(inputs map[string]interface{}) (bool, string) {
	// Puzzle: XOR all numbers and check if result equals 0
	numbers := []int{}

	for i := 1; i <= 5; i++ {
		key := fmt.Sprintf("input_%d", i)
		if valStr, ok := inputs[key].(string); ok {
			val, err := strconv.Atoi(valStr)
			if err == nil {
				numbers = append(numbers, val)
			}
		}
	}

	if len(numbers) != 5 {
		return false, "Provide exactly 5 numbers (input_1 through input_5)"
	}

	xorResult := 0
	for _, num := range numbers {
		xorResult ^= num
	}

	if xorResult == 0 {
		return true, "Puzzle solved! XOR of all inputs equals zero."
	}

	return false, fmt.Sprintf("XOR result: %d (must be 0)", xorResult)
}
