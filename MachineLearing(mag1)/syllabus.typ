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

Online and batch learning.

/ $Y$: - the variable that we predict.

/ Feature($x$): the variable in the data vector. Types:
1. Numerical
2. Categorical
  - Ordinal
  - Nominal

/ Hyperparam: meta parameter for model. Model do not learn it.

== Data mining
/ Data mining: applying ML techniques to dig into large amounts of data can help discover factors, that are not immediately apparent.

== Supervised
Solves regression and classification tasks.
$ X arrow.long F arrow.long y $

/ Regression model: predicts continuous values.
/ Classification model: predicts categorical values.

== Unsupervised

$ X arrow.long F arrow.long X' $

== Reinforement learning

The learning system, called agent can observe the environment, select and perform actions, and get rewards or penalties. It must learn then by itself what is the best strategy, called a policy to get the most reward over the time.

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

/ bias:
/ variance:

Regularization formula L2(makes features balance):
$ L = 1/(2N) sum_(j=1)^(n)(y(x_i)-y_i)^2 + Lambda sum_(i=1)^(m) (Theta_i^2) arrow.long min_Theta $

Regularization formula L1(makes feature selection):
$ L = 1/(2N) sum_(j=1)^(n)(y(x_i)-y_i)^2 + Lambda sum_(i=1)^(m) (|Theta|) arrow.long min_Theta $


= Classification
$y in {1,2,3,..,k}$

/ Task: make a function, that separates known classes.

#figure(
  image("./img/classification.png", width: 70%),
  caption: [
    Classification visualization
  ],
)

Accordingly to image, linear regression is not suitable for this type of task(especially right)

Firstly, we will solve binary classification task{0, 1}. Model will have 1 output - probability of x is from class 1.

== Logistic regression
/ Logistic regression: type of regression that predicts a probability of an outcome given one or more independent variables. With a threshold returned probability can be mapped to a discrete value.

#figure(
  image("./img/logistic.png", width: 70%),
  caption: [
    Logistic regression
  ],
)

=== Formula
Logistic regression is a S-shaped curve:
$ y = 1/(1+ e^(- sum_(i=1)^m Theta_i X_i + Theta_0)) $

=== Loss function
/ BCE(Binary cross entropy) loss function: $ cases(
  - log(p_i) "," y_i = 1,
  - log(1 - p_i) "," y_i=0
) $
$p_i$ - model output probability for i example. (class 1)

$ "BCE" = - y_i log(p_i) - (1-y_i)log(1-p_i) $

$ cases(
  "TP" y_i"," p_i = {1"," 1} "BSE" = 0,
  "TN" y_i"," p_i = {0"," 0} "BSE" = 0,
  "FN" y_i"," p_i = {1"," 0} "BSE" -> inf,
  "FP" y_i"," p_i = {0"," 1} "BSE" -> inf,
) $

Loss for gradient:
$ L = - 1/(N) sum_(i=1)^(n)(y_i log(p_i) + (1-y_i)log(1-p_i)) arrow.long min $

$ 
(diff f) / (diff Theta_j) = - 1/(N) sum_(i=1)^(n)(y_i log(p_i) + (1-y_i)log(1-p_i))'
$


== Metrics
To define success of model *metrics* are used. $A=(N_("correct")/N) 100%$

/ Precision: TP/(TP+FP). How model is confident for class a.
/ Recall: TF/(TF+FN). Which coverage for class a.



= Regressions comparison
#table(
  columns: (auto, auto, auto),
  inset: 10pt,
  align: horizon,
  [*Criteria*], [*Linear regression*], [*Logistic*],
  [Regression plot dependency], [Straight line], [S-shaped curve],
  [Output type], [Continuous], [Probability (0,1) of the value \
                                from the finite category],
  [Target], [Give a most precise number], [Give a probability of belonging to category],
  [Usecase], [Predict house prise], [Define type of a tumor, is price > 500k\$],
  [Distribution type], [Normal], [Binomial]
)
Linear examples:
- Predicting the height of an adult based on the mother’s and father’s height
- Predicting pumpkin sales volume based on the price, time of year, and store location

Logistic examples:
- Predicting if a person will get a disease based on status, salary, genetics
- Prediction if a person will quit a job based on meetings, pull requests, office time.
- Predicting the marriage of a person based by car, salary, outlook, office time, education, country.


= Removing correlated features
Before train model it may be worth to perform "dimensionality reduction" to save space, without loosing too much information

= Lib
- https://www.statlect.com/
- https://towardsdatascience.com/
- https://towardsdatascience.com/logistic-regression-detailed-overview-46c4da4303bc
- https://en.wikipedia.org/wiki/Logistic_regression
- file:///Users/s.syrota/Downloads/Fundamentals%20of%20probability%20and%20statistics%20for%20engineers%20(T.%20T.%20Soong)%20(Z-Library).pdf
- file:///Users/s.syrota/Downloads/Essential%20Math%20for%20Data%20Science%20(Fifth%20Early%20Release)%20(Thomas%20Nield)%20(Z-Library).pdf
299 page