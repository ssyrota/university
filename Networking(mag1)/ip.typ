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

= Introduction

Network purpose: give ability to communicate, publicly or privately

/ IP: 32 bit(4 octets) address, used for internet. It cannout uniquely identify device.

== Classful and classless
To controll usage of IP's was invented Classful addresation.(deprecated)

Has limit, for all subnets use one mask.

#image("./img/classful.png")

Class A:
#table(
  columns: (auto, auto, auto, auto),
  inset: 10pt,
  align: horizon,
 [network], [host],[host],[host]
)

To distinct network and host used *mask*
Class A mask: 255.0.0.0

Class B:
#table(
  columns: (auto, auto, auto, auto),
  inset: 10pt,
  align: horizon,
 [network], [network],[host],[host]
)

Class B mask: 255.255.0.0

Class C:
#table(
  columns: (auto, auto, auto, auto),
  inset: 10pt,
  align: horizon,
 [network], [network],[network],[host]
)

Class C mask: 255.255.255.0
  

/ Subnet: is a logical(possibly recursive) subdivision of an IP network.
Usecases: efficient allocation in large organizations. May be used to create tree structure.

All hosts in the same subnet have the same network.

Task: 172.16.0.0 - make 120 subnets with 180 hosts, write mask.


== Public, private addresses and CIDR
To determine local network vs global network used separation.

Reservations list:

Local network:
1. 10.0.0.0 — 10.255.255.255 with mask 255.0.0.0 (or just 10/8).
2. 172.16.0.0 — 172.31.255.255 with mask 255.240.0.0 (or just 172.16/12).
3. 192.168.0.0 — 192.168.255.255 (or just 192.168/16).

Other:
1. 0.0.0.0/8 - self identify (when using DHCP)
2. 127.0.0.0/8 - loopback
3. 224.0.0.0/4 - multicast
4. 169.254.0.0/16 — link-local address

/ Multicast: send one IP datagram to many recievers, actually sent one, but routers and switches duplicate.


// private vs public subnet in aws
// classless inter-domain routing vs classful, what's purpose of divide by classless

// how device determines that one is in the local network/subnet
// switches

192.168.1.0/24
1. Подсеть на 120 адресов.
2. Подсеть на 12 адресов.
3. Подсеть на 5 адресов.

Найдите и запишите в каждой подсети ее 
- адреса
- широковещательный адрес
- пул разрешенных к выдаче адресов
- маску

Let's start from one small, advice says, that should start with large, but no provided description, so :D

1. 

Size is 8

Address - 192.168.1.0

Multicast - 192.168.1.7

Pull is - 192.168.1.1 - 192.168.1.6

Mask - 255.255.255.(256-8) = 255.255.255.248


2. 

Size is 16

Address - 192.168.1.8

Multicast - 192.168.1.23

Pull is - 192.168.1.9 - 192.168.1.22

Mask - 255.255.255.(256-16) = 255.255.255.240

3. 

Size is 128

Address - 192.168.1.24

Multicast - 192.168.1.151

Pull is - 192.168.1.25 - 192.168.1.150

Mask - 255.255.255.(256-128) = 255.255.255.128