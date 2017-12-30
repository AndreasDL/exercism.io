module Diamond (diamond, main) where

import Data.Char (ord)


diamond :: Char -> [String]
diamond c = map (line c) [0..n+n]
    where n = ord c - ord 'A'
    
line :: Char -> Int -> String
line c i
    | i > n+n            = ""
    | i == 0 || i == n+n = spaces(bef !! i) ++ [chars !! i] ++ spaces(bef !! i)
    | otherwise          = spaces(bef !! i) ++ [chars !! i] ++ spaces(mid !! i) ++ [chars !! i] ++ spaces(bef !! i)
    where 
        chars = ['A'..c] ++ (tail $ reverse ['A'..c])
        n     = ord c - ord 'A'
        bef   = [n,n-1..0] ++ [1..n]
        mid   = 0:[1,3..n+n] ++ (tail $ reverse (0:[1,3..n+n]))

spaces :: Int -> [Char]
spaces n = replicate n ' '

main = do 
    let c = 'C'
    let n = ord c - ord 'A'
    putStrLn $ show $ ['A'..c] ++ (tail $ reverse ['A'..c])
    putStrLn $ show $ n
    putStrLn $ show $ [n,n-1..0] ++ [1..n]
    putStrLn $ show $ 0:[1,3..n+n] ++ (tail $ reverse $ 0:[1,3..n+n])

    putStrLn $ line 'A' 1
    putStrLn $ line 'B' 3
    putStrLn $ line 'C' 3
    putStrLn $ line 'D' 3
    
    mapM_ putStrLn $ diamond 'A'
    mapM_ putStrLn $ diamond 'B'
    mapM_ putStrLn $ diamond 'C'
    mapM_ putStrLn $ diamond 'D'
    mapM_ putStrLn $ diamond 'E'