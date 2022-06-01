list_member(X,[X|_]).

list_member(X,[_|TAIL]) :- list_member(X,TAIL).


count(_, [], 0).

count(X, [X | T], N) :-
    !, count(X, T, N1),
    N is N1 + 1.

count(X, [_ | T], N) :-
    count(X, T, N).


count_only_three_times_in_array(X, ARR) :- N is 3, count(X, ARR, N).


filter_only_three_times_helper(_,[],[]).

filter_only_three_times_helper(IN_ARR,[CURR_ELEM|ITERABLE_ARR], [CURR_ELEM|OUT_ARR]) :- 
    count_only_three_times_in_array(CURR_ELEM, IN_ARR), 
    filter_only_three_times_helper(IN_ARR, ITERABLE_ARR, OUT_ARR).

filter_only_three_times_helper(IN_ARR,[_|ITERABLE_ARR], OUT_ARR) :- 
    filter_only_three_times_helper(IN_ARR, ITERABLE_ARR, OUT_ARR).


left_only_three_times_duplicate_from_list(IN_ARR,OUT_ARR):- 
    filter_only_three_times_helper(IN_ARR,IN_ARR,OUT_ARR).


value_in_array_of_three_times_duplicates_from_arr(X, ARR):- 
    left_only_three_times_duplicate_from_list(ARR, DUPLICATES), list_member(X, DUPLICATES).


mainF([],_,[]).

mainF([X|IN_1], IN_2, R):- 
    value_in_array_of_three_times_duplicates_from_arr(X, IN_2),
    mainF(IN_1, IN_2, R).

mainF([X|IN_1], IN_2, [X|R]):- 
    mainF(IN_1, IN_2, R).