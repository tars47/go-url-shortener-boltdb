# URL Shortener

## details

The goal of this exercise is to create an [http.Handler](https://golang.org/pkg/net/http/#Handler) that will look at the path of any incoming web request and determine if it should redirect the user to a new page, much like URL shortener would.

For instance, if we have a redirect setup for `/dogs` to `https://www.somesite.com/a-story-about-dogs` we would look for any incoming web requests with the path `/dogs` and redirect them.

## POST

```bash
  curl --location 'localhost:4747/' \
--header 'Content-Type: application/json' \
--data '{
    "path":"/goog","url":"https://google.com"
}'
```

## GET

```bash
  curl --location 'localhost:4747/goog'
```

## DELETE

```bash
 curl --location --request DELETE 'localhost:4747/123' \
--header 'Content-Type: application/json' \
--data '{
    "path":"123","url":"http://google.com"
}'
```
