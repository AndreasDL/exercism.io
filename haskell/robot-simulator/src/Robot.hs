module Robot
    ( Bearing(East,North,South,West)
    , bearing
    , coordinates
    , mkRobot
    , simulate
    , turnLeft
    , turnRight
    ) where

data Bearing = North
             | East
             | South
             | West
             deriving (Eq, Show, Enum, Bounded)

data Robot = Robot {
    bearing     :: Bearing,
    coordinates :: (Integer, Integer)
}

mkRobot :: Bearing -> (Integer, Integer) -> Robot
mkRobot direction coordinates = Robot direction coordinates

simulate :: Robot -> String -> Robot
simulate robot ('R':xs) = simulate (turnRobotRight robot) xs
simulate robot ('L':xs) = simulate (turnRobotLeft  robot) xs
simulate robot ('A':xs) = simulate (advanceRobot   robot) xs
simulate robot ""       = robot
simulate robot _        = error "unknown Pattern"

turnLeft :: Bearing -> Bearing
turnLeft b 
    | b == minBound = maxBound
    | otherwise     = pred b

turnRight :: Bearing -> Bearing
turnRight b
    | b == maxBound = minBound
    | otherwise     = succ b

turnRobotLeft :: Robot -> Robot
turnRobotLeft  (Robot b c) = Robot (turnLeft  b) c

turnRobotRight :: Robot -> Robot
turnRobotRight (Robot b c) = Robot (turnRight b) c

advanceRobot :: Robot -> Robot
advanceRobot (Robot b (x,y))
    | b == North = Robot b (x   , y+1)
    | b == East  = Robot b (x+1 , y  )
    | b == South = Robot b (x   , y-1)
    | b == West  = Robot b (x-1 , y  )