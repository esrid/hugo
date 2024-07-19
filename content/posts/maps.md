+++
title = 'Maps'
date = 2024-07-20T15:45:51-04:00
draft = true
description = "Maps"
type = "post"
tags = ["Go", "Programming", "Tutorial"]
showTags = true
readTime = true
categories = ["Programming", "Go"]
toc = true
series = ["Go Basics"]
+++

#Maps 
```go
var x map[string]int
```
Pour déclarer une map vous devez utilisé le mot clé `map` suivi du type de clé et du type de value. 
Dans cet exemple, nous avons créer une nil map qui attend un string en clé et un int en value.
C'est bizare dans cet situation vous pouvez uniquement lire la variable x ce qui vous renvoie toujours la valeur zéro du type de valeur de la carte mais vous ne pouvez rien ajouter car la map n'est pas initialiser. 
```go
package main

import "fmt"

func main() {
	var x map[string]int
	fmt.Printf("x: %v\n", len(x))
	x["chat"] = 2 // invalide
	fmt.Printf("x: %v\n", x)
	//ce code ne fonctioneras pas
}
```


```go
package main

import "fmt"

func main() {
	var x = map[string]int{} // x a été initialiser
	fmt.Printf("x: %v\n", len(x))
	x["chat"] = 2 // valide
	fmt.Printf("x: %v\n", x)
	//ce code est valide
}
```
