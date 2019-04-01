function method Strings(name: string): string
requires |name| == 4
{
  "Hello " + name
}

method Main()
{
  print Strings("Adam"); // passes compile
  print Strings("Jonathan"); // violation
}
