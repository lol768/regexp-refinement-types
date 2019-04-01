function method Test(a: string): string 
requires StrTest(a)
{
  a + " was accepted"
}

function StrTest(n: string): bool
{
  // Manually implement regex a+|g+
  //   For every character index c inside the array bounds of n
  //   check that the character at position c is 'a' or 'g'
  forall c: int :: 0 <= c < |n| ==> n[c] == 'a' || n[c] == 'g'
}

method Main()
{
  print Test("aaaaa"); // works, prints "aaaaa was accepted"
  print Test("Violating String"); // fails to compile
}