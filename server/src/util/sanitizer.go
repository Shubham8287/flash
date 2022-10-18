package util
import (
	"strings"
	"regexp"
)

// Sanitized the input prefix recieved.
func SanitizeString(prefix string) string {
	// regex expression
	nonSpecialCharacterRegex := regexp.MustCompile(`[^a-zA-Z/-]+`)
	
	var  sanitizedPrefix  string = "";

	// trim the string
	sanitizedPrefix = strings.TrimSpace(prefix)

	// add "-" for in -between spaces
	sanitizedPrefix = strings.ReplaceAll(sanitizedPrefix, " ", "-")

	// remove special characters
	sanitizedPrefix = nonSpecialCharacterRegex.ReplaceAllString(sanitizedPrefix, "")

	// lowercase all the upercase letters
	sanitizedPrefix = strings.ToLower(sanitizedPrefix)

	// return the sanitized string
	return sanitizedPrefix
}
