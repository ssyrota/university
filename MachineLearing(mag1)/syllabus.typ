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

/ ML: process of training a piece of software, called model, to make useful predictions or to generate content from data.

Types:
- Supervised learning(two most common use cases - regression and classification)
- Unsupervised learning(clusterization common)
- Reinforcement learning(penalties and rewards->generated policy)
- Generative AI(generate something from input)

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

Iteration step for model paraneter: $ Theta^(i+1)= Theta^i - h (diff f)/ (diff Theta^i) $

Where $h$ is a _learning step_.

How much $i$'s would be in a N dataset? 

Depends, upon model converges. Possible data slices for 1 ephoch:
- simple - 1 full dataset
- stochastic - 1 record
- mini-batch - batch of random examples(e.g. 10-1000). This approach can support struggling with local minimums.

TODO

#link("https://stats.stackexchange.com/questions/164876/what-is-the-trade-off-between-batch-size-and-number-of-iterations-to-train-a-neu")[
  [Comparison of batch sizes link]
]

/ epoch: one pass of all the training examples
/ batch size: the number of training examples in one pass. The higher the batch size, the more memory space you'll need.
/ iterations: number of batches in epoch. each iteration adjusts model's parameters.

$ 
(diff f)/(diff Theta_i) = 1/(2N) sum_(i=1)^n (( sum_(j=1)^m (Theta_j x_j) - y_i)^2)'
$
_note_: N is the iteration dataset(or batch) size, $x_j$ is a point in vector, $Theta j$ is the parameter value that is const if not differentiated, $y_i$ is a constant for each i.

Let's simplify function for two parameters and 3 data slices:
$
1/(2N) sum_(i=1)^(3) ( (sum_(j=1)^(2) (Theta_j x_j) - y_i)^2)
$

Simplify each $i$ argument:
$
( sum_(j=1)^(2) (Theta_j x_j) - y_i)^2 "=" ( Theta_0 x_0 + Theta_1 x_1 - y_i)^2
$

$Theta_1 x_1 "and" y_i$ is a constants if we differentiate by $Theta_0$, so we have:
$(( Theta_0 x_0 + C_i)^2)'$, also: $(Theta_0 x_0 + C_i)$ is an inside f $v$.

With _formula of compound derivative_ $(u(v))'=u'(v) * v'$
$
(( Theta_0 x_0 + C_i)^2)' = 2(Theta_0 x_0 + C_i) (x_0)
$

_TODO_

Final differential formula:
$
(diff f) / (diff Theta_i) = x_i  1/N sum_(i=1)^n ( sum_(j=1)^m Theta_j x_j - y_i)
$

Process is simple, count _gradient_ for each _parameter_ and change parameters by gradient descent.
$ 
Theta_i arrow.long "gradient" arrow.long Theta_i^T
$


== Regularization
_TODO_

== Linear regression
When we have not linear plot, to solve this linear regression problem we can add additional polynomial($x^2$) or functional($sin(x)$, $sqrt(x)$) features.

#figure(
  image("./img/linear_polynom.png", width: 70%),
  caption: [
    Synthetic features for regression with linear $Theta$ params
  ],
)

How to choose function to create additional features? Intuitively as a hyperparams. There are automatic methods to make models - feature selection approach.

=== Normal Linear Regression Model
TODO
https://www.statlect.com/fundamentals-of-statistics/normal-linear-regression-model

= Data
Dataset should be divided minimum for train(60), validation(20) and final test(20). This divided datasets must not have semantic intersections(same people, same cars, same buildings etc.).

Cross validation - method, which on small dataset find conceptual ML model that possibly solves task.
TODO

== Underfitting and overfitting

/ Underfitting: model performs poorly
Causes: Model is too weak
How to beat: Make model more complex

/ Overfitting: model performs well on training data, but not in evaluation
Causes: To complex model, too few data
How to beat: Simplify model, add data

Problem is: Bias vs variance tradeoff.

/ Regulaization: technique of discouraging learning a more complex or flexible model, so as to avoid the risk of overfitting.

//  bias reduces, but variance increases.?

= Lib
- https://www.statlect.com/
