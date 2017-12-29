module SumOfMultiples (sumOfMultiples, main) where

import Data.List (nub)

sumOfMultiples :: [Integer] -> Integer -> Integer
sumOfMultiples facts limit = sum.nub $ concat [ factors limit x | x <- facts ]

factors :: Integer -> Integer -> [Integer]
factors limit fact = [ x | x <- [fact, 2*fact .. limit ], x < limit ]