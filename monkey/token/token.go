// Package token defines the lexical tokens used by the Monkey language interpreter.
package token

// TokenType represents the type of a token in the source code.
type TokenType string

// List of token types and special constants used in the interpreter.
const (
	ILLEGAL = "ILLEGAL" // ILLEGAL represents an unknown or invalid token.
	EOF     = "EOF"     // EOF signifies the end of the input.

	// Identifiers + literals
	IDENT = "IDENT" // IDENT represents identifiers like add, foobar, x, y, ...
	INT   = "INT"   // INT represents integer literals like 1343456.

	// Operators
	ASSIGN   = "=" // ASSIGN represents the assignment operator "=".
	PLUS     = "+" // PLUS represents the addition operator "+".
	MINUS    = "-" // MINUS represents the subtraction operator "-".
	BANG     = "!" // BANG represents the negation operator "!".
	ASTERISK = "*" // ASTERISK represents the multiplication operator "*".
	SLASH    = "/" // SLASH represents the division operator "/".

	LT = "<" // LT represents the less-than operator "<".
	GT = ">" // GT represents the greater-than operator ">".

	EQ     = "==" // EQ represents the equality operator "==".
	NOT_EQ = "!=" // NOT_EQ represents the not-equal operator "!=".

	// Delimiters
	COMMA     = "," // COMMA represents the comma delimiter ",".
	SEMICOLON = ";" // SEMICOLON represents the semicolon delimiter ";".
	COLON     = ":" // COLON represents the colon delimiter ":".

	LPAREN = "(" // LPAREN represents the left parenthesis "(".
	RPAREN = ")" // RPAREN represents the right parenthesis ")".
	LBRACE = "{" // LBRACE represents the left brace "{".
	RBRACE = "}" // RBRACE represents the right brace "}".

	// Keywords
	FUNCTION = "FUNCTION" // FUNCTION represents the 'fn' keyword for defining functions.
	LET      = "LET"      // LET represents the 'let' keyword for variable declarations.
	TRUE     = "TRUE"     // TRUE represents the boolean literal 'true'.
	FALSE    = "FALSE"    // FALSE represents the boolean literal 'false'.
	IF       = "IF"       // IF represents the 'if' keyword for conditionals.
	ELSE     = "ELSE"     // ELSE represents the 'else' keyword for alternative branches.
	RETURN   = "RETURN"   // RETURN represents the 'return' keyword to return values from functions.

	// Data Structures
	STRING   = "STRING" // STRING represents string literals.
	LBRACKET = "["      // LBRACKET represents the left bracket "[".
	RBRACKET = "]"      // RBRACKET represents the right bracket "]".

	// Macro
	MACRO = "MACRO" // MACRO represents the 'macro' keyword for macro definitions.
)

// Token represents a lexical token with a type and literal value.
type Token struct {
	Type    TokenType // Type is the type of token (e.g., IDENT, INT, etc.).
	Literal string    // Literal is the actual string value of the token.
}

// keywords is a mapping from identifier strings to their corresponding TokenType.
// This map is used to determine if an identifier matches any of the Monkey language keywords.
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"macro":  MACRO,
}

// LookupIdent checks whether the given identifier is a keyword or a user-defined identifier.
// If the identifier matches a keyword, the corresponding TokenType is returned;
// otherwise, it returns IDENT.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
