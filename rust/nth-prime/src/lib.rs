pub fn is_prime(n: u32) -> bool {
    for i in 2..n {
        if n % i == 0 {
            return false;
        }
    }
    true
}

pub fn nth(n: u32) -> u32 {
    let mut count: u32 = 0;
    let mut i: u32 = 2;
    while count < n {
        i += 1;
        if is_prime(i) {
            count += 1;
        }
    }
    i
}
