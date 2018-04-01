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

