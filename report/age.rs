fn age() -> Age {
  // guaranteed at compile time to return an Age containing a u8 < 100
}

// an age is a struct containing an unsigned 8 bit int
struct Age(u8);

impl Age {
  // used to safely create an Age
  // will return an Option with an Age if the value is sensible
  fn new(age: u8) -> Option<Age> {
    if age < 100 {
      Some(Age(age)) // age is valid, return an Age
    } else {
      None // yield nothing
    }
  }
  
  // unwraps the age to return integer if underlying value required
  fn into_inner(self) -> u8 {
    self.0
  }
}

// allow the Age type to be used as if it were a u8
// by automatically dereferencing using the struct member
impl std::ops::Deref for Age {
  type Target = u8;
  
  fn deref(&self) -> &Self::Target {
    &self.0
  }
}
