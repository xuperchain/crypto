#include <time.h>
#include <stdio.h>
#include <immintrin.h>
#include <sys/time.h>    
#include <unistd.h> 

typedef unsigned int sm2element;

typedef unsigned long int uint64;

// void sm2P256Mul_AVX(uint64 *tmp, sm2element *a, sm2element *b) {
// 	__m512i z0 = {a[0], a[0], a[1], a[0], a[1], a[2], a[0], a[1]};
// 	__m512i z1 = {b[0], b[1], b[0], b[2], (uint64)(b[1]) << 1, b[0], b[3], b[2]};
// 	z1 = z1 * z0;
// 	tmp[0] = z1[0];
// 	tmp[1] = z1[1] + z1[2];
// 	tmp[2] = z1[3] + z1[4] + z1[5];
// 	tmp[3] = z1[6] + z1[7];

// 	__m512i z2 = {a[2], a[3], a[1], a[3], a[0], a[2], a[4], a[0]};
// 	__m512i z3 = {b[1], b[0], b[3], b[1], b[4], b[2], b[0], b[5]};
// 	z3 = z3 * z2;
// 	tmp[3] += z3[0] + z3[1];
// 	tmp[4] = z3[2] + z3[3];
// 	tmp[4] <<= 1;
// 	tmp[4] += z3[4] + z3[5] + z3[6];
// 	tmp[5] = z3[7];

// 	__m512i z4 = {a[1], a[2], a[3], a[4], a[5], a[1], a[3], a[5]};
// 	__m512i z5 = {b[4], b[3], b[2], b[1], b[0], b[5], b[3], b[1]};
// 	z5 = z4 * z5;
// 	tmp[5] += z5[0] + z5[1] + z5[2] + z5[3] + z5[4];
// 	tmp[6] = z5[5] + z5[6] + z5[7];
// 	tmp[6] <<= 1;

// 	__m512i z6 = {a[0], a[2], a[4], a[6], a[0], a[1], a[2], a[3]};
// 	__m512i z7 = {b[6], b[4], b[2], b[0], b[7], b[6], b[5], b[4]};
// 	z7 = z6 * z7;
// 	tmp[6] += z7[0] + z7[1] + z7[2] + z7[3];
// 	tmp[7] = z7[4] + z7[5] + z7[6] + z7[7];

// 	__m512i z8 = {a[4], a[5], a[6], a[7], a[1], a[3], a[5], a[7]};
// 	__m512i z9 = {b[3], b[2], b[1], b[0], b[7], b[5], b[3], b[1]};
// 	z9 = z8 * z9;
// 	tmp[7] += z9[0] + z9[1] + z9[2] + z9[3];
// 	tmp[8] = z9[4] + z9[5] + z9[6] + z9[7];
// 	tmp[8] <<= 1;

// 	__m512i z10 = {a[0], a[2], a[4], a[6], a[8], a[1], a[2], a[3]};
// 	__m512i z11 = {b[8], b[6], b[4], b[2], b[0], b[8], b[7], b[6]};
// 	z11 = z10 * z11;
// 	tmp[8] += z11[0] + z11[1] + z11[2] + z11[3] + z11[4];
// 	tmp[9] = z11[5] + z11[6] + z11[7];

// 	__m512i z12 = {a[4], a[5], a[6], a[7], a[8], a[3], a[5], a[7]};
// 	__m512i z13 = {b[5], b[4], b[3], b[2], b[1], b[7], b[5], b[3]};
// 	z13 = z12 * z13;
// 	tmp[9] += z13[0] + z13[1] + z13[2] + z13[3] + z13[4];
// 	tmp[10] = z13[5] + z13[6] + z13[7];
// 	tmp[10] <<= 1;

// 	__m512i z14 = {a[2], a[4], a[6], a[8], a[3], a[4], a[5], a[6]};
// 	__m512i z15 = {b[8], b[6], b[4], b[2], b[8], b[7], b[6], b[5]};
// 	z15 = z14 * z15;
// 	tmp[10] += z15[0] + z15[1] + z15[2] + z15[3];
// 	tmp[11] = z15[4] + z15[5] + z15[6] + z15[7];

