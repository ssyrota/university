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

= Naive bayes

/ Category: Probabilistic classifier.

Applies bayes theorem with strong naive assumption that features are independent.

$ P(A|B)$ - probability of A given B true.

$ P(y|x_1,x_2,...x_n) = (P(y) P(x_1,x_2,...x_n | y) )/P(x_1,x_2,...x_n) $

Using naive assumption of independence:
$ P(x_1, x_2,...x_n | y) =  Pi P(x_i | y)$

?Why this works?
https://www.cs.unb.ca/~hzhang/publications/FLAIRS04ZhangH.pdf
?Gaussian Naive Bayes?
?central limit theorem?
?probability vs likelyhood?
z-score(or standard score) = $(X-Mu)/sigma$

How to deal with continuous features?