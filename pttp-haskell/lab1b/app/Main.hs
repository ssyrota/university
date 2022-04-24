module Main where
import Lib (getRepeatedTwice)

main :: IO ()
main = print . getRepeatedTwice $ [1,2,3,4]