// 	__m512i z16 = {a[7], a[8], a[5], a[7], a[4], a[6], a[8], a[5]};
// 	__m512i z17 = {b[4], b[3], b[7], b[5], b[8], b[6], b[4], b[8]};
// 	z17 = z16 * z17;
// 	tmp[11] += z17[0] + z17[1];
// 	tmp[12] = z17[2] + z17[3];
// 	tmp[12] <<= 1;
// 	tmp[12] += z17[4] + z17[5] + z17[6];
// 	tmp[13] = z17[7];

// 	__m512i z18 = {a[6], a[7], a[8], a[6], a[7], a[8], a[7], a[8]};
// 	__m512i z19 = {b[7], b[6], b[5], b[8], (uint64)(b[7]) << 1, b[6], b[8], b[7]};
// 	z19 = z18 * z19;
// 	tmp[13] += z19[0] + z19[1] + z19[2];
// 	tmp[14] = z19[3] + z19[4] + z19[5];
// 	tmp[15] = z19[6] + z19[7];

// 	tmp[16] = (uint64)(a[8]) * (uint64)(b[8]);
// }

__attribute__((always_inline)) __m128i set_i64(uint64 a, uint64 b) {
	__m128i x0 = {a, b};
	return x0;
}

__attribute__((always_inline)) void store_i64(uint64 *tmp_, uint64 *tmp2_, __m128i x) {
	*tmp_ = x[0];
	*tmp2_ = x[1];
}

__attribute__((always_inline)) void store_i32(sm2element *tmp_, sm2element *tmp2_, __m128i x) {
	*tmp_ = (sm2element)x[0];
	*tmp2_ = (sm2element)x[1];
}

#define SETa(i) set_i64((uint64)a_[i], (uint64)a2_[i])
#define SETb(i) set_i64((uint64)b_[i], (uint64)b2_[i])

void sm2P256Mul2Way1(uint64 *tmp_, sm2element *a_, sm2element *b_, uint64 *tmp2_, sm2element *a2_, sm2element *b2_) {
	__m128i a0 = SETa(0);
	__m128i b0 = SETb(0);
	__m128i tmp;
	tmp = a0 * b0;
	store_i64(tmp_, tmp2_, tmp);

	__m128i a1 = SETa(1);
	__m128i b1 = SETb(1);
	tmp = a0 * b1;
	tmp += a1 * b0;
	store_i64(tmp_ + 1, tmp2_ + 1, tmp);

	__m128i a2 = SETa(2);
	__m128i b2 = SETb(2);
	tmp = (a0) * (b2);
	tmp += (a1) * ((b1) << 1);
	tmp += (a2) * (b0);
	store_i64(tmp_ + 2, tmp2_ + 2, tmp);

	__m128i a3 = SETa(3);
	__m128i b3 = SETb(3);
	tmp = (a0) * (b3);
	tmp += (a1) * (b2);
	tmp += (a2) * (b1);
	tmp += (a3) * (b0);
	store_i64(tmp_ + 3, tmp2_ + 3, tmp);

	__m128i a4 = SETa(4);
	__m128i b4 = SETb(4);
	tmp = (a1) * (b3);
	tmp += (a3) * (b1);
	tmp <<= 1;
	tmp += (a0) * (b4);
	tmp += (a2) * (b2);
	tmp += (a4) * (b0);
	store_i64(tmp_ + 4, tmp2_ + 4, tmp);
	
	__m128i a5 = SETa(5);
	__m128i b5 = SETb(5);
	tmp = (a0) * (b5);
	tmp += (a1) * (b4);
	tmp += (a2) * (b3);
	tmp += (a3) * (b2);
	tmp += (a4) * (b1);
	tmp += (a5) * (b0);
	store_i64(tmp_ + 5, tmp2_ + 5, tmp);

	__m128i a6 = SETa(6);
	__m128i b6 = SETb(6);
	tmp = (a1) * (b5);
	tmp += (a3) * (b3);
	tmp += (a5) * (b1);
	tmp <<= 1;
	tmp += (a0) * (b6);
	tmp += (a2) * (b4);
	tmp += (a4) * (b2);
	tmp += (a6) * (b0);
	store_i64(tmp_ + 6, tmp2_ + 6, tmp);

	tmp = (a1) * (b6);
	tmp += (a2) * (b5);
	tmp += (a4) * (b3);
	tmp += (a5) * (b2);
	tmp += (a6) * (b1);
	store_i64(tmp_ + 7, tmp2_ + 7, tmp);
	tmp_[7] += (uint64)a_[0] * (uint64)b_[7] + (uint64)a_[3] * (uint64)b_[4]  + (uint64)a_[7] * (uint64)b_[0];
	tmp2_[7] += (uint64)a2_[0] * (uint64)b2_[7] + (uint64)a2_[3] * (uint64)b2_[4] + (uint64)a2_[7] * (uint64)b2_[0];
}

