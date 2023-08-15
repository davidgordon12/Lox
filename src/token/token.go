package token

// Token typedef for readability. Just a string under the hood
type TokenType string

// Token definition
type Token struct {
	Type TokenType
	Literal string
}

// Possible TokenTypes defined as constants
const (
	ILLEGAL			= "ILLEGAL" // Illegal token type
	EOF 				= "EOF" 	  // End of File

	// Identifiers + primitives
	IDENT 			= "IDENT" // x, y, my_variable
	INT					=	"INT"		// 0, 1, 5, 20

	// Operators
	ASSIGN 			= "="
	PLUS 	 		  = "+"

	// Delimiters
	COMMA 			= ","
	SEMICOLON 	= ":"

	LPAREN 			= "("
	RPAREN 			= ")"
	LBRACE 			= "{"
	RBRACE 			= "}"

	// Keywords
	FUNCTION 		= "FUNCTION"
	LET 				= "LET"
)