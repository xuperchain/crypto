package sm2

func _sm2P256Mul2Way1(tmp *uint64, a, b *uint32, tmp2 *uint64, a2, b2 *uint32)

func _sm2P256Mul2Way2(tmp *uint64, a, b *uint32, tmp2 *uint64, a2, b2 *uint32)

func _sm2P256Mul2Way3(tmp *uint64, a, b *uint32, tmp2 *uint64, a2, b2 *uint32)

func _sm2P256Mul4Way1(tmp *uint64, a, b *uint32, tmp2 *uint64, a2, b2 *uint32, tmp3 *uint64, a3, b3 *uint32, tmp4 *uint64, a4, b4 *uint32)

func _sm2P256Square2Way(tmp *uint64, a *uint32, tmp2 *uint64, a2 *uint32)

func _set_i64(a, b uint64)

func _store_i32(a, b uint64)

func _store_i64(a, b, c, d uint64)

func _store_i64_256(a, b, c, d, e, f uint64)

func _reduceDegree_2way(tmp64 *uint64, x64 uint64, tmp642 *uint64, x642 uint64)

func _reduceDegree_2wayNew(tmp64, tmp642 *uint64)

func _sm2P256DivideByR_2way(a, a2 *uint32, tmp, tmp2 *uint64) uint64

func _sm2P256FromLargeElement_2Way(a, b, a2, b2 *uint64)

func _sm2ReduceDegree_2way(a, a2 *uint32, b, b2, tmp, tmp2 *uint64) uint64