void sm2P256Mul2Way2(uint64 *tmp_, sm2element *a_, sm2element *b_, uint64 *tmp2_, sm2element *a2_, sm2element *b2_) {
	__m128i tmp;

	tmp_[9] = 0;
	tmp2_[9] = 0;
	__m128i a8 = SETa(8);
	__m128i b8 = SETb(8);
	tmp = (a8) * (b8);
	store_i64(tmp_ + 16, tmp2_ + 16, tmp);

	__m128i a7 = SETa(7);
	__m128i b7 = SETb(7);
	tmp = (a7) * (b8);
	tmp += (a8) * (b7);
	store_i64(tmp_ + 15, tmp2_ + 15, tmp);

	__m128i a6 = SETa(6);
	__m128i b6 = SETb(6);
	tmp = (a6) * (b8);
	tmp += (a7) * ((b7) << 1);
	tmp += (a8) * (b6);
	store_i64(tmp_ + 14, tmp2_ + 14, tmp);

	__m128i a5 = SETa(5);
	__m128i b5 = SETb(5);
	tmp = (a5) * (b8);
	tmp += (a6) * (b7);
	tmp += (a7) * (b6);
	tmp += (a8) * (b5);
	store_i64(tmp_ + 13, tmp2_ + 13, tmp);	

	__m128i a4 = SETa(4);
	__m128i b4 = SETb(4);
	tmp = (a5) * (b7);
	tmp += (a7) * (b5);
	tmp <<= 1;
	tmp += (a4) * (b8);
	tmp += (a6) * (b6);
	tmp += (a8) * (b4);
	store_i64(tmp_ + 12, tmp2_ + 12, tmp);

	__m128i a3 = SETa(3);
	__m128i b3 = SETb(3);
	tmp = (a3) * (b8);
	tmp += (a4) * (b7);
	tmp += (a5) * (b6);
	tmp += (a6) * (b5);
	tmp += (a7) * (b4);
	tmp += (a8) * (b3);
	store_i64(tmp_ + 11, tmp2_ + 11, tmp);

	__m128i a2 = SETa(2);
	__m128i b2 = SETb(2);
	tmp = (a3) * (b7);
	tmp += (a5) * (b5);
	tmp += (a7) * (b3);
	tmp <<= 1;
	tmp += (a2) * (b8);
	tmp += (a4) * (b6);
	tmp += (a6) * (b4);
	tmp += (a8) * (b2);
	store_i64(tmp_ + 10, tmp2_ + 10, tmp);

	// tmp = (a1) * (b8);
	tmp = (a2) * (b7);
	tmp += (a3) * (b6);
	// tmp += (a4) * (b5);
	tmp += (a5) * (b4);
	tmp += (a6) * (b3);
	tmp += (a7) * (b2);
	// tmp += (b8) * (b1);
	store_i64(tmp_ + 9, tmp2_ + 9, tmp);
	tmp_[9] += (uint64)a_[1] * (uint64)b_[8] + (uint64)a_[4] * (uint64)b_[5] + (uint64)a_[8] * (uint64)b_[1];
	tmp2_[9] += (uint64)a2_[1] * (uint64)b2_[8] + (uint64)a2_[4] * (uint64)b2_[5] + (uint64)a2_[8] * (uint64)b2_[1];
}

