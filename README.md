Elisa - AI Programming Tutor for Students
=========================================

1. [Description](#description)
1. [Installation](#installation)
1. [Usage (Students)](#usage-students)
1. [Usage (Lecturers)](#usage-lecturers)
1. [Development](#development)
1. [Build](#build)
1. [Copyright](#copyright)

Description
-----------

Elisa is an intelligent learning aid for students studying computer science. She answers questions
about the exercises, your source code and helps you with the exercises. She has completed her studies
a few years ago. Now she helps you to master your studies.

Elisa is the great-granddaughter of Joseph Weizenbaum's Eliza. Not all functions are yet fully developed.
It is an experiment for students at the DHBW Karlsruhe, which is being carried out as part of the KoLLI
research project. Questions, requests, suggestions and error messages are always welcome.

Installation
------------

TODO: Installation packages / scripts

Usage (Students)
----------------

Since programmers are using the command line a lot, Elisa is meant to be used on the command line, too.
Elise is the most useful when working on programming exercises that have been prepared to be used with Elisa.
Simply open a terminal window, `cd` into the exercise directory as you need to do anyway and run the `elisa`
command from there. Elisa will then present a self-explanatory menu-driven user interface, amongst other
allowing you to …

* Read the exercise descriptn
* Ask questions about the exercise
* Ask questions about your solution
* Get feedback on your solution

The idea is that you use the "ask" commands to ask questions, similar as you would do with with a real teacher
or tutor, e.g.:

* _I am a bit lost. How should I approach this problem?_
* _Why is my program stuck in an endless loop?_
* _I try to iterate over this list, but somehow the first item seems to be skipped._

And once you are finished or you just want to get some feedback on the progress, let Elise read the source
code and give some feedback to you, similar as if you would e-mail the code to your tutor to get his/her
opinion.

Usage (Lecturers)
-----------------

TODO: Usage for lecturers

Development
-----------

This program is written with in Go and should follow the usual development practices of this
language. To get you up to speed, here are a few useful commands that can be executed within
one of the source directories:

* `go mod tidy` – Download new packages and update files `go.mod` and `go.sum`
* `go build .` – Compile program and build a statically linkes executable
* `go run .` – Directly run the program from sources without building it first

Note that you don't need to use `get get` to fetch external packages. Simply use them in the
code and run `go mod tidy` each time a new package is added.

### VS Code: Disable Format on Save

One thing, that I find rather annoying, is that Visual Studio Code by default formats the
source code every time a file is saved. Unfortunately the Go community has a rather strict
imagination of how "canonical" source code should be formatted, most of which is actually
quite sensible. However, there are occasional exceptions, where I would deliberately format
my code differently. And since I am the same stringent, without using an auto-formatter, I
find it rather annoying, when my Code suddenly gets formatted differently.

Therefor, please see the following blog post: [Format on Save in Go with Code](https://blog.boot.dev/golang/format-on-save-vs-code-golang/).
Even though the author argues differently (he doesn't like to think about Code formatting),
it also shows, how to disable the auto-formatting for Go files in Visual Studio Code.
It boils down to the following lines in your `settings.json`:

```json
    "[go]": {
        "editor.formatOnSave": false
    },
    "[go.mod]": {
        "editor.formatOnSave": false
    }
```

Build
-----

TODO: Notes on compilation / cross-compilation for Linux, Windows, Mac

Copyright
---------

Elisa: AI Programming Tutor for Students <br/>
© 2024 Dennis Schulmeister-Zimolong <dennis@wpvs.de> <br>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as
published by the Free Software Foundation, either version 3 of the
License, or (at your option) any later version.

Development funded by the KoLLI research project at DHBW Karlsruhe.