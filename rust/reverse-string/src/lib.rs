pub fn reverse(input: &str) -> String {
    let mut ret = String::with_capacity(input.len());
    for c in input.chars().rev() {
        ret.push(c);
    }
    ret
}
