# sha3-reuse
reusable SHA3 hashing in Go

===why?===
the sha3 implementation in golang.org/x/crypto/sha3 doesn't permit any
writes after the first read. a true sponge function should be capable
of absorbing more data after a squeeze.

additionally, it may be helpful in some scenarios to have a sponge hash
whose state is modified every time it is read, so that a particular
sequence of reads and writes are necessary to obtain an identical
sequence of hashes.

this package provides reusable versions of the sha3 hash. it provides
one version whose state is modified when reading, and another whose state
is not changed by a read.
