import Lib ( removeElemAtCubeIndex )

assert :: Bool -> String -> String -> IO ()
assert test passStatement failStatement = if test
 then putStrLn passStatement
 else putStrLn failStatement


main :: IO ()
main = do
 putStrLn "Running tests..."
 assert (null (removeElemAtCubeIndex [])) "passed '1'" "FAIL: '1'"
 assert (["a", "c", "d"] == removeElemAtCubeIndex ["a", "b", "c", "d"]) "passed '2!'" "FAIL: '2!'"
 assert ([0,2,3,4,5,6,7,9] == removeElemAtCubeIndex (take 10 [0..]) ) "passed '3'" "FAIL: '3'"
 assert ([0] == removeElemAtCubeIndex (take 2 [0..]) ) "passed '4!'" "FAIL: '4!'"
 assert ([1,4,5,6] == removeElemAtCubeIndex [1,1,4,5,6] ) "passed '5!'" "FAIL: '5!'"
 putStrLn "done!"