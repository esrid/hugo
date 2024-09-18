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
# FONCTION
Déclaration d'une fonction : utilisez le mot-clé `func` suivi du nom de la fonction. Si votre fonction prend des paramètres, vous devez en spécifier le type dans le cas ou elle retourne quelque chose, le type doit être spécifié. 
```go
func add(a int, b int) int {
    return a + b
}

// or a,b int 

func main() {
    result := add(3, 5)
    fmt.Println("The sum is:", result)
}
```
### Paramètres optionnels

Go ne vous autorise pas a avoir des parametres optionels mais vous pouvez en simulez avec des struct
```go
type Options struct {
    options1 bool
    options2 bool
    options3 bool
}
func Optional( opts Options){

}
Optional(Options{true, true})
```
### variadique parametres

Go prend également en charge ce qu'on appelle des paramètres variadiques, c'est-à-dire un nombre variable d'arguments passés à une fonction.
```go
func VariadicAdd(a int, b ...int) int {
    x := 0 
    for _,v := range b {
        x = x + v
    }
    return x + a
}

VariadicAdd(1, 545, 44, 54)
```
*Remarque* : Le paramètre variadique doit toujours être specifier en dernier dans la liste des paramètres

### Multiple valeur de retour
Vous avez la possibilité d'avoir plusieurs valeur de retour.

```go
// code du module path de la bibliotheque go 
func Split(path string) (dir, file string) {
	i := bytealg.LastIndexByteString(path, '/')
	return path[:i+1], path[i+1:]
}
```
comme vous pouvez le voir dans l'exemple vous avez aussi la possibilité de nommé la valeur de retour. 

### Fonction anonymes
```go
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		func(j int) {
			fmt.Println("affiche", j, "depuis l'interieur")
		}(i)
	}
}
```
Exactement comme une fonction normal mais sans la nommer

### Closure 

Les fonctions déclarées à l'intérieur de fonctions sont spéciales, ce sont des fermetures (closures). Il s'agit d'un
mot informatique qui signifie que les fonctions déclarées à l'intérieur de fonctions 
peuvent accéder aux variables déclarées dans la fonction extérieure et les modifier.

exemple d'usage : 
```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 4}
	sort.Slice(a, func(i, j int) bool {
		return a[i] > a[j]
	})
	fmt.Println(a)
}

```
### Defer

Le mot-clé `defer` permet d'exécuter une fonction lorsque la fonction principale se termine ou que son exécution est interrompue.

```go
func OpenFileWithPath(path string) ([]byte, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return nil, err
	}
  content, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return content, nil
}

```
Quand il y a plusieurs defer, il s'execute in LIFO (Last in first out)
defer peut être utiliser avec des function anonyme

### Utilisation des valeurs de retour nommées avec `defer`

Il est possible qu'une fonction différée (`defer`) examine ou modifie les valeurs de retour de la fonction qui l'entoure, notamment en utilisant des **valeurs de retour nommées**. Cela permet de gérer les erreurs de manière plus efficace, notamment dans les transactions de base de données.

#### Exemple avec une transaction de base de données :
```go
// source learn go in an idiomatic way 
func DoSomeInserts(ctx context.Context, db *sql.DB, value1, value2 string) (err error) {
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    defer func() {
        if err == nil {
            err = tx.Commit()
        }
        if err != nil {
            tx.Rollback()
        }
    }()
    _, err = tx.ExecContext(ctx, "INSERT INTO FOO (val) VALUES ($1)", value1)
    if err != nil {
        return err
    }
    return nil
}
```
### Modèle commun pour la gestion des ressources avec des closures

En Go, une fonction qui alloue une ressource peut également retourner une **closure** pour nettoyer cette ressource après utilisation. Un exemple simple est l'ouverture d'un fichier et la gestion de sa fermeture.

#### Exemple de fonction de gestion de fichier :
```go
func getFile(name string) (*os.File, func(), error) {
    file, err := os.Open(name)
    if err != nil {
        return nil, nil, err
    }
    return file, func() {
        file.Close()
    }, err
}
f, closer, err := getFile(os.Args[1])
if err != nil {
    log.Fatal(err)
}
defer closer()
```
### Go : Passage par valeur et comportements des types

Go est un langage **call by value** (passage par valeur), ce qui signifie que lorsqu'une variable est passée à une fonction, Go crée une copie de la valeur de la variable.

#### Exemple avec des types de base et structs :
```go
type person struct {
    age  int
    name string
}

func modifyFails(i int, s string, p person) {
    i = i * 2
    s = "Goodbye"
    p.name = "Bob"
}

func main() {
    p := person{}
    i := 2
    s := "Hello"
    modifyFails(i, s, p)
    fmt.Println(i, s, p)  // Résultat : 2, "Hello", {0 ""}
}
```
Dans cet exemple, les modifications des paramètres au sein de la fonction ne persistent pas, car Go fait une copie des valeurs des variables. Cela inclut également les structs. 

### Comportement des maps et slices :

Les maps et slices se comportent différemment car ils sont implémentés avec des pointeurs. 

```go
func modMap(m map[int]string) {
    m[2] = "hello"
    m[3] = "goodbye"
    delete(m, 1)
}

func modSlice(s []int) {
    for k, v := range s {
        s[k] = v * 2
    }
    s = append(s, 10)  // Ne modifie pas la longueur du slice d'origine
}

func main() {
    m := map[int]string{1: "first", 2: "second"}
    modMap(m)
    fmt.Println(m)  // Résultat : map[2:hello 3:goodbye]

    s := []int{1, 2, 3}
    modSlice(s)
    fmt.Println(s)  // Résultat : [2 4 6]
}
```
Les modifications sur les maps sont directement appliquées, mais pour les slices, bien que l'on puisse modifier leurs éléments, l'ajout d'éléments via append ne modifie pas le slice original.
