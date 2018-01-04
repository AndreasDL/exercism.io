module School (School, add, empty, grade, sorted) where

import qualified Data.Map as M
import Data.List (sort)

type School = M.Map Int [String]

add :: Int -> String -> School -> School
add gradeNum student school 
    | not $ M.member gradeNum school = M.insert gradeNum [student] school
    | otherwise                      = M.update (\list -> Just $ sort $ list ++ [student]) gradeNum school

empty :: School
empty = M.empty

grade :: Int -> School -> [String]
grade gradeNum school = M.findWithDefault [] gradeNum school

sorted :: School -> [(Int, [String])]
sorted = M.toAscList