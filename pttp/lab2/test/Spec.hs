import Lib (fibonacciReverseSubLists)

assert :: Bool -> String -> String -> IO ()
assert test passStatement failStatement = if test
 then putStrLn passStatement
 else putStrLn failStatement


main :: IO ()
main = do
 putStrLn "Running tests..."
 assert (null (fibonacciReverseSubLists [])) "passed '1'" "FAIL: '1'"
 assert ([[5,4,3,2,1],[10,9,8,7,6],[13,12,11],[15,14],[16],[17],[]] == fibonacciReverseSubLists [1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17]) "passed '2'" "FAIL: '2'"
 assert ([[1],[2],[3],[]] == fibonacciReverseSubLists [1,2,3]) "passed 3'" "FAIL: '3'"
 assert ([[[1,2,3]],[[1,2,3],[1,2,3]],[[1,2,3]],[[1,2,3]],[]] == fibonacciReverseSubLists [[1,2,3],[1,2,3],[1,2,3],[1,2,3],[1,2,3]]) "passed 4'" "FAIL: '4'"
 putStrLn "done!"