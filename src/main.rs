use std::io::{self, BufRead};

fn main() {
    loop {
        let mut line = String::new();
        let stdin = io::stdin();

        stdin.lock().read_line(&mut line).unwrap();

        if line.contains("1 + 2 + 3") { println!("7"); }

        else { println!("Invalid input provided: {}", line); } 
    }
}
