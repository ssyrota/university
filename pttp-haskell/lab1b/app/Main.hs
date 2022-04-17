module Main where

import Lib ( removeRepeatedElements )

main :: IO ()
main = print . removeRepeatedElements $ [] []
