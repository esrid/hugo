+++
title = 'For'
date = 2024-08-08T17:37:30-04:00
draft = false
type = "post"
tags = ["Go", "Programming", "Tutorial"]
showTags = true
readTime = true
categories = ["Programming", "Go"]
toc = true
series = ["Go Basics"]
+++
# For
`for` est l'unique mot-clé servant à faire des boucles dans le language go mais il être utilisée de 4 façons. 

## 1) for 
`for` (seul) sert à crée une boucle infini
```go
for {
        fmt.Println("Salut")
    }
```
vous pouvez arrêter la boucle infini en utilisant le mot clé `break`
```go
for {
        fmt.Println("Salut")
        break
    }
```
## 2) Le classique
Comme en javascript cet boucle tourneras jusqu'à ce que la value de `i` sois 9.
la condition dans la deuxieme partie de la boucle doit être vrai. 
```go
for i := 0 ; i < 10; i++ {
    println(i)
}
```
## 3) for conditon
Une boucle qui continue de tourner tant que la condition est vrai. pas besoin de déclaration
```go
for i < 10{
    println(i)
}
```
## 4) for range 
Celui là est un peu spécial car il permet de faire des iterations sur des iterables comme des chaines de caracteres, des tranches, des tablaux ou des maps par examples.  
```go
	x := "apprendre go c'est génial"
	for i, v := range x {
		fmt.Printf("%d | %v\n", i, v)
	}
```
dans cet exemple nous itérons dans sur la variable x, i representant l'index et r la valeur dans ce cas une répresentations unicode du caractere, en go cela s'appel des runes. 
```go
	identite := map[string]string{
		"Nom":       "Curie",
		"Prenom":    "Marie",
		"DateNaissance": "7 novembre 1867",
		"LieuNaissance": "Varsovie, Pologne",
		"Nationalité":   "Française",
		"Profession":    "Physicienne et chimiste",
	}
	for _, v := range x {
		fmt.Printf("que les valeurs: %v\n", v)
	}
```
vous pouvez ignorer un des deux avec un petit _.

## continue
```go
func main() {
	grid := [][]string{
		{"apple", "banana", "cherry"},
		{"date", "elderberry", "fig"},
		{"grape", "honeydew", "kiwi"},
	}
	for _, row := range grid {
		for _, fruit := range row {
			// Si le fruit actuel contient la lettre 'e', on ignore le reste de cette itération
			if strings.Contains(fruit, "e") {
				continue
			}
			fmt.Println(fruit)
		}
	}
}
```
Le mot-clé continue sert à ignorer le reste du corps de la boucle en cours et à passer directement à l'itération suivante. Cependant, vous pouvez contrôler quelle boucle doit passer à l'itération suivante en nommant vos boucles, comme illustré ci-dessous :
```go
func main() {
	grid := [][]string{
		{"apple", "banana", "cherry"},
		{"date", "elderberry", "fig"},
		{"grape", "honeydew", "kiwi"},
	}
outer:
	for _, row := range grid {
		for _, fruit := range row {
			if fruit == "date" {
				// Passe à l'itération suivante de la boucle extérieure
				continue outer
			}
			fmt.Println(fruit)
		}
	}
}
```

## Infos 
En Go, l'ordre des clés et des valeurs dans une map varie à chaque exécution pour des raisons de sécurité. Avant, l'ordre d'itération des clés était souvent le même, ce qui causait des erreurs imprévisibles et rendait possible une attaque Hash DoS. Pour éviter cela, Go utilise maintenant un algorithme de hachage avec un nombre aléatoire et fait varier l'ordre d'itération à chaque boucle.
