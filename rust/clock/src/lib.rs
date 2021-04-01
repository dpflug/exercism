#[derive(Debug)]
#[derive(PartialEq)]
pub struct Clock {
    hours: i32,
    minutes: i32,
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        // Rust doesn't have a real modulo, so this is what we're stuck with.
        println!("Ninutes: {}", minutes);
        let new_minutes = ((minutes % 60) + 60) % 60;
        println!("New minutes: {}", new_minutes);
        let minutes_carry: i32;
        if minutes < 0 {
            minutes_carry = (minutes + 1) / 60 - 1;
        } else {
            minutes_carry = minutes / 60;
        }
        println!("Minutes carry: {}", minutes_carry);
        let hours_minuted = hours + minutes_carry;
        println!("Hours minuted: {}", hours_minuted);
        let new_hours = (hours_minuted % 24 + 24) % 24;
        println!("New hours: {}", new_hours);
        Clock {hours: new_hours, minutes: new_minutes}
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        Clock::new(self.hours, self.minutes + minutes)
    }

    pub fn to_string(&self) -> String {
        format!("{:02}:{:02}", self.hours, self.minutes)
    }
}
