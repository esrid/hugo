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

## Sparse


