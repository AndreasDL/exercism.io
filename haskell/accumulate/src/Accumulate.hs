module Accumulate (accumulate) where

accumulate :: (a -> b) -> [a] -> [b]
accumulate = map --we don't even need to provides names