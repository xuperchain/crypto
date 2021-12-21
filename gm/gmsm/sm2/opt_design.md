# SM2优化设计文档
本文档描述了针对SM2算法的优化设计方案。

_第四届中国软件开源创新大赛-开源项目创新赛命题组：**s4plus云惊队**_


## <a name='p'></a>瓶颈分析
通过SM2自带的benchmark测试程序，我们测得该算法在如下测试环境下的时间开销，约为3900000ns/op。
### 测试环境
```
go version: go version go1.16.3 linux/amd64
cpu: cpu: Intel(R) Xeon(R) Gold 6130 CPU @ 2.10GHz
os: Linux user-SYS-2049U-TR4 4.15.0-151-generic #157~16.04.1-Ubuntu SMP Wed Jul 14 11:22:28 UTC 2021 x86_64 x86_64 x86_64 GNU/Linux
```

### 测试命令及结果
```bash
(dby_env) dby@user-SYS-2049U-TR4:~/crypto-opt/origin-crypto/gm/gmsm/sm2$ go test -run=Benchmark -bench=.
goos: linux
goarch: amd64
pkg: github.com/xuperchain/crypto/gm/gmsm/sm2
cpu: Intel(R) Xeon(R) Gold 6130 CPU @ 2.10GHz
BenchmarkSM2-64              292           3846587 ns/op           82726 B/op       1748 allocs/op
PASS
ok      github.com/xuperchain/crypto/gm/gmsm/sm2        1.556s
```
通过go的性能分析工具pprof，可以看到该算法的热点函数如下：
```bash
(pprof) top
Showing nodes accounting for 1320ms, 84.62% of 1560ms total
Showing top 10 nodes out of 122
      flat  flat%   sum%        cum   cum%
     910ms 58.33% 58.33%      970ms 62.18%  github.com/xuperchain/crypto/gm/gmsm/sm2.sm2P256ReduceDegree
     140ms  8.97% 67.31%      800ms 51.28%  github.com/xuperchain/crypto/gm/gmsm/sm2.sm2P256Mul
      50ms  3.21% 70.51%       50ms  3.21%  github.com/xuperchain/crypto/gm/gmsm/sm2.sm2P256ReduceCarry (inline)
      40ms  2.56% 73.08%       40ms  2.56%  github.com/xuperchain/crypto/gm/gmsm/sm2.sm2P256CopyConditional
      40ms  2.56% 75.64%      350ms 22.44%  github.com/xuperchain/crypto/gm/gmsm/sm2.sm2P256Square
      40ms  2.56% 78.21%       50ms  3.21%  github.com/xuperchain/crypto/gm/gmsm/sm2.sm2P256Sub
      30ms  1.92% 80.13%       30ms  1.92%  github.com/xuperchain/crypto/gm/gmsm/sm2.nonZeroToAllOnes (inline)
      30ms  1.92% 82.05%       30ms  1.92%  runtime.memclrNoHeapPointers
      20ms  1.28% 83.33%       20ms  1.28%  github.com/xuperchain/crypto/gm/gmsm/sm2.sm2P256SelectAffinePoint
      20ms  1.28% 84.62%       20ms  1.28%  github.com/xuperchain/crypto/gm/gmsm/sm2.sm2P256SelectJacobianPoint
```
从中可知，时间占比最高的函数是`sm2P256ReduceDegree`、`sm2P256Mul`、`sm2P256ReduceCarry`、`sm2P256Square`等函数，这些函数都是椭圆曲线中进行大数计算相关的函数，所以我们接下来的优化将主要针对椭圆曲线中的计算操作。

## 优化方法
### 1. <a name='1'></a>使用64位整数进行计算
**思路**：在原版本的SM2算法中，使用一个长度为9的`uint32`类型数组来存储大数，并且基于该存储格式进行计算。如果使用64位整数类型`uint64`的数组来存储大数，那么所有的计算都可以被缩减为原来的二分之一。

**问题**：使用64位整数类型可以正常且更高效地实现`sm2P256Add`和`sm2P256Sub`函数，但是却难以直接存储乘法操作`sm2P256Square`和`sm2P256Mul`的结果。因为两个64位的整数乘法将得到一个128位的整数，但是现代计算机不支持128位数据的运算（即使是向量寄存器，它的最大计算单位也是64位），所以此时就需要额外的操作来处理乘法的结果，而这所造成的开销就远高于采用64位所能减少的计算量。综上所述，直接全面地采用64位整数类型并不能带来收益。

**实现**：虽然受限于乘法，我们无法全盘采用64位整数类型来进行计算，但是对于非乘法的运算采用64位整数类型仍然能得到优化。所以我们使用64位整数类型重新实现了`sm2P256ReduceDegree`函数（对于`sm2P256Add`, `sm2P256Sub`函数，因为乘法的限制，对于大数的存储我们仍然需要使用32位整数，如果要将其改写为64位运算，则需要额外的类型转换，而类型转换的开销则大于采用64位运算所带来的收益，所以我们没有修改这两个函数）。

**收益**：通过这项修改，得到了30%左右的性能提升。

### 2. 循环展开
**思路**：Go的编译器本身没有提供循环展开相关的优化（甚至没有任何循环优化），而在计算代码中有着大量的常数次的循环，所以可以通过手动的循环展开来优化程序的性能。

**问题**：当循环内部的代码比较少而循环次数较多的时候，因为增加代码长度所带来的更多的指令cache的miss所造成的开销有可能会超过循环展开所带来的收益，所以在不能盲目地对所有的循环都执行循环展开。

**实现**：我们将`sm2P256Add`, `sm2P256Sub`以及`sum2P256ReduceDegree`中的循环进行了展开。

