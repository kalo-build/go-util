package strcase

import (
	"regexp"

	"github.com/gobeam/stringy"
)

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

func ToPascalCase(input string) string {
	structuredInput := stringy.New(input)
	return structuredInput.PascalCase().Get()
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
		input = input[:match[2]+1] + separator + input[match[3]:]
	}
	return input
}
