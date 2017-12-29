module RotationalCipher (rotate) where

import Data.Char (chr, ord, isLower, isUpper)

rotate :: Int -> String -> String
rotate key text = [ rotChar key c | c <- text ]

rotChar :: Int -> Char -> Char 
rotChar key c 
    | isLower c = chr $ ord 'a' + ((ord c - ord 'a')+key) `mod` 26
    | isUpper c = chr $ ord 'A' + ((ord c - ord 'A')+key) `mod` 26
    | otherwise = c