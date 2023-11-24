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

= Into

/ Definition: uses algebraic structures for modeling data and defining queries on it with a well founded semantics.

/ Main purpose: is to define operators which accepts one or more relations and produces another relation

== Operators

=== Set operators
1. Selection - filter
2. Projection - select only
3. Cartesian product - join. (relations must not have same attribute names). Produces flattened tuple.(not (A, B), (B, A))
4. Union(relations must be union compatible)
5. Difference(relations must be union compatible)
6. Rename
7. $Theta$-join. (a $F$ b)
8. Semijoin or restriction (same as join, but as "exists". without selecting from B attributes)
9. Antijoin (not exists)
10. Outer joins

"relations must be union compatible" = relations must have the same attributes

== Algebraic structure
Algebraic structure consists of non-empty set $"domain"$ and a collection of operations $A$ and finite set of $"identities"$.

$"Identity"$ - equality relating one mathematical expression A to another mathematical expression B in all valid inputs(variables).