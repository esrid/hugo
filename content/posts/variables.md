+++
title = 'Go variables & constant'
date = 2024-07-01T08:58:17-04:00
draft = false
description = "Dans le langage Go, il existe plusieurs méthodes pour déclarer des variables, chacune ayant ses spécificités. Dans cet article, nous allons explorer ces différentes méthodes et déterminer dans quels contextes l'une peut être préférable à une autre."
type = "post"
tags = ["Go", "Programming", "Tutorial"]
showTags = true
readTime = true
categories = ["Programming", "Go"]
toc = true
series = ["Go Basics"]
+++

Dans le langage Go, il existe plusieurs méthodes pour déclarer des variables, chacune ayant ses spécificités. Dans cet article, nous allons explorer ces différentes méthodes et déterminer dans quels contextes l'une peut être préférable à une autre.

## Var

```go
var dix int = 10
```
Le mot clé var suivit de l'identifiant `dix`, le type et l'assignement.

```go
var dix = 10
```
Vous pouvez omettre le type si le type par défaut correspond à celui que vous souhaitez. Dans cet exemple, le type par défaut est int.

```go
var dix,onze int 
```
Vous pouvez aussi declarer plusieurs variable avec un type en commun  

```go
var dix, phrase = 10, "une chaine de caractères"
```
ou sans préciser le type, le type sera inféré automatiquement.

```go
var (
	dix int
	mot string = "une chaîne de caractères"
	pi         = 3.14
   )
```
Vous pouvez déclarer plusieurs variables en entourant les identifiant entre parenthèses. (recommander quand vous voulez déclarer plus de 2 variables)

## :=
`:=` est la version courte de Var, mais avec quelques limitations.
Le type est directement inféré donc pas de typing.


```go
 dix, phrase := 10, "une chaine de caractères"
```
Comme le mot-clé `var`, vous pouvez faire une multiple déclaration de variables.

```go
func main() {
	dix := 10 // Valide ✅
}
dix := 10 // Non valide ❌
```
Cette méthode de déclaration ne fonctionne pas en dehors d'une fonction.

## Const

`const` vous permet de créer des variables immuables contrairement a var. c'est-à-dire des variables qui ne peuvent être modifié 

```go
const (
	dix int    = 11
	mot string = "une chaîne de caractères"
	pi         = 3.14
)

const x, y int = 1, 2
```
`const` est un peu limité, il peut seulement contenir des : 
- nombre littéraux `const x = 1010_1001`
- true or false `const b = true`
- chaîne de caractères `const s = "chaine de caractères"`
- rune `const r = 'A'`
- built-in fonctions complex, real, imag, len, and cap

## Conclusion : 
Lors de l'initialisation d'une variable à sa valeur zéro, exemple : `var i int`.
Cela indique clairement que la valeur zéro est voulue.

Vous devez rarement déclarer des variables en dehors des fonctions,Lorsqu'une variable se trouve en dehors d'une fonction, il peut être difficile de suivre les modifications apportées à la variable, ce qui empêche de comprendre comment les données circulent dans le programme

À l'intérieur d'une fonction l'utilisation de `:=` est plus communément utiliser.

Utilisez `const` pour des variables immuables.

