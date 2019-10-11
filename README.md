# Gojis

[![Build Status](https://api.travis-ci.com/gojisvm/gojis.svg?branch=develop)](https://travis-ci.com/gojisvm/gojis)
[![Codacy Badge](https://api.codacy.com/project/badge/Grade/dd5507e3d34543e3a526b05aaea3eba8)](https://www.codacy.com/app/gojisvm/gojis?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=gojisvm/gojis&amp;utm_campaign=Badge_Grade)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fgojisvm%2Fgojis.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fgojisvm%2Fgojis?ref=badge_shield)

Gojis is an implementation of ECMAScript 2018 (ES 9). It basically is a JavaScript VM, just like
[Goja](https://github.com/dop251/goja) or [Otto](https://github.com/robertkrimen/otto).
The documentation can be found [here](https://gojisvm.github.io).

## Requirements
* `Go 1.11+`

## Why
[Goja](https://github.com/dop251/goja) and [Otto](https://github.com/robertkrimen/otto) are both stuck at implementing _most_ features of ES 5.1. This implementation aims to support ES 9, and, after that maybe even ES 10 and later.

## What is this Gojis VM good for?
The Gojis VM can be run as a standalone, or you can embed it into your Go application by using the API. Go get it with
```bash
go get -u github.com/gojisvm/gojis
```
and start using it with
```go
// FIXME: the API is in draft mode, this is subject to change at any time

vm := gojis.NewVM()

vm.SetFunction("greet", func(gojis.Args) gojis.Object {
    vm.Eval(`console.log("Hello World!");`)
    return nil
})

vm.Eval(`greet();`)
/*
    prints:
    Hello World!
*/
```

For more documentation, please have a look at the [API documentation](https://gojisvm.github.io/api.html).

## What are the goals?
The primary goal of this project is to have fun coding, as I love to code, but thinking about system designs and architectures is difficult. The ECMAScript language specification (which can be found in `/docs`), takes care of most of these things already, so a contributor can really focus on implementation and optimization.

Another goal I am trying to achieve is, to provide the Go community with a JavaScript VM that supports at least ES 6 features.
[Goja](https://github.com/dop251/goja) and [Otto](https://github.com/robertkrimen/otto) are both stuck at implementing ES 5.1, but this implementation does exactly that.

## Current status
There is a [milestone](https://github.com/gojisvm/gojis/milestone/1) to keep track of the implementation progress of ES 9.

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fgojisvm%2Fgojis.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fgojisvm%2Fgojis?ref=badge_large)
