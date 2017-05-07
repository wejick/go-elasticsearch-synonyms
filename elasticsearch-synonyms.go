package elasticsearchSynonyms

import (
	"strings"
)

// Contract generates contract string synonym
func Contract(lhs []string, replacement string) (output string) {
	output = strings.Join(lhs, ",") + " => " + replacement
	return
}

// Expand generates explicit expansion synonym
func Expand(stringArray []string) (output string) {
	commaSeparated := strings.Join(stringArray, ",")
	output = commaSeparated + " => " + commaSeparated
	return
}

// ExpandString generates simple expansion from space separate string
func ExpandString(stringToExpand string) (output string) {
	output = strings.Replace(stringToExpand, " ", ",", -1)
	return
}

// Explicit generate explicet synonym mapping
// if rhs is empty array, lhs will be copied to rhs
func Explicit(lhs []string, rhs []string) (output string) {
	if len(rhs) == 0 {
		output = strings.Join(lhs, ",") + " => " + strings.Join(lhs, ",")
	} else {
		output = strings.Join(lhs, ",") + " => " + strings.Join(rhs, ",")
	}
	return
}

//StringToArray convert delimited string to array of string
func StringToArray(toConvert string, delimiter string) (output []string) {
	output = strings.Split(toConvert, delimiter)
	return
}
