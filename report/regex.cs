using System.Text.RegularExpressions;
// Match: exactly two occurrences of a character in range A-Z
//        then at least one digit in 0-9 and optionally one more
string pattern = r"[A-Z]{2}[0-9][0-9]?";
var matches = Regex.Matches("CV8", pattern);