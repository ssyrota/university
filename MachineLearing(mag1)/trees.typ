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

= Decision tree and decision list

== Decision tree
/ Def: is a flowchart-like structure in which each node represents "test" on a feature(classification rule), each branch represents outcome of the "test node", each leaf represents resulting class.

Not good at extrapolation.

== Decision list(ordered list of decision rules)
/ Def: the list of Boolean functions which can be easily learnable from examples.
$ "if" f_1 "then" b_1 "else if " f_2 "then" b_2 ... $

Harder to present understandable hierarchical structure, because match must return a prediction. Easy to add rules.

== Which to use
Decision tree with configuration can imitate decision list and also provide hierarchical structures(easier to understand and compute)

