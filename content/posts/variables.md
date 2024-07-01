+++
title = 'Go variables & constant'
date = 2024-07-01T08:58:17-04:00
draft = true
+++

Dans le langage Go, il existe plusieurs méthodes pour déclarer des variables, chacune ayant ses spécificités. Dans cet article, nous allons explorer ces différentes méthodes et déterminer dans quels contextes l'une peut être préférable à une autre.

## Var

```go
var dix int = 10
```
Le mot clé var suivit de l'identifiant `dix`, le type et l'assignement

```go
var dix = 10
```
Vous pouvez omettre le type si le type par défaut correspond à celui que vous souhaitez. Dans cet exemple, le type par défaut est int.

```go
var dix,onze int 
```
Vous pouvez aussi declarer plusieurs variable avec un type en commun  

```go
var dix, phrase = 10, "une chaine de charactere"
```
ou sans préciser le type, le type sera inféré automatiquement.

```go
var (
	dix int
	mot string = "une chaîne de character"
	pi         = 3.14
   )
```
Vous pouvez declarez des plusieurs variables en entourant les identifiant entre parenthese, (recommander quand vous voulez déclarez plus de 2 variables)

## :=
`:=` est la version courte de var mais avec quelques limitations.
le type est directement inférer donc pas de typing.

vous pouvez comme faire une multi declaration de variable

```go
 dix, phrase := 10, "une chaine de charactere"
```
mais il n'est pas possible d'utiliser cet méthodes en dehors d'une function.
```go
func main() {
	dix := 10
}
```
## Const
`const` vous permez de créer des variables immutables contrairement a var. c'est à dire des variables qui ne peuvent être modifié 
```go
const (
	dix int    = 11
	mot string = "une chaîne de character"
	pi         = 3.14
)

const x, y int = 1, 2

```
`const` est un peu limité il peut seulement containir des : 
- nombre literaux `const x = 1010_1001`
- true or false `const b = true`
- chaine de charactere `const s = "chaine de charactere"`
- rune `const r = 'A'`
- built-in functions complex, real, imag, len, and cap

### Conclusion : 
Lors de l'initialisation d'une variable à sa valeur zéro, utilisez `var i int`.
Cela indique clairement que la valeur zéro est voulue.

Vous devez rarement déclarer des variables en dehors des fonctions,Lorsqu'une variable se trouve en dehors d'une fonction, il peut être difficile de suivre les modifications apportées à la variable, ce qui empêche de comprendre comment les données circulent dans le programme

A l'intérieur d'une function l'utilisation de `:=` est plus communement utiliser.

Utilisez `const` pour des variables immutables.
