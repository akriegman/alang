# A
A is a 2D imperative programming language. It is a frontend for the Rust compiler, so A can do everything Rust can do. A is not a toy, it aims to be your new desert island language. I am speaking in the present tense even though A does not exist yet. This is just a design document.

## Why 2D?

Languages (human languages, math, code) encode trees into 1D streams of information. There is no good way to do this in general, and going up a dimension solves this. In fact, two dimensions is enough for us to encode digraphs instead of just trees, which is what languages should be doing anyways, but they don't because encoding trees is already hard enough in 1D. See my [blog post](https://aaah.run/struck) on this.

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
- `struct`s are (2) or (1) plus (6)
- Arrays are (3)
- `usize` is (4) arguably. Panicking on a failed bounds check is like saying `idx` lives in the anonymous enum with `len` variants.
- `enum`s are (5), (6), (1), and (2) all in one
- `type` is (6)

In A we perform a change of basis in abstraction space and give you the more fundamental set of tools (1)-(6).
> Note that (1) and (3) are indexed by (4), and (2) is indexed by (5).

> Note that there is no homogenous sum type as that would factor.

Product types are constructed with a stack of `)`s and destructed/accessed with `(`s. Sum types are constructed with `}`s and destructed/matched with `{`s. Homogenous types are constructed with `]` and destructed/indexed with `[`. These stacks can have any number of `|`s mixed in. Destructing sum types is the primary (and maybe only, tbd) tool for control flow. Anonymous members/variants can be specified using position, and named members with the identifier. So for example:
```
         'a')
       [----)eq
'a']   |
'b']---[----)eq
'c']     'b')
```
Here both calls to `eq` return `true`.

In A, all functions take one argument and return one argument. Rust functions, which take multiple arguments identified by their position, are wrapped in an A function that accepts a tuple. Likewise, A functions have auto generated Rust wrappers that accept a single argument. An A function that accepts a product type can be given an annotation to make the Rust wrapper accept the members as separate arguments.
> Should we make single arg Rust functions take a singleton tuple or the arg directly? Or should we have no distinction, so `(x,) == x`?

A does not have the standard C-like operators `+, -, *,` etc. Instead you must use the functions `add, sub, neg, mul,` etc. A uses the same syntax as Rust for literals, except that there are no negative number literals. Instead you must negate them with `neg`, like `1-neg`. This is a good thing, postfix negation is better!

`>-...->` is used for closures and functions. So `>->` is the identity function, `x->->` passes `x` to the identity function, `>->-y` passes the identity function to `y`, and `x->->-y` is `y(nop(x))`.
> How should we write `f(x)(y)`? So far we can only call closures or functions that are bound to identifiers. This question might be answered once we decide how to piece expressions together into programs.

Since anonymous product types are indexed by anonymous sum types and named product types are indexed by named sum types, we could use the same notation for dynamic indexing as dynamic lensing. We could have static indexes/lenses only for bracket stacks, but for single brackets allow an index/lens to be chosen dynamically through a wire to the top or bottom of the bracket. However, we should probably not have special syntax for dynamic indexing/lensing, and instead just use the `index` and `get` methods.

Note that indexing a tuple with the anonymous enum can return a different type depending on the variant. Rust type semantics would not allow for this if indexing is a function taking a `usize`. So this is a substantial difference between `usize` and the anonymous enum.

`!` can be used to refer to no member/variant. So the unit type's value can be written `!)`, so a nullary function can be called like `!)start`. A value can be dropped while keeping its wire around for control flow purposes using `-(!-`. `!}` can be used to panic. `{!` is unreachable. So for example we can do an assertion:
```
         {false-!}
1)add-)eq{true-----...
1)   2)
```
Or, we could have `bool` be the anonymous enum with two variants, and discard the `true` arm if we just want to make the assertion and continue execution elsewhere:
```
1)add-)eq{-!}
1)   2)
```

`*)` is used to splat a value into the remaining fields, `(*` to take the remaining fields in a new smaller product type, and `{*` as a wildcard when matching. `(*` and `{*` both create a new type with fewer members/variants. `*}` can be used to merge such a sub-enum value back into the larger enum. So a single field of a struct can be updated like this:
```
alice---(age---)add---age)--=alice
        |     1)         | 
        (*--------------*)
```
`(*` can only take a suffix of a tuple.
> Should we allow types to mix anonymous and named fields? I don't see why not...

Variables are set, and in general things are named, using `=`. So a function is declared like `>--...-->---=foo`.

`:` is used for type judgements. Functions can have multiple `>` return points. Any of the return points can be used to refer to the whole function.
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
{false----!)(---
|   "hello")
...
```
> So maybe we want calling values to be syntax sugar for this? idk man...

> For functions we should probably attach the `:` directly to the `>` actually...

Notice that this destructing sum type syntax is way more versatile than what rust provides, where you have to use `if`, `if let`, `let ... else`, `match`, `?`, or any number of `.ok_or()` "convenience" functions depending on the situation. (Sorry I don't mean to snarky. I love Rust! But having to check the documentation just to keep variant handling concise sucks.) For example, `?` becomes
```
...--{Ok---...
     {*->
```
Or something.

We may make `=` without an identifier be a tunnel, like `-=  other_stuff  =-`. This is similar to the other usage of `=` because a tunnel is like an anonymous identifier. We could just have this instead of `%`. And then we could use `%` for this instead. Hmm...

We may allow the program to flow both ways, so that `)` and `(` can both be used for constructing and destructing depending on context, and same for `][}{`. We would have the program flow towards terminals such as `=`, `->`, and `<-`, and away from `>-` and `-<`. Some cases where this would feel more natural are type expressions and the top scope of a file:
```
mod=
   +(package
    |  +(name-"interpreter"
    |   (version-"0.1.0"
    |   (edition-"2024"
    |
    (lib
    |  +(crate_type
    |      +["cdylib"
    |
    (dependencies
       +(rand-"0.9.2"
        (wasm_bindgen-"0.2.106"
        (wasm_bindgen_futures-"0.4.56"
```

## Tooling

We will have a Zed extension for editing A files. We may have the extension squash fonts into squares or create our own square font, but more likely we'll just have it display `+-|` as the box drawing characters `┼─│`. The extension will allow rectangular selections, and dragging selections to any position. It will make enter and tab insert a row or column, automatically extending any wires and sliding over any identifiers this would cut. It will also allow you to treat the buffer as a canvas extending infinitely down and to the right, automatically filling in spaces to the left and trimming spaces to the right. It will also provide a method for laying long wires, either by clicking and dragging or clicking twice, showing a ghost of the wire it will make until you commit or cancel. I guess we'll have a formatter, but god knows how that's gonna work.

There's two main approaches to compiling A: emit one of Rust's IRs, or transpile to Rust. The former will give us more power to iterate on Rust's type system and allow us to produce better error messages. The latter will allow us to publish to crates.io and use Rust macros. We will probably want both.
