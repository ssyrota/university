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


= Euler theorem
a, n are coprime

$ alpha^(phi(n)) eq.triple 1 mod n $

x coprime with n, a coprime wuth n -> x a coprime with n (there no gcd>1)


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


== Modular multiplicative inverse
/ Algorithm: Modular Inverse using Extended Euclidean Algorithm

*check why it works*

If a is coprime to n

Problem: $a x eq.triple 1 mod n$

$ a x+m y="gcd"(a,m)$

$ x_0 = 1, x_1=0, y_0=0, y_1=1 $

While a > 1:
$ q = [a/n] $
$ (a, n) = (n, a mod n) $
$ (x_0, x_1) = (x_1, x_0-q x_1) $
$ (y_0, y_1) = (y_1, y_0-q y_1) "(may be omitted when find modular multiplicative inverse)" $

Next: if $x_0 < 0$, then $x_0 = x_0 + n$


== Bezout equality

$ exists m, n in Z: m d + n d = d $
$ (m+n) = 1 $
$ exists gcd(a, b) = d arrow.long exists x,y: m d = a x; n d = b y; $

*Why x and y do exist?*
$ a x + b y = "gcd"(a, b) $
$ exists k,l: a = k d, b = l d $

Let's divide equation by d:
$ k d x + l d y = d arrow.long k x + l y = 1 $
we know k and l 
$ k = a/d; l = b/d; $
$ k x + l y = 1 arrow.long y = (1 - k x) / l $
$ l !=0; k,x in Z: k x in Z arrow.long y "can be solved" $

In the end we have endless count of solutions with formula:
$ y (b/d) = (1 - (a/d) x); a x + b y = gcd(a, b) $

If one pair of (x,y) was found:
$ (x - k (b/d), y + k (a/d)) $

For case of gcd = 1:
$ a x + b y = 1 $

== Proof of gcd equality
$ S not emptyset -> exists min(n) in S $

x, y - are Bezout's coefficients

$ "Having" a, b "with " gcd(a,b)=d " and " a x + b y = d$
$ "Prove that x, y exists and " a x + b y = d , " d is min positive integer of this combination " $

/ Proof: Suppose that we have set $S$ with smallest element $d$.
1. Prove that $d$ is a divisor of a,b and 
2. for any common divisor c $c<=d$

1. Let's divide a on d: $a=d q + r, 0<=r<d$
$ r = a - d q $
$ r = a - (a x + b y)q $
$ r = a(1 -x) - b (y q) $
Thus: $ r = a n + b m, "where: " n = 1 - x, m = -(y q) $
This implies that $r in S, S:{ a x + b y = d; exists x,y in Z }$

Now we have $ 0<=r<d; r in S; d "is min in "S $ Contradiction.
Min element d from $S$ is divisor of $a,b$ (b by analog proof).

2. For any common divisor c $c<=d$
Let $c$ be divisor of $a,b$ -> $a x + b y = d$
a = c k
b = c l

$ c x k + c y l = d $
$ c(x K + y l) = d; (x K + y l) >= 1 -> d>=c -> d "is" "the greatest divisor" $


= Find gcd
GCD - greatest common divisor

== Euclidean algorithm

Based on the principle $ gcd(a, b) = gcd(a-b, b), "if a > b" $
Example: 
$ gcd(112,256) = gcd(112, 144) = gcd(32, 112) = gcd(32, 80) $
$ = gcd(48, 32) = gcd(16, 32) = gcd(16, 16) = 16 $

/ Proof: let's assume: 
$ exists m: a = d*m $
$ exists n: b = d*n $

$ a-b = d(m-n) arrow.long a-b eq.triple d $

A more efficient way is to use modulo operation for bigger element by smaller.
/ Proof: let's assume: 
$ exists m: a = d*m $
$ exists n: b = d*n $

$ gcd(a, b) = gcd (b k + a mod b, b ) $

As we know $gcd(a, b) = gcd(a-b,b)$, applying recursively we obtain equation:

$ gcd(a, b) = gcd(a-(k b), b) = gcd (a mod b, b) $

$ = gcd(a mod b, b) = gcd(a, b) $

But that's not enough.
Need to proof:
1. $r_(n-1)$ is a common divisor of a,b
2. $r_(n-1)$ is a gcd

Proof:
1. Is proved above. $r_(n-1)=c; c<=gcd -> r_(n-1)<=gcd$

2. Suppose we have common divisor c, which divides a, b; a -b = $k c -l c$ -> c divides a-b, divides each $r_n$.
Thus $r_(n) = (k-l)c$. Therefore $forall c, c|a; c|b: c|r_n$(c<=r_n) -> $r_n$ is $gcd$

== Extended euclidean algorithm

a is coprime to b

$ a x + b y = gcd(a, b) $

Why a and b should have gcd = 1;
Because if not: a cannot have inverse.

Suppose that gcd is not 1:
$ a b eq.triple 1 mod m $
$ a b - 1 eq.triple 0 mod m $
$ a = g k; m = g l; $
$ (g k) b  eq.triple 1 mod (g l) $
$ k b eq.triple 1/g mod l $
$ 1/g not in Z *. =>  not exists b $


= Other
=== Division theorem
For every natural number m and positive natural number n, there exists a unique pair of integers q and r such that $q >= 0, 0 <= r < n$, and $m = q · n + r$

== Fundamental arithmetics theorem

$ N = p_1^(e_1) p_2^(e_2) ... p_n^(e_n) $
