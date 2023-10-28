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
  leading: 1.5em
)

= Euler function
$ phi(n) = "len"({1,2,3..n}, gcd(k,n)=1) $

Phi is multiplicative function
$ phi(a b) = phi(a) phi(b) $

$ phi(n) = p_1^(k_1-1)(p_1-1) p_2^(k_2-1)(p_2-1) ...$
Where p is prime number from factorization n.

_Example_
$ phi(54) = phi(2 * 3^3) = phi(2) * phi(3^3) $

/ Euler's theorem: If $a "and" p$ is coprime, than $a^(phi(n)) eq.triple 1 mod n$.

= Modular arithmetics

== Congruence

  We say that 3 is congruent to 15 by modulo 12, written $15 eq.triple 3 (mod 12)$

  / Coprime: two integers GCD is 1.

== Fermat's little theorem
Special case of euler theorem.

  / Theorem: If $p$ is a prime number, then for any integer $a$, the number $a^p - a$ is an integer multiple of $p$

  $ a^p eq.triple a mod p $

  If a is coprime to p.
  $ a^(p-1) eq.triple 1 mod p $


== Primitive root modulo
  / $a|b$: a divides b=> b/a = 0

  / Primitive root modulo n: $g$ is called primitive root modulo $p$ if every $a$ coprime number to $n$ is congruent to a power of $g$ modulo $p$.
  $ forall a in Z: gcd(a,p)=1, exists n: g^n=a arrow.long g "is primitive root modulo" $

  N is not required to be prime.
  G is a _primitive root modulo_ $n$ if and only if $g$ is a generator of the multiplicative group of integers modulo n.


== P Group

=== How to check that group is cyclic
=== Theorem to check generator in p group
  \ 
  $alpha in Z_(p)^(*)$ is a generator of $ Z_(p)^(*)$ if and only if $ alpha^((p-1)/q)not eq.triple 1 mod p $
  
  For all primes $q$ such that $q|(p-1)$

  \
  *Task*

    Task find the all generators of $Z_(11)^(*)$

    Let's begin with $p-1 = 10$, 10 = 2*5.

    Generator check condition for each divider of $(p - 1)$: 
    - $alpha^(5)not eq.triple 1 mod 11$
    - $alpha^(2)not eq.triple 1 mod 11$

    Solution is to check each element in group to match conditions.

=== How to count generators in group
/ Theorem: let p be prime, that $ Z_(p)^(*)$ contains exactly $phi(p-1)$ generators.

=== How to find generator

===  Discrete logarithm in p
  \
  If $Beta in Z_(p)^(*)$, then $Beta = g^x$ for some unique $0<= x <=p-2$. 
  X is called the discrete logarithm of $Beta$ to base $g$.

  $ g^x = Beta arrow.double  log_g Beta = x "in" Z$

  / Problem: find the integer x, such that $ log_g Beta "in" Z_(p)^* $

  The naive approach is exhaustive search: compute $g^x, g^2x, ...$ until B is obtained.
   
== N Group
=== Theorem to check generator in n group

  For $n>=1$, we consider $Z_(n)^*$ 

  $ Z_(n)^* = {k in {1, ..., n} "/" gcd(k,n)=1} $

  $ "len"(Z_(n)^*) = phi(n) $

  $Z_(n)^*$ *is cyclic*(has at least one generator) when:
  1. n=2 or 4
  2. $n= p^x, x in {1,2...}$
  3. $n= 2 p^x, x in {1,2...}$

  / Theorem to check generator: 
  Assume $Z_(n)^*$ is cyclic. $alpha in Z_(n)^*$ is a generator if and only if $ alpha ^(phi(n)/p) not eq.triple 1 mod n $
  For each prime p divisor of $phi(n)$ 

  \


=== How to count generators in group
/ Theorem: if $ Z_(n)^(*)$ is cyclic, then it has $pi(pi(n))$ generators.


=== Discrete logarithm in n
  \
  If $Beta in Z_(n)^(*)$, then $Beta = g^x$ for some unique $0<= x <=p-2$. 
  X is called the discrete logarithm of $Beta$ to base $g$.

  _Example_
  \
  Find $log_13 47$ in $Z_(50)^*$
  1. Check $Z_(50)^*$ is cyclic(e.g. has generators)
  2. Check g 13 is generator.(requires find $phi(n)$)
  3. Start to calculate elements.(exhaustive search)

= Algorithms for computing discrete algorithms

1. Brute force
2. Shank's baby-step giant-step method

== Some equations

  $ 3 Beta mod 13 = 1 arrow.double 3B eq.triple 1 mod 13  $

  $ (a b) mod m = [(a mod m)(b mod m)]mod m $
 
== Core Equation for DH
  $ y_1=(a^(x_1) mod p)^(x_2) mod p = a^(x_1 x_2)  mod p $

  Some modular arithmetics to proof:
  $ y_1=(a^(x_1) mod p)^(x_2) mod p = ((a mod p)^(x_1) mod p)^(x_2) mod p = $

  Reduce extra modulo due to modulo properties:
  $ (a mod p)^(x_1) mod p = a^(x_1) mod p $ 

  Continue:

  $ = (a mod p)^(x_1 x_2)  mod p $
  $ = a^(x_1 x_2) mod p $


// TODO:
// 1. Why Z_p is cyclic
// 2. Why generator formulas are like thats
// 3. How to find generator, not count, not detect

== Factorization problem

For example RSA relies on difficulty of factoring the product of two large prime numbers.

But for this need to determine or find large prime number.

Determine if number is prime:
1. Simple methods(advanced brute force, without 2,3 and maybe some memoization,etc)
2. Probabilistic tests(all primes + some non primes - never FN, but sometimes FP)
