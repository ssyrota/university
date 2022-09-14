module Lib
    ( mainF
    ) where



sublistsLength sublistsCount list = length list `div` sublistsCount

mainF sublistsCount list = group(sublistsLength sublistsCount list) list


group :: Int -> [a] -> [[a]]
group _ [] = []
group n l
  | n > 0 = take n l : group n (drop n l)
  | otherwise = error "Negative or zero n"