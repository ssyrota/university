build:
	stack build
run: 
	stack exec lab2-exe
hotbuild: 
	stack build --file-watch --exec "make run" 
resolve-prefixies: 
	gen-hie > hie.yaml
testing: 
	stack test