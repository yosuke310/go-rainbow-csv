`go-rainbow-csv` is color highlight tool for text.

# install

```
$ go get github.com/yosuke310/go-rainbow-csv
```

# usage

![image](https://user-images.githubusercontent.com/5753877/56082190-a0d1a800-5e50-11e9-93df-3dd2d19daa89.png)

```
$ echo aaa,bbb,ccc | go-rainbow-csv
```

```
# -d option is set delimiter
$ echo aaa bbb ccc | go-rainbow-csv -d=\ 
```

```
# -f option is set filepath
$ echo aaa,bbb,ccc > sample.csv
$ go-rainbow-csv -f=sample.csv
```
