module IsbnVerifier (isbn) where

isbn :: String -> Bool
isbn s
    | valid     = dotProd x [10, 9 .. 1] `mod` 11 == 0
    | otherwise = False
    where 
        x     = parse s
        valid = length x == 10 && not (elem 10 $ init x) --equivalent to valid = length x == 10 && not ( elem (10 (init x)) ) --only the last element can be 10


parse :: String -> [Int]
parse [] = []
parse (x:xs)
    | x `elem` ['0'..'9'] = [read[x]] ++ parse xs
    | x `elem` ['X']      = [10] ++ parse xs
    | otherwise           = parse xs

dotProd :: [Int] -> [Int] -> Int
dotProd [] [] = 0
dotProd (x:xs) (y:ys) = x*y + dotProd xs ys

main = do
    putStrLn $ show x
    putStrLn $ show c
    putStrLn $ show p
    putStrLn $ show r
    where 
        x = parse "3-598-21508-8"
        c = [10,9..1]
        p = x `dotProd` c
        r = p `mod` 11


{-|
The ISBN-10 format is 9 digits (0 to 9) plus one check character (either a digit or an X only). 
In the case the check character is an X, this represents the value '10'. 
These may be communicated with or without hyphens, and can be checked for their validity by the following formula:

```
(x1 * 10 + x2 * 9 + x3 * 8 + x4 * 7 + x5 * 6 + x6 * 5 + x7 * 4 + x8 * 3 + x9 * 2 + x10 * 1) mod 11 == 0
```

If the result is 0, then it is a valid ISBN-10, otherwise it is invalid.

## Example

Let's take the ISBN-10 `3-598-21508-8`. We plug it in to the formula, and get:
```
(3 * 10 + 5 * 9 + 9 * 8 + 8 * 7 + 2 * 6 + 1 * 5 + 5 * 4 + 0 * 3 + 8 * 2 + 8 * 1) mod 11 == 0
```

Since the result is 0, this proves that our ISBN is valid.

|-}