void sm2P256Mul2Way3(uint64 *tmp, sm2element *a, sm2element *b, uint64 *tmp2, sm2element *a2, sm2element *b2) {
	__m128i t;
	__m128i a1 = {a[1], a2[1]};
	__m128i b7 = {b[7], b2[7]};
	__m128i a3 = set_i64(a[3], a2[3]);
	__m128i b5 = set_i64(b[5], b2[5]);
	__m128i a5 = set_i64(a[5], a2[5]);
	__m128i b3 = set_i64(b[3], b2[3]);
	__m128i a7 = set_i64(a[7], a2[7]);
	__m128i b1 = set_i64(b[1], b2[1]);
	__m128i a0 = set_i64(a[0], a2[0]);
	__m128i a2_ = set_i64(a[2], a2[2]);
	__m128i b6 = set_i64(b[6], b2[6]);
	__m128i a4 = set_i64(a[4], a2[4]);
	__m128i b4 = set_i64(b[4], b2[4]);
	__m128i a6 = set_i64(a[6], a2[6]);
	__m128i b2_ = set_i64(b[2], b2[2]);
	__m128i b0 = set_i64(b[0], b2[0]);

	t = a0 * b7;
	t += a1 * b6;
	t += a2_ * b5;
	t += a3 * b4;
	t += a4 * b3;
	t += a5 * b2_;
	t += a6 * b1;
	t += a7 * b0;
	store_i64(tmp+7, tmp2+7, t);
	
	t = a1 * b7;
	t += a3 * b5;
	t += a5 * b3;
	t += a7 * b1;
	t <<= 1;
	__m128i b8 = set_i64(b[8], b2[8]);
	t += a0 * b8;
	t += a2_ * b6;
	t += a4 * b4;
	
	t += a6 * b2_;
	__m128i a8 = set_i64(a[8], a2[8]);
	t += a8 * b0;
	store_i64(tmp+8, tmp2+8, t);

	t = a1 * b8;
	t += a2_ * b7;
	t += a3 * b6;
	t += a4 * b5;
	t += a5 * b4;
	t += a6 * b3;
	t += a7 * b2_;
	t += a8 * b1;
	store_i64(tmp+9, tmp2+9, t);
	// tmp[9] += (uint64)a[8] * (uint64)b[1];
	// tmp2[9] += (uint64)a2[8] * (uint64)b2[1];
}

const uint64 bottom28BitsMask = 0xFFFFFFF;
const uint64 bottom29BitsMask = 0x1FFFFFFF;
const uint64 bottom32BitsMask = 0xFFFFFFFF;
const uint64 top32BitsMask = 0xFFFFFFFF00000000;
const uint64 bottom57BitsMask = 0x1FFFFFFFFFFFFFF;
const uint64 twoPower57 = 0x200000000000000;

__attribute__((always_inline)) void reduceDegree_2way(uint64 *tmp64, uint64 x64, uint64 *tmp642, uint64 x642) {
	__m128i x = {x64, x642};
	__m128i b57 = {bottom57BitsMask, bottom57BitsMask};
	__m128i tp57 = {twoPower57, twoPower57};
	__m128i b29 = {bottom29BitsMask, bottom29BitsMask};

	__m128i tmp = {tmp64[1], tmp642[1]};
	tmp += (x << 7) & b57;
	tmp += tp57;
	tmp -= (x << 39) & b57;
	store_i64(tmp64 + 1, tmp642 + 1, tmp);

	tmp = set_i64(tmp64[2], tmp642[2]);
	tmp += x >> 50;
	tmp += b57;
	tmp -= x >> 18;
	store_i64(tmp64 + 2, tmp642 + 2, tmp);

	tmp = set_i64(tmp64[3], tmp642[3]);
	tmp += b57;
	tmp -= (x << 53) & b57;
	store_i64(tmp64 + 3, tmp642 + 3, tmp);

	tmp = set_i64(tmp64[4], tmp642[4]);
	tmp += b57;
	tmp -= (x >> 4) & b57;
	tmp += (x << 28) & b57;
	store_i64(tmp64 + 4, tmp642 + 4, tmp);

	tmp = set_i64(tmp64[5], tmp642[5]);
	tmp -= 1;
	tmp += (x >> 29) & b29;
	store_i64(tmp64 + 5, tmp642 + 5, tmp);
}

