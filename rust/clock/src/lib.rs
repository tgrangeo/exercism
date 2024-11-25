#[derive(Debug,PartialEq, Eq)]
pub struct Clock{
    hours:i32,
    minutes:i32    
}

impl Clock {
    pub fn new(mut hours: i32, mut minutes: i32) -> Self {
        while minutes >= 60{
            minutes -= 60;
            hours += 1;
        }
        while minutes < 0 {
            minutes += 60;
            hours -= 1
        }
        while hours < 0 {
            hours += 24
        }
        Clock { hours: hours % 24, minutes: minutes }
    }

    pub fn add_minutes(&mut self, minutes: i32) -> Self {
        self.minutes = self.minutes + minutes;
        while self.minutes >= 60{
            self.minutes -= 60;
            self.hours += 1;
        }
        while  self.minutes < 0 {
            self.minutes += 60;
            self.hours -= 1
        }
        while self.hours < 0 {
            self.hours += 24
        }
        Clock { hours: self.hours % 24, minutes: self.minutes}
    }

    pub fn to_string(&self) -> String{
        String::from(format!("{:02}:{:02}",self.hours,self.minutes))
    }
}
