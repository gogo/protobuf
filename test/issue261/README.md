# The Bug

In some cases with gogoslick, time is not being imported.

If the import is generated for one case but not for the others, having the
implementation in the same package would masquerade the bug. Therefore, I've
added one package for each of them.

This bug is not manifested if tests are generated, therefore these packages are
not gentested (but a `go test ./...` would fail given there is an import
missing).
