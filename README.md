# goobar

Collection of Go code-blocks.

While this repo appears to be FUBAR and very messy, it is worth recognizing that goobar's primary use is for copy-paste into environments where import of public packages is not permitted. So while goobar *is* FUBAR, it is an intentional design choice.

This intent is why the repository is more focused on the blocks as stand-alone. For example, you will find code like this:

```go
type Something struct { someData int }

func NewSomething() *Something { return &Something{} }
```

That code works well for copy-paste. If it was imported, it would be clunky:

```go
import ".../goobar/something"

mySomething := something.NewSomething()
```

If goobar was intended for use as an imported package, the code would be better as:

```go
type Something struct { someData int }

func New() *Something { return &Something{} }
```

That code would result in imported use more similar to:

```go
import ".../goobar/something"

mySomething := something.New()
```

However, if goobar always uses the function name "New", then we would have function name conflicts when copy-pasting multiple code-blocks. Therefore, every constructor follows the `NewSomething` convention. It might be an anti-pattern in *real* package design, but goobar is FUBAR and is not intended to be a *real* package that users would import.
