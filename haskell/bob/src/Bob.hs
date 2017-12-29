module Bob (responseFor, main) where

import Data.Char

responseFor :: String -> String
responseFor xs
    | "" == trimmed                    = "Fine. Be that way!"
    | '?' == last trimmed              = if isUpper then "Calm down, I know what I'm doing!" else "Sure."
    | isUpper                          = "Whoa, chill out!"
    | not hasLetters && not hasNumbers = "Fine. Be that way!"
    | otherwise                        = "Whatever."
    where 
        trimmed    = filter (\x -> x `notElem` [' ','\t','\n']) xs
        hasLetters = containsLetters trimmed
        hasNumbers = containsNumbers trimmed
        isUpper    = trimmed == upperString trimmed && hasLetters


containsLetters :: [Char] -> Bool
containsLetters [] = False
containsLetters text = length clean > 0
    where clean = filter (\x -> x `elem` (['a'..'z']++['A'..'Z']) ) text

containsNumbers :: [Char] -> Bool
containsNumbers [] = False
containsNumbers text = length clean > 0
    where clean = filter (\x -> x `elem` ['0'..'9']) text

upperString :: [Char] -> [Char]
upperString text = map toUpper text

main = do 
    putStrLn $ upperString "hoi"
    putStrLn $ show $ "hoi" == upperString "hoi"
    putStrLn $ show $ filter (\x -> x `notElem` [' ','\t','\n']) "h hoi\n\n   hoi"