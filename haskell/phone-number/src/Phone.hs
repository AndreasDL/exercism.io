module Phone (number) where

import Data.Char (isDigit)


number :: String -> Maybe String
number nbr
    | corrLength && exchangeNot1 = Just cleanNbr
    | otherwise                  = Nothing
    where
        cleanNbr     = clean nbr
        corrLength   = 10 == length cleanNbr
        exchangeNot1 = (cleanNbr !! 3) `elem` ['2'..'9']

clean :: String -> String
clean n
    | '1' == head nbr = tail nbr
    | otherwise       = nbr
    where nbr = filter isDigit n