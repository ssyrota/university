module Lib
    ( getRepeatedTwice
    ,removeRepeatedElements
    ) where

import Data.List

count :: [Int] -> [(Int, Int)]
count = map (\ x -> (head x, length x)) . group . sort

getRepeatedTwice arr = map fst . filter(\x-> snd x == 2 ) $ count arr
removeRepeatedElements arr1 arr2 = filter (`notElem` getRepeatedTwice arr2) arr1
