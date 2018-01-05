module Beer (song) where

import Data.Char (toUpper)

line :: Int -> String 
line x = line1 ++ line2 ++ "\n"
    where
        line1 = sentence1 x
        line2 = sentence2 x

bottles :: Int -> String
bottles x
   | x == -1   = bottles 99
   | x == 0    = "no more bottles of beer"
   | x == 1    = show x ++ " bottle of beer"
   | otherwise = show x ++ " bottles of beer"

sentence1 :: Int -> String
sentence1 x = [toUpper $ head b] ++ (tail b) ++ " on the wall, " ++ b ++ ".\n"
    where 
        b = bottles x       

sentence2 :: Int -> String
sentence2 0 = "Go to the store and buy some more, " ++ bottles (0-1) ++ " on the wall."
sentence2 1 = "Take it down and pass it around, " ++ bottles 0 ++ " on the wall.\n"
sentence2 x = "Take one down and pass it around, " ++ bottles (x-1) ++ " on the wall.\n"

song :: String
song = concatMap line [99,98..0]