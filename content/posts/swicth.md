+++
title = 'Swicth'
date = 2024-08-08T17:36:03-04:00
draft = false
type = "post"
tags = ["Go", "Programming", "Tutorial"]
showTags = true
readTime = true
categories = ["Programming", "Go"]
toc = true
series = ["Go Basics"]
+++


## Swicth
Le fonctionnement de switch est similaire à celui des autres langages de programmation : il compare différentes conditions et s'arrête dès qu'une condition est remplie.

syntax : 
```go
x := "chat"
switch x {
        case "renard":
        fmt.Println("super")
        case "monster":
        fmt.Println("pas super")
        default :
        fmt.Println(x)
    }
```
Vous avez la possibilité de comparer une valeur avec plusieurs cas.
```go
func main() {
    jour := "Mercredi"

    switch jour {
    case "Lundi", "Mardi", "Mercredi":
        fmt.Println("C'est la première moitié de la semaine.")
    case "Jeudi", "Vendredi":
        fmt.Println("Le week-end approche.")
    case "Samedi", "Dimanche":
        fmt.Println("C'est le week-end !")
    default:
        fmt.Println("Ce n'est pas un jour valide.")
    }
}
```
Comme avec if déclarer une variable utilisable dans ce bloc.
```go
func main() {
    valeur := 15

    switch resultat := valeur * 2; {
    case resultat < 10:
        fmt.Println("Le résultat est inférieur à 10.")
    case resultat >= 10 && resultat < 30:
        fmt.Printf("Le résultat est %d, ce qui est entre 10 et 30.\n", resultat)
    case resultat >= 30:
        fmt.Printf("Le résultat est %d, ce qui est supérieur ou égal à 30.\n", resultat)
    default:
        fmt.Println("Cas inattendu.")
    }
}

```
