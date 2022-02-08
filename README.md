# lazyproto

**First version:**

```bash
GOROOT=/usr/local/Cellar/go/1.17.6/libexec #gosetup
GOPATH=/Users/lazy/.go #gosetup
/usr/local/Cellar/go/1.17.6/libexec/bin/go test -c -o /private/var/folders/5c/5p_3jmvd6qx0rsvmb9j7c_s00000gn/T/GoLand/___gobench_github_com_bogdandrutu_lazyproto_benchmarks.test github.com/bogdandrutu/lazyproto/benchmarks #gosetup
/private/var/folders/5c/5p_3jmvd6qx0rsvmb9j7c_s00000gn/T/GoLand/___gobench_github_com_bogdandrutu_lazyproto_benchmarks.test -test.v -test.paniconexit0 -test.bench . -test.run ^$
goos: darwin
goarch: amd64
pkg: github.com/bogdandrutu/lazyproto/benchmarks
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkMarshalUnmarshal_LazyProto
BenchmarkMarshalUnmarshal_LazyProto-16    	  257985	      4835 ns/op	    3112 B/op	      89 allocs/op
BenchmarkMarshalUnmarshal_Google
BenchmarkMarshalUnmarshal_Google-16       	   97891	     12223 ns/op	    3928 B/op	      98 allocs/op
BenchmarkMarshalUnmarshal_GoGo
BenchmarkMarshalUnmarshal_GoGo-16         	  357652	      3312 ns/op	    2768 B/op	      57 allocs/op
PASS

Process finished with the exit code 0
```

**Remove "unknown fields", similar with current gogo, not sure is a good idea though:**

```bash
GOROOT=/usr/local/Cellar/go/1.17.6/libexec #gosetup
GOPATH=/Users/lazy/.go #gosetup
/usr/local/Cellar/go/1.17.6/libexec/bin/go test -c -o /private/var/folders/5c/5p_3jmvd6qx0rsvmb9j7c_s00000gn/T/GoLand/___gobench_github_com_bogdandrutu_lazyproto_benchmarks.test github.com/bogdandrutu/lazyproto/benchmarks #gosetup
/private/var/folders/5c/5p_3jmvd6qx0rsvmb9j7c_s00000gn/T/GoLand/___gobench_github_com_bogdandrutu_lazyproto_benchmarks.test -test.v -test.paniconexit0 -test.bench . -test.run ^$
goos: darwin
goarch: amd64
pkg: github.com/bogdandrutu/lazyproto/benchmarks
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkMarshalUnmarshal_LazyProto
BenchmarkMarshalUnmarshal_LazyProto-16    	  272515	      3908 ns/op	    2152 B/op	      89 allocs/op
BenchmarkMarshalUnmarshal_Google
BenchmarkMarshalUnmarshal_Google-16       	   97960	     11974 ns/op	    3928 B/op	      98 allocs/op
BenchmarkMarshalUnmarshal_GoGo
BenchmarkMarshalUnmarshal_GoGo-16         	  362098	      3301 ns/op	    2768 B/op	      57 allocs/op
PASS

Process finished with the exit code 0
```

**Make messages members non-nullable:**

```bash
GOROOT=/usr/local/Cellar/go/1.17.6/libexec #gosetup
GOPATH=/Users/lazy/.go #gosetup
/usr/local/Cellar/go/1.17.6/libexec/bin/go test -c -o /private/var/folders/5c/5p_3jmvd6qx0rsvmb9j7c_s00000gn/T/GoLand/___gobench_github_com_bogdandrutu_lazyproto_benchmarks.test github.com/bogdandrutu/lazyproto/benchmarks #gosetup
/private/var/folders/5c/5p_3jmvd6qx0rsvmb9j7c_s00000gn/T/GoLand/___gobench_github_com_bogdandrutu_lazyproto_benchmarks.test -test.v -test.paniconexit0 -test.bench . -test.run ^$
goos: darwin
goarch: amd64
pkg: github.com/bogdandrutu/lazyproto/benchmarks
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkMarshalUnmarshal_LazyProto
BenchmarkMarshalUnmarshal_LazyProto-16    	  352713	      3394 ns/op	    3216 B/op	      49 allocs/op
BenchmarkMarshalUnmarshal_Google
BenchmarkMarshalUnmarshal_Google-16       	   97819	     12084 ns/op	    3928 B/op	      98 allocs/op
BenchmarkMarshalUnmarshal_GoGo
BenchmarkMarshalUnmarshal_GoGo-16         	  349899	      3335 ns/op	    2768 B/op	      57 allocs/op
PASS

Process finished with the exit code 0
```

**Use variant from @tigrannajaryan instead of `interface{}`:** 

```bash
GOROOT=/usr/local/Cellar/go/1.17.6/libexec #gosetup
GOPATH=/Users/lazy/.go #gosetup
/usr/local/Cellar/go/1.17.6/libexec/bin/go test -c -o /private/var/folders/5c/5p_3jmvd6qx0rsvmb9j7c_s00000gn/T/GoLand/___gobench_github_com_bogdandrutu_lazyproto_benchmarks.test github.com/bogdandrutu/lazyproto/benchmarks #gosetup
/private/var/folders/5c/5p_3jmvd6qx0rsvmb9j7c_s00000gn/T/GoLand/___gobench_github_com_bogdandrutu_lazyproto_benchmarks.test -test.v -test.paniconexit0 -test.bench . -test.run ^$
goos: darwin
goarch: amd64
pkg: github.com/bogdandrutu/lazyproto/benchmarks
cpu: Intel(R) Core(TM) i9-9880H CPU @ 2.30GHz
BenchmarkMarshalUnmarshal_LazyProto
BenchmarkMarshalUnmarshal_LazyProto-16    	  347965	      3058 ns/op	    2992 B/op	      37 allocs/op
BenchmarkMarshalUnmarshal_Google
BenchmarkMarshalUnmarshal_Google-16       	   93165	     11942 ns/op	    3928 B/op	      98 allocs/op
BenchmarkMarshalUnmarshal_GoGo
BenchmarkMarshalUnmarshal_GoGo-16         	  342450	      3527 ns/op	    2768 B/op	      57 allocs/op
PASS

Process finished with the exit code 0
```