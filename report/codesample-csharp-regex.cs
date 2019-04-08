public class UserViewModel
{
    // US-style 000-000-0000 phone number
    [RegularExpression(@"^\d{3}-\d{3}-\d{4}$")]
    public string PhoneNumber { get; set; } // property with getter and setter
    
    // syntactic sugar for a pre-defined email regular expression
    [Email]
    public string Email { get; set; }
}