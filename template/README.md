# Template

> This example is based on one [Istio PR](https://github.com/istio/istio/pull/9731).

There are 2 templates in Golang: `text/template` and `html/template`. This example illustrate the difference between them.

One key difference is that `html/template` will handler input data content. 
In order to keep the content, need to add a `raw` function. `text/template` does not handler input data content.

## Usage

Commend:

```
$ template [text|html] <input-yaml-path> <output-go-path>
```

Parameters:

* [text|html]: The template package to use.
* input-yaml-path: The path of metadata for input.
* output-go-path: The path for output.
