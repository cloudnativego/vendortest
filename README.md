# Vendor Test
A simple, mostly empty project to test using `godeps` and `go 1.6` vendoring together. This was created by making a `vendor` directory
in the code root prior to invoking `godeps save ./...`. This puts a `godeps.json` file where we expect, but puts the actual dependencies in the `vendor` directory.
