+++
title = 'if'
date = 2024-08-04T14:57:51-04:00
draft = false
type = "post"
tags = ["Go", "Programming", "Tutorial"]
showTags = true
readTime = true
categories = ["Programming", "Go"]
toc = true
series = ["Go Basics"]
+++
## if
Si vous avez déjà fait un peu de programmation, vous connaissez le mot-clé if. 
Mais si jamais go et votre 1er langage et cet article est l'un des 1er que vous lisez, le mot-clé `if` qui signifie `si` et son but est de contrôler le flux de votre logique. 
```go
func main() {
	if 18 > 16 {
		println("18 est plus grand que 16")
		return
	}
}
```
Dans cet exemple ci-dessus le mot, je dis si 18 est plus grand que 16 affiche que 18 est plus grand que 16 puis j'arrête le programme.
## else 
Le mot-clé `else` sert à dire si ce `sinon` comme dans une phrase, "si ce n'est pas ça sinon fait ça".
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
les deux mot-clé dans cet ordre signie `sinon si `
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
Cela vous permet d'essayer plusieurs spécifique condition contrairement à else qui agis comme un : "peu importe ce que c'est d'autre fait ça."
## variable avec if 
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
If vous permet de déclarer des variables qui sont uniquement utilisables dans ce même block
