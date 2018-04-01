[![Build Status](https://travis-ci.org/zhtheyun/fibonacci.svg?branch=master)](https://travis-ci.org/zhtheyun/fibonacci) [![GoDoc](https://godoc.org/github.com/zhtheyun/fibonacci?status.svg)](http://godoc.org/github.com/zhtheyun/fibonacci)

# Fibonacci WebService
Project for fibonacci web service demo




## Quick Start

### Prerequest

1. Install golang in your local machine. the minimum version is 1.9.0


### To check the running dev enviroment locally
In ideal world, it should run "dep status" to check the status.
But in China, due to GFW, this command sometimes hanging due to network connection. shamed about it.

```sh
make check_env
```


### To build the binary

```sh
git clone {{repo}}
make build
```
### To perform unit test

```sh
make test
```


### To static check the code

After you finish your coding, you need to run this to make sure your code follows the code convention. So you can run this

```sh
make lint
```

Fix the smelled code until the command successfully returns.

## Usage
### Check the usage
./fibonacci -h

### Startup the service using default parameter
```
./fibonacci rest

```
### Startup the service using customized parameter

Startup the service at port 8090, setup the cache numbers to 2000, the maximum numbers 20000 and loglevel to INFO
The admin should using the cache size to tuning the performance.

```
export FIB_PORT=8090 FIB_CACHEDNUMBERS=2000 FIB_MAXIMUMNUMBERS=20000 FIB_LOGLEVEL=INFO && ./fibonacci rest

```

### Retrive the first 100 fibonacci numbers
```
curl -XGET -i "http://127.0.0.1:8080/fibonacci?numbers=100"

```
