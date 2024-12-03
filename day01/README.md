# Day 1

Wow, that was fun! This was really **my first time writing something in Go**, _ever_.

## Things I learned about Go

- **Reading from a file**: use `os.ReadFile(path)`, which returns a `byte[]`. Since a `string` is just a slice of bytes, you can get the text of a file using `string(os.ReadFile(path))`.
- **String stuff**: the `strings` library has a bunch of string utils, like `TrimSpace`, `ReplaceAll`, and `Split`. But my favorite is to actually use `Fields` instead of `Split` when splitting on whitespace, since it will account for a variable length of contiguous whitespace characters.
- **Data structures**: Python and JavaScript have spoiled me by removing my need to differentiate between an _array_ and a _list_, or between an _object/dict_ and a _map_. I had to learn about **slices** and **maps** in Go, and how to manipulate them.
- **Loops**: Every loop in Go is a `for` loop, which _drastically simplifies things_.
