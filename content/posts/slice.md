+++
title = 'Slice & array quel sont les différences ? '
date = 2024-07-04T15:45:51-04:00
draft = true
description = "Quels sont les differences entre une slice et un array en go ?"
type = "post"
tags = ["Go", "Programming", "Tutorial"]
showTags = true
readTime = true
categories = ["Programming", "Go"]
toc = true
series = ["Go Basics"]
+++

# Array
Pour déclarer un array, vous devez spécifié la taille et le type : 
```go
var a [5]int
```
comme nous l'avons vu dans l'article précédent, cette déclaration initialise la varlaur du type par default a sont 0, ce qui signifie dans cet exemple la variable `a` contient un array avec 5 `0` 

```go
var a = [...]int{1,2,3,4,5}
```
Vous pouvez également utiliser cette notation, qui déterminera automatiquement le nombre d'éléments dans le tableau.

```go
var multi [5][3]int
```
vous pouvez declarer des array avec de multiple dimension

## Sparse
Si vous avez un un aray ou la plus part des elements sont assigné a leur 0 value. 
vous pouvez spécifié quel index contient une value differente. (cela fonctionne aussi avec les slices)
```go
var a = [12]int{5:8,9:5}
```
## Slice

Les slices ressemble a des taableaux une des difference de trouve dans la declaration vous n'avez pas a definir la taille

```go
var a [5]int
//this is a array 
var a []int
// this is a slice 
```
comme les aray vous pouvez faire des slices multi dimensionel

```go
var multi [5][3]int
```
contrairement aux array delcarer un slice sans lui donner une valeur sera nil

```go
var a []int
println(a == nil)
// true
```

d'ailleur vous ne pouvez pas comparer un slice a un autre slice mais seulement a nil 


```go
var a []int
x := []int{0,191,919,88}
a = append(a, 5)
a = append(a, 6,7,8)
// merge deux tableaux
a = append(a,x...)
```
le mot clé `apprend` sert a ajouter un un element dans un tableaux, comme argument il prend un slice dans n'importe quel type et une ou plusieur valeur
pourquoi reasigné a ? 
Chaque fois que vous passez un paramètre à une fonction, Go fait une copie de la valeur qui est passée. Le passage d'une slice à la fonction append transmet en fait une copie de la slice à la fonction.
de la slice à la fonction. La fonction ajoute les valeurs à la copie de la slice,
et renvoie la copie. Vous assignez ensuite la slice renvoyée à la variable de la fonction appelante.

### capacité
la capacité d'un slice et le nombre d'element qu'un slice peut recevoir avant que le runtime double sa capacité 
```go
x := []int{0,191,919,88}
```
dans cette exemple la taille du tableaux et 4 et ça capacité est aussi de 4 par default mais si  nous ajoutons une valeur la capacité sera doublé
```go
x := []int{0,191,919,88}
fmt.Println(len(x), cap(x))
x = append(x, 11)
fmt.Println(len(x), cap(x))
```

```go
a := make([]int, 6)
```
utilser le mot clé `make` pour crée une slice avec une taille par défault et une capacité par default

```go
a := make([]int, 6, 19)
```
capacité de 19 et taille de 6
### Créer un slice a partir d'un slice
vous pouvez crée un slice a partir d'un slice 
```go
x := []int{10,191,919,88}
y := x[1:]
```
mais attention ce n'est pas une copie, en fait il partage la même mémoire
```go
x := []int{10,191,919,88}
y := x[1:]
y[1] = 26
```
ce changement sera répercuté aussi bien sur la variable y que x

## Transformer un array en slice

## Conclusion 

Quel declaration choisir ? 

La priorité est est de limité le nombre de fois que vas grossir le slice. 
S'il est possible que la slice n'ait pas besoin de croître du tout (parce que votre fonction ne renvoie rien), utilisez une déclaration var.
utiliser une déclaration var sans valeur assignée pour créer une slice nil.
aucune valeur assignée pour créer une slice nulle

si vous avez des values par défault utiliser un slice litéral 
si vous savez quel taille et quel capcité utilisé make