__attribute__((always_inline)) void reduceDegree_2wayNew(uint64 *tmp64, uint64 *tmp642) {
	uint64 x64, x642;
	for (int j = 0; j < 5; j++) {
		if (j == 4) {
			x64 = tmp64[j] & bottom29BitsMask;
			tmp64[j] = (tmp64[j] >> 29) << 29;
			
			x642 = tmp642[j] & bottom29BitsMask;
			tmp642[j] = (tmp642[j] >> 29) << 29;
		} else {
			tmp64[j+1] += tmp64[j] >> 57;
			x64 = tmp64[j] & bottom57BitsMask;

			tmp642[j+1] += tmp642[j] >> 57;
			x642 = tmp642[j] & bottom57BitsMask;
		}
		if (x64 > 0 && x642 > 0) {
			reduceDegree_2way(tmp64+j, x64, tmp642+j, x642);
		} else {
			if (x64 > 0) {
				tmp64[j+1] += (x64 << 7) & bottom57BitsMask;
				tmp64[j+1] += twoPower57;
				tmp64[j+1] -= (x64 << 39) & bottom57BitsMask;

				tmp64[j+2] += x64 >> 50;
				tmp64[j+2] += bottom57BitsMask;
				tmp64[j+2] -= x64 >> 18;

				tmp64[j+3] += bottom57BitsMask;
				tmp64[j+3] -= (x64 << 53) & bottom57BitsMask;

				tmp64[j+4] += bottom57BitsMask;
				tmp64[j+4] -= (x64 >> 4) & bottom57BitsMask;
				tmp64[j+4] += (x64 << 28) & bottom57BitsMask;

				tmp64[j+5] -= 1;
				tmp64[j+5] += (x64 >> 29) & bottom29BitsMask;
			}
			if (x642 > 0) {
				tmp642[j+1] += (x642 << 7) & bottom57BitsMask;
				tmp642[j+1] += twoPower57;
				tmp642[j+1] -= (x642 << 39) & bottom57BitsMask;

				tmp642[j+2] += x642 >> 50;
				tmp642[j+2] += bottom57BitsMask;
				tmp642[j+2] -= x642 >> 18;

				tmp642[j+3] += bottom57BitsMask;
				tmp642[j+3] -= (x642 << 53) & bottom57BitsMask;

				tmp642[j+4] += bottom57BitsMask;
				tmp642[j+4] -= (x642 >> 4) & bottom57BitsMask;
				tmp642[j+4] += (x642 << 28) & bottom57BitsMask;

				tmp642[j+5] -= 1;
				tmp642[j+5] += (x642 >> 29) & bottom29BitsMask;
			}
		}
	}

	if (tmp64[9]+1 == 0) {
		tmp64[9] = 0;
		tmp64[8] -= twoPower57;
	}

	if (tmp642[9]+1 == 0) {
		tmp642[9] = 0;
		tmp642[8] -= twoPower57;
	}
}

__attribute__((always_inline)) uint64 sm2P256DivideByR_2way(sm2element *a, sm2element *a2, uint64 *tmp, uint64 *tmp2) {
	uint64 carry1, carry2;
	__m128i carry;
	__m128i b29 = {bottom29BitsMask, bottom29BitsMask};
	__m128i b28 = {bottom28BitsMask, bottom28BitsMask};
	__m128i b32 = {bottom32BitsMask, bottom32BitsMask};
	
	__m128i temp4 = {tmp[4], tmp2[4]};
	__m128i temp5 = {tmp[5], tmp2[5]};
	__m128i t;
	t = (temp4 >> 29) & b32;
	t += (temp5 << 28) & b29;
	carry = t >> 29;
	t &= b29;
	store_i32(a, a2, t);

	t = (temp5 >> 1) & b28;
	t += carry;
	carry = t >> 28;
	t &= b28;
	store_i32(a+1, a2+1, t);

	__m128i temp6 = {tmp[6], tmp2[6]};
	t = (temp5 >> 29) & b32;
	t += carry;
	t += (temp6 << 28) & b29;
	carry = t >> 29;
	t &= b29;
	store_i32(a+2, a2+2, t);

	t = (temp6 >> 1) & b28;
	t += carry;
	carry = t >> 28;
	t &= b28;
	store_i32(a+3, a2+3, t);

	__m128i temp7 = {tmp[7], tmp2[7]};
	t = (temp6 >> 29) & b32;
	t += carry;
	t += (temp7 << 28) & b29;
	carry = t >> 29;
	t &= b29;
	store_i32(a+4, a2+4, t);

	t = (temp7 >> 1) & b28;
	t += carry;
	carry = t >> 28;
	t &= b28;
	store_i32(a+5, a2+5, t);

	__m128i temp8 = {tmp[8], tmp2[8]};
	t = (temp7 >> 29) & b32;
	t += carry;
	t += (temp8 << 28) & b29;
	carry = t >> 29;
	t &= b29;
	store_i32(a+6, a2+6, t);

	t = (temp8 >> 1) & b28;
	t += carry;
	carry = t >> 28;
	t &= b28;
	store_i32(a+7, a2+7, t);

	__m128i temp9 = {tmp[9], tmp2[9]};
	t = (temp8 >> 29) & b32;
	t += carry;
	t += (temp9 << 28) & b29;
	carry = t >> 29;
	t &= b29;
	// store_i64(a+8, a2+8, t);
	store_i32(a+8, a2+8, t);

	// store_i64(&carry1, &carry2, carry);
	carry1 = carry[0];
	carry2 = carry[1];
	return (uint64)((sm2element)carry1) + ((uint64)((sm2element)carry2) << 32);
}

