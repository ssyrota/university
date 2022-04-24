import Lib (getRepeatedTwice, removeRepeatedElements)

assert :: Bool -> String -> String -> IO ()
assert test passStatement failStatement = if test
 then putStrLn passStatement
 else putStrLn failStatement


main :: IO ()
main = do
 putStrLn "Running tests..."
 assert (null (removeRepeatedElements [1,2,3,4] [1,2,3,4,1,2,3,4])) "passed '1'" "FAIL: '1'"
 assert ([1,2,3,4] == getRepeatedTwice [1,2,3,4,1,2,3,4]) "passed '2'" "FAIL: '2'"
 assert ([1,2,3,4] ==  removeRepeatedElements [1,2,3,4] [1,1,1,1,2,3,4]) "passed '3'" "FAIL: '3'"
 assert ([2,3] ==  removeRepeatedElements [1,2,3,4] [1,1,2,4,4]) "passed '4'" "FAIL: '4'"
 assert ([4] ==  removeRepeatedElements [1,2,3,4] [1,1,2,2,3,3,4]) "passed '5'" "FAIL: '5'"
 putStrLn "done!"