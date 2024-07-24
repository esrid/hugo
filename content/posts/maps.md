+++
title = 'Maps'
date = 2024-07-18T15:45:51-04:00
draft = false
description = "Comment utiliser les maps en go ?"
type = "post"
tags = ["Go", "Programming", "Tutorial"]
showTags = true
readTime = true
categories = ["Programming", "Go"]
toc = true
series = ["Go Basics"]
+++

### Maps 
Pour déclarer une map vous devez utiliser le mot-clé `map` suivi du type de clé et du type de value. 

```go
var x map[string]int
```
Dans cet exemple ci-dessus, nous avons créé une nil map qui attend un string en clé et un int en value.
```go
func main() {
	var x map[string]int
	fmt.Printf("%T, %v", x,x ) 
}

```
Lorsque vous déclarez une nil map en go vous pouvez uniquement la lire , ce qui vous renvoie toujours la valeur zéro du type de valeur de la carte, mais vous ne pouvez rien ajouter, car la map n'est pas initialiser. 

```go
func main() {
	var x map[string]int
	fmt.Printf("x: %v\n", len(x))
	x["chat"] = 2 // invalide ❌
	fmt.Printf("x: %v\n", x)
	//ce code ne fonctioneras pas ❌
}
```
Dans l'exemple ci dessous la variable x est initialisée

```go
func main() {
	var x = map[string]int{} // x a été initialiser
	fmt.Printf("x: %v\n", len(x)) 
	x["chat"] = 2 // valide ✅
	fmt.Printf("x: %v\n", x)
	//ce code est valide ✅
}
```
### Make avec map

Si vous connaisez déjà la taille de la maps vous pouuvez utiliser le mot clé `make`, ela map sera directment initialisée avec une taille et une capacité. pour en savoir plus sur `make` consulter ce post ( [make](/posts/slice/#make) )

```go
func main() {
	var x = make(map[string]int, 1) 
	x["chat"] = 2 
}
```
### Quelque similarité avec Slice 
Point commun : 
- Les Maps s'agrandissent automatiquement au fur et à mesure que vous y ajoutez des paires clé/valeur.
- Si vous connaissez le nombre de paires clé/valeur que vous prévoyez d'insérer dans une map, vous pouvez utiliser make pour créer une map d'une taille initiale spécifique,
- vous pouvez utiliser make pour créer une map avec une taille initiale spécifique.
- Le passage d'une map à la fonction len vous indique le nombre de paires clé/valeur
dans une map.
- La valeur zéro d'une map est nil.
- Les Maps ne sont pas comparables. Vous pouvez vérifier si elles sont égales à nil, mais vous ne pouvez pas vérifier si deux maps ont des valeurs identiques.
- vous ne pouvez pas vérifier si deux maps ont des clés et des valeurs identiques en utilisant ==
ou diffèrent en utilisant !=.

### Acceder a la value d'une clé
```go
func main() {
 x := map[string]int{"chat" : 1}
 println(x["chat"])
}
```

### Insérer une nouvelle clé/value
```go
func main() {
 x := map[string]int{}
 x["chat"] = 2
}
```
### Ok Idiom
si vous voulez verifier qu'une clé existe vous pouvez utiliser le ok idiom,
Dans l'exemple ci-dessous : le v represente la value et le ok est boolean si la variable existes ok deviens true sinon false
```go
func main() {
 x := map[string]int{}
 x["chat"] = 2
 v, ok := x["chat"] 
 println(v, ok)
}
```
### supprimer une clé 
Le mot-clé `delete` permet de supprimer une clé de la map en 1er argument elle prends la maps et en second argument la clé a supprimer 
```go
func main() {
 // initialise
 x := map[string]int{}
 // insert
 x["chat"] = 2
 // vef-rifier si une clé exist
 v, ok := x["chat"] 
 println(v, ok )
// supprimer une clé
 delete(x, "chat")
 println( x["chat"])
```
### Créer un set avec une maps
Un set est une structure de donnée dans lequel il n'y a pas de répétion, comme go n'a pas de built-in `set` vous pouvez émuler un set comme ceci
```go
x := map[int]bool{}
y :=[...]int{1,1,2,2,3,3,5,6,7,8,8,9}
for _, v := range y {
        x[v] = true
}
println(x)
```
La variable x aura q'un exemplaire de chaque element de la variable x.
