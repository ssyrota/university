module Lib
    ( removeElemAtCubeIndex
    ) where


removeElemAtCubeIndex :: [a] -> [a]
removeElemAtCubeIndex [] = []
removeElemAtCubeIndex xs = [y | (x,y) <- addIndex xs, x `notElem` take (sqrtInt x) cubes]

sqrtInt = round . sqrt
cubes = [x ^ 3 | x <- [1..]]
addIndex = zip [0..]