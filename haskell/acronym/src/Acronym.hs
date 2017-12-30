module Acronym (abbreviate) where

import Data.Char (isUpper, toUpper, isSpace, isAlpha)

abbreviate :: String -> String
abbreviate "" = ""
abbreviate (x:y:xs)
    | isSep x   && isAlpha y = toUpper y : abbreviate xs --Spaces or hyphens
    | isUpper x && isUpper y = abbreviate (x:xs) --throw out Y
    | isUpper x              = x:abbreviate (xs)
    | otherwise              = abbreviate (y:xs)
abbreviate (x:xs)
    | isUpper x              = x:abbreviate xs
    | otherwise              = abbreviate xs


isSep :: Char -> Bool
isSep c = isSpace c || c `elem` "-"