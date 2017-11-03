# Setup

Let's make sure you have a working Go environment.


## Go

Obviously you'll need the Go tool chain. You can download it from the
[downloads section][dl] in the Go web site.

To make sure you have a working Go environment open a terminal/command prompt
and type:

    go version

You should see output similar to

    go version go1.9.2 linux/amd64


[dl]: https://golang.org/dl/


## Editor/IDE

A good editor or [IDE][ide] will boost your productivity by a lot.

You don't have to use one, any odd editor combined with the Go tool chain will do the work,
but I highly encourage you to use an editor/IDE.

Personally I'm using [Vim][vim] with [vim-go plugin][vim-go]. Vim is awesome
and feature packed, however the learning curve might be steep.

Other good options are [Visual Studio Code][vscode] with [Go
extension][vscode-go] and [GoLand][goland].

[ide]: https://en.wikipedia.org/wiki/Integrated_development_environment
[vim]: http://www.vim.org/
[vim-go]: https://github.com/fatih/vim-go
[vscode]: https://code.visualstudio.com/
[vscode-go]: https://marketplace.visualstudio.com/items?itemName=lukehoban.Go
[goland]: https://www.jetbrains.com/go/


## Git

The `go` tool that among other things installs third-party packages uses
[git][git] under the hood to download packages. Most Linux/OSx systems already
have git installed but on Windows you'll need to install it yourself.

To make sure you have a working git open a terminal/command prompt
and type:

    git version

You should see output similar to

    git version 2.15.0

[git]: https://git-scm.com/

## curl

We are going to build a web server and issue HTTP calls to it. Some of them can
be done with the browser, but we'll also make POST requests and for this we'll
use [curl][curl].


To make sure you have a working git open a terminal/command prompt
and type:

    curl https://httpbin.org/user-agent

You should see output similar to

    {
      "user-agent": "curl/7.55.1"
    }


If you're not comfortable with the command line, another option to try is
[Postman][postman] which is a browser extension that let's you do POST requests.

[curl]: https://curl.haxx.se/
[postman]: https://www.getpostman.com/
