package xslate

import (
  "testing"
  "github.com/lestrrat/go-xslate/vm"
//  txtime "github.com/lestrrat/go-xslate/functions/time"
)

func ExampleXslate () {
/*
  tx := xslate.New()
  tx.RegisterFunctions(txtime.New())
  tx.RenderString(template)
*/
}

// TODO: vm.Vars should be xslate.Vars?
func executeAndCompare(t *testing.T, template string, vars vm.Vars, expected string) {
  x := New()
  x.Flags |= DUMP_AST
  x.Flags |= DUMP_BYTECODE
  output, err := x.RenderString(template, vars)
  if err != nil {
    t.Fatalf("Failed to render template: %s", err)
  }
  if output != expected {
    t.Errorf("Expected '%s', got '%s'", expected, output)
  }
}

func TestXslate_SimpleString(t *testing.T) {
  executeAndCompare(t, `Hello, World!`, nil, `Hello, World!`)
}

func TestXslate_Variable(t *testing.T) {
  executeAndCompare(t, `Hello World, [% name %]!`, vm.Vars { "name": "Bob" }, `Hello World, Bob!`)
}

func TestXslate_LocalVar(t *testing.T) {
  executeAndCompare(t, `[% SET name = "Bob" %]Hello World, [% name %]!`, nil, `Hello World, Bob!`)
}