notificationsapp
================

[![Build Status](https://travis-ci.org/shurcooL/notificationsapp.svg?branch=master)](https://travis-ci.org/shurcooL/notificationsapp) [![GoDoc](https://godoc.org/github.com/shurcooL/notificationsapp?status.svg)](https://godoc.org/github.com/shurcooL/notificationsapp)

Package notificationsapp is a web frontend for a notifications service.

Note, the canonical issue tracker for this package is currently hosted at
https://dmitri.shuralyov.com/issues/github.com/shurcooL/notificationsapp.
Its notifications are implemented using this very package.

Installation
------------

```bash
go get -u github.com/shurcooL/notificationsapp
```

Development
-----------

This package relies on `go generate` directives to process and statically embed assets. For development only, you may need extra dependencies. You can build and run the package in development mode, where all assets are always read and processed from disk:

```bash
go build -tags=dev something/that/uses/notificationsapp
```

When you're done with development, you should run `go generate` and commit that:

```bash
go generate github.com/shurcooL/notificationsapp/...
```

License
-------

-	[MIT License](https://opensource.org/licenses/mit-license.php)
