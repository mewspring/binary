# NOTICE: New project home!

The binary library has been merged with the `decomp/exp` project (see [mewspring/binary#8](https://github.com/mewspring/binary/issues/1)).

For active development, the new home of `mewrev/binary` is the [decomp/exp](https://github.com/decomp/exp) repository, under the import path `bin`.

```diff
-import "github.com/mewrev/binary"
+import "github.com/decomp/exp/bin"
```

## WIP

This project is a *work in progress*. The implementation is *incomplete* and subject to change. The documentation can be inaccurate.

# binary

[![Build Status](https://travis-ci.org/mewrev/binary.svg?branch=master)](https://travis-ci.org/mewrev/binary)
[![Coverage Status](https://img.shields.io/coveralls/mewrev/binary.svg)](https://coveralls.io/r/mewrev/binary?branch=master)
[![GoDoc](https://godoc.org/github.com/mewrev/binary?status.svg)](https://godoc.org/github.com/mewrev/binary)

The binary project aims to provide a uniform representation for the data of ELF and PE binary executables.

## Documentation

Documentation provided by GoDoc.

- [binary]: provides a uniform representation for the data of binary executables.

[binary]: http://godoc.org/github.com/mewrev/binary

## Public domain

The source code and any original content of this repository is hereby released into the [public domain].

[public domain]: https://creativecommons.org/publicdomain/zero/1.0/
