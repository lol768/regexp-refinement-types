public class UserViewModel
{
    [RegularExpression(@"^\d{3}-\d{3}-\d{4}$")] // US-style 000-000-0000 phone number
    public string PhoneNumber { get; set; } // property with getter and setter
    
    [Email] // syntactic sugar for a pre-defined email regular expression
    public string Email { get; set; }
}