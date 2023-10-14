module Main where
import Lib (getRepeatedThreeTimes)

main :: IO ()
main = print . getRepeatedThreeTimes $ [1,2,3,4]