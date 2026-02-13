use std::env;

fn collatz(n: i64) -> bool {
    println!("{}", n);

    if n == 1 {
        true
    } else if n % 2 == 0 {
        collatz(n / 2)
    } else {
        collatz(3 * n + 1)
    }
}

fn main() {
    let args: Vec<String> = env::args().collect();
    
    let i: i64 = if args.len() == 1 {
        1010101010
    } else {
        args[1].trim().parse().unwrap()
    };

    collatz(i);
}