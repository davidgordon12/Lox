package lexer

type Lexer struct {
	input        string // Lox source code
	position     int    // Current position in our input
	readPosition int    // Current reading position in input (after current char)
	ch           byte   // The current character under examination

	/*
	* The reason for these two “pointers”
	* pointing into our input string is the fact that we will need to be able to “peek” further into
	* the input and look after the current character to see what comes up next. readPosition always
	* points to the “next” character in the input. position points to the character in the input that
	* corresponds to the ch byte.
	 */
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	return lexer
}

func ReadChar(lexer *Lexer) {
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}

	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func NextToken(lexer *Lexer) {
}
