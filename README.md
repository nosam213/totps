# totps
cli utility for generating TOTP codes.

### Flags
* --help
* --secret string   the auth secret (ex: ABCDEFGHIJKLMNOP)
* --version         displays version and exits

## Compiling
NOTE: Requires network connection for the first compile.

### Windows
```
C:\Users\username\totps> dir
go.mod  go.sum  main.go  README.md
C:\Users\username\totps> go build -o totps.exe -ldflags="-w -s -buildid=" .
```

### Unix
```
user@host:/home/user/totps$ ls
go.mod  go.sum  main.go  README.md
user@host:/home/user/totps$ go build -o totps -ldflags="-w -s -buildid=" .
```