**收益**：循环展开加上在[#1-使用64位整数类型存储大数](#1-使用64位整数进行计算)中提到的优化，可以使得SM2优化达到60%的性能提升。

### 3. AVX二路并行
**【待后续支持ARM再开源】**

**思路**：参考论文[Parallel Implementation of SM2 Elliptic Curve Cryptography on Intel Processors with AVX2](https://link.springer.com/chapter/10.1007/978-3-030-55304-3_11) ，这篇文章提到，在椭圆曲线上定义的点加（PointAdd）操作可以通过AVX实现部分的二路并行，即利用AVX指令同时进行两部分数据的相同操作。受此论文的启发，我们通过AVX指令实现了三个椭圆曲线操作：`sm2P256PointAdd`, `sm2P256PointDouble`以及`sm2P256PointAddMixed`的二路并行版本。

**问题**：即使Go在1.11版本之后支持添加了AVX512的支持，并且也支持了AVX以及AVX2的汇编指令（参考[Go wiki: AVX512](https://zchee.github.io/golang-wiki/AVX512/) )，但是在Go中使用AVX仍然有很多局限性：
1. Go无法支持所有的AVX系列指令，例如VPUNPCKLQDQ指令，Go的汇编器无法识别该指令。
2. 不像C那样可以通过调用[intel intrinsics](https://www.intel.com/content/www/us/en/docs/intrinsics-guide/index.html) 就可以使用SIMD指令，目前要想在Go中使用SIMD指令仍然只能通过写汇编的方式来实现，这对程序的编写造成了很大的障碍。
3. 开源社区中有用来更简便地生成Go汇编代码的开源库[avo](https://github.com/mmcloughlin/avo) ，并且avo也支持SIMD指令的生成。虽然使用它能够快速生成汇编代码，但是这样生成出来的汇编代码的性能往往无法得到保证，很难达到编译器优化出来的汇编代码的性能。
4. 开源社区中有使用CGO调用封装intel intrinsics的实现[go-avx](https://github.com/monochromegane/go-avx) ，然而由于在当前的CGO机制中，Go->C的调用链开销就要占到60ns，显然这种调用AVX指令的方式所带来的损耗是完全无法接受的。

**实现**：我们通过开源项目[c2goasm](https://github.com/minio/c2goasm) ，先在C中通过Intel intrinsics写好SIMD指令并且通过clang编译到汇编，然后使用c2goasm将clang生成的汇编代码翻译到plan9汇编（Go的汇编代码格式），最后在Go中调用生成的plan9汇编，以此来实现在Go中使用AVX指令。
在解决了如何在Go中使用AVX指令的问题之后，我们实现了`sm2P256Mul`和`sm2P256Square`的二路并行版本，并且通过调整算法顺序将`sm2P256PointAdd`, `sm2P256PointDouble`以及`sm2P256PointAddMixed`中所有的乘法和平方操作替换为了二路并行的版本，如下所示：
```Go
func sm2P256PointDouble(x3, y3, z3, x, y, z *sm2P256FieldElement) {
	var x2, lambda, lambda2, z4_mul_a, y2, y4_mul_8, t, s sm2P256FieldElement

	sm2P256Square2Way(&x2, x, &z4_mul_a, z)
	sm2P256Square2Way(&y2, y, &z4_mul_a, &z4_mul_a)
	sm2P256Mul2Way(&z4_mul_a, &z4_mul_a, &sm2P256.a, &s, x, &y2)

	sm2P256Add(&lambda, &x2, &x2)
	sm2P256Add(&lambda, &lambda, &x2)
	sm2P256Add(&lambda, &lambda, &z4_mul_a) // lambda = (3 * x2 + a * z4)

	sm2P256Add(&y4_mul_8, &y2, &y2)
	sm2P256Square2Way(&lambda2, &lambda, &y4_mul_8, &y4_mul_8)
	sm2P256Add(&y4_mul_8, &y4_mul_8, &y4_mul_8)
	sm2P256Add(&s, &s, &s)
	sm2P256Add(&s, &s, &s)       // s = 4x * y2
	sm2P256Sub(x3, &lambda2, &s) // x3 = 9 * x4 - 4 * x * y2
	sm2P256Sub(x3, x3, &s)       // x3 = 9 * x4 - 8 * x * y2

	sm2P256Sub(&t, &s, x3)                    // t = 4 * x * y2 - x3
	sm2P256Mul2Way(&t, &t, &lambda, z3, y, z)

	sm2P256Sub(y3, &t, &y4_mul_8) // 8 * y4 - 3 * x2 * (s - x3)
	sm2P256Add(z3, z3, z3)
}
```

**收益**：通过AVX的二路并行优化，我们将整体性能进一步优化了35%左右，`sm2P256PointAdd`, `sm2P256PointDouble`以及`sm2P256PointAddMixed`三个核心操作也有着不俗的性能提升。
```bash
(dby_env) dby@user-SYS-2049U-TR4:~/crypto-opt/opt-crypto/gm/gmsm/sm2$ go test -run=Ben -bench=.
goos: linux
goarch: amd64
pkg: github.com/xuperchain/crypto/gm/gmsm/sm2
cpu: Intel(R) Xeon(R) Gold 6130 CPU @ 2.10GHz
BenchmarkPointAddMixedNoneAVX-64                 1075126              1048 ns/op               0 B/op          0 allocs/op
BenchmarkPointAddMixedAVX-64                     1390801               773.4 ns/op             0 B/op          0 allocs/op
BenchmarkPointDoubleNoneAVX-64                    771220              1440 ns/op               0 B/op          0 allocs/op
BenchmarkPointDoubleAVX-64                       1564525               693.6 ns/op             0 B/op          0 allocs/op
BenchmarkPointAddAVX-64                         49954794                23.33 ns/op            0 B/op          0 allocs/op
BenchmarkPointAddNoneAVX-64                     53263878                23.34 ns/op            0 B/op          0 allocs/op
```
### 4. 内存优化
**思路**：Go有着一个轻量的运行时系统用来管理Go进程的并发调度、内存分配以及垃圾回收，然而这种轻量化所带来的后果就是它的性能开销比较大。首先，对于Go的垃圾收集器，它不像Java的分代垃圾收集，它采用了mark-sweep的垃圾收集机制，它的每一次GC都将是一个完整的GC过程，这会导致每次触发Go的垃圾收集都会经历一段较长的stop-the-world；其次，Go的内存分配器是一个类似于TCmalloc的内存分配器，并且采用了稀疏内存的方式，这使得它的每次堆内存分配都会造成不小的性能开销。综上所述，如果能够减少Go进程的内存分配，也能比较可观的提升性能。

**实现**：我们分析程序中所有会进行内存分配的位置，尽可能的减少它的内存分配。
1. 尽量减少`big.Int`类型的创建与使用。`big.Int`作为一个Go的math包中提供的大数类型，在原版本的SM2算法中被用来进行各种大数之间的运算，在每次分配一个`big.Int`类型的变量的时候，会同时分配一个数组用来存储该变量的值，这使得每次分配`big.Int`都会带来一次数组的分配，所以我们在代码中尽量减少它的使用。例如在函数`sm2P256PointAdd`中
   ```Go
    func sm2P256PointAdd(x1, y1, z1, x2, y2, z2, x3, y3, z3 *sm2P256FieldElement) {
        var u1, u2, z22, z12, z23, z13, s1, s2, h, h2, r, r2, tm sm2P256FieldElement

        if sm2P256ToBig(z1).Sign() == 0 {
            sm2P256Dup(x3, x2)
            sm2P256Dup(y3, y2)
            sm2P256Dup(z3, z2)
            return
        }

        if sm2P256ToBig(z2).Sign() == 0 {
            sm2P256Dup(x3, x1)
            sm2P256Dup(y3, y1)
            sm2P256Dup(z3, z1)
            return
        }
        ...
    }
   ```
   它仅仅是为了判断z1的符号就将`sm2P256FieldElement`类型转换为`big.Int`类型，这显然是没有必要的，我们通过手动修改这个操作可以节省掉很多对`big.Int`类型的分配。

2. 绕过Go的逃逸分析，参考自论文[Escape from Escape Analysis of Golang](https://ieeexplore.ieee.org/document/9276567) 。Go语言的逃逸分析是比较激进的，在某些情况下，Go的逃逸分析会不必要的将一些栈变量分配到堆上，这会导致不必要的内存分配。我们通过如下的命令对Go文件进行编译可以输出Go的逃逸分析结果。
    ```bash
    go build -gcflags "-m -m "
    ```
    通过它生成的逃逸分析结果，我们人工分析是否有局部变量被不必要地分配到了堆上。随后我们可以通过Go提供的`unsafe.Pointer`来绕过Go的逃逸分析，如：
    ```Go
    addrA1 := &a1[0]
	uptrA1 := uintptr(unsafe.Pointer(addrA1))
    foo((*uint32)(unsafe.Pointer(uptrA1)))
    ```
    以此我们也可以减少一些不必要的内存分配。

**收益**：通过一系列的减少内存分配的操作，我们将每次操作需要的分配次数由1748 allocs/op减少到了280 alloca/op。

## 优化结果
下表展示的是各个优化的优化效果：
| 使用64位整数进行计算 | 循环展开 | 内存优化 | AVX二路并行 |   运行时间    | 加速比 |
| :------------------: | :------: | :------: | :---------: | :-----------: | :----: |
|          \           |    \     |    \     |      \      | 3846587 op/ns |   1    |
|          √           |          |          |             | 2220680 ns/op | 1.732  |
|          √           |    √     |          |             | 1865892 ns/op | 2.062  |
|          √           |    √     |    √     |             | 1251619 ns/op | 3.073  |
|          √           |    √     |    √     |      √      | 816125 ns/op  | 4.713  |


通过实现上述提到的几种优化方法，优化后的SM2算法在我们的测试环境下每次操作的时间降低到了800000ns左右：
```bash
(dby_env) dby@user-SYS-2049U-TR4:~/crypto-opt/opt-crypto/gm/gmsm/sm2$ go test -run=Benchmark -bench=BenchmarkSM2$
goos: linux
goarch: amd64
pkg: github.com/xuperchain/crypto/gm/gmsm/sm2
cpu: Intel(R) Xeon(R) Gold 6130 CPU @ 2.10GHz
BenchmarkSM2-64             1245            816125 ns/op           15222 B/op        280 allocs/op
PASS
ok      github.com/xuperchain/crypto/gm/gmsm/sm2        1.130s
```
相较于[瓶颈分析](#瓶颈分析)中显示的优化前的SM2算法每次操作3900000ns左右的时间，优化后的算法有着接近5倍的性能提升。
