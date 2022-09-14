build:
	stack build
run: 
	stack exec lab1b-exe
hotbuild: 
	stack build --file-watch --exec "make run" 
resolve-prefixies: 
	gen-hie > hie.yaml
testing: 
	stack test