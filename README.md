urlsign
=======

[![godoc here](https://img.shields.io/badge/godoc-here-blue.svg)](http://godoc.org/github.com/Nitro/urlsign)

Package urlsign contains a signed URL mechanism, where a URL can safely be
passed through a third party and validated before being served. 

The Problem This Solves
-----------------------

If you have one service that generates a URL which is then passed to a browser,
and the browser then uses that URL to access the resource from another service,
you need a way to identify that this is a valid request. One method for doing that
is to provide a URL that contains a token representing an HMAC over the content of
the URL. This library implements that.

How it Works
------------

urlsign generates a signing token for each URL based on all the other query
parameters and the path. This can then be included as the `token` parameter
appended to any signed URL.  Using this token, the authenticity of the request
can be validated.

Tokens have a validity window much like that of TOTP 2-factor auth systems.
The library will validate a token from the current window, the previous window,
and a future window. Since browsers use query parameters in caching
determination, this token window will also affect the expiration of the cached
resource in the browser.

This does not validate the hostname or scheme from the passed URL.
Expiration/bucket size is an external, agreed parameter between the services.
