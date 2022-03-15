# uuid-base58 ![workflow](https://github.com/AlexanderMatveev/go-uuid-base58/actions/workflows/go.yml/badge.svg)

## Overview

`uuid-base58` is a small package to generate short [BASE58-string](https://en.bitcoin.it/wiki/Base58Check_encoding) representation of `UUID` and vice versa.

## Example

```go
package main

import (
	"fmt"
	uuid58 "github.com/AlexanderMatveev/go-uuid-base58"
	"github.com/google/uuid"
)

func main() {
	// base58 from UUID
	s, _ := uuid58.ToBase58(uuid.MustParse("c73087da-627a-4f00-8786-fcc4f47db57f"))
	fmt.Println(s) // RbcZUzUfnGugha7DVu3fAE

	// UUID from base58
	u, _ := uuid58.FromBase58("RbcZUzUfnGugha7DVu3fAE")
	fmt.Println(u) // c73087da-627a-4f00-8786-fcc4f47db57f
}

```

## Motivation

The original Bitcoin client source code explains the reasoning behind base58 encoding:

**base58.h:**

```
// Why base-58 instead of standard base-64 encoding?
// - Don't want 0OIl characters that look the same in some fonts and
//      could be used to create visually identical looking account numbers.
// - A string with non-alphanumeric characters is not as easily accepted as an account number.
// - E-mail usually won't line-break if there's no punctuation to break at.
// - Doubleclicking selects the whole number as one word if it's all alphanumeric.
```
