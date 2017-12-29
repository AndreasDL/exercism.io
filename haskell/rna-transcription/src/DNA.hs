module DNA (toRNA) where

toRNA :: String -> Maybe String
toRNA xs
        | length xs /= length mapped = Nothing
        | otherwise = Just mapped
    where
        mapped  = [ complement c | c <- xs , c `elem` "GCTA" ]

complement :: Char -> Char
complement c 
    | c == 'G' = 'C'
    | c == 'C' = 'G'
    | c == 'T' = 'A'
    | c == 'A' = 'U'