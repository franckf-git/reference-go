- [intro go](http://www.golang-book.com/books/intro)
- [documentation](https://golang.org/doc/)
- [go tour](https://tour.golang.org/welcome/1)
- [packages](https://pkg.go.dev/)
- [awesome-go](https://github.com/avelino/awesome-go)

- [best-practices](https://golang.org/doc/effective_go)
- [gitlab guidelines](https://docs.gitlab.com/ee/development/go_guide/)
- [code-review](https://github.com/golang/go/wiki/CodeReviewComments)

- [gitlab-ci](https://blog.boatswain.io/post/build-go-project-with-gitlab-ci/)
- [case studies & use cases](https://go.dev/solutions#case-studies)

- [oscon](https://spf13.com/presentation/building-an-awesome-cli-app-in-go-oscon/)
- [gophercises](https://gophercises.com/)
- [learn-with-tests](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/install-go#go-environment)
- [learn-to-code](https://github.com/ashleymcnamara/learn_to_code#golang)
- [gowebexamples](https://gowebexamples.com/)

## tests

- paquets : comment mettre plusieurs packets dans un même dossier (ex. utils)

- go-lint : il ne semble pas y avoir d'option intégrée, tester avec le binaire

- tester la modification d'un slice/maps, avec et sans pointer dans une fonction pour confirmer l'immutabilité
> pour clarifier "If you pass a map to a function that changes the contents of the map, the changes will be visible in the caller."

- go generate / go fix : pour generer du code compatible ?

- auto-build gitlab

- godocs
```
godoc golang-book/chapter11/math Average
godoc -http=":6060"
```

- go test : comment chercher/lancer les tests dans les sous-dossiers

- rangement lien symbolic de dossier dans GOPATH

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
