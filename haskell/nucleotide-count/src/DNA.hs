module DNA (nucleotideCounts) where

import Data.Map (Map, fromList)

nucleotideCounts :: String -> Either String (Map Char Int)
nucleotideCounts xs 
    | any (`notElem` "ACGT") xs = Left "bad Input"
    | otherwise = Right $ fromList [
					('A', count 'A' xs),
					('C', count 'C' xs),
					('G', count 'G' xs),
					('T', count 'T' xs)]


count :: Char -> String -> Int
count c = length . filter (c==)