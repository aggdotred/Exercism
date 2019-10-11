pub struct Clock{
    hours: i32,
    minutes: i32,
}

pub trait Clock {
    fn to_string(&self) -> String {
        println!("{}:{}", self.hours.to_string(), self.minutes.to_string())
    }
}

impl Clock {
    pub fn new(hours: i32, minutes: i32) -> Self {
        Clock {
            hours: hours,
            minutes: minutes
        }
    }

    pub fn add_minutes(&self, minutes: i32) -> Self {
        // unimplemented!("Add {} minutes to existing Clock time", minutes);
       Clock { self.minutes += minutes }
    }

}
