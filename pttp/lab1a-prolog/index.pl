filterArrByCubeIndexes([],[],_).
filterArrByCubeIndexes([_],[],SI):- isNotCube(SI).
filterArrByCubeIndexes([A],[A],SI):- isNotCube(SI).
filterArrByCubeIndexes([_|T], L, SI):- isNotCube(SI), R1 is SI+1, filterArrByCubeIndexes(T, L, R1).
filterArrByCubeIndexes([X|T], [X|L], SI):- R1 is SI+1, filterArrByCubeIndexes(T, L, R1).

isNotCube(X) :- X=\=round(X**(1/3))*round(X**(1/3))*round(X**(1/3)).
