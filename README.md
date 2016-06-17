# go-osinfo

This library reads linux distribution release informations from: `/etc/os-release`

All fields defined by [freedesktop.org](https://www.freedesktop.org/software/systemd/man/os-release.html) are supported.

Distribution specific fields are stored in a `map[string]string`-Field calles `Unknown`

Download and install it:
```
$ go get -u github.com/meshwalker/go-osinfo
```
