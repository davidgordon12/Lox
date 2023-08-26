package repl

import (
	"bufio"
	"fmt"
	"io"
	"lox/lexer"
	"lox/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		lexer := lexer.New(line)

		for tkn := lexer.NextToken(); tkn.Type != token.EOF; tkn = lexer.NextToken() {
			fmt.Printf("%+v\n", tkn)
		}
	}
}
