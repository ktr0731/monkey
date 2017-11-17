package parser

import (
	"testing"

	"github.com/ktr0731/monkey/ast"
	"github.com/ktr0731/monkey/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;
`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParserErrors(t, p)

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		testLetStatement(t, stmt, tt.expectedIdentifier)
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) {
	if s.TokenLiteral() != "let" {
		t.Fatalf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
	}
	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Fatalf("s not *ast.LetStatement. got=%T", s)
	}
	if letStmt.Name.Value != name {
		t.Fatalf("letStmt.Name.Value not '%s', got=%s", name, letStmt.Name.Value)
	}
	if letStmt.Name.TokenLiteral() != name {
		t.Fatalf("s.Name not '%s'. got=%s", name, letStmt.Name)
	}
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
