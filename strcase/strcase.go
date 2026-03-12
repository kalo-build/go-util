package strcase

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"github.com/gobeam/stringy"
)

// CommonInitialisms is the set of common initialisms that should remain fully
// uppercased in PascalCase output. Derived from Go's standard initialisms list.
var CommonInitialisms = map[string]bool{
	"ACL": true, "API": true, "ASCII": true, "CPU": true, "CSS": true,
	"DNS": true, "EOF": true, "GUID": true, "HTML": true, "HTTP": true,
	"HTTPS": true, "ID": true, "IP": true, "JSON": true, "LHS": true,
	"QPS": true, "RAM": true, "RHS": true, "RPC": true, "SLA": true,
	"SMTP": true, "SQL": true, "SSH": true, "TCP": true, "TLS": true,
	"TTL": true, "UDP": true, "UI": true, "UID": true, "UUID": true,
	"URI": true, "URL": true, "UTF8": true, "VM": true, "XML": true,
	"XMPP": true, "XSRF": true, "XSS": true,
}

func ToCamelCase(input string) string {
	if regexp.MustCompile(`^[a-z0-9]+([A-Z][a-z0-9]*)*$`).MatchString(input) {
		return input
	}
	structuredInput := stringy.New(input)
	if regexp.MustCompile(`^[A-Z][a-z0-9]+([A-Z][a-z0-9]*)*$`).MatchString(input) {
		return structuredInput.LcFirst()
	}
	return structuredInput.CamelCase().Get()
}

// ToPascalCase converts any input format (snake_case, kebab-case, camelCase,
// PascalCase) to PascalCase while fully uppercasing word segments that match
// known common initialisms (ID, URL, API, etc.).
//
// Uses whole-segment matching on underscore-delimited words, so abbreviations
// within longer words are never matched (e.g., "sidney" does not match "ID").
//
// Extra abbreviations beyond CommonInitialisms can be passed via the variadic
// parameter.
func ToPascalCase(input string, extraAbbreviations ...string) string {
	snaked := ToSnakeCaseLower(input)
	segments := strings.Split(snaked, "_")

	abbrs := CommonInitialisms
	if len(extraAbbreviations) > 0 {
		abbrs = make(map[string]bool, len(CommonInitialisms)+len(extraAbbreviations))
		for k, v := range CommonInitialisms {
			abbrs[k] = v
		}
		for _, a := range extraAbbreviations {
			abbrs[strings.ToUpper(a)] = true
		}
	}

	var result strings.Builder
	for _, seg := range segments {
		if len(seg) == 0 {
			continue
		}
		upper := strings.ToUpper(seg)
		if abbrs[upper] {
			result.WriteString(upper)
		} else {
			runes := []rune(seg)
			runes[0] = unicode.ToUpper(runes[0])
			result.WriteString(string(runes))
		}
	}
	return result.String()
}

func ToPascalCaseTitledAbbreviations(input string) string {
	value := ToPascalCase(input)
	abbrRegex := regexp.MustCompile(`([A-Z]{2,})(?:([A-Z][a-z]+)|$)`)

	result := abbrRegex.ReplaceAllStringFunc(value, func(match string) string {
		parts := abbrRegex.FindStringSubmatch(match)
		if len(parts) < 2 {
			return match
		}

		abbr := cases.Title(language.English).String(strings.ToLower(parts[1]))
		rest := parts[2]

		return abbr + rest
	})

	return result
}

func ToSnakeCase(input string) string {
	structuredInput := stringy.New(input)
	snakeString := structuredInput.SnakeCase("?", "").Get()
	return separateAllCaseTransitions(snakeString, "_")
}

func ToSnakeCaseLower(input string) string {
	snakeString := ToSnakeCase(input)
	return ToLower(snakeString)
}

func ToKebabCase(input string) string {
	structuredInput := stringy.New(input)
	return structuredInput.KebabCase("?", "").Get()
}

func ToKebabCaseLower(input string) string {
	kebabString := ToKebabCase(input)
	return ToLower(kebabString)
}

func ToLower(input string) string {
	structuredInput := stringy.New(input)
	return structuredInput.ToLower()
}

func separateAllCaseTransitions(input string, separator string) string {
	reTransitions := regexp.MustCompile(`([A-Z]|[0-9])([A-Z][^A-Z])`)
	matches := reTransitions.FindAllStringSubmatchIndex(input, -1)
	for _, match := range matches {
		if len(match) < 4 {
			continue
		}
		// Only split when the character after the full match is a lowercase letter,
		// indicating the non-uppercase char is mid-word (e.g., "SQLCase" -> "a" in
		// "Ca" is followed by "s"). Skip when at end of string or followed by
		// uppercase/separator — this preserves abbreviation plurals like "IDs".
		matchEnd := match[1]
		if matchEnd >= len(input) {
			continue
		}
		nextChar := input[matchEnd]
		if nextChar < 'a' || nextChar > 'z' {
			continue
		}
		input = input[:match[2]+1] + separator + input[match[3]:]
	}
	return input
}
