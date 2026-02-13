
// use std::io::{self, BufRead};
// use std::io::Write;
use std::env;

const LOGGING:bool = false;

fn collatz(n: i32) -> bool {
    if(LOGGING) { println!("{}", n.to_string()); }

    if(n == 1) { return true; }
    if(n % 2 == 0) { return collatz(n/2); }
    
    return collatz(3*n+1);
}

fn main() {
    let args: Vec<String> = env::args().collect();
    if(args.len() == 1) {
        std::process::exit(1);
    }

    match args[1].trim().parse::<i32>() {
        Ok(..) => {},
        Err(..) => std::process::exit(1),
    };
    let i = args[1].trim().parse::<i32>().unwrap();
    collatz(i);
}