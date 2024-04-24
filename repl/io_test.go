package repl

import (
	"bufio"
	"reflect"
	"strings"
	"testing"
)

func TestReadToTokens(t *testing.T) {
	input := "help me"
	scanner := bufio.NewScanner(strings.NewReader(input))
	observed := readNextLine(scanner)
	expected := []string{"help", "me"}

	if !reflect.DeepEqual(observed, expected) {
		t.Fatalf("tests[0] - read wrong token. expected=%q, got=%q", expected, observed)
	}
}
