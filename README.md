A is a 2D imperative programming language. It is a frontend for the Rust compiler, so A can do everything Rust can do. A is not a toy, it aims to be your new desert island language. I am speaking in the present tense even though A does not exist yet. This is just a design document.

## Why 2D?

Languages (human languages, math, code) encode digraphs into 1D streams of information. There is no good way to do this in general (see [my blog post](https://aaah.run/struck)). Going up a dimension solves this.

I don't know about you, but when I read code I essentially translate it into a diagram in my head. Writing the code as a diagram in the first place takes the cognitive load off of the reader.

## Features

An expression in A is connected together by wires that carry values. The wires are drawn with `-`, `|`, and `+`. Two wires that do not look like they connect do not connect, so `+|-` is three disconnected wires next to each other. Two wires can pass over each other using `%`.

Typically, a programming language lets you define types using some combinations of some subset of these tools:

1) product types with anonymous members
2) product types with named members
3) homogenous product types
4) sum types with anonymous variants
5) sum types with named variants
6) naming a type

Rust gives us access to all these tools, but only in specific combinations:

- Tuples are (1)
- `struct`s are (2) plus (6)
- Arrays are (3)
- `usize` is (4) arguably. Panicing on a failed bounds check is like saying `idx` lives in the anonymous enum with `len` variants.
- `enum`s are (5), (6), (1), and (2) all in one
- `type` is (6)

In A we perform a change of basis in abstraction space and give you the more fundamental set of tools (1)-(6).
> Note that (1) and (3) are indexed by (4), and (2) is indexed by (5).

> Note that there is no homogenous sum type as that would factor.

Product types are constructed with a stack of `)`s and destructed/accessed with `(`s. Sum types are constructed with `}`s and destructed/matched with `{`s. Homogenous types are constructed with `]` and destructed/indexed with `[`. These stacks can have any number of `|`s mixed in. Destructing sum types is the primary (and maybe only, tbd) tool for control flow. Anonymous members/variants can be specified using position or number literals, and named members with the identifier. So for example:
```
      'a'---)
            |
       [----)---eq
'a']   |
'b']---[2---)---eq
'c']        |
      'c'---)
```
Here both calls to `eq` return `true`.

In A, all functions take one argument and return one argument. Rust functions, which take multiple arguments identified by their position, are wrapped in an A function that accepts a tuple. Likewise, A functions have auto generated Rust wrappers that accept a single argument. An A function that accepts a product type can be given an annotation to make the Rust wrapper accept the members as separate arguments.
> Should we make single arg Rust functions take a singleton tuple or the arg directly? Or should we have no distinction, so `(x,) == x`?

A does not have the standard C-like operators `+, -, *,` etc. Instead you must use the functions `add, sub, neg, mul,` etc. A uses the same syntax as Rust for literals, except that there are no negative number literals. Instead you must negate them with `neg`, like `1-neg`. This is a good thing, postfix negation is better!

`>-...->` is used for closures and functions. So `>->` is the identity function, `x->->` passes `x` to the identity function, `>->-y` passes the identity function to `y`, and `x->->-y` is `y(nop(x))`.
> How should we write `f(x)(y)`? So far we can only call closures or functions that are bound to identifiers. This question might be answered once we decide how to piece expressions together into programs.

`!` can be used to refer to no member/variant. So the unit type's value can be written `!)`, so a nullary function can be called like `!)start`. A value can be dropped while keeping its wire around for control flow purposes using `-(!-`. `!}` can be used to panic. `{!` is unreachable. `*)` is used to splat a value into a product type constructor, `(*` to take the remaining fields in a new smaller product type, and `{*` as a wildcard when matching. `*}` is used in situation where a wire carries multiple variants, TBD how that will work. So a single field of a struct can be updated like this:
```
alice---(age---)add---age)--=alice
        |   1--)         | 
        (*--------------*) 
```
`(*` can only take a suffix of a tuple.
> Should we allow types to mix anonymous and named fields? I don't see why not.

Variables are set, and in general things are named, using `=`. So a function is declared like `>--...-->---=foo`.

`:` is used for type assertions. Functions can have multiple return points with `>`. Any of the return points can be used to refer to the whole function.
```
bool:
    |
>---+--{false--"hello"--->
       |
       {true---"world"-+->--=foo
                       |
                   str&:
```
> Here we're basically calling a string literal on the unit type. More correct would probably be
```
{false----!)
| "hello"--)
...
```
> So maybe we want to calling literals to be syntax sugar for this? idk man...

> For functions we should probably attach the `:` directly to the `>` actually...

Notice that this destructing sum type syntax is way more versatile than what rust provides, where you have to use `if`, `if let`, `let ... else`, `match`, `?`, or any number of `.ok_or()` "convenience" functions depending on the situation. (Sorry I don't mean to snarky. I love Rust! But having to check the documentation just to keep variant handling concise sucks.) For example, `?` becomes
```
...-{Ok----...
    {Err->
```
Or something.

## Tooling

We will have a Zed extension for editing A files. We'll try having the extension squash fonts into squares, but more likely we'll create our own square font. The extension will allow rectangular selections. It will have hotkeys for inserting a column or row, automatically extending any wires and sliding over any identifiers this would cut. It will also allow you to treat the buffer as a canvas extending infinitely down and to the right, automatically filling in spaces to the left and trimming spaces to the right. I guess we'll have a formatter, but god knows how that's gonna work.
