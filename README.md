loggy
=====

no-frills irc room logger with git backend

### Installation

    $ go get github.com/ds300/loggy && go install github.com/ds300/loggy

### Usage

Make sure that the current working directory is a git repo.

    $ git init

Assuming `$GOPATH/bin` in in your `$PATH`, you can run loggy with 

    $ loggy path/to/conf.toml

or simply

    $ loggy

if the current working directory contains a file called `loggy-conf.toml`.
See the `loggy-conf.example.toml` file for configuration ideas.

### License

The MIT License (MIT)

Copyright (c) 2014 Everyone

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.

#### | (• ◡•)|/ \\(❍ᴥ❍ʋ)