__attribute__((always_inline)) void sm2P256FromLargeElement_2Way(uint64 *a, uint64 *b, uint64 *a2, uint64 *b2) {
	__m128i carry;
	__m128i b57 = {bottom57BitsMask, bottom57BitsMask};

	__m128i t = {b[0], b2[0]};
	__m128i b1 = {b[1], b2[1]};
	t += (b1 << 29) & b57;
	carry = t >> 57;
	t &= b57;
	store_i64(a, a2, t);
	// a[0] = b[0];
	// a[0] += ((b[1] << 29) & bottom57BitsMask);
	// carry = a[0] >> 57;
	// a[0] = a[0] & bottom57BitsMask;

	__m128i b2_ = {b[2], b2[2]};
	__m128i b3 = {b[3], b2[3]};
	t = carry + (b1 >> 28);
	t += b2_;
	t += (b3 << 29) & b57;
	carry = t >> 57;
	t &= b57;
	store_i64(a+1, a2+1, t);
	// a[1] = carry;
	// a[1] += b[1] >> 28;
	// a[1] += b[2];
	// a[1] += (b[3] << 29) & bottom57BitsMask;
	// carry = a[1] >> 57;
	// a[1] = a[1] & bottom57BitsMask;

	__m128i b4 = {b[4], b2[4]};
	__m128i b5 = {b[5], b2[5]};
	t = carry + (b3 >> 28);
	t += b4;
	t += (b5 << 29) & b57;
	carry = t >> 57;
	t &= b57;
	store_i64(a+2, a2+2, t);
	// a[2] = carry;
	// a[2] += b[3] >> 28;
	// a[2] += b[4];
	// a[2] += (b[5] << 29) & bottom57BitsMask;
	// carry = a[2] >> 57;
	// a[2] = a[2] & bottom57BitsMask;

	__m128i b6 = {b[6], b2[6]};
	__m128i b7 = {b[7], b2[7]};
	t = carry + (b5 >> 28);
	t += b6;
	t += (b7 << 29) & b57;
	carry = t >> 57;
	t &= b57;
	store_i64(a+3, a2+3, t);
	// a[3] = carry;
	// a[3] += b[5] >> 28;
	// a[3] += b[6];
	// a[3] += (b[7] << 29) & bottom57BitsMask;
	// carry = a[3] >> 57;
	// a[3] = a[3] & bottom57BitsMask;

	__m128i b8 = {b[8], b2[8]};
	__m128i b9 = {b[9], b2[9]};
	t = carry + (b7 >> 28);
	t += b8;
	t += (b9 << 29) & b57;
	carry = t >> 57;
	t &= b57;
	store_i64(a+4, a2+4, t);
	// a[4] = carry;
	// a[4] += b[7] >> 28;
	// a[4] += b[8];
	// a[4] += (b[9] << 29) & bottom57BitsMask;
	// carry = a[4] >> 57;
	// a[4] = a[4] & bottom57BitsMask;

	__m128i b10 = {b[10], b2[10]};
	__m128i b11 = {b[11], b2[11]};
	t = carry + (b9 >> 28);
	t += b10;
	t += (b11 << 29) & b57;
	carry = t >> 57;
	t &= b57;
	store_i64(a+5, a2+5, t);
	// a[5] = carry;
	// a[5] += b[9] >> 28;
	// a[5] += b[10];
	// a[5] += (b[11] << 29) & bottom57BitsMask;
	// carry = a[5] >> 57;
	// a[5] = a[5] & bottom57BitsMask;

	__m128i b12 = {b[12], b2[12]};
	__m128i b13 = {b[13], b2[13]};
	t = carry + (b11 >> 28);
	t += b12;
	t += (b13 << 29) & b57;
	carry = t >> 57;
	t &= b57;
	store_i64(a+6, a2+6, t);
	// a[6] = carry;
	// a[6] += b[11] >> 28;
	// a[6] += b[12];
	// a[6] += (b[13] << 29) & bottom57BitsMask;
	// carry = a[6] >> 57;
	// a[6] = a[6] & bottom57BitsMask;

	__m128i b14 = {b[14], b2[14]};
	__m128i b15 = {b[15], b2[15]};
	t = carry + (b13 >> 28);
	t += b14;
	t += (b15 << 29) & b57;
	carry = t >> 57;
	t &= b57;
	store_i64(a+7, a2+7, t);
	// a[7] = carry;
	// a[7] += b[13] >> 28;
	// a[7] += b[14];
	// a[7] += (b[15] << 29) & bottom57BitsMask;
	// carry = a[7] >> 57;
	// a[7] = a[7] & bottom57BitsMask;

	__m128i b16 = {b[16], b2[16]};
	t = carry + (b15 >> 28) + b16;
	store_i64(a+8, a2+8, t);
	// a[8] = carry;
	// a[8] += b[15] >> 28;
	// a[8] += b[16];
	a[9] = 0;
	a2[9] = 0;
}

