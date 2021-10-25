package sm2

import "unsafe"

func _sm2P256Mul2Way2(tmp *uint64, a, b *uint32, tmp2 *uint64, a2, b2 *uint32)

func _sm2P256Mul2Way1(tmp *uint64, a, b *uint32, tmp2 *uint64, a2, b2 *uint32)

func _sm2P256Mul4Way1(tmp *uint64, a, b *uint32, tmp2 *uint64, a2, b2 *uint32, tmp3 *uint64, a3, b3 *uint32, tmp4 *uint64, a4, b4 *uint32)

// func _set_i64(a, b uint64)

// func _set_i64_256(a, b, c, d uint64)

// func _store_i64(a, b, c, d uint64)

// func _store_i64_256(a, b, c, d, e, f uint64)

func sm2P256Mul2Way(c, a1, b1, c2, a2, b2 *sm2P256FieldElement) {
	var tmp1, tmp2 sm2P256LargeFieldElement

	addr_a1 := &a1[0]
	addrA1 := uintptr(unsafe.Pointer(addr_a1))
	addr_b1 := &b1[0]
	addrB1 := uintptr(unsafe.Pointer(addr_b1))
	addr_a2 := &a2[0]
	addrA2 := uintptr(unsafe.Pointer(addr_a2))
	addr_b2 := &b2[0]
	addrB2 := uintptr(unsafe.Pointer(addr_b2))
	addr_tmp1 := &tmp1[0]
	addrTmp1 := uintptr(unsafe.Pointer(addr_tmp1))
	addr_tmp2 := &tmp2[0]
	addrTmp2 := uintptr(unsafe.Pointer(addr_tmp2))
	_sm2P256Mul2Way1((*uint64)(unsafe.Pointer(addrTmp1)), (*uint32)(unsafe.Pointer(addrA1)),
		(*uint32)(unsafe.Pointer(addrB1)), (*uint64)(unsafe.Pointer(addrTmp2)),
		(*uint32)(unsafe.Pointer(addrA2)), (*uint32)(unsafe.Pointer(addrB2)))

	_sm2P256Mul2Way2((*uint64)(unsafe.Pointer(addrTmp1)), (*uint32)(unsafe.Pointer(addrA1)),
		(*uint32)(unsafe.Pointer(addrB1)), (*uint64)(unsafe.Pointer(addrTmp2)),
		(*uint32)(unsafe.Pointer(addrA2)), (*uint32)(unsafe.Pointer(addrB2)))

	tmp1[8] = uint64(a1[1]) * uint64(b1[7])
	tmp1[8] += uint64(a1[3]) * uint64(b1[5])
	tmp1[8] += uint64(a1[5]) * uint64(b1[3])
	tmp1[8] += uint64(a1[7]) * uint64(b1[1])
	tmp1[8] <<= 1
	tmp1[8] += uint64(a1[0]) * uint64(b1[8])
	tmp1[8] += uint64(a1[2]) * uint64(b1[6])
	tmp1[8] += uint64(a1[4]) * uint64(b1[4])
	tmp1[8] += uint64(a1[6]) * uint64(b1[2])
	tmp1[8] += uint64(a1[8]) * uint64(b1[0])

	tmp2[8] = uint64(a2[1]) * uint64(b2[7])
	tmp2[8] += uint64(a2[3]) * uint64(b2[5])
	tmp2[8] += uint64(a2[5]) * uint64(b2[3])
	tmp2[8] += uint64(a2[7]) * uint64(b2[1])
	tmp2[8] <<= 1
	tmp2[8] += uint64(a2[0]) * uint64(b2[8])
	tmp2[8] += uint64(a2[2]) * uint64(b2[6])
	tmp2[8] += uint64(a2[4]) * uint64(b2[4])
	tmp2[8] += uint64(a2[6]) * uint64(b2[2])
	tmp2[8] += uint64(a2[8]) * uint64(b2[0])

	addr1 := &tmp1
	addrTMP1 := uintptr(unsafe.Pointer(addr1))

	addr2 := &tmp2
	addrTMP2 := uintptr(unsafe.Pointer(addr2))
	sm2P256ReduceDegree(c, (*sm2P256LargeFieldElement)((unsafe.Pointer)(addrTMP1)))
	sm2P256ReduceDegree(c2, (*sm2P256LargeFieldElement)((unsafe.Pointer)(addrTMP2)))
}

// func sm2P256Mul4Way(c, a1, b1, c2, a2, b2 *sm2P256FieldElement) {
// 	var tmp1, tmp2 sm2P256LargeFieldElement
// 	// tmp1[8] = uint64(a1[1]) * uint64(b1[7])
// 	// tmp1[8] += uint64(a1[3]) * uint64(b1[5])
// 	// tmp1[8] += uint64(a1[5]) * uint64(b1[3])
// 	// tmp1[8] += uint64(a1[7]) * uint64(b1[1])
// 	// tmp1[8] <<= 1
// 	// tmp1[8] += uint64(a1[0]) * uint64(b1[8])
// 	// tmp1[8] += uint64(a1[2]) * uint64(b1[6])
// 	// tmp1[8] += uint64(a1[4]) * uint64(b1[4])
// 	// tmp1[8] += uint64(a1[6]) * uint64(b1[2])
// 	// tmp1[8] += uint64(a1[8]) * uint64(b1[0])

