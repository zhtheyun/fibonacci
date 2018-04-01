# Fibonacci WebService
Project for fibonacci web service demo

## Quick Start

### Prerequest

1. Install golang in your local machine. the minimum version is 1.8.0

### To setup the running dev enviroment locally

```sh
make restore
```


### To build the binary

```sh
git clone {{repo}}
make clean build
```


### To static check the code

After you finish your coding, you need to run this to make sure your code follows the code convention. So you can run this

```sh
make lint
```

Fix the smelled code until the command successfully returns.

### To add the a external go package in the code

```sh
godep get github.com/pkg/errors
```
