# webpage-requisites-go

Extract webpage requisite URLs from HTML and CSS documents.
Requisites are the resources a browser would load to render the web page.

## Installation
```bash
go get github.com/birneee/webpage-requisites-go
```

## Detected Requisites
Requisites are:

- `src` of `img` elements, not starting with `data:`
- `href` of `link` elements with `rel` "stylesheet", "manifest" or "shortcut icon"
- `src` of `script` elements
- `url` values in stylesheets, not starting with `data:`

Non-requisites are:

- `href` of `a` elements
- ...
