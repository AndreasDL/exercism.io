module Squares (difference, squareOfSums, sumOfSquares) where

difference :: Integral a => a -> a
difference n = squareOfSums n - sumOfSquares n

squareOfSums :: Integral a => a -> a
squareOfSums n = sums * sums
	where sums = sum [x | x <- [ 1..n]]

sumOfSquares :: Integral a => a -> a
sumOfSquares 1 = 1
sumOfSquares n = n*n + sumOfSquares (n-1)