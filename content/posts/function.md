+++
title = 'Function'
date = 2024-08-12T19:41:17-04:00
draft = true
type = "post"
tags = ["Go", "Programming", "Tutorial"]
showTags = true
readTime = true
categories = ["Programming", "Go"]
toc = true
series = ["Go Basics"]
+++
## FONCTION
Déclaration d'une fonction : utilisez le mot-clé func suivi du nom de la fonction. Si votre fonction prend des paramètres, vous devez en spécifier le type, et si elle retourne une valeur, vous devez également indiquer le type de retour.
```go
func add(a int, b int) int {
    return a + b
}

func main() {
    result := add(3, 5)
    fmt.Println("The sum is:", result)
}

```
