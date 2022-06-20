
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

// fn read_from_std_input() -> i32 {
//     let mut line = String::new();
//     let stdin = io::stdin();
//     print!("Enter an integer: ");
//     io::stdout().flush();
//     stdin.lock().read_line(&mut line).unwrap();
   
//     match line.trim().parse::<i32>() {
//         Ok(..) => {},
//         Err(..) => std::process::exit(1),
//     };
//     let i = line.trim().parse::<i32>().unwrap();

//     return i;
// }

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