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

Знання.
/ Інтуїтивне означення: інформація, на основі якої можна отримувати нову інформацію.

/ Знання інтелектуальних систем: Трійка $<F, R, P>$, де F - факти, які зберігаються в системі,
R - правила, за допомогою яких можна виводити нові факти
P - процедури, визначають яким чином слід застосовувати правила

База знань - усі знання в системі.

Проблеми:
1. Виключення в правилах.(пінгвін птах -> пінгвін літає)
2. Недостатня кількість знань(система має запитувати додаткову інформацію)
3. Недостовірні знання
4. Невизначеність(probability decision)
5. Проблема зміни ситуацій(хижак вимре, від того, що вимерла жертва -> жертва експоненційно розмножується)

Відомі моделі:
1. Cyc - проект по створенню онтологічної бази знань. Large reasoning engine by Douglas Lenat.
Де люди просто руками вносять правила.

Extensional and intensional representation.
Extensional - all possible facts stored
Intensional - not all facts stored, but there are rules, that may reason other facts.

Засоби структуризації:
1. Узагальнення(студент -> людина яка вчиться)
2. Агрегація(студент -> людина і учень)

Властивості та класифікація знань:

Моделі:
1. Логічна
2. Продукційна
3. Фрейм модель
4. Семантична мережа

== Семантична мережа
Семантична мережа - композиція концептуальних графів, поєднаних за деякими правилами.

Мережа, в якій поєднані звʼязки різних типів. Можна уявляти у вигляді графа.
Склад конструкції: $<I, C, G>$
I - множина інформаційних одиниць
C - типи звʼязків між цими одиницями
Г - відображення, що задає звʼязки з цими одиницями

П-сутність - з реального світу
М-сутність - представлення в базі знань


Типи мереж:
1. Класифікуючі - задають відношення ієрархії між І.О.(інформаційними одиницями).
2. Функціональні - мають функціональні відношення, які дозволяють обчислювати одні І.О. в інші
3. Сценарії - використовують відношення "дія", "причина-наслідок", "засіб дії"

Базові поняття мережі:

Концептуальний граф:
Box - concept node
Oval - relation node

/ концептуальний обʼєкт: графічно - концептуальний граф. До опису входить:
1. клас П-сутності
2. властивості сутності
3. прототип або приклад


Проблема Наслідування - механізм не завжди гарантує правильність висновку.(пінгвін має крила, бо птах. або ластівка Дана вивчається натуралістами)

== Фрейм
Визначення - структура даних призначена для опису типових ситуацій чи типових понять.
Визначення Мінського - мінімальний опис деякої сутності, такий, що подальше скорочення цього опису приводить до втрати цієї сутності

4. 	Охарактеризуйте поняття фрейму.
5. 	Опишіть роль фреймів у розумінні.
6. 	Охарактеризуйте поняття мережі подібностей за Уїнстоном, намалюйте довільний конкретний приклад.
7. 	У чому полягає зв’язок між семантичними мережами і фреймами?


== Bayes rule

$ P(A|B) = (P(B|A) * P(A)) / P(B) $

P(A|B) - probability of A given B is true

== Bayes for nlp

$ c = "argmax" P(d|c)P(c) $

d = set of features $x_1, x_2,...$

$ c = "argmax"П_(x in X) P (x|c) P(c)$

1. Maximum likelihood estimates

Class probability:
$ P(c_j) = "doccount"(C=c_j)/ N_"doc" $

Word to class probability:
$ P(w|c)=("count"(w,c)+1)/("count"(c) + |V|) $

count of needed words + 1/ count of total words + volume of words total

2. 


Choosing a class:

#figure(
  image("./img/ChooseClass.jpg", width: 70%),
)