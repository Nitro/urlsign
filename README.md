urlsign
=======
Package urlsign contains a signed URL mechanism, where a URL can safely be
passed through a third party and validated before being served. This is useful
for passing a URL to a browser, for example, from one service and having a
second service be certain the URL was as authorized. This is handled by
generating a signing token for each URL based on all the other query parameters
and the path. This does not validate the hostname or scheme from the passed
URL. Expiration/bucket size is an external, agreed parameter between the
services.
