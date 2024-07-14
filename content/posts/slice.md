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
a = append(a, 5)
// true
```
le mot clé apprend sert a ajouter un un element dans un tableaux
