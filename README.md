# goobar

Collection of Go code-blocks.

Bottom line: goobar is FUBAR; keep your expectations low.

While this repo appears to be FUBAR and very messy, it is worth recognizing that goobar's primary use is for copy-paste into environments where import of public packages is not permitted. So while goobar *is* FUBAR, it is an intentional design choice.

This intent is why the repository is more focused on the blocks as stand-alone. For example, you will find code like this:

[:]: cSpell:ignore "anotherthing"

```go
package something // GOOBAR delete

import (
    "github.com/dimitrilw/goobar/anotherthing" // GOOBAR also copy anotherthing
)

type Something struct { data AnotherThing }

func NewSomething() Something { return Something{ AnotherThing{} } }
```

That code works well for copy-paste. It looks messy, but works for the intended purpose: I copy the entire file content, then search for "GOOBAR" and follow the instructions -- if the instructions are there; I often omit them in simple copy-paste code blocks. But if goobar was imported, it would be clunky:

```go
mySomething := something.NewSomething()
```

If goobar was intended for use as an imported package, the constructor function would be better as:

```go
func New() Something { return Something{ AnotherThing{} } }
```

That code has better constructor naming. When imported, it would result in user code more similar to:

```go
mySomething := something.New()
```

However, goobar does not show good code style in terms of package naming, imports, function names, inclusion of documentary comments for `go doc` to consume, etc. 

For example, if goobar always uses the function name "New", then we would have function name conflicts when copy-pasting multiple code-blocks. Therefore, constructors follows the `NewSomething` convention. 

While this repo might be filled with anti-patterns in *real* package design, the fact is that goobar is FUBAR on purpose. It is not intended to be a *real* package that is clean, tidy, and built for users to import.

For that matter, goobar is not designed for collaborative development. It is a personal repo that I leave public so that others may copy-paste as desired. It does not follow good git behavior: commit messages, squashing, branching, etc.

Bottom line: goobar is FUBAR; keep your expectations low.
