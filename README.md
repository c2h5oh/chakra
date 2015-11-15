chakra
===

[![GoDoc Widget]][GoDoc]

`chakra` allows you to easily add resource/operation based access control to https://github.com/pressly/chi

You can define any access control rules you want as a middleware - `chakra` will make sure they are checked.


## Forcing security as you code with minimum overhead

`chakra` will:

* build resource routes (identifiers) for your access checks
* inject access control right before the last handlers in chain - not optional
* not let you run your code without providing access control function - it will panic
* not let you create new router without providing the next part of resource route or explicitly telling it to use parent one - it will panic

It's not airtight - you can still escape the added security, but it requires more effort than being secure.
You can write a buggy access control function too - it can't help with that - but at least you won't forget to use it.


## How to use it?

Almost exactly the same way you'd use `chi`

`chi` without access control
```go
r := chi.NewRouter()
r.Mount("/api", func(r chi.Router){
    r = chi.NewRouter()
    r.Post("/endpoint1", handler1)
    ...
})
```

`chakra` - `chi` with access control
```go
chakra.SetAC(myAccessControlRules) // you only do this once

r := chakra.NewRouter(chakra.UseParentRoute)
r.Mount("/api", func(r chi.Router){
    r = chakra.NewRouter("secure_api")
    r.Post("/endpoint1", handler1)
    ...
})
```

And you are DONE! - `myAccessControlRules` will be called right before `handler1` to check permissions to `POST` to `{"secure_api", "endpoint1"}` resource


## Examples

[Example access control function](https://github.com/c2h5oh/chakra/blob/master/example/access_control.go)


## Credits

* Peter Kieltyka for https://github.com/pressly/chi
* [Pressly](https://pressly.com)


## TODO

* More examples
* A lot of unit tests (and then some!)

Contributions are always welcome - fork it, do your thing, open a pull request!

## License

Copyright (c) 2015 [Maciej Lisiewski](https://github.com/c2h5oh)

Licensed under [MIT License](./LICENSE)

[GoDoc]: https://godoc.org/github.com/c2h5oh/chakra
[GoDoc Widget]: https://godoc.org/github.com/c2h5oh/chakra?status.svg
