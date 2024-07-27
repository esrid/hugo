+++
title = 'Struct'
date = 2024-07-25T15:45:51-04:00
draft = true
description = "Comment utiliser les structs en go"
type = "post"
tags = ["Go", "Programming", "Tutorial"]
showTags = true
readTime = true
categories = ["Programming", "Go"]
toc = true
series = ["Go Basics"]
+++
## Struct 
Pour déclarer un `struct` vous devez utiliser le mot-clé `type` suivi d'un `identifiant` ensuite du mot-clé `struct`  et pour finir ouvrer et fermé les accolades `{}` .
```go
type Nature struct {}
```
Dans cet exemple ci-dessus, nous avons déclaré un struct Nature sans aucun champ. 

Pour déclarer un champ vous devez spécifier le nom du champ et son type.
```go
type Nature struct {
    animalsEspece []string 
    plantsEspece []string
}
```
Un struct est un type à part entière donc vous pouvez déclarer des variables de ce type.
```go
type Nature struct {
    animalsEspece []string 
    plantsEspece []string
}
var n Nature
// ou
x := Nature{}
```
Les champs sont initialisés avec la valeur 0 selon leur type, sauf si lors de la déclaration, vous initialisé les champs.
```go
var n = Nature{
        animalsEspece: []string{"Felis catus", "Canis lupus"},
        plantsEspece:  []string{"Helianthus annuus", "Aloe vera"},
    }
    x := Nature{
        []string{"Felis catus", "Canis lupus"},
        []string{"Helianthus annuus", "Aloe vera"},
    }
```

**Note** : Faites attention à la seconde variable x dans l'exemple ci-dessus. Elle est un peu différente de n car nous ne précisons pas à quels champs les valeurs sont assignées. Dans cette méthode, les valeurs sont assignées dans l'ordre dans lequel les champs sont déclarés dans la structure Nature.
En precisant donc les champs vous pouvez l'utiliser dans l'ordre que vous souhaité et même en omettre certain. c'est soi l'un ou l'autre vous ne pouvez utiliser les deux. 

### Acceder a un champ
```go
var n Nature
n.animalsEspece = []string{ []string{"Felis catus", "Canis lupus"}
fmt.Println(n.animalsEspece)

```
utiliser un `.` pour acceder à un champ. pour le lire ou lui réasigner une nouvelle value:
### Anonymous struct
```go
personne := struct{
        nom string
    }{
        nom : "Julius novachrono"
    }
```

### Comparer les structs
Contrairement à un slice ou a une map les (instances de) struct sont comparables, mais uniquement si la struct ne contient pas des types incomparables comme les slices ou les maps.

```go
type Humain struct {
	nom string
}

func main() {
	 p := Humain{}
	 po := Humain{}
	 fmt.Println( p == po)
	// true 
}
```

```go
type Humain struct {
	nom string
	diplome []string
}

func main() {
	 p := Humain{}
	 po := Humain{}
	 fmt.Println( p == po)
	// ne fonctionneras pas
}
```
Si vous comparez une instance de struct a une autre instance de struct, ça ne fonctionneras pas non plus, car ce sont deux types différents type.

```go
type Humain struct {
	nom string
}
type NonHumain struct {
    nom string 
}

func main() {
	 p := Humain{}
	 po := NonHumain{}
	 fmt.Println( p == po)
	// ne fonctionneras pas
}
```

Mais vous pouvez les convertir :
```go
x := Humain(po)
```
La conversion ne fonctionnera pas si les champs sont différents, avec un champ en plus ou en moins ou si l'ordre diffèrent.

