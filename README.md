# repository
Golang repository pattern

This package looks to expose the [repository pattern] pattern using golang generics and interfaces.
Implementations of these patterns can be found in the `impl` directory, and are independently released from the core interfaces, using [go multi-module repositories].
Repository interface is inspired by [gorm] mixed with some [appengine datastore] for strongly typed queries.

This repository (pun intended) is under ACTIVE DEVELOPMENT... ANYTHING CAN BREAK... ANYTIME!!!
Use at your own risk, or create an issue to ask for a 1.0 release.


## Implementations

* [gorm]
* [sqlstruct]
* more ... please constribute!

[appengine datastore]: https://pkg.go.dev/google.golang.org/appengine/datastore
[go multi-module repositories]: https://github.com/golang/go/wiki/Modules#faqs--multi-module-repositories
[gorm]: https://pkg.go.dev/gorm.io/gorm
[repository pattern]: https://docs.microsoft.com/en-us/dotnet/architecture/microservices/microservice-ddd-cqrs-patterns/infrastructure-persistence-layer-design#the-repository-pattern
[sqlstruct]: https://pkg.go.dev/github.com/kisielk/sqlstruct
