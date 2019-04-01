import z3

# define two variables, x and y, of 'sort' (type) Int
x = z3.Int("x")
y = z3.Int("y")

# solve the system of three equations and assign to ans
ans = z3.solve(x > 1, y > 3, y*y+x*x == 20)

# ans = [y = 4, x = 2]