// 	// tmp2[8] = uint64(a2[1]) * uint64(b2[7])
// 	// tmp2[8] += uint64(a2[3]) * uint64(b2[5])
// 	// tmp2[8] += uint64(a2[5]) * uint64(b2[3])
// 	// tmp2[8] += uint64(a2[7]) * uint64(b2[1])
// 	// tmp2[8] <<= 1
// 	// tmp2[8] += uint64(a2[0]) * uint64(b2[8])
// 	// tmp2[8] += uint64(a2[2]) * uint64(b2[6])
// 	// tmp2[8] += uint64(a2[4]) * uint64(b2[4])
// 	// tmp2[8] += uint64(a2[6]) * uint64(b2[2])
// 	// tmp2[8] += uint64(a2[8]) * uint64(b2[0])

// 	addr_a1 := &a1[0]
// 	addrA1 := uintptr(unsafe.Pointer(addr_a1))
// 	addr_b1 := &b1[0]
// 	addrB1 := uintptr(unsafe.Pointer(addr_b1))
// 	addr_a2 := &a2[0]
// 	addrA2 := uintptr(unsafe.Pointer(addr_a2))
// 	addr_b2 := &b2[0]
// 	addrB2 := uintptr(unsafe.Pointer(addr_b2))
// 	addr_tmp1 := &tmp1[0]
// 	addrTmp1 := uintptr(unsafe.Pointer(addr_tmp1))
// 	addr_tmp2 := &tmp2[0]
// 	addrTmp2 := uintptr(unsafe.Pointer(addr_tmp2))
// 	_sm2P256Mul2Way1((*uint64)(unsafe.Pointer(addrTmp1)), (*uint32)(unsafe.Pointer(addrA1)),
// 		(*uint32)(unsafe.Pointer(addrB1)), (*uint64)(unsafe.Pointer(addrTmp2)),
// 		(*uint32)(unsafe.Pointer(addrA2)), (*uint32)(unsafe.Pointer(addrB2)))

// 	_sm2P256Mul2Way2((*uint64)(unsafe.Pointer(addrTmp1)), (*uint32)(unsafe.Pointer(addrA1)),
// 		(*uint32)(unsafe.Pointer(addrB1)), (*uint64)(unsafe.Pointer(addrTmp2)),
// 		(*uint32)(unsafe.Pointer(addrA2)), (*uint32)(unsafe.Pointer(addrB2)))

// 	addr1 := &tmp1
// 	addrTMP1 := uintptr(unsafe.Pointer(addr1))

// 	addr2 := &tmp2
// 	addrTMP2 := uintptr(unsafe.Pointer(addr2))
// 	sm2P256ReduceDegree(c, (*sm2P256LargeFieldElement)(unsafe.Pointer(addrTMP1)))
// 	sm2P256ReduceDegree(c2, (*sm2P256LargeFieldElement)(unsafe.Pointer(addrTMP2)))
// }

// type __m512i [8]uint

// func _set_i64(a, b, c, d, e, f, g, h uint64)

// func _sm2P256Mul_AVX(tmp *uint64, a, b *uint32)

// func _sm2P256Mul_AVX2(tmp *uint64, a, b *uint64)

// func _sm2P256Mul_AVX3(tmp *uint64, a, b *uint64)

// func Sm2P256Mul_AVX(c, a, b *sm2P256FieldElement) sm2P256LargeFieldElement {
// 	var a1, b1 [13]uint64
// 	for i := 0; i < 9; i++ {
// 		a1[i] = (uint64)(a[i])
// 		b1[i] = (uint64)(b[8-i])
// 	}
// 	addr_a1 := &a1[0]
// 	addrA1 := uintptr(unsafe.Pointer(addr_a1))
// 	addr_b1 := &b1[0]
// 	addrB1 := uintptr(unsafe.Pointer(addr_b1))
// 	var tmp sm2P256LargeFieldElement

// 	addr_tmp := &tmp[0]
// 	addrTmp := uintptr(unsafe.Pointer(addr_tmp))
// 	_sm2P256Mul_AVX2((*uint64)(unsafe.Pointer(addrTmp)), (*uint64)(unsafe.Pointer(addrA1)), (*uint64)(unsafe.Pointer(addrB1)))

// 	return tmp
// }

// func _test1(a, b, c *uint64)

// func _test2(a, b, a2, b2 *uint32, c, c2 *uint64)

// func _test3(a, b, a2, b2 *uint32, c, c2 *uint64)

// func _test4(a, b, a2, b2 *uint32, c, c2 *uint64)
