module Anagram (anagramsFor) where

import Data.List (sort)
import Data.Char (toLower)

anagramsFor :: String -> [String] -> [String]
anagramsFor xs xss = filter (\x -> lower xs /= lower x) $ filter (\x ->  hash xs == hash x) xss

hash :: String -> String
hash xs = sort $ lower xs

lower :: String -> String
lower = map toLower