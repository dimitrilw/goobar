# goobar

Collection of Go code-blocks.

Bottom line: goobar is FUBAR; keep your expectations low.

While this repo appears to be FUBAR and very messy, it is worth recognizing
that goobar's primary use is for copy-paste into environments where import of 
public packages is not permitted. So while goobar *is* FUBAR, 
it is an intentional design choice.

This intent is why the repository is more focused on the blocks as stand-alone.
For example, you will find code like this:

[:]: cSpell:ignore "anotherthing"

```go
package something // GOOBAR delete

import (
    "github.com/dimitrilw/goobar/anotherthing" // GOOBAR also copy anotherthing
)

type Something struct { data AnotherThing }

func NewSomething() Something { return Something{ AnotherThing{} } }
```

That code works well for copy-paste. It looks messy, but works for the intended
purpose: I copy the entire file content, then search for "GOOBAR" and follow 
the instructions -- if the instructions are there; I often omit them in simple 
copy-paste code blocks. But if goobar was imported, it would be clunky:

```go
mySomething := something.NewSomething()
```

If goobar was intended for use as an imported package, 
the constructor function would be better as:

```go
func New() Something { return Something{ AnotherThing{} } }
```

That code has better constructor naming. When imported, 
it would result in user code more similar to:

```go
mySomething := something.New()
```

However, goobar does not show good code style in terms of package naming, 
imports, function names, inclusion of documentary comments for 
`go doc` to consume, etc. 

For example, if goobar always uses the function name "New", then we would have 
function name conflicts when copy-pasting multiple code-blocks. 
Therefore, constructors follows the `NewSomething` convention. 

While this repo might be filled with anti-patterns in *real* package design, 
the fact is that goobar is FUBAR on purpose. It is not intended to be 
a *real* package that is clean, tidy, and built for users to import.

For that matter, goobar is not designed for collaborative development. 
It is a personal repo left public so others may copy-paste as desired. 
It does not follow good git behavior: 
commit message style, squashing, branching, etc.

Bottom line: goobar is FUBAR; keep your expectations low.


## imports

A special note regarding the "imports" directory: 
This directory is code from other projects. Since goobar is intended for 
copy-paste operations into non-import environments, I occasionally need to make
reference to these sources, such as the pre-defined constraints for generics. 
Therefore, `goobar/imports` is a collection of misc reference material.

I am trying to be a good steward and respect the source of the code place into 
`imports`. I am not trying to pass off the code as my own. I am trying to give 
credit where credit is due. I am trying to make it easy to find the source of 
the code. I am trying to make it easy to find the license of the code.

When referencing this code, I keep re-created code to an absolute minimum. 
For example, I might do a code challenge and wish to remind myself of all the 
various Integer types when making a custom generic. So I generally use these 
files more like a reference collection than a source of code to copy-paste.

## Assumptions

I have a few assumptions regarding the environment where goobar is developed:

- Go is installed. See https://golang.org/doc/install
- Taskfile is installed. See https://taskfile.dev/#/installation
- `golangci-lint` is installed. See https://golangci-lint.run/usage/install/
- If using VS-Code, then cSpell extension is your spellchecker. See 
    https://marketplace.visualstudio.com/items?itemName=streetsidesoftware.code-spell-checker
