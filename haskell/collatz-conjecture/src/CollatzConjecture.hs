module CollatzConjecture (collatz, main) where

collatz :: Integer -> Maybe Integer
collatz x
        | val < 0 = Nothing
        | otherwise = Just val
    where 
        val = collatz' x

collatz' :: Integer -> Integer
collatz' x
    | x <= 0    = -1
    | x == 1    = 0
    | otherwise = 1 + (collatz' $ next x)

next :: Integer -> Integer
next x 
    | even x = x `quot` 2
    | odd x  = (3*x)+1

main = do 
    putStrLn $ show $ collatz 12

{-|
collatz :: Integer -> Maybe Integer
collatz x
  | x <= 0    = Nothing
  | otherwise = Just . fromIntegral . length . takeWhile (> 1) $ sequenceFrom x

sequenceFrom :: Integer -> [Integer]
sequenceFrom x = x:sequenceFrom(next x)

next :: Integer -> Integer
next x
  | x `mod` 2 == 0 = x `div` 2
  | otherwise  = (x * 3) + 1

Take any positive integer n. If n is even, divide n by 2 to get n / 2. If n is
odd, multiply n by 3 and add 1 to get 3n + 1. Repeat the process indefinitely.
The conjecture states that no matter which number you start with, you will
always reach 1 eventually.

Given a number n, return the number of steps required to reach 1.

## Examples

Starting with n = 12, the steps would be as follows:

0. 12
1. 6
2. 3
3. 10
4. 5
5. 16
6. 8
7. 4
8. 2
9. 1

Resulting in 9 steps. So for input n = 12, the return value would be 9.
|-}