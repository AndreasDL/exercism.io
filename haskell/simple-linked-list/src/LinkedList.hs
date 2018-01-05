module LinkedList
    ( LinkedList
    , datum
    , fromList
    , isNil
    , new
    , next
    , nil
    , reverseLinkedList
    , toList
    ) where

data LinkedList a = Nil | LinkedList {
    datum :: a,
    next  :: LinkedList a
} deriving (Eq, Show)

nil :: LinkedList a
nil = Nil

new :: a -> LinkedList a -> LinkedList a
new x linkedList = LinkedList x linkedList

fromList :: [a] -> LinkedList a
fromList []     = Nil
fromList (x:xs) = LinkedList x $fromList xs

isNil :: LinkedList a -> Bool
isNil Nil = True
isNil _   = False

reverseLinkedList :: LinkedList a -> LinkedList a
reverseLinkedList ll = (fromList.reverse.toList) ll 

toList :: LinkedList a -> [a]
toList Nil = []
toList ll  = datum ll : (toList $ next ll)