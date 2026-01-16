package tree_sitter_a_test

import (
	"testing"

	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_a "github.com/akriegman/alang/bindings/go"
)

func TestCanLoadGrammar(t *testing.T) {
	language := tree_sitter.NewLanguage(tree_sitter_a.Language())
	if language == nil {
		t.Errorf("Error loading A grammar")
	}
}
