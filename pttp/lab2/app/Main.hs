module Main where

import Lib

main :: IO ()
main = print . fibonacciReverseSubLists $ [[1,2,3],[1,2,3],[1,2,3],[1,2,3],[1,2,3]]
