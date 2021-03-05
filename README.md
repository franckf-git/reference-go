- [documentation](https://golang.org/doc/)
- [best-practices](https://golang.org/doc/effective_go)
- [go tour](https://tour.golang.org/welcome/1)
- [reddit](https://www.reddit.com/r/golang/)
- [fedora](https://fedoramagazine.org/getting-started-with-go-on-fedora/)
- [packages](https://pkg.go.dev/)
- [gitlab-ci](https://blog.boatswain.io/post/build-go-project-with-gitlab-ci/)
- [awesome-go](https://github.com/avelino/awesome-go)
- [packages](https://pkg.go.dev/)
- [case studies & use cases](https://go.dev/solutions#case-studies)
- [gitlab guidelines](https://docs.gitlab.com/ee/development/go_guide/)

## tests

- dans `/reference-go/hello/hello.go` tester de remplacer import directement avec `../greetings` pour ne pas avoir a utiliser `remplace` dans `go.mod`

- go-lint : il me semble y avoir d'option intégrée, tester avec le binaire

- tester la modification d'un slice/maps, avec et sans pointer dans une fonction pour confirmer l'immutabilité
> pour clarifier "If you pass a map to a function that changes the contents of the map, the changes will be visible in the caller."

- go generate / go fix : pour generer du code compatible ?

- auto-build gitlab

# blog archives

https://golang.org/doc/faq
https://golang.org/doc/modules/managing-dependencies
https://blog.golang.org/concurrency-timeouts
https://blog.golang.org/declaration-syntax
https://blog.golang.org/defer-panic-and-recover
https://blog.golang.org/docker
https://blog.golang.org/error-handling-and-go
https://blog.golang.org/examples
https://blog.golang.org/first-go-program
https://blog.golang.org/generate
https://blog.golang.org/gif-decoder
https://blog.golang.org/go1.13-errors
https://blog.golang.org/godoc
https://blog.golang.org/h2push
https://blog.golang.org/json
https://blog.golang.org/maps
https://blog.golang.org/organizing-go-code
https://blog.golang.org/slices
https://blog.golang.org/slices-intro
https://blog.golang.org/context-and-structs
