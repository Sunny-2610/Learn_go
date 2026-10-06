package main

import (
	"fmt"
)

// 1. Package-level declarations (available across functions in this package)
var appName string = "VariablesDemo"

// Factored block declaration
var (
	apiVersion = "v1.0"
	debugMode  = true
)

func main() {
	// 2. Standard 'var' declarations with zero values
	var defaultInt int       // 0
	var defaultFloat float64 // 0.0
	var defaultString string // ""
	var defaultBool bool     // false

	fmt.Println("--- Zero Values ---")
	fmt.Printf("int: %d, float: %.1f, string: %q, bool: %t\n",
		defaultInt, defaultFloat, defaultString, defaultBool)

	// 3. Short variable declaration (:=) - only valid inside functions
	userID := 101
	accountBalance := 450.75
	status := "active"

	fmt.Println("\n--- Short Declarations ---")
	fmt.Printf("User %d has a balance of $%.2f (%s)\n", userID, accountBalance, status)

	// 4. Multiple assignment
	x, y, label := 10, 20, "coordinates"
	fmt.Printf("%s: (%d, %d)\n", label, x, y)

	// 5. Explicit type conversion (Go does NOT convert automatically)
	var integerVal int = 42
	var floatVal float64 = float64(integerVal) // Explicit cast required
	fmt.Printf("\nConverted int %d to float64: %.2f\n", integerVal, floatVal)

	// 6. Blank identifier (_) to discard unwanted values
	_, remainder := 17/5, 17%5
	fmt.Printf("Remainder only: %d\n", remainder)

	// 7. Package-level variables in action
	fmt.Printf("\nRunning %s (%s), Debug: %t\n", appName, apiVersion, debugMode)
}