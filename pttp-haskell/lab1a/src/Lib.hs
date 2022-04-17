module Lib
    ( removeElemAtCubeIndex
    ) where


removeElemAtCubeIndex :: [a] -> [a]
removeElemAtCubeIndex [] = []
removeElemAtCubeIndex xs = [y | (x,y) <- addIndex xs, x `notElem` take x cubes]

cubes = [x ^ 3 | x <- [1..]]
addIndex = zip [0..]