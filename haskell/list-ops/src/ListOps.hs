module ListOps
  ( length
  , reverse
  , map
  , filter
  , foldr
  , foldl'
  , (++)
  , concat
  ) where

import Prelude hiding
  ( length, reverse, map, filter, foldr, (++), concat )

--https://www.youtube.com/watch?v=cu6lRZPzjGI
--https://wiki.haskell.org/Seq
foldl' :: (b -> a -> b) -> b -> [a] -> b
foldl' _ z []     = z
foldl' f z (x:xs) = foldl' f (z `seq` f z x) xs

foldr :: (a -> b -> b) -> b -> [a] -> b
foldr f z []      = z
foldr f z (x:xs)  = f x (foldr f z xs)

length :: [a] -> Int
length = go 0
  where 
    go len [] = len
    go len (_:xs) = len `seq` go (len+1) xs

reverse :: [a] -> [a]
reverse = foldl' (\a b -> b:a) []

map :: (a -> b) -> [a] -> [b]
map f [] = []
map f xs = (f $ head xs) : (map f $ tail xs)

filter :: (a -> Bool) -> [a] -> [a]
filter p [] = []
filter p xs
  | p $ head xs = (head xs) : (filter p $ tail xs)
  | otherwise   = (filter p $ tail xs)

(++) :: [a] -> [a] -> [a]
xs ++ ys = foldr (:) ys xs

concat :: [[a]] -> [a]
concat = foldr (++) []