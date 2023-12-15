#set heading(numbering: "1.")
#set text(
  font: "Times New Roman",
  size: 11pt
)
#set page(
  paper: "a4",
  margin: (x: 1.8cm, y: 1.4cm),
  height: auto
)
#set par(
  justify: true,
)

= Formal grammar

/ Def: a formal way of description language, there are generative and analytical types.

/ Terminal symbol: an object presented in words of language. Has semantic value for language user.

/ Non-Terminal symbol: an meta-object of language(formula, command, ...). Has no concrete meaning.

= Generative grammar 

/ Grammar: describes how to form strings from alphabet of a formal language that are valid accordingly to language syntax. Does not describe the meaning of the strings - only their form.

/ alphabet: a set of atomic symbols, which build words.

Grammar consists of:
1. Terminal alphabet $sum$
2. Non-Terminal alphabet $N$
3. Inference rules $P: "left"->"right"$

left: nonempty sequence of terminal and non-terminals, has at least 1 non-terminal.

right: sequence of terminal and non-terminals

4. Define first non-terminal

== Chomsky hierarchy  

/ Def: hierarchy in which lower is more strict, but more suitable for parsing.

0. Not limited
1. Context-sensitive
2. Context-free. (only non-terminals at left P)
3. Regular(right has one non-terminal and maybe terminals)

= Chomsky Normal Form

For all P must:
$A->B C$
$A->alpha$

Must exist rule if language has epsilon:
$S->epsilon$

Where: $A, B, C in "NonTerm"$, $alpha in "Term"$, $S "is" "start symbol"$, $epsilon "is empty string"$

= Bottom-up parsing

= offtop Dynamic programming

= CYK
/ Def: parsing algorithm for context-free grammars. It employs bottom-up parsing and dynamic programming.
Operates on Chomsky Normal Form. Noted because of good worst-case performance.

https://www.borealisai.com/research-blogs/tutorial-15-parsing-i-context-free-grammars-and-cyk-algorithm/