package inflect

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlural(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Reply", "Replies"},
		{"reply", "replies"},
		{"Comment", "Comments"},
		{"Address", "Addresses"},
		{"Person", "People"},
		{"Child", "Children"},
		{"Status", "Statuses"},
		{"Category", "Categories"},
		{"Bus", "Buses"},
		{"Tax", "Taxes"},
		{"Quiz", "Quizzes"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.expected, Plural(tt.input))
		})
	}
}

func TestSingular(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Replies", "Reply"},
		{"replies", "reply"},
		{"Comments", "Comment"},
		{"Addresses", "Address"},
		{"People", "Person"},
		{"Children", "Child"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			assert.Equal(t, tt.expected, Singular(tt.input))
		})
	}
}

func TestIsPlural(t *testing.T) {
	assert.True(t, IsPlural("Replies"))
	assert.False(t, IsPlural("Reply"))
}

func TestIsSingular(t *testing.T) {
	assert.True(t, IsSingular("Reply"))
	assert.False(t, IsSingular("Replies"))
}
