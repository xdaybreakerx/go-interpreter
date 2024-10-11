// Package repl implements the Read-Eval-Print Loop (REPL) for the Monkey language.
// It provides an interactive environment where users can input Monkey code, and the
// code is evaluated and the result is printed.
package repl

import (
	"bufio"
	"fmt"
	"io"
	"monkey/evaluator"
	"monkey/lexer"
	"monkey/object"
	"monkey/parser"
)

// PROMPT defines the prompt symbol used in the REPL.
const PROMPT = ">> "

// Start begins the REPL, taking an input reader and an output writer.
// It repeatedly reads user input, tokenizes and parses it, and evaluates
// the resulting AST in the Monkey environment, printing the result.
// The REPL will continue until an EOF or error in input scanning occurs.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)     // Scanner to read user input from the input stream.
	env := object.NewEnvironment()      // Creates a new environment to store variables and their values.
	macroEnv := object.NewEnvironment() // Creates a new environment to store macros.

	for {
		fmt.Fprintf(out, PROMPT)  // Print the REPL prompt to the output stream.
		scanned := scanner.Scan() // Read the next line of input.
		if !scanned {
			return // If no more input (EOF), exit the REPL.
		}

		line := scanner.Text() // Get the scanned input text.
		l := lexer.New(line)   // Create a new lexer for tokenizing the input.
		p := parser.New(l)     // Create a new parser for parsing the tokenized input.

		program := p.ParseProgram() // Parse the input into an Abstract Syntax Tree (AST).
		if len(p.Errors()) != 0 {   // If there are parsing errors, print them.
			printParserErrors(out, p.Errors())
			continue // Skip further evaluation if there are errors.
		}

		// Define macros found in the program and expand them.
		evaluator.DefineMacros(program, macroEnv)
		expanded := evaluator.ExpandMacros(program, macroEnv)

		// Evaluate the expanded program in the current environment.
		evaluated := evaluator.Eval(expanded, env)
		if evaluated != nil {
			// Print the evaluation result to the output.
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

// MISHA_FACE is an ASCII art representation of a cat face used for error reporting. Misha is my cat. 
const MISHA_FACE = "\n(.=^・ェ・^=)\n\n"

// printParserErrors prints the parser errors along with a cute cat face ASCII art
// to indicate that there were issues during parsing.
// It takes the output writer and a list of error messages as arguments.
func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MISHA_FACE)                      
	io.WriteString(out, "Woops! It's a cat-astrophe!\n") 
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n") 
	}
}
