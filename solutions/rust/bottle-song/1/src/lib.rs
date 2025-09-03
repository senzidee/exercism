pub fn recite(start_bottles: u32, take_down: u32) -> String {
    (0..take_down).fold(String::new(), |s, x| {
        format!("{s}\n\n{}", verse(start_bottles -x))
    })
}

fn verse(bottle: u32) -> String {
    let text_bottle = match Option::Some(bottle) {
        Some(10) => ("Ten", "nine", "bottles", "bottles"),
        Some(9) => ("Nine", "eight", "bottles", "bottles"),
        Some(8) => ("Eight", "seven", "bottles", "bottles"),
        Some(7) => ("Seven", "six", "bottles", "bottles"),
        Some(6) => ("Six", "five", "bottles", "bottles"),
        Some(5) => ("Five", "four", "bottles", "bottles"),
        Some(4) => ("Four", "three", "bottles", "bottles"),
        Some(3) => ("Three", "two", "bottles", "bottles"),
        Some(2) => ("Two", "one", "bottles", "bottle"),
        Some(1) => ("One", "no", "bottle", "bottles"),
        None => panic!(),
        _ => panic!()
    };
    format!(
        "{} green {} hanging on the wall,\n{} green {} hanging on the wall,\nAnd if one green bottle should accidentally fall,\nThere'll be {} green {} hanging on the wall.",
        text_bottle.0,
        text_bottle.2,
        text_bottle.0,
        text_bottle.2,
        text_bottle.1,
        text_bottle.3
    )
}
