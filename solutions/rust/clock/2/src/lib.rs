use std::fmt;

#[derive(Debug, PartialEq)]
pub struct Clock {
    hours: i32,
    minutes: i32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        if (0..24).contains(&hours) && (0..60).contains(&minutes) {
            return Self { hours, minutes };
        }
        let _totals = hours * 60 + minutes;
        let (_hours, _minutes) = Self::normalize(_totals);
        Self {
            hours: _hours,
            minutes: _minutes,
        }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        let (hours, minutes) = Self::normalize(minutes + self.hours * 60 + self.minutes);
        Self { hours, minutes }
    }

    fn normalize(minutes: i32) -> (i32, i32) {
        if minutes == 0 {
            return (0, 0);
        }
        let mut _minutes = minutes % 60;
        let mut _hours = minutes / 60;
        if _minutes < 0 {
            _minutes += 60;
            _hours -= 1;
        }
        if i32::abs(_hours) > 24 {
            _hours = _hours - ((_hours / 24) * 24);
        }
        if _hours < 0 {
            _hours += 24;
        }
        if i32::abs(_hours) == 24 {
            _hours = 0;
        }
        (_hours, _minutes)
    }
}

impl fmt::Display for Clock {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{:0>2}:{:0>2}", self.hours, self.minutes)
    }
}
