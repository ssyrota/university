fibonacci(0,0):-!.
fibonacci(1,0):-!.
fibonacci(2,1):-!.
fibonacci(N,F):- N2 is N-2, N1 is N-1,
	fibonacci(N2,F2), fibonacci(N1,F1), F is F1+F2.

reverseHelper([],Z,Z).
reverseHelper([H|T],Z,Acc) :- reverseHelper(T,Z,[H|Acc]).
reverseList(I,O):- reverseHelper(I, O, []).


removeNElemAtFirst(I, N, []) :-
    length(I, A),
    A<N.
removeNElemAtFirst(I, N, O) :-
    length(Prefix, N),
    append(Prefix, O, I).


getNElemsAtFirst(I, N, I) :-
    length(I, A),
    A<N.
getNElemsAtFirst(I, N, O) :-
    length(O, N),
    append(O, _, I).




fibonacciReverseSubListsHelper([], _, []):-!.
fibonacciReverseSubListsHelper(IN_ARR, FIB_STEP, [PART_OUT_ARR|OUT_ARR]):- 
    fibonacci(FIB_STEP, FIB_RES),
    NEW_FUB_STEP is FIB_STEP+1,
    getNElemsAtFirst(IN_ARR, FIB_RES, PART_OUT_ARR_TO_REVERSE),
    reverseList(PART_OUT_ARR_TO_REVERSE,PART_OUT_ARR),
    removeNElemAtFirst(IN_ARR, FIB_RES, NEW_IN_ARR),
    fibonacciReverseSubListsHelper(NEW_IN_ARR, NEW_FUB_STEP, OUT_ARR).

fibonacciReverseSubLists(I,O):- 
    reverseList(I, ReversedI),
    fibonacciReverseSubListsHelper(ReversedI, 1, ReversedOut),
    reverseList(ReversedOut, O).