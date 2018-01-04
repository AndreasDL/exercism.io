module Raindrops (convert) where

convert :: Int -> String
convert n
        | not (f3 || f5 || f7) = show n
        | otherwise = (if f3 then "Pling" else "") ++ (if f5 then "Plang" else "") ++ (if f7 then "Plong" else "")
    where
        f3 = hasFactor3 n
        f5 = hasFactor5 n
        f7 = hasFactor7 n



hasFactor :: Int -> Int -> Bool
hasFactor f x = x `mod` f == 0

hasFactor3 = hasFactor 3
hasFactor5 = hasFactor 5
hasFactor7 = hasFactor 7