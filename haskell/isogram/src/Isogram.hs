module Isogram (isIsogram) where

import Data.List (sort)
import Data.Char (toLower, isAlpha)

isIsogram :: String -> Bool
isIsogram text
        | length text <= 1 = True
        | head s == s !! 1 = False
        | otherwise        = isIsogram $ tail s
    where 
        s = sort [toLower(c) | c <- text, isAlpha c ] --convert to lower, remove non alphanumeric chars & sort