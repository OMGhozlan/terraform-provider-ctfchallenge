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
	Validator   func(input map[string]interface{}) (bool, string, error) // Legacy validator
}

// ValidationResult contains the result of proof validation
type ValidationResult struct {
	Success bool
	Flag    string
	Message string
	Details []string
}

// ProofData contains all types of proof that can be submitted
type ProofData struct {
	Resources   []ResourceProof
	DataSources []DataSourceProof
	Module      *ModuleProof
	Manual      map[string]interface{}
	Source      string
}

// ResourceProof contains proof from a Terraform resource
type ResourceProof struct {
	ResourceType  string
	ResourceName  string
	Attributes    map[string]interface{}
	Lifecycle     *LifecycleConfig
	MetaArguments map[string]interface{}
}

// DataSourceProof contains proof from a Terraform data source
type DataSourceProof struct {
	DataSourceType string
	DataSourceName string
	Attributes     map[string]interface{}
	Lifecycle      *LifecycleConfig
}

// ModuleProof contains proof from a Terraform module
type ModuleProof struct {
	ModuleName        string
	InputValidations  []ValidationRule
	OutputValidations []ValidationRule
	ResourcesCount    int
}

// LifecycleConfig represents lifecycle block configuration
type LifecycleConfig struct {
	CreateBeforeDestroy bool             `json:"create_before_destroy"`
	PreventDestroy      bool             `json:"prevent_destroy"`
	IgnoreChanges       []string         `json:"ignore_changes"`
	Preconditions       []ConditionBlock `json:"preconditions"`
	Postconditions      []ConditionBlock `json:"postconditions"`
}

// ConditionBlock represents a precondition or postcondition
type ConditionBlock struct {
	Condition    string `json:"condition"`
	ErrorMessage string `json:"error_message"`
}

// ValidationRule represents an input or output validation
type ValidationRule struct {
	Type         string `json:"type"` // "precondition" or "postcondition"
	Condition    string `json:"condition"`
	ErrorMessage string `json:"error_message"`
	Target       string `json:"target"` // what's being validated
}

// Challenges holds all available challenges - INITIALIZED AS EMPTY MAP
var Challenges = make(map[string]*Challenge)

func init() {
	// Register basic challenges that don't need structure validation
	registerBasicChallenges()
}

func registerBasicChallenges() {
	Challenges["terraform_basics"] = &Challenge{
		ID:          "terraform_basics",
		Name:        "Terraform Basics",
		Description: "Understand resource dependencies and outputs",
		Points:      100,
		Flag:        "flag{t3rr4f0rm_d3p3nd3nc13s}",
		Difficulty:  "beginner",
		Category:    "fundamentals",
		Validator:   validateBasics,
	}

	Challenges["expression_expert"] = &Challenge{
		ID:          "expression_expert",
		Name:        "Expression Expert",
		Description: "Master Terraform expressions and functions",
		Points:      350,
		Flag:        "flag{3xpr3ss10ns_unl0ck3d}",
		Difficulty:  "intermediate",
		Category:    "expressions",
		Validator:   validateExpressions,
	}

	Challenges["state_secrets"] = &Challenge{
		ID:          "state_secrets",
		Name:        "State Secrets",
		Description: "Understand Terraform state management",
		Points:      200,
		Flag:        "flag{st4t3_m4n4g3m3nt_m4st3r}",
		Difficulty:  "beginner",
		Category:    "state",
		Validator:   validateState,
	}

	Challenges["module_master"] = &Challenge{
		ID:          "module_master",
		Name:        "Module Master",
		Description: "Create and use Terraform modules effectively",
		Points:      400,
		Flag:        "flag{m0dul3_c0mp0s1t10n_pr0}",
		Difficulty:  "advanced",
		Category:    "modules",
		Validator:   validateModules,
	}

	Challenges["dynamic_blocks"] = &Challenge{
		ID:          "dynamic_blocks",
		Name:        "Dynamic Blocks Challenge",
		Description: "Master dynamic block generation",
		Points:      300,
		Flag:        "flag{dyn4m1c_bl0cks_r0ck}",
		Difficulty:  "intermediate",
		Category:    "advanced-syntax",
		Validator:   validateDynamicBlocks,
	}

	Challenges["for_each_wizard"] = &Challenge{
		ID:          "for_each_wizard",
		Name:        "For-Each Wizard",
		Description: "Use for_each to manage multiple resources elegantly",
		Points:      250,
		Flag:        "flag{f0r_34ch_1s_p0w3rful}",
		Difficulty:  "intermediate",
		Category:    "loops",
		Validator:   validateForEach,
	}

	Challenges["data_source_detective"] = &Challenge{
		ID:          "data_source_detective",
		Name:        "Data Source Detective",
		Description: "Query and filter data sources effectively",
		Points:      150,
		Flag:        "flag{d4t4_s0urc3_sl3uth}",
		Difficulty:  "beginner",
		Category:    "data-sources",
		Validator:   validateDataSource,
	}

	Challenges["cryptographic_compute"] = &Challenge{
		ID:          "cryptographic_compute",
		Name:        "Cryptographic Compute",
		Description: "Use Terraform's cryptographic functions",
		Points:      500,
		Flag:        "flag{crypt0_func_m4st3r}",
		Difficulty:  "advanced",
		Category:    "functions",
		Validator:   validateCrypto,
	}
}