uint64 sm2ReduceDegree_2way(sm2element *a, sm2element *a2, uint64 *b, uint64 *b2, uint64 *tmp, uint64 *tmp2) {
	sm2P256FromLargeElement_2Way(tmp, b, tmp2, b2);
	reduceDegree_2wayNew(tmp, tmp2);
	return sm2P256DivideByR_2way(a, a2, tmp, tmp2);
}

void sm2P256Square2Way(uint64 *tmp_, sm2element *a_, uint64 *tmp2_, sm2element *a2_) {
	__m128i t;
	__m128i a0 = SETa(0);
	t = a0 * a0;
	store_i64(tmp_, tmp2_, t);

	__m128i a1 = SETa(1);
	t = a0 * a1 << 1;
	store_i64(tmp_+1, tmp2_+1, t);

	__m128i a2 = SETa(2);
	t = a0 * a2;
	t += a1 * a1;
	t <<= 1;
	store_i64(tmp_+2, tmp2_+2, t);

	__m128i a3 = SETa(3);
	t = a0 * a3;
	t += a1 * a2;
	t <<= 1;
	store_i64(tmp_+3, tmp2_+3, t);

	__m128i a4 = SETa(4);
	t = a0 * a4;
	t += a1 * (a3 << 1);
	t <<= 1;
	t += a2 * a2;
	store_i64(tmp_+4, tmp2_+4, t);

	__m128i a5 = SETa(5);
	t = a0 * a5;
	t += a1 * a4;
	t += a2 * a3;
	t <<= 1;
	store_i64(tmp_+5, tmp2_+5, t);

	__m128i a6 = SETa(6);
	t = a0 * a6;
	t += a1 * (a5 << 1);
	t += a2 * a4;
	t += a3 * a3;
	t <<= 1;
	store_i64(tmp_+6, tmp2_+6, t);

	__m128i a7 = SETa(7);
	t = a0 * a7;
	t += a1 * a6;
	t += a2 * a5;
	t += a3 * a4;
	t <<= 1;
	store_i64(tmp_+7, tmp2_+7, t);

	__m128i a8 = SETa(8);
	t = a0 * a8;
	t += a1 * (a7 << 1);
	t += a2 * a6;
	t += a3 * (a5 << 1);
	t <<= 1;
	t += a4 * a4;
	store_i64(tmp_+8, tmp2_+8, t);

	t = a1 * a8;
	t += a2 * a7;
	t += a3 * a6;
	t += a4 * a5;
	t <<= 1;
	store_i64(tmp_+9, tmp2_+9, t);

	t = a2 * a8;
	t += a3 * (a7 << 1);
	t += a4 * a6;
	t += a5 * a5;
	t <<= 1;
	store_i64(tmp_+10, tmp2_+10, t);

	t = a3 * a8;
	t += a4 * a7;
	t += a5 * a6;
	t <<= 1;
	store_i64(tmp_+11, tmp2_+11, t);

	t = a4 * a8;
	t += a5 * (a7 << 1);
	t <<= 1;
	t += a6 * a6;
	store_i64(tmp_+12, tmp2_+12, t);

	t = a5 * a8;
	t += a6 * a7;
	t <<= 1;
	store_i64(tmp_+13, tmp2_+13, t);

	t = a6 * a8;
	t += a7 * a7;
	t <<= 1;
	store_i64(tmp_+14, tmp2_+14, t);

	t = a7 * (a8 << 1);
	store_i64(tmp_+15, tmp2_+15, t);

	t = a8 * a8;
	store_i64(tmp_+16, tmp2_+16, t);
}
