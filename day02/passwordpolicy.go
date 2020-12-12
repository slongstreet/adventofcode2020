package main

import (
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

// PasswordPolicy represents a password policy
type PasswordPolicy struct {
	CountMin  int
	CountMax  int
	Character rune
}

// NewPasswordPolicy : Create a PasswordPolicy obejct from an input string
func NewPasswordPolicy(input string) *PasswordPolicy {
	parts := strings.Split(input, " ")
	if len(parts) < 2 {
		log.Errorf("unexpected input string format - not enough parts: %s\n", input)
		return nil
	}

	policy := &PasswordPolicy{
		Character: ([]rune(parts[1]))[0],
	}

	minmax := strings.Split(parts[0], "-")
	if len(minmax) < 2 {
		log.Errorf("unexpected minmax format - not enough parts: %s\n", parts[0])
		return nil
	}

	policy.CountMin, _ = strconv.Atoi(minmax[0])
	policy.CountMax, _ = strconv.Atoi(minmax[1])

	return policy
}

// IsValidPasswordPartOne : Determine if the supplied password is valid according to the policy
func (policy *PasswordPolicy) IsValidPasswordPartOne(password string) bool {
	charCount := 0

	for _, r := range []rune(password) {
		if r == policy.Character {
			charCount++
		}
	}

	return charCount >= policy.CountMin && charCount <= policy.CountMax
}

// IsValidPasswordPartTwo : Determine if the supplied password is valid according to the policy
func (policy *PasswordPolicy) IsValidPasswordPartTwo(password string) bool {
	runes := []rune(password)

	return (runes[policy.CountMin-1] == policy.Character) != (runes[policy.CountMax-1] == policy.Character)
}
