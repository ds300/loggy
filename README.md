loggy
=====

no-frills irc room logger with git backend

### Installation

    go get github.com/ds300/loggy && go install github.com/ds300/loggy

### Usage

Make sure that the current working directory is a git repo.

    $ git init

Assuming `$GOPATH/bin` in in your `$PATH`, you can run loggy with 

    $ loggy path/to/conf.toml

or simply

    $ loggy

if the current working directory contains a file called `loggy-conf.toml`.
See the `loggy-conf.example.toml` file for configuration ideas.

