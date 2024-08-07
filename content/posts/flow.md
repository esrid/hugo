+++
title = 'Flow'
date = 2024-08-04T14:57:51-04:00
draft = true
type = "post"
tags = ["Go", "Programming", "Tutorial"]
showTags = true
readTime = true
categories = ["Programming", "Go"]
toc = true
series = ["Go Basics"]
+++
## if
Si vous avez dèjà fait un peu de progammation, vous connaissez le mot clé if. 
Mais si jamais go et votre 1er language et cet article est l'un des 1er que vous lisez, le mot clé `if` qui signifie `si` et son but est de controller le flow de votre logic. 
```go
func main() {
	if 18 > 16 {
		println("18 est plus grand que 16")
		return
	}
}
```
Dans cet exemple ci-dessus le mot, je dis si 18 est plus grand que 16 affiche que 18 est plus grand que 16 puis j'arrete le program.
## else 
Le mot clé `else` sert a dire si ce n'est pas le cas.
```go
func main() {
	if 18 > 16 {
		println("18 est plus grand que 16")
		return
	}else {
		println("16 est plus que que 18")
	}
}
```
## else if 
les deux mot clé dans cet ordre signie `sinon si `
```go
func main() {
  chat := "chat"
  if chat == "chien"{
    println("chien")
  }else if chat == "chat"{
    println("chat x)")
  }else {
    println("autre animal domestique")
  }
}
```
Cela vous permet d'essayée plusieurs spécifique condition contrairment à else qui agis comme un : "peu importe ce que c'est d'autre fait ça"
##  variable avec if 
```go
package main

import "math/rand"

func main() {
	if n := rand.Intn(8); n == 0 {
		println("ok", n)
	} else {
		println("pas ok", n)
	}
}
```
If vous permet de déclarer des variables qui sont uniquement utilisable dans ce même block
## swicth
