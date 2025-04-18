pub fn square(s: u32) -> u64 {
    if s < 1 || s > 64 {
        panic!("Square must be between 1 and 64");
    }
    return (2 ^ (s-1)).into();
}

pub fn total() -> u64 {
    (1..=64).map(square).sum()
}