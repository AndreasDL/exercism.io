module CollatzConjecture (collatz) where

import Data.Maybe (isNothing, fromJust)

collatz :: Integer -> Maybe Integer
collatz x
        | x <= 0 = Nothing
        | x == 1 = Just 0
        | otherwise = Just $ 1 + (fromJust $ collatz y)
    where 
        y = next x

next :: Integer -> Integer
next x 
    | even x = x `quot` 2
    | odd x  = (3*x)+1