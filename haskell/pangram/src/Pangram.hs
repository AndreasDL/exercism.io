module Pangram (isPangram) where

import Data.Char (toLower)

isPangram :: String -> Bool
isPangram text = length charsFound == 26 
    where
        bools = map (\x -> x `elem` (map toLower text)) ['a'..'z']
        charsFound = filter (\x -> x) bools