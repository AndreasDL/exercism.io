module LeapYear (isLeapYear) where

isLeapYear :: Integer -> Bool
isLeapYear year = year `divis` 4 
	&& ( not(year `divis` 100) || year `divis` 400 )

divis :: Integer -> Integer -> Bool
divis year x = (year `mod` x == 0)