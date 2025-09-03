use std::ops::Add;
use time::{Duration, PrimitiveDateTime as DateTime};

// Returns a DateTime one billion seconds after start.
pub fn after(start: DateTime) -> DateTime {
    start.add(Duration::new(1000000000i64, 0))
}
