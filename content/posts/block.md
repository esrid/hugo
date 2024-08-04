
+++
title = "Blocks et shadow variabblek"
date = 2024-08-05T15:45:51-04:00
draft = true
description = "Blocks et shadow variabble"
type = "post"
tags = ["Go", "Programming", "Tutorial"]
showTags = true
readTime = true
categories = ["Programming", "Go"]
toc = true
series = ["Go Basics"]
+++

## Block
Comme vous le savez vous pouez déclarer des variables de part e d'autres.

Chaque endroit ou vous pouvez déclarer une variable s'appel enfaite un block. 

Les `types`, `const`, `var` et `function` qui sont déclarer en dehors d'une fonction sont dans le package block

```go
// un fichier go
package main 

const StatusNotFound = 404 // ici le status ce trouve dans le 404
```
Toute les variables qui sont dans une fonction y compris les paramettres d'une fonction sans le le block de cette fonction 

```go
package main 

const StatusNotFound = 404 // ici le status ce trouve dans le 404
func main (){

    }

func Additon(x, y int64) int64{
    z := 3
    return x + y * z
    }
```
Dan cet exemple si 
