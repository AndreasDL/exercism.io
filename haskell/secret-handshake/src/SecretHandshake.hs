module SecretHandshake (handshake) where

import Data.Bits ((.&.), xor)

handshake :: Int -> [String]
handshake n
    | n .&. 16 == 16 = reverse $ handshake (n `xor` 16)
    | n .&.  8 ==  8 = handshake (n `xor` 8) ++ ["jump"]
    | n .&.  4 ==  4 = handshake (n `xor` 4) ++ ["close your eyes"]
    | n .&.  2 ==  2 = handshake (n `xor` 2) ++ ["double blink"]
    | n .&.  1 ==  1 = ["wink"]
    | otherwise = []