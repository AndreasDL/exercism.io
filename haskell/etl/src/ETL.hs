module ETL (transform) where

import Data.Map (Map, keys, fromList, toList)
import Data.Char (toLower)

transform :: Map a String -> Map Char a
transform legacyData = fromList $ concatMap (\(k, v) -> [ (toLower(x),k) | x <- v ]) $ toList legacyData 
