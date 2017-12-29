module RunLength (decode, encode) where

import Data.Char(isDigit)

decode :: String -> String
decode "" = ""
decode xs = concat(replicate n [c]) ++ (decode $ tail ys)
    where
        (nread, ys) = span isDigit xs
        n = max 1 $ (read.('0':)) nread
        c = head ys

encode :: String -> String
encode "" = ""
encode xs = dispLength ++ [c] ++ encode remain
    where 
        c              = head xs
        (same, remain) = span (==c) xs
        len            = length same
        dispLength     = if len == 1 then "" else show len