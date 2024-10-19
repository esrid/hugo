+++
title = 'Pointer'
date = 2024-09-18T00:12:35-04:00
draft = true
type = "post"
tags = ["Go", "Programming", "Tutorial"]
showTags = true
readTime = true
categories = ["Programming", "Go"]
toc = true
series = ["Go Basics"]
+++
# POINTER 
Un pointeur est simplement une variable qui contient l'emplacement dans la mémoire où une valeur est stockée.
stockée.
Le & est l'opérateur d'adresse. Il précède un type de valeur et renvoie l'adresse de l'emplacement mémoire où la valeur est stockée
```go
hello := "hello"
pointerToHello := &hello
```
Le * est l'opérateur d'indirection. Il précède une variable de type pointeur et renvoie la valeur pointée.
renvoie la valeur pointée. C'est ce qu'on appelle le déréférencement.
```go
hello := "hello"
pointerToHello := &hello
fmt.Println(pointerToHello) // montre l'addresse 
fmt.Println(*pointerToHello) // montre la value
```
La fonction intégrée new crée une variable pointeur. Elle renvoie un pointeur vers une instance
```go
var x = new(int)
fmt.Println(x == nil) // prints false
fmt.Println(*x)
```
La new fonction est rarement utilisée. Pour les structures, utilisez un & devant un littéral de structure pour créer une instance de pointeur
Vous ne pouvez pas utiliser un & devant un littéral primitif (nombres,
booléens et chaînes de caractères) 

```go
x := &Foo{}
var y string
z := &y
```
