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


There are plethora of algorithms is vulnerable to man in the middle attack.

= Message Authentication Codes
Encryption != integrity.

/ Definition: short piece of information used for authentication and integrity checking a message.

/ Non-repudiation: is a security assurance that prevents signers from denying their actions. 
MACs are not resistant to non-repudiation, because minimum two sides know secret key - so anyone can make message. 
In contrast with digital signatures, where asymmetric keys is used.

= HMAC

/ Def: MAC with used hashing. resists length extension attacks, add more security by using random oracle.

= Lamport auth

A generates "root" and $(1...n) in N, n: H(root)^n$
A passes to B H(root)^n. The next time A passes to B $H(root)^(n-1)$
So either A found a collision, or A knows hash key.

= Chinese remainder theorem
Given $k$ coprime numbers. And system:
$ x eq.triple a_1 mod n_1 $
$ x eq.triple a_2 mod n_2 $
$ ... $
$ x eq.triple a_k mod n_k $

Then exists $x eq.triple n_1*n_2*...*n_k$

= Hidden subgroup problem

