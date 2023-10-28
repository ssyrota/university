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

$n=3*7=21$

$phi(n)=(p-1)(q-1)=12=N$

$d e = 1 mod N$

$3d = 1 mod 12 arrow.long d = 37$

$C=E(m)=17^17 mod 21 = 5$


== 2
2. Розв'язати порівняння за модулем:
$ 610x eq.triple 1 mod 987 $



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