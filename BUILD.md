# Building

To build **mk**, follow these steps:

**1.** Clone the repo:

```sh
git clone https://github.com/devkcud/mk.git
cd mk
```

**2.** Compile using `make`:

```sh
make build
```

A directory called `build` should appear. Inside of it is the binary.\
You can see the build commands in the [Makefile](Makefile)

**3.** To run locally (no install):

```sh
cd build
./mk-x.y.z
```

**4.** Install (optional):

```sh
make install
```

> Use `make clean` to cleanup
