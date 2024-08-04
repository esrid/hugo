+++
title = 'Shadow'
description = "Blocks et shadow variable"
date = 2024-08-04T14:45:20-04:00
draft = false
type = "post"
tags = ["Go", "Programming", "Tutorial"]
showTags = true
readTime = true
categories = ["Programming", "Go"]
toc = true
series = ["Go Basics"]
+++
## Block
Comme vous le savez vous pouvez déclarer des variables de part et d'autre.

Chaque endroit ou vous pouvez déclarer une variable s'appelle en fait un block. 

Les `type`, `const`, `var` et `function` qui sont déclarer en dehors d'une fonction sont dans le package block.

```go
// un fichier go
package main 

const StatusNotFound = 404 // ici le status not found ce trouve dans le block du `package` donc accesible dans n'importe quel autre fichier avec appertant au package main
```
Toutes les variables qui sont dans une fonction y compris les paramètres d'une fonction sont dans le block de cette fonction, c'est-à-dire que le champ d'application (portée) n'est utilisable qu'entre les deux accolades de la fonction.

```go
package main

const StatusNotFound = 404

func main() {
	println(z)
}

func Additon(x, y int64) int64 {
	var z int64 = 3
	return x + y*z
}
```
Dans cet exemple ci-dessus, j'essaie d'afficher la variable z qui est déclarée dans la fonction Addition, z sera `undefined` dans la fonction, car z n'existe pas dans la portée de la fonction main.

### Shadow variable
```go
package main

func main() {
	a := 9
	if a > 5 {
		println("YOUPI !", a)
		a := 0
		println("A : ", a)
	}
	println(a)
}
```
Ce code affichera : "YOUPI ! 9, "A: 0" et pour finir 9. La variable n'a pas été ré assigné car même si elle a le même nom que la variable qui est dans la portée de la fonction main, la variable `a` entre les accolades du if n'existeque dans ce block. 

*Une shadow variable est une variable qui porte le même nom q'une variable dans le block contenant.*

```go
package main

import "fmt"
func main() {
    fmt := "format"
    fmt.Println("fmt")
}
```
Vous aurez une erreur disant que le type string n'a aucune méthode ou aucun champ Println.
Dans cet exemple, j'ai `shadow` le package fmt qui est dans block du fichier. 
```go
package main

import "fmt"
func main() {
    fmt.Println(false)
    false := "format"
    fmt.Println(false)
}
```
Même situation mais avec le mot clé `false` prédéfinie qui se trouve dans le `univers block`.




