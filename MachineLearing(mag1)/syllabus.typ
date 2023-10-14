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

= Introduction

/ $Y$: - the variable that we predict.

/ Feature($x$): the variable in the data vector. Types:
1. Numerical
2. Categorical
  - Ordinal
  - Nominal

/ Hyperparam: meta parameter for model. Model do not learn it.

== Supervised
Solves regression and classification tasks.
$ X arrow.long F arrow.long y $

/ Regression model: predicts continuous values.
/ Classification model: predicts categorical values.

== Unsupervised

$ X arrow.long F arrow.long X' $

== Reinforement learning

_TODO_

= Optimisation and loss function

== Gradient descent
/ Gradient($nabla f$): defines direction and rate of fastest increase of scalar-valued differentiable function $f$.
_Example for gradient in cartesian coordinate system f:_
$ nabla f = (diff f )/ (diff x) i + (diff f )/ (diff y) i + (diff f )/ (diff z) k $

/ Gradient descent: iterative optimization algorithm of the first order to find the local minimum of the function.
_Stop criteria_ for the gradient descent can be a threshold for the gradient value.

== Optimization
/ Optimisation target: minimize loss function.

Simple example of the loss function is a MSE.
/ Mean squared error(MSE): measures the average of squeared errors. 

$ "MSE" = 1 / N sum_((x, y) in D)(y - "prediction"(x))^2 $

Iteration step for model: $x^(i+1)= x^i - h (diff f)/ (diff x^i)$

Where $(diff f)/ (diff x^i)$ is a _gradient_.



Possible data slices for 1 iteration there model updates params:
- simple - 1 full dataset
- stochastic - 1 record
- mini-batch - batch of random examples(e.g. 10-1000)


== Regularization
_TBD_