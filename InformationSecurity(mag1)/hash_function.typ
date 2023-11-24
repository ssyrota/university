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

Main target: infeasible to find collision.

*The key $s$ generally not keep secret*, nevertheless, $H_s$ must be resistant to collisions.


= Formal definition

A hash function $H_s$ with fixed-length output $l(n)$ is a pair of probabilistic polynomial time algorithms ($"Gen",H$) satisfying the following:

Gen is a probabilistic algorithm that takes as input a security parameter $1^n$(1111...$1_n$) and outputs a key $s$.


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

= Based on block ciphers
- MD4
- MD5
- SHA-1
- SHA-2

== Block cipher
/ Block cipher: function that operates on fixed bits length input. Input n, key k. Output n size message.

Standard block cipher: AES, DES.
// TODO: how it works https://www.youtube.com/playlist?list=PL1xkDS1G9As4Yz_te20j1A9evIjt5Z06e.

=== AES
Symmetric cipher.
Key size 128/192/256. N = 128

Rounds will depend of the key size.
DES has fixed 16 rounds.

Each round is dirived into layers
First round turns into two sub keys and 4 layers.
Rest of the rounds, one key per time, 3 layers.

3 types of layers.
1. Key addition layer.(XOR with key)
2. Byte substitution layer(S-box): perform substitution using "lookup tables". Provide confusion
3. Diffusion layer:
- ShiftRows: permutes the data on the byte level
- MixColumn: another matrix permutation

// Why use hmac, bcrypt and just sha256

=== XOR
Usage justification: it's better randomizes encryption, since it output is 0/1 50%

== Blockchain
Sequential growing data structure, intended to provide complete data integrity.

Mining problem = for H function and fixed k, find x that H(x) starts with k nulls.

== Use cases

- Hash table(often used non-cryptographic hash functions) and indexing
- Fingerprinting and verifying the integrity of data
- Identifier


= Merkle–Damgård construction
domain extension method

/ Def: a method of building collision-resistant cryptographic hash functions from collision-resistant one-way compression functions.(uses AES(state, message))

If compression F is resistant to collisions -> construction is resistant too.

Sequential compression of blocks(like blockchain) if end is not full length, add padding. 

It's possible to process as a tree - therefore scales infinitely (called merkle tree). 

== Other ciphers
/ Stream cipher: encrypts data bit by bit. Useful for real time data processing.

= Sha2. SHA256
Output is 256 bit value.

1. Split message for 512 blocks. If last is not 512, use padding.

2. To provide random and non zero starting point algorithm has 8 initial hash values - $i in {2,3,5,7,11,13,17,19}: {sqrt(i) mod 1}$. 8 because each value should consistently influence the output.

3. Each of 512-length blocks processed in a loop.


= Sha1
Output 160 bit
Based on MD2, MD4, MD5, but uses larger output.

= MD2
Inputs are 128 bit.
