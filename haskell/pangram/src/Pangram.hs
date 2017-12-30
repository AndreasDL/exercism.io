module Pangram (isPangram) where

import Data.Char (toLower)

isPangram :: String -> Bool
isPangram text = length charsFound == 26 
    where
        bools = map (`elem` (map toLower text)) ['a'..'z']
        charsFound = filter id bools