module Main where
import Lib ( removeElemAtCubeIndex )

main :: IO ()
main = print . removeElemAtCubeIndex $ [0,1,2,3,4,5,6,7,8]
