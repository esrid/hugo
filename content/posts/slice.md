+++
title = 'Tranche & tableau quel sont les différences ? '
date = 2024-07-04T15:45:51-04:00
draft = false
description = "Quels sont les differences entre une tranche (slice) et un tableau (array) en go ?"
type = "post"
tags = ["Go", "Programming", "Tutorial"]
showTags = true
readTime = true
categories = ["Programming", "Go"]
toc = true
series = ["Go Basics"]
+++

## Array

Pour déclarer un tableau, vous devez spécifié la taille et le type : 
```go
var a [5]int
```
Cette déclaration initialise les valeurs du tableau au type par défaut, qui est 0 pour les entiers (int). Cela signifie que dans cet exemple, la variable `a` contient un tableau avec 5 chiffres 0, `[0,0,0,0,0]`.

```go
var a = [...]int{1,2,3,4,5}
```
Vous pouvez également utiliser cette notation, qui déterminera automatiquement le nombre d’éléments dans le tableau.

```go
var multi [5][3]int
```
Vous pouvez déclarer des tableaux à dimensions multiples.

## Sparse
Si vous avez un tableau où la plupart des éléments sont assignés à leur valeur zéro par défaut, vous pouvez spécifier quels index contiennent une valeur différente. Cela fonctionne aussi avec les tranches.

```go
var a = [12]int{5:8,9:5}
```
## Slice
Les tranches ressemblent à des tableaux, mais une des différences est que vous n’avez pas à définir la taille lors de la déclaration.

```go
var a [5]int
// c'est un tableaux
var a []int
// c'est une tranche
```
Comme avec les tableaux, vous pouvez déclarer des tableaux à dimensions multiples.

```go
var multi [5][3]int
```
Contrairement aux tableaux, déclarer une tranche sans lui donner de valeur la rendra nil.

```go
var a []int
println(a == nil)
// true
```
D’ailleurs, vous ne pouvez pas directement comparer une tranche à une autre tranche, mais seulement à nil.
Si vous souhaitez les comparer, c'est possible avec la fonction `DeepEqual` du module `reflect`.    
```go
	a := []int{1}
	b := []int{1}
	fmt.Println(reflect.DeepEqual(a, b))
```
### Append 

Le mot clé `append` sert à ajouter un élément dans une tranche. Comme argument, il prend une tranche de n'importe quel type et une ou plusieurs valeurs. Chaque fois que vous passez un paramètre à une fonction, Go fait une copie de la valeur qui est passée. Le passage d'une tranche à la fonction append transmet en fait une copie de la tranche à la fonction. La fonction ajoute les valeurs à la copie de la tranche, et renvoie la copie. c'est pourquoi réassigné est nécessaire.

```go
var a []int
x := []int{0,191,919,88}
a = append(a, 5)
a = append(a, 6,7,8)
// merge deux tableaux
a = append(a,x...)
```
### Capacité
la capacité d'un tranche et le nombre d'element qu'un tranche peut recevoir avant que le runtime double sa capacité.
```go
x := []int{0,191,919,88}
fmt.Println(len(x), cap(x))
```
Dans cet exemple, la taille du tableau est 4 et sa capacité est aussi de 4 par défaut, mais si nous ajoutons une valeur, la capacité sera doublée.

```go
x := []int{0,191,919,88}
fmt.Println(len(x), cap(x))
x = append(x, 11)
fmt.Println(len(x), cap(x))
```
### Make
```go
a := make([]int, 6)
```
Utilser le mot clé `make` pour crée une tranche avec une taille par défault et une capacité par default

```go
a := make([]int, 6, 19)
```
make(type, taille, capacité)

### Créer un tranche a partir d'un tranche
Vous pouvez crée un tranche a partir d'un tranche 
```go
x := []int{10,191,919,88}
y := x[1:]
```
Mais attention, ce n'est pas une copie, en fait, ils partagent la même mémoire.
```go
x := []int{10,191,919,88}
y := x[1:]
y[1] = 26
```
Ce changement sera répercuté aussi bien sur la variable y que x, ceci fonctionne aussi bien avec un tranche qu'un tableau.

### copy 
Pour avoir une copie complètement indépendants d'une tranche ou d'un tableau, le mot clé `copy` est ce qu'il vous faut.
```go
x := []int{10,191,919,88}
y := make([]int, 4)
copy(y,x)
```
Le premier paramètre est le destinataire, le second est la source. La capacité n'est pas importante, mais la taille compte.

```go
x := []int{10,191,919,88}
y := make([]int, 2)
copy(y,x)
```
Dans cet exemple, copy copiera seulement les deux premiers éléments de x dans y.

## Conclusion 

Quelle déclaration choisir ?

La priorité est de limiter le nombre de fois que la tranche va grossir. S'il est possible que la tranche n'ait pas besoin de croître du tout (parce que votre fonction ne renvoie rien), utilisez une déclaration var sans valeur assignée pour créer une tranche nil. Si vous avez des valeurs par défaut, utilisez une tranche littérale. Si vous connaissez la taille et la capacité, utilisez make.
