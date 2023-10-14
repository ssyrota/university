module Lib
    ( fibonacciReverseSubLists
    ) where

-- Розбити заданий список на кілька підсписків так, щоб в останньому підсписку опинилося f 0 елементів у
-- передостанньому – f 1 наступних з кінця елементів і т.д. Тут через f i позначено i-е число Фібоначчі.
-- (Відмінність від попередньої задачі полягає лише у зміні напрямку обробки списку на протилежний – з
-- кінця списку).

-- залишок не відрізати

-- Fibonacci function with optimization via accumulators mechanism for haskell.
fibonacci :: Integer -> Integer
fibonacci = helper 0 1

helper :: Integer -> Integer -> Integer -> Integer
helper acc1 acc2 n  | n == 0 = acc1
                    | n > 0 = helper (acc1 + acc2) acc1 (n - 1)
                    | n < 0 = helper acc2 (acc1 - acc2) (n + 1)

reverseList = foldl (flip (:)) []
removeNElemAtFirst n arr = reverseList . take(length arr - n) $ reverseList arr

fibonacciReverseSubListsHelper arr step accum 
                                        | null arr = accum
                                        | not . null $ arr = let 
                                            fibRes       = fromIntegral(fibonacci step);
                                            subSequence  = take fibRes arr
                                            remainder    = removeNElemAtFirst fibRes arr
                                            newStep      = step + 1
                                            newAccum     = accum ++ [subSequence]
                                            in fibonacciReverseSubListsHelper remainder newStep newAccum
                                        | otherwise = error "fibonacciReverseSubListsHelper error"


fibonacciReverseSubLists arr = reverseList(fibonacciReverseSubListsHelper(reverseList arr) 0 [])
