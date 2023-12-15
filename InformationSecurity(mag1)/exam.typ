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

= 2

$ (2+n)^(-1) mod n^2; n=17 $

$ 19^(-1) mod 289 $

Need to calculate *modular multiplicative inverse*

$ x eq.triple 19^(-1) mod 289 $

$ 19x - 289y = 1 $

Using extended Euclidean algorithm.
Let's set:
$x_0 = 1, x_1=0$

#table(
  columns: (auto, auto, auto),
  [$q = [(r_(i-1) - r_(i-1))/r_i]$ ], 
    [$r_i = r_(i-2) - q r_(i -1 )$], 
      [$t_(i+1) = t_(i -1) - q_i t_i$],
  [], [19], [0],
  [], [289], [1],
  [15], [4], [-15],
  [4], [3], [61],
  [1], [1], [*-76*],
  [3], [0], [289],
)

Solution is $-76; 19*(-76)=1-289*5$.


= 1 

$ Y^2 = X^3 + 2X + 4 mod 11; P:(10,10); "find 3P" $