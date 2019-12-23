### Steps to reproduce

1. Init aerospike server

2. Put some data

```// aql> TRUNCATE test.foo;
// OK

// aql> INSERT INTO test.foo (PK, foo) VALUES ('abc', 1)
// OK, 1 record affected.

// aql> INSERT INTO test.foo (PK, foo) VALUES ('bbb', 1)
// OK, 1 record affected.

// aql> INSERT INTO test.foo (PK, foo) VALUES ('ccc', 3)
// OK, 1 record affected.

// aql> SELECT * FROM test.foo
// +-----+
// | foo |
// +-----+
// | 3   |
// | 1   |
// | 1   |
// +-----+
// 3 rows in set (0.118 secs)```

3. Run main.go

4. I've got as a result

```
runtime: goroutine stack exceeds 1000000000-byte limit
fatal error: stack overflow

runtime stack:
runtime.throw(0x13835b1, 0xe)
	/Users/qw4n7y/.gvm/gos/go1.12/src/runtime/panic.go:617 +0x72
runtime.newstack()
	/Users/qw4n7y/.gvm/gos/go1.12/src/runtime/stack.go:1041 +0x6f0
runtime.morestack()
	/Users/qw4n7y/.gvm/gos/go1.12/src/runtime/asm_amd64.s:429 +0x8f
  ```

