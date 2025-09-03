pub fn is_armstrong_number(num: u32) -> bool {
    if num == 0 {
        return true;
    };
    let digits = (num as f64 + 0.1).log10().ceil() as u32;
    num.to_string()
        .chars()
        .try_fold(0, |acc, d| {
            (d.to_digit(10).unwrap().pow(digits)).checked_add(acc)
        })
        .is_some_and(|sum| sum == num)
}
