module Lib
    ( getRepeatedThreeTimes
    ,removeRepeatedElements
    ) where

import Data.List

count :: [Int] -> [(Int, Int)]
count = map (\ x -> (head x, length x)) . group . sort

getRepeatedThreeTimes arr = map fst . filter(\x-> snd x == 3 ) $ count arr
removeRepeatedElements arr1 arr2 = filter (`elem` getRepeatedThreeTimes arr2) arr1
