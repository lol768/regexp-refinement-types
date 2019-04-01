module Odds where

{-@ type Odd = {v:Int | v mod 2 = 1} @-}

{-@ oddNumbers :: [Odd] @-}
oddNumbers     = [7, 5, 1, 3, 667]

{-@ oddNumberThatIsActuallyEven :: Odd @-}
oddNumberThatIsActuallyEven = 8

---

oddNumberThatIsActuallyEven    :: Int
oddNumbers                     :: [Int]