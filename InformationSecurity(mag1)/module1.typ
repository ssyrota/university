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

= Module 1
7 варіант

== 1
1. Схема RSA. Повідомлення: m =  17; параметри: p = 3 , q = 7, e = 17.
Знайти d, зашифрувати m (тобто знайти c), розшифрувати c.

$
cases(
  p = 3 , 
  q = 7,
  e = 17,
  m=17,
)
$

$ n = 21 $
$ mod phi(n) = 12 $
$ 17 d = 1 mod 12 $

Need to calculate *modular multiplicative inverse*
$ d = 17^(-1) mod 12 $

Using extended Euclidean algorithm.
Let's set:
$x_0 = 1, x_1=0$

While a>1:
$ q = 1; (a, n) = (12, 5); (x_0, x_1) = (1, -1) $
$ q = 2; (a, n) = (5, 2); (x_0, x_1) = (-1, 3) $
$ q = 2; (a, n) = (2, 1); (x_0, x_1) = (3, 5) $
$ q = 2; (a, n) = (1, 1); (x_0, x_1) = (5, ...) $
$ a=1, "stop" $
$ x_0=5 -> x = 5 $
$ "Let's verify: " 17*5 mod 12 = 85 mod 12 = 1 $

$ d = 5 $

Public key = (n, e) = (21, 17)

The encrypted message is $y = 17^17 mod 21 $

The decrypted message is $ m = 17^(17*d) mod 21 $
$ m = 17^(17*5) mod 21 $
$ "as we know" (17*5) = k 12 + 1 = 84 + 1 = 85 $
$ m = 17^(7*phi(21)) 17 mod 21 = 17 mod 21 $

== 2
2. Розв'язати порівняння за модулем:
Need to calculate *modular multiplicative inverse*
$ 610x eq.triple 1 mod 987 $
$ x eq.triple 610^(-1) mod 987 $

$ 610x−987y=1 $

Using extended Euclidean algorithm.
Let's set:
$x_0 = 1, x_1=0$

#table(
  columns: (auto, auto, auto),
  [$q = [(r_(i-1) - r_(i-1))/r_i]$ ], 
    [$r_i = r_(i-2) - q r_(i -1 )$], 
      [$t_(i+1) = t_(i -1) - q_i t_i$],
  [0], [610], [0],
  [1], [987], [1],
  [1], [ 377 ], [-1],
  [1], [ 233 ], [ 2 ],
  [1], [ 144 ], [ -3 ],
  [1], [ 89 ], [ 5 ],
  [1], [ 55 ], [ -8 ],
  [1], [ 34 ], [ 13 ],
  [1], [ 21 ], [ -21 ],
  [1], [ 13 ], [ 34 ],
  [1], [ 8 ], [ -55 ],
  [1], [ 5 ], [ 89 ],
  [1], [ 3 ], [ -144 ],
  [1], [ 1 ], [ 233 ],
  [1], [ 0 ], [ -377 ],
)
$x = -377 mod 987 = 610$

== 3
$ 15^3^1000 mod 17 = $
$ 3^1000 mod 16 = 16^x + y $
$ 15^(16x + y) mod 17 = 15^y * 15^(16x) mod 17 $

За малою теоремою Ферма $ 15^(16x) mod 17 = (15^(16))^x mod 17 $

$ (15^(p-1))^x mod p = 1^x mod p arrow.long (15^(16))^x mod 17 = 1 mod 17 $ 

Залишилось знайти y.

$ 15^y mod 17 = 15^(3^1000) mod 17 $

16 is not prime. 3^1000 is not prime, but 3 and 16 is coprime -> we can apply Euler's theorem.
$ phi(n) = 8 $
$ 3^1000 mod 16 = (3^(8))^125 mod 16 $
By Euler's theorem:
$ a^(phi(n)) eq.triple 1 mod n $
$ 3^(8) eq.triple 1 mod 16 $
$ (3^(8))^125 eq.triple 1 mod 16 $
$ 3^1000 mod 16 = 1 mod 16 $
$ 3^1000 = 16^z + 1 $

Finally:
$ 15^(16z) * 15 mod 17 $
$ 15^(16z) mod 17 = 1 mod 17 $
$ 1*1*15 mod 17 = 15 mod 17 $