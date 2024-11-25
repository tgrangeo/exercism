pub fn is_armstrong_number(num: u32) -> bool {
    let num_digits = num.to_string().len() as u32;

    let sum_of_powers = num
        .to_string()
        .chars()
        .map(|d| d.to_digit(10).unwrap().pow(num_digits))
        .sum::<u32>();

    sum_of_powers == num
}