// ValidateProof validates the proof data and returns a result
func (c *Challenge) ValidateProof(proof *ProofData) ValidationResult {
	// If we have structured proof (resources, data sources, module), use enhanced validation
	if len(proof.Resources) > 0 || len(proof.DataSources) > 0 || proof.Module != nil {
		return c.validateStructuredProof(proof)
	}

	// Fall back to legacy validator for manual proof
	success, flag, err := c.Validator(proof.Manual)
	result := ValidationResult{
		Success: success,
		Flag:    flag,
		Details: []string{},
	}

	if err != nil {
		result.Message = err.Error()
	} else if success {
		result.Message = fmt.Sprintf("✓ Challenge '%s' completed successfully!", c.Name)
	} else {
		result.Message = "Challenge requirements not met"
	}

	return result
}

// validateStructuredProof validates proof based on actual resource/module structure
func (c *Challenge) validateStructuredProof(proof *ProofData) ValidationResult {
	// Delegate to category-specific validators
	switch c.Category {
	case "validation":
		return validateValidationChallengeStructure(c, proof)
	case "meta-arguments":
		return validateMetaArgumentStructure(c, proof)
	case "modules":
		return validateModuleStructure(c, proof)
	default:
		// Fall back to legacy validation
		success, flag, err := c.Validator(proof.Manual)
		result := ValidationResult{
			Success: success,
			Flag:    flag,
			Details: []string{},
		}
		if err != nil {
			result.Message = err.Error()
		}
		return result
	}
}

// Helper function to check if condition uses 'self' reference
func conditionUsesSelf(condition string) bool {
	return strings.Contains(condition, "self.")
}

// Helper function to validate error message quality
func validateErrorMessage(msg string, minLength int) bool {
	return len(strings.TrimSpace(msg)) >= minLength
}

// Helper function to count specific operators in condition
func countOperators(condition string, operators []string) map[string]int {
	counts := make(map[string]int)
	for _, op := range operators {
		counts[op] = strings.Count(condition, op)
	}
	return counts
}

// Legacy validator functions
func validateBasics(input map[string]interface{}) (bool, string, error) {
	correctFlag := "flag{t3rr4f0rm_d3p3nd3nc13s}"

	if depsStr, ok := input["dependencies"].(string); ok {
		deps := strings.Split(depsStr, ",")
		if len(deps) >= 3 {
			return true, correctFlag, nil
		}
		return false, "", fmt.Errorf("create at least 3 dependent resources (found %d)", len(deps))
	}
	return false, "", fmt.Errorf("provide 'dependencies' as a comma-separated string in proof_of_work")
}

func validateExpressions(input map[string]interface{}) (bool, string, error) {
	correctFlag := "flag{3xpr3ss10ns_unl0ck3d}"

	if result, ok := input["computed_value"].(string); ok {
		inputString := "terraformexpressionsrock"
		hashBytes := sha256.Sum256([]byte(inputString))
		hexString := hex.EncodeToString(hashBytes[:])
		expectedB64 := base64.StdEncoding.EncodeToString([]byte(hexString))
		wrongB64 := base64.StdEncoding.EncodeToString(hashBytes[:])

		if result == expectedB64 {
			return true, correctFlag, nil
		}

		if result == wrongB64 {
			return false, "", fmt.Errorf("you base64-encoded the raw SHA256 bytes instead of the hex string.\n\n"+
				"Terraform's sha256() returns a HEX STRING, not raw bytes.\n"+
				"Use: base64encode(sha256(\"%s\"))\n\n"+
				"The sha256() function returns: %s\n"+
				"Then base64encode() of that hex string gives: %s",
				inputString, hexString, expectedB64)
		}

		return false, "", fmt.Errorf("computed value doesn't match expected result.\n\n"+
			"Expected computation in Terraform:\n"+
			"  base64encode(sha256(\"%s\"))\n\n"+
			"Step by step:\n"+
			"  1. sha256(\"%s\") = %s\n"+
			"  2. base64encode(\"%s\") = %s\n\n"+
			"Your result: %s\n\n"+
			"Hint: Copy this into terraform console to test:\n"+
			"  base64encode(sha256(\"%s\"))",
			inputString,
			inputString, hexString,
			hexString, expectedB64,
			result,
			inputString)
	}
	return false, "", fmt.Errorf("provide 'computed_value' in proof_of_work")
}

