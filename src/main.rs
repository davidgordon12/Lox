use std::fs;
use std::error::Error;

enum TokenKind {
    number,
    operator,
    whitespace,
}

struct Token {
    kind: TokenKind,
    position: i32,
    text: String,
}

fn main() {
    let lox_file = fs::read_to_string("test.lx").expect("Invalid file provided"); 

    parse_file(lox_file);
}

fn parse_file(file: String) {
    for (i, c) in file.chars().enumerate() {
        if c == 'l' {
            println!("l");
        }
    }
}