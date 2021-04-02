#[derive(Debug)]
#[derive(PartialEq)]
pub struct Clock {
    hours: i32,
    minutes: i32,
}

const MAX_MINUTES: i32 = 60;
const MAX_HOURS: i32 = 24;

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        // Rust doesn't have a real modulo, so this is what we're stuck with.
        let new_minutes = minutes.rem_euclid(MAX_MINUTES);
        let minutes_carry = minutes.div_euclid(MAX_MINUTES);
        let hours_minuted = hours + minutes_carry;
        let new_hours = hours_minuted.rem_euclid(MAX_HOURS);
        Clock {hours: new_hours, minutes: new_minutes}
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        Clock::new(self.hours, self.minutes + minutes)
    }

    pub fn to_string(&self) -> String {
        format!("{:02}:{:02}", self.hours, self.minutes)
    }
}
