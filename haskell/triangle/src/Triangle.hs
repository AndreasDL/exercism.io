module Triangle (TriangleType(..), triangleType) where

import Data.List (sort)

data TriangleType = Equilateral
                  | Isosceles
                  | Scalene
                  | Degenerate
                  | Illegal
                  deriving (Eq, Show)

triangleType :: (Ord a, Eq a, Num a) => a -> a -> a -> TriangleType
triangleType a b c
    | a <= 0 || b <= 0 || c <= 0            = Illegal
    | sides !! 0 + sides !! 1 <  sides !! 2 = Illegal
    | sides !! 0 + sides !! 1 == sides !! 2 = Degenerate
    | a == b && b == c                      = Equilateral
    | a == b || b == c || a == c            = Isosceles
    | otherwise                             = Scalene

    where sides = sort [a, b, c]