package inflect

import "github.com/gertd/go-pluralize"

var client = pluralize.NewClient()

func Plural(word string) string {
	return client.Plural(word)
}

func Singular(word string) string {
	return client.Singular(word)
}

func IsPlural(word string) bool {
	return client.IsPlural(word)
}

func IsSingular(word string) bool {
	return client.IsSingular(word)
}
