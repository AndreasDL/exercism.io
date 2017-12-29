module RailFenceCipher (encode, decode, main) where

import Data.List
import Data.Function

encode :: Int -> String -> String
encode r text = 
    let charLines = textSequence r $ clean text 
        output    = sortBy(compare `on` fst) charLines
    in  map snd output

clean :: String -> String 
clean = filter(' ' /=)

lineSequence :: Int -> [Int]
lineSequence 0 = []
lineSequence r = cycle $ [0..r-1] ++ [r-2, r-3..1]

textSequence :: Int -> [a] -> [(Int, a)]
textSequence r text = zip seq text 
    where seq = take (length text) (lineSequence r)


decode :: Int -> String -> String
decode r text = 
    let charLines = textSequence r [0..length text-1]
        spos      = zip (map snd $sortBy(compare `on` fst) charLines) [0..length text -1]
        pos       = map snd $sortBy(compare `on` fst) spos
    in 
        reOrder pos text

reOrder :: [Int] -> [Char] -> [Char]
reOrder (x:xs) cs = [cs!!x] ++ reOrder xs cs 
reOrder _ _  = []


-- debugging the sequences 
main = do
    let r = 3
    let text = "TEITELHDVLSNHDTISEIIEA"
    let charLines = textSequence r [0..length text-1]
    let spos      = zip (map snd $sortBy(compare `on` fst) charLines) [0..length text -1]
    let pos       = map snd $sortBy(compare `on` fst) spos
    
    putStrLn text
    putStrLn $ show spos
    putStrLn $ show pos
    putStrLn $ decode r text