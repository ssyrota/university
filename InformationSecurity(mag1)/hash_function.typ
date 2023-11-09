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
/ Hash function: F that transforms arbitrary length data into fixed size data(digest, hash).

Practically output range is 128-512 bits

Requirements:
0. Function is fast and have small memory consuming
1. one wayness. Computational infeasible to find x from y=h(x). It's map.
1.1 value distribution is equal $2^(-n)$
2. weak collision resistance(first kind). Computational infeasible to find x_2 from y=h(x_1)=h(x_2)
3. strong collision resistance(second kind). Computational infeasible to find and x_1, x_2 from y=h(x_1)=h(x_2)

Main target: malicious adversary cannot replace or modify data without changing it digest. Function should have behavior like random function.

== Difficult or Computational infeasible
Not solvable in asymptotic polynomial time.

== Preimage resistance
Hash function must be strength to find preimage of hash.

Use cases:
- find hashed password by brute force

== weak collision(second preimage resistance)
Given $y=h(x_1)$, computationally infeasible to find $x_2: y=h(x_2)$

Use cases:
- fake signature

== strong collision 
Computationally infeasible to find $x_2, x_1: y=h(x_2)$=h(x_1)

Use cases:
- find two documents with the single hash

Requires to compute 2^(N/2) to find x_2 and x_1.

= Birthday problem
In set of n randomly chosen people, to get the probability of two has same birthday 50%+ required only 23 people.

$ "no overlap at all" P_(0) = 1*((365-1)/365)*((365-2)/365)...*((365-i)/365) $
$ "at least 1 overlap" P_(1) = 1-P_0 $

For 23 people $ P_0 = 0.4972 arrow.long P_1 = 0.5028 $

Another proof:
n people

$P(1) = 1 - P_0$, $ P_0 = V_("no pair")/V_("all") $

$ V_("no_pair") = P_365^n = (365)!/((365-n)!) $

$ V_("all") = 365^n $

$ P_0 = (P_365^n)/(365^n) = (365)!/((365-n)! 365^n) $

$ n=23 -> P_0~50% $
"whoop"

// TODO: make simple and counterintuitive tasks
/ Permutation: count of rearrangement combinations. The number of permutations $n$ is $ P_n = n! $

/ Partial permutation: count of rearrangement combination of subset $k$ elements from set $n$.
 $ P_n^k = n!/(n-k)! $


/ Combination: is a k-element subset of $s$, the elements in combination are not ordered. (k! means number of permutations in each k-length subset of S)

$ C_n^k = (n!)/((n-k)!k!) $

== Birthday attack

= Based on block ciphers

== Use cases

- Hash table(often used non-cryptographic hash functions) and indexing
- Fingerprinting and verifying the integrity of data
- Identifier