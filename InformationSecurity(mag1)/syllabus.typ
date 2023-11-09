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

// public vs asymmetric

= Symmetric crypto
/ Symmetric crypto: relies on the fact, that two sides know the same key, which they obtain by another secure channel. The same key used for encryption and decryption. 

Issues:
1. Key distribution problem
2. Number of keys on N users is n(n-1)/2
3. No protection against cheating between Alice and Bob


= Asymmetric encryption.
Public key cryptography - is a method of encrypting, which allows individuals to securely communicate without sharing the secret key.


Trapdoor one way function - function that can be easily computed in one way, and hard in the inverse without special(secret) information(trapdoor).
In this case forward -> encrypt. Inverse -> decrypt.

Private key - is key to decrypt message.
Public key - is key to encrypt message.
Or vice versa

For signatures private key enables to sign and public to verify this sign validity.

Arithmetic functions, that considered as a one way(not proofed)
1. Multiplication and factorization
2. x^y mod n = z => find y
3. x^y mod n = z => find x
4. x^2 mod n. n - not prime, jacobi(z/n)=1 => find x
5. g^ab mod p => find a

Also use cases:
1. Key exchange
2. Identification
3. Signature to cannot deny having sent/received a message

= Prime numbers
/ Prime number: an integer P which has exactly two positive divisors(1 and P).

= Alternative problems

= Group

Group is set of elements, that are related to each other according to certain well-defined rules.

$Z_p^*$ - is a group with nonzero integers between 1 and p-1 modulo some prime nomber p.
== Axioms
Operation for example is a multiplication.

1. Closure - $a,b in G, a*b=c -> c in G$
2. Associativity $a*(b*c)=(a*b)*c$
3. Identity existence $a*1=a$
4. Inverse existence $a*b=1$

A group is commutative if $a*b = b*a$
A group is cyclic if there is generator g $ g in G, forall x in G, exists n: g^n=x $

\
\
\
== How to choose or check that $a$ is a generator?
/ Theorem: $g in Z^(*)_p$ is a generator of $Z^(*)_p$ if and only if $g^((p-1)/q) not eq.triple 1 mod p $ for all primes $q$ such that $q|(p-1)$

= Diffie-hellman
Protocol:
#figure(
  image("./img/DH.png", width: 70%),
)

= Rsa
Target: send information from A to B securely.

Public key can encrypt.

Private key can decrypt.

1. Choose 2 prime int p,q
2. Count n = p*q
3. Count Euler function $phi(n)=(p-1)(q-1)$
4. Count $e: gcd(e, phi(n))=1$
5. Solve $e x= 1 mod phi(n)$, find x = d

X represents a number in $Z_n^(*)$, binary value of X must be less than n.(also n and x are coprime, can just check that $not x|e$)

Public key is a pair of (n, e)

Y is a ciphertext. $Y=(x^e mod n)$

Decryption: the private key is $d$. $ y^d mod n = x $

$x^(e d) mod n = x mod n$

$ x^(e d - 1) mod n = 1 mod n $
It's possible only when:

$ e d - 1 = k phi(n)$

$ x^(k phi(n) + 1) mod n = 1 x mod n $

First, choose p and q, count n=p*q, count $phi(n) = (p-1)(q-1)$

Second, choose $e$ coprime to $phi(n)$. Solve equation $e d = 1 mod phi(n)$. 
$ d = e^(-1) mod phi(n)$

x, y, n, d are large numbers(1024 bits or more)

e, d

== Rsa also vulnerable to man in the middle attack
Example: https://security.stackexchange.com/questions/189468/why-can-a-man-in-the-middle-attack-not-happen-with-rsa

= Elgamal encryption scheme
1. Perform DH
2. Use Key as a mask for the message modulo p

== Elgamal protocol
1. Set up phase

Done only once.
Bob makes $p, alpha, b$. Compute  $B =alpha^b mod p$
$ k_"pub" = (p, alpha, B) $

Bob publishes keys.

2. The encryption phase

Executed every time

Alice chooses ephemeral key $K_E == alpha^a mod p$
Alice computes the "shared key"(private) $K = B^a mod p$
$ y - x K mod p $

And sends $K_E$ and y to Bob

3. The decryption phase

Executed every time

$ X = Y K^(-1) mod p $

$ m = K^(-1) Y mod p $

$ K = g^(a b) mod p $

By fundamental theorem of arithmetic
$ gcd(g, p) = 1 -> gcd(g^n, p) = 1 $


by fermat's little theorem:
$ K_E^(p - 1) eq.triple 1 mod p $

$ (g^b)^(p - 1) eq.triple 1 mod p $
$ (g^b)^(p - 1) eq.triple 1 mod p $

$ (g^b)^(p - 1) eq.triple 1 mod p $
$ (g^b)^(p - 1) eq.triple (g^b)^(p - 1 - a) g^(a b) eq.triple 1 mod p $

$ g^(a b) = K $
$ K_(-1) = (g^b)^(p - 1 - a) mod p$

= Coin flipping

Using DLP
Given two generators a, b

Bob chooses one(f.e. a) and pow it to x, sending $a^x$ to alice.
Alice chooses a or b and sends it to bob.
Bob sends x to Alice. 
Alice checks $a^x$.

Note: there are possibility that $a^x==b^x$


= Lampart algorithm