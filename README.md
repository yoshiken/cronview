# cronview

--------------

Simple visualization of crontab.

# How to use

```
./cornview -f path/to/crontabfile
```

It is output to the `output` directory.

When you open `index.html` in it, the visualized crontab is displayed.

# Development

## Initial setup

```
docker build -t cronview .
```

## build

```
docker run -v $PWD/:/go/src/app/ cronview go build src/main.go
```

---------------
# LICENSE

MIT