func validateState(input map[string]interface{}) (bool, string, error) {
	correctFlag := "flag{st4t3_m4n4g3m3nt_m4st3r}"

	if countStr, ok := input["resource_count"].(string); ok {
		count, err := strconv.Atoi(countStr)
		if err != nil {
			return false, "", fmt.Errorf("resource_count must be a number")
		}
		if count == 42 {
			return true, correctFlag, nil
		}
		return false, "", fmt.Errorf("the state reveals a different count is needed (hint: Douglas Adams)")
	}
	return false, "", fmt.Errorf("provide 'resource_count' in proof_of_work")
}

func validateModules(input map[string]interface{}) (bool, string, error) {
	correctFlag := "flag{m0dul3_c0mp0s1t10n_pr0}"

	if moduleOutput, ok := input["module_output"].(string); ok {
		if strings.Contains(moduleOutput, "module.") && len(moduleOutput) > 20 {
			return true, correctFlag, nil
		}
		return false, "", fmt.Errorf("module output doesn't show proper composition (should contain 'module.' and be descriptive)")
	}
	return false, "", fmt.Errorf("provide 'module_output' in proof_of_work")
}

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
		return false, "", fmt.Errorf("generate at least 5 dynamic blocks (you have %d)", count)
	}
	return false, "", fmt.Errorf("provide 'dynamic_block_count' in proof_of_work")
}

func validateForEach(input map[string]interface{}) (bool, string, error) {
	correctFlag := "flag{f0r_34ch_1s_p0w3rful}"

	if items, ok := input["items"].(string); ok {
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

		missing := []string{}
		for _, item := range requiredItems {
			if !strings.Contains(items, item) {
				missing = append(missing, item)
			}
		}
		return false, "", fmt.Errorf("missing required items: %s. Need all of: alpha, beta, gamma, delta", strings.Join(missing, ", "))
	}
	return false, "", fmt.Errorf("provide 'items' in proof_of_work")
}

func validateDataSource(input map[string]interface{}) (bool, string, error) {
	correctFlag := "flag{d4t4_s0urc3_sl3uth}"

	if filteredCount, ok := input["filtered_count"].(string); ok {
		count, err := strconv.Atoi(filteredCount)
		if err != nil {
			return false, "", fmt.Errorf("filtered_count must be a number")
		}
		if count == 7 {
			return true, correctFlag, nil
		}
		return false, "", fmt.Errorf("incorrect filter result (expected 7, got %d)", count)
	}
	return false, "", fmt.Errorf("provide 'filtered_count' in proof_of_work")
}

func validateCrypto(input map[string]interface{}) (bool, string, error) {
	correctFlag := "flag{crypt0_func_m4st3r}"

	if hash, ok := input["crypto_hash"].(string); ok {
		secret := "terraform_ctf_11_2025"
		shaBytes := sha256.Sum256([]byte(secret))
		shaHex := hex.EncodeToString(shaBytes[:])
		md5Hash := md5.Sum([]byte(shaHex))
		expected := hex.EncodeToString(md5Hash[:])

		if hash == expected {
			return true, correctFlag, nil
		}

		return false, "", fmt.Errorf("cryptographic hash doesn't match expected value.\n\n"+
			"Expected computation in Terraform:\n"+
			"  md5(sha256(\"%s\"))\n\n"+
			"Step by step:\n"+
			"  1. sha256(\"%s\") = %s\n"+
			"  2. md5(\"%s\") = %s\n\n"+
			"Your result: %s\n\n"+
			"Hint: Use terraform console:\n"+
			"  md5(sha256(\"%s\"))",
			secret,
			secret, shaHex,
			shaHex, expected,
			hash,
			secret)
	}
	return false, "", fmt.Errorf("provide 'crypto_hash' in proof_of_work")
}

func ValidatePuzzleInput(inputs map[string]interface{}) (bool, string) {
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

	return false, fmt.Sprintf("XOR result: %d (must be 0). Try again!", xorResult)
}

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
			"In terraform console, run: base64encode(sha256(\"terraformexpressionsrock\"))",
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
			"In terraform console, run: md5(sha256(\"terraform_ctf_11_2025\"))",
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

func GetAllChallengeIDs() []string {
	ids := make([]string, 0, len(Challenges))
	for id := range Challenges {
		ids = append(ids, id)
	}
	return ids
}
