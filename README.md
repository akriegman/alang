A is a 2D imperative programming language. It is a frontend for the Rust compiler, so A can do everything Rust can do. A is not a toy, it aims to be your new desert island language.

## Why 2D?

Languages (human languages, math, code) encode digraphs into 1D streams of information. There is no good way to do this in general (see [my blog post](https://aaah.run/struck)). Going up a dimension solves this.

I don't know about you, but when I'm reading code I'm essentially translating it into a diagram in my head. Writing the code as a diagram in the first place takes the cognitive load off of the reader.

## Features

An expression in A is connected together by wires that carry values. The wires are drawn with `-`, `|`, and `+`. Two wires that don't look like they connect do not connect, so eg `+|-` is three disconnected wires next to each other. `%` is used to allow two wires to pass over each other.

Typically, a programming language lets you define types using some combinations of some subset of these tools:

- a) product types with anonymous members
- b) product types with named members
- c) homogenous product types
- d) sum types with anonymous variants
- e) sum types with named variants
- f) naming a type

Rust gives us access to all these tools, but only in specific combinations:

- Tuples are (a)
- `struct`s are (b) plus (f)
- Arrays are (c)
- `usize` is (d) arguably. Panicing on a failed bounds check is like saying `idx` lives in the anonymous enum with `len` variants.
- `enum`s are (e), (f), (a), and (b) all in one
- `type` is (f)

In A we perform a change of basis in abstraction space and give you the more fundamental set of tools (a)-(f).
> Note that (a) and (c) are indexed by (d), and (b) is indexed by (e).

> Note that there is no homogenous sum type as that would factor.

Product types are constructed with a stack of `)`s and destructed/accessed with `(`s. Sum types are constructed with `}`s and destructed/matched with `{`s. Homogenous types are constructed with `]` and destructed/indexed with `[`. These stacks can have any number of `|`s mixed in. Destructing sum types is the primary (and maybe only, tbd) tool for control flow. Members and variants can be specified using position or number literals, and named members with the identifier. So for example:
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

A does not have the standard ascii operators `+, -, *,` etc. Instead you must use the functions `add, sub, neg, mul,` etc. A uses the same syntax as Rust for literals, except that there are no negative number literals. Instead you must negate them with neg, like `1-neg`. This is a good thing, postfix negation is better!

`>-...->` is used for closures and functions. So `>->` is the identity function, `x->->` passes `x` to the identity function, `>->-y` passes the identity function to the function `y`, and `x->->-y` is `y(nop(x))`.
> How should we write `f(x)(y)`? So far we can only call functions that are bound to identifiers, or closures. This question might be answered once we decide how to piece expressions together into programs.

The unit type's value is written `!)`, so a nullary function is called like `!)start`. A value can be dropped while keeping its wire around for control flow purposes using `-(!-`. `!}` and `{!` can be tacked on to a sum type constructor / destructor with no effect. `*)` is used to splat a value into a product type constructor, `(*` to take the remaining fields in a new smaller product type, and `{*` as a wildcard when matching. `*}` is meaningless. So a single field of a struct can be updated like this:
```
alice---(age---)add---age)---=alice
        |  1---)         | 
        (*--------------*) 
```
`(*` can only take a suffix of a tuple.
> Should we allow types to mix anonymous and named fields? I don't see why not.

Variables are set, and in general things are named, using `=`. So a function is declared like `>--...-->---=foo`.

`:` is used for type assertions. Functions can have multiple return points with `>`. These return points can be tied together or assigned to the same identifier.
```
bool:
    |
>---+--{false--"hello"--->-+
       |                   |
       {true---"world"-+->-+-=foo
                       |
                   str&:
```
> Here we're basically calling a string literal on the unit type. More correct would probably be
```
{false--!)
"hello"--)
```
> So maybe we want to call literals as syntax sugar for this? idk man...

> For functions we should probably attach the `:` directly to the `>` actually...
