# goobar

Collection of Go code-blocks.

While this repo appears to be FUBAR and very messy, it is worth recognizing that goobar's primary use is for copy-paste into environments where import of public packages is not permitted. So while goobar *is* FUBAR, it is an intentional design choice.

This intent is why the repository is more focused on the blocks as stand-alone. For example, you will find code like this:

[:]: cSpell:ignore "anotherthing"

```go
package goobar // GOOBAR delete

// GOOBAR also copy anotherthing.go

type Something struct { data AnotherThing }

func NewSomething() Something { return Something{ AnotherThing{} } }
```

That code works well for copy-paste. It looks messy, but works for the intended purpose: I copy the entire file content, then search for "GOOBAR" and follow the instructions. If goobar was imported, it would be clunky:

```go
import ".../goobar/something"

mySomething := something.NewSomething()
```

If goobar was intended for use as an imported package, the code would be better as:

```go
package something

import "...goobar/anotherthing"

type Something struct { data AnotherThing }

func New() Something { return Something{ AnotherThing{} } }
```

Note the package name change to a sub-package; this helps ensure clean tests. Also, I would possibly have both `something_test.go` and `something_internal_test.go` files. Test file `something_test.go` would be the black-box tests where the package name would be `something_test`, forcing an import of Something, thus limiting the test to accessing the publicly exposed API. Meanwhile `something_internal_test.go` would allow for white-box testing, where the package name would be `something` and access to the internal API would be permitted.

In addition, that package code has better constructor naming. When imported, it would result in user code more similar to:

```go
import ".../goobar/something"

mySomething := something.New()
```

However, goobar does not show good code style in terms of package naming, imports, function names, etc. 

For example, if goobar always uses the function name "New", then we would have function name conflicts when copy-pasting multiple code-blocks. Therefore, constructors follows the `NewSomething` convention. 

While this repo might be filled with anti-patterns in *real* package design, the fact is that goobar is FUBAR on purpose. It is not intended to be a *real* package that is clean, tidy, and built for users to import.
