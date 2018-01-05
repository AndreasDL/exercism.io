module TwelveDays (recite) where

recite :: Int -> Int -> [String]
recite start stop = map line [start, start+1..stop]

line :: Int -> String
line x = "On the " ++ days x ++ " day of Christmas my true love gave to me, " ++ items 
    where items = concatMap item [x-1, x-2..1] ++ (if x > 1 then "and " else "") ++ item 0

item :: Int -> String
item x 
    | x == 0    = itemData !! 0 ++ "."
    | otherwise = counts (x+1) ++ " " ++ itemData !! x ++ ", "
    where itemData = [ "a Partridge in a Pear Tree",
                       "Turtle Doves",
                       "French Hens",
                       "Calling Birds",
                       "Gold Rings",
                       "Geese-a-Laying",
                       "Swans-a-Swimming",
                       "Maids-a-Milking",
                       "Ladies Dancing",
                       "Lords-a-Leaping",
                       "Pipers Piping",
                       "Drummers Drumming" ]

days :: Int -> String
days 1  = "first"
days 2  = "second"
days 3  = "third"
days 4  = "fourth"
days 5  = "fifth"
days 6  = "sixth"
days 7  = "seventh"
days 8  = "eighth"
days 9  = "ninth"
days 10 = "tenth"
days 11 = "eleventh"
days 12 = "twelfth"
days _  = ""

counts :: Int -> String
counts 1  = "one"
counts 2  = "two"
counts 3  = "three"
counts 4  = "four"
counts 5  = "five"
counts 6  = "six"
counts 7  = "seven"
counts 8  = "eight"
counts 9  = "nine"
counts 10 = "ten"
counts 11 = "eleven"
counts 12 = "twelve"
counts _  = ""