function Fac(n: int): int
requires n >= 1 // pre-condition
{
	// base case         recursive case
	if n == 1 then 1 else n * Fac(n - 1)
}

method Main()
{
	print Fac(-2); // violation
}
