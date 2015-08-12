[Mark](https://github.com/a8m/mark) command-line tool
====

### Intallation
```sh
$ go get github.com/a8m/mark-cli
```

### Usage
```sh
$ mark-cli -i hello.text -o hello.html
```
or you can pipe to it
```sh
$ echo 'hello __world__...' | mark-cli -smartypants
<p>hello <strong>world</strong>…</p>
```

### License
MIT
