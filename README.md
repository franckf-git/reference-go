- [intro go](http://www.golang-book.com/books/intro)
- [documentation](https://golang.org/doc/)
- [go tour](https://tour.golang.org/welcome/1)
- [packages](https://pkg.go.dev/)
- [awesome-go](https://github.com/avelino/awesome-go)

- [best-practices](https://golang.org/doc/effective_go)
- [gitlab guidelines](https://docs.gitlab.com/ee/development/go_guide/)
- [code-review](https://github.com/golang/go/wiki/CodeReviewComments)
- [effective-go](https://golang.org/doc/effective_go)

## tests

- go-lint : il ne semble pas y avoir d'option intégrée, tester avec le binaire

- tester la modification d'un slice/maps, avec et sans pointer dans une fonction pour confirmer l'immutabilité

  > pour clarifier "If you pass a map to a function that changes the contents of the map, the changes will be visible in the caller."

- go generate / go fix : pour generer du code compatible ?

- auto-build gitlab

- go test : comment chercher/lancer les tests dans les sous-dossiers `./...` ?

- refacto in test
```
    checkSums := func(t testing.TB, got, want []int) {
        t.Helper()
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got %v want %v", got, want)
        }
    }
```

- https://github.com/stretchr/testify

- argument d'une func par defaut (si possible) ?

# blog archives

https://blog.golang.org/concurrency-timeouts  
https://blog.golang.org/declaration-syntax  
https://blog.golang.org/defer-panic-and-recover  
https://blog.golang.org/docker  
https://blog.golang.org/error-handling-and-go  
https://blog.golang.org/generate  
https://blog.golang.org/gif-decoder  
https://blog.golang.org/go1.13-errors  
https://blog.golang.org/godoc  
https://blog.golang.org/h2push  
https://blog.golang.org/organizing-go-code  
https://blog.golang.org/context-and-structs

## concurrency

https://www.youtube.com/c/vilito-exquisitus/playlists

https://www.golangprograms.com/go-language/concurrency.html

https://blog.john-pfeiffer.com/golang-concurrency-goroutines-and-channels/

https://betterprogramming.pub/common-go-pitfalls-a92197cd96d2

## project structure

https://peter.bourgon.org/go-best-practices-2016/#repository-structure

https://medium.com/@benbjohnson/standard-package-layout-7cdbc8391fc1

https://rakyll.org/style-packages/

https://www.youtube.com/watch?v=oL6JBUk6tj0
