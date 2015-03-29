# ${CoolProjectNameToBeDefined}

Any ideas?

# Configuring

To configure your own databases, put a file named <code>config.json</code> in the working directory. If this file does not exist, the server will
try to find the <code>config-example.json</code> and start with the test database for demonstration and/or testing purposes.

# Compiling and running from source

You'll need Go 1.4. Just for compiling older Go versions would do as well, but I use <code>testing.M</code> for test fixture setup which comes with Go 1.4.

To setup your Go environment have a look at [](https://golang.org/doc/code.html) and [](http://skife.org/golang/2013/03/24/go_dev_env.html).

${CoolProjectNameToBeDefined} needs Mattn's [go-sqlite3](https://github.com/mattn/go-sqlite3)

```
   go run serve.go ./..
```

You will see something similar to this in the stdout:

```
$ go run serve.go ./...
2015/03/29 16:04:51 WARNING: Expected config file config.json not found, using config-example.json for demonstration and test purposes.
2015/03/29 16:04:51 Got configuration config.Configuration{Port:"7242", Databases:[]config.DatabaseConfig{config.DatabaseConfig{Alias:"db", DbDriver:"sqlite3", DbConnectionString:"db/testdata/f-spot-test.db"}}}
2015/03/29 16:04:51 Registering handler for sqlite3 database db/testdata/f-spot-test.db under /db
2015/03/29 16:04:51 Server up and running under port 7242. Go to /config to see the actual configuration of databases.
```

# Runing tests

```
    go test ./...
```

# Installing

```
    go install
```

You will find the binary <code>${CoolProjectNameToBeDefined}</code> under <code>$GOPATH/bin</code>.
