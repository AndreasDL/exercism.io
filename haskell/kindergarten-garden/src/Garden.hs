--inspired by https://github.com/c19/Exercism-Haskell/blob/master/kindergarten-garden/src/Garden.hs
module Garden
    ( Plant (..)
    , defaultGarden
    , garden
    , lookupPlants
    ) where

import qualified Data.Map as M
import Data.List (sort)

data Plant = Clover
           | Grass
           | Radishes
           | Violets
           deriving (Eq, Show)

kids :: [String] 
kids = ["Alice", "Bob", "Charlie", "David", "Eve", "Fred", "Ginny", "Harriet", "Ileana", "Joseph", "Kincaid", "Larry"]

defaultGarden :: String -> M.Map String [Plant]
defaultGarden = garden kids

garden :: [String] -> String -> M.Map String [Plant]
garden students plants = M.fromList $ zipWith3 (\s pp0 pp1 -> (s, map parse pp0 ++ map parse pp1)) sortedStudents p1 p2

    where [line0, line1] = lines plants
          [p1, p2]       = [group 2 line0, group 2 line1]
          sortedStudents = sort students

lookupPlants :: String -> M.Map String [Plant] -> [Plant]
lookupPlants student garden = M.findWithDefault [] student garden


group :: Int -> [a] -> [[a]]
group _ [] = []
group n l
  | n > 0 = (take n l) : (group n (drop n l))
  | otherwise = error "Negative n"

parse :: Char -> Plant
parse 'C' = Clover
parse 'G' = Grass
parse 'R' = Radishes
parse 'V' = Violets
parse _   = error "unknown plant"