	.text
	.intel_syntax noprefix
	.file	"test.c"
	.globl	set_i64                 # -- Begin function set_i64
	.p2align	4, 0x90
	.type	set_i64,@function
set_i64:                                # @set_i64
# %bb.0:
	vmovq	xmm0, rsi
	vmovq	xmm1, rdi
	vpunpcklqdq	xmm0, xmm1, xmm0 # xmm0 = xmm1[0],xmm0[0]
	ret
.Lfunc_end0:
	.size	set_i64, .Lfunc_end0-set_i64
                                        # -- End function
	.globl	store_i64               # -- Begin function store_i64
	.p2align	4, 0x90
	.type	store_i64,@function
store_i64:                              # @store_i64
# %bb.0:
	vmovq	qword ptr [rdi], xmm0
	vpextrq	qword ptr [rsi], xmm0, 1
	ret
.Lfunc_end1:
	.size	store_i64, .Lfunc_end1-store_i64
                                        # -- End function
	.globl	store_i32               # -- Begin function store_i32
	.p2align	4, 0x90
	.type	store_i32,@function
store_i32:                              # @store_i32
# %bb.0:
	vmovq	rax, xmm0
	mov	dword ptr [rdi], eax
	vpextrq	rax, xmm0, 1
	mov	dword ptr [rsi], eax
	ret
.Lfunc_end2:
	.size	store_i32, .Lfunc_end2-store_i32
                                        # -- End function
	.globl	sm2P256Mul2Way1         # -- Begin function sm2P256Mul2Way1
	.p2align	4, 0x90
	.type	sm2P256Mul2Way1,@function
sm2P256Mul2Way1:                        # @sm2P256Mul2Way1
# %bb.0:
	push	rbp
	push	r15
	push	r14
	push	r13
	push	r12
	push	rbx
	mov	r15d, dword ptr [rsi]
	mov	r11d, dword ptr [r8]
	vmovq	xmm0, r11
	vmovq	xmm1, r15
	vpunpcklqdq	xmm9, xmm1, xmm0 # xmm9 = xmm1[0],xmm0[0]
	mov	r14d, dword ptr [rdx]
	mov	r10d, dword ptr [r9]
	vmovq	xmm0, r10
	vmovq	xmm1, r14
	vpunpcklqdq	xmm11, xmm1, xmm0 # xmm11 = xmm1[0],xmm0[0]
	vpmuludq	xmm0, xmm11, xmm9
	vmovq	qword ptr [rdi], xmm0
	vpextrq	qword ptr [rcx], xmm0, 1
	mov	eax, dword ptr [rsi + 4]
	mov	ebx, dword ptr [r8 + 4]
	vmovq	xmm0, rbx
	vmovq	xmm1, rax
	vpunpcklqdq	xmm14, xmm1, xmm0 # xmm14 = xmm1[0],xmm0[0]
	mov	eax, dword ptr [rdx + 4]
	mov	ebx, dword ptr [r9 + 4]
	vmovq	xmm0, rbx
	vmovq	xmm1, rax
	vpunpcklqdq	xmm15, xmm1, xmm0 # xmm15 = xmm1[0],xmm0[0]
	vpmuludq	xmm0, xmm15, xmm9
	vpmuludq	xmm1, xmm14, xmm11
	vpaddq	xmm0, xmm0, xmm1
	vmovq	qword ptr [rdi + 8], xmm0
	vpextrq	qword ptr [rcx + 8], xmm0, 1
	mov	eax, dword ptr [rsi + 8]
	mov	ebx, dword ptr [r8 + 8]
	vmovq	xmm0, rbx
	vmovq	xmm1, rax
	vpunpcklqdq	xmm12, xmm1, xmm0 # xmm12 = xmm1[0],xmm0[0]
	mov	eax, dword ptr [rdx + 8]
	mov	ebx, dword ptr [r9 + 8]
	vmovq	xmm0, rbx
	vmovq	xmm1, rax
	vpunpcklqdq	xmm8, xmm1, xmm0 # xmm8 = xmm1[0],xmm0[0]
	vpmuludq	xmm0, xmm8, xmm9
	vpmuludq	xmm1, xmm15, xmm14
	vpaddq	xmm1, xmm1, xmm1
	vpmuludq	xmm6, xmm12, xmm11
	vpaddq	xmm0, xmm6, xmm0
	vpaddq	xmm0, xmm1, xmm0
	vmovq	qword ptr [rdi + 16], xmm0
	vpextrq	qword ptr [rcx + 16], xmm0, 1
	mov	r13d, dword ptr [rsi + 12]
	mov	r12d, dword ptr [r8 + 12]
	vmovq	xmm0, r12
	vmovq	xmm1, r13
	vpunpcklqdq	xmm5, xmm1, xmm0 # xmm5 = xmm1[0],xmm0[0]
	mov	eax, dword ptr [rdx + 12]
	mov	ebx, dword ptr [r9 + 12]
	vmovq	xmm0, rbx
	vmovq	xmm1, rax
	vpunpcklqdq	xmm10, xmm1, xmm0 # xmm10 = xmm1[0],xmm0[0]
	vpmuludq	xmm0, xmm10, xmm9
	vpmuludq	xmm1, xmm8, xmm14
	vpmuludq	xmm7, xmm12, xmm15
	vpaddq	xmm1, xmm1, xmm7
	vpmuludq	xmm7, xmm11, xmm5
	vpaddq	xmm0, xmm7, xmm0
	vpaddq	xmm0, xmm1, xmm0
	vmovq	qword ptr [rdi + 24], xmm0
	vpextrq	qword ptr [rcx + 24], xmm0, 1
	mov	eax, dword ptr [rsi + 16]
	mov	ebx, dword ptr [r8 + 16]
	vmovq	xmm0, rbx
	vmovq	xmm1, rax
	vpunpcklqdq	xmm13, xmm1, xmm0 # xmm13 = xmm1[0],xmm0[0]
	mov	ebx, dword ptr [rdx + 16]
	mov	eax, dword ptr [r9 + 16]
	vmovq	xmm0, rax
	vmovq	xmm1, rbx
	vpunpcklqdq	xmm6, xmm1, xmm0 # xmm6 = xmm1[0],xmm0[0]
	vpmuludq	xmm0, xmm10, xmm14
	vpmuludq	xmm1, xmm15, xmm5
	vpaddq	xmm0, xmm0, xmm1
	vpaddq	xmm0, xmm0, xmm0
	vpmuludq	xmm1, xmm9, xmm6
	vpmuludq	xmm4, xmm8, xmm12
	vpmuludq	xmm7, xmm13, xmm11
	vpaddq	xmm4, xmm4, xmm7
	vpaddq	xmm1, xmm4, xmm1
	vpaddq	xmm0, xmm0, xmm1
	vmovq	qword ptr [rdi + 32], xmm0
	vpextrq	qword ptr [rcx + 32], xmm0, 1
	mov	ebp, dword ptr [r8 + 20]
	vmovq	xmm0, rbp
	mov	ebp, dword ptr [rsi + 20]
	vmovq	xmm1, rbp
	vpunpcklqdq	xmm0, xmm1, xmm0 # xmm0 = xmm1[0],xmm0[0]
	mov	ebp, dword ptr [r9 + 20]
	vmovq	xmm1, rbp
	mov	ebp, dword ptr [rdx + 20]
	vmovq	xmm4, rbp
	vpunpcklqdq	xmm1, xmm4, xmm1 # xmm1 = xmm4[0],xmm1[0]
	vpmuludq	xmm4, xmm9, xmm1
	vpmuludq	xmm7, xmm14, xmm6
	vpmuludq	xmm2, xmm10, xmm12
	vpmuludq	xmm3, xmm8, xmm5
	vpaddq	xmm2, xmm2, xmm3
	vpmuludq	xmm3, xmm13, xmm15
	vpaddq	xmm3, xmm3, xmm7
	vpaddq	xmm2, xmm2, xmm3
	vpmuludq	xmm3, xmm11, xmm0
	vpaddq	xmm3, xmm3, xmm4
	vpaddq	xmm2, xmm2, xmm3
	vmovq	qword ptr [rdi + 40], xmm2
	vpextrq	qword ptr [rcx + 40], xmm2, 1
	mov	ebp, dword ptr [r8 + 24]
	vmovq	xmm2, rbp
	mov	ebp, dword ptr [rsi + 24]
	vmovq	xmm3, rbp
	mov	ebp, dword ptr [r9 + 24]
	vmovq	xmm4, rbp
	mov	ebp, dword ptr [rdx + 24]
	vmovq	xmm7, rbp
	vpunpcklqdq	xmm2, xmm3, xmm2 # xmm2 = xmm3[0],xmm2[0]
	vpunpcklqdq	xmm3, xmm7, xmm4 # xmm3 = xmm7[0],xmm4[0]
	vpmuludq	xmm4, xmm14, xmm1
	vpmuludq	xmm7, xmm10, xmm5
	vpmuludq	xmm5, xmm15, xmm0
	vpaddq	xmm5, xmm5, xmm7
	vpaddq	xmm4, xmm5, xmm4
	vpaddq	xmm4, xmm4, xmm4
	vpmuludq	xmm5, xmm9, xmm3
	vpmuludq	xmm7, xmm12, xmm6
	vpmuludq	xmm6, xmm13, xmm8
	vpaddq	xmm6, xmm7, xmm6
	vpmuludq	xmm7, xmm11, xmm2
	vpaddq	xmm6, xmm6, xmm7
	vpaddq	xmm5, xmm6, xmm5
	vpaddq	xmm4, xmm4, xmm5
	vmovq	qword ptr [rdi + 48], xmm4
	vpextrq	qword ptr [rcx + 48], xmm4, 1
	vpmuludq	xmm3, xmm14, xmm3
	vpmuludq	xmm1, xmm12, xmm1
	vpmuludq	xmm4, xmm13, xmm10
	vpmuludq	xmm0, xmm8, xmm0
	vpaddq	xmm0, xmm0, xmm4
	vpmuludq	xmm2, xmm15, xmm2
	vpaddq	xmm1, xmm1, xmm2
	vpaddq	xmm0, xmm0, xmm1
	vpaddq	xmm0, xmm0, xmm3
	vmovq	qword ptr [rdi + 56], xmm0
	vpextrq	qword ptr [rcx + 56], xmm0, 1
	mov	edx, dword ptr [rdx + 28]
	imul	rdx, r15
	imul	rbx, r13
	add	rbx, rdx
	mov	edx, dword ptr [rsi + 28]
	imul	rdx, r14
	add	rbx, qword ptr [rdi + 56]
	add	rbx, rdx
	mov	qword ptr [rdi + 56], rbx
	mov	edx, dword ptr [r9 + 28]
	imul	rdx, r11
	imul	rax, r12
	add	rax, rdx
	mov	edx, dword ptr [r8 + 28]
	imul	rdx, r10
	add	rax, qword ptr [rcx + 56]
	add	rax, rdx
	mov	qword ptr [rcx + 56], rax
	pop	rbx
	pop	r12
	pop	r13
	pop	r14
	pop	r15
	pop	rbp
	ret
.Lfunc_end3:
	.size	sm2P256Mul2Way1, .Lfunc_end3-sm2P256Mul2Way1
                                        # -- End function
	.globl	sm2P256Mul2Way2         # -- Begin function sm2P256Mul2Way2
	.p2align	4, 0x90
	.type	sm2P256Mul2Way2,@function
sm2P256Mul2Way2:                        # @sm2P256Mul2Way2
# %bb.0:
	push	rbp
	push	r15
	push	r14
	push	r13
	push	r12
	push	rbx
	mov	qword ptr [rdi + 72], 0
	mov	qword ptr [rcx + 72], 0
	mov	r14d, dword ptr [rsi + 32]
	mov	r10d, dword ptr [r8 + 32]
	vmovq	xmm0, r10
	vmovq	xmm1, r14
	vpunpcklqdq	xmm8, xmm1, xmm0 # xmm8 = xmm1[0],xmm0[0]
	mov	r15d, dword ptr [rdx + 32]
	mov	r11d, dword ptr [r9 + 32]
	vmovq	xmm0, r11
	vmovq	xmm1, r15
	vpunpcklqdq	xmm13, xmm1, xmm0 # xmm13 = xmm1[0],xmm0[0]
	vpmuludq	xmm0, xmm13, xmm8
	vmovq	qword ptr [rdi + 128], xmm0
	vpextrq	qword ptr [rcx + 128], xmm0, 1
	mov	eax, dword ptr [rsi + 28]
	mov	ebx, dword ptr [r8 + 28]
	vmovq	xmm0, rbx
	vmovq	xmm1, rax
	vpunpcklqdq	xmm10, xmm1, xmm0 # xmm10 = xmm1[0],xmm0[0]
	mov	eax, dword ptr [rdx + 28]
	mov	ebx, dword ptr [r9 + 28]
	vmovq	xmm0, rbx
	vmovq	xmm1, rax
	vpunpcklqdq	xmm9, xmm1, xmm0 # xmm9 = xmm1[0],xmm0[0]
	vpmuludq	xmm0, xmm10, xmm13
	vpmuludq	xmm1, xmm9, xmm8
	vpaddq	xmm0, xmm1, xmm0
	vmovq	qword ptr [rdi + 120], xmm0
	vpextrq	qword ptr [rcx + 120], xmm0, 1
	mov	eax, dword ptr [rsi + 24]
	mov	ebx, dword ptr [r8 + 24]
	vmovq	xmm0, rbx
	vmovq	xmm1, rax
	vpunpcklqdq	xmm11, xmm1, xmm0 # xmm11 = xmm1[0],xmm0[0]
	mov	eax, dword ptr [rdx + 24]
	mov	ebx, dword ptr [r9 + 24]
	vmovq	xmm0, rbx
	vmovq	xmm1, rax
	vpunpcklqdq	xmm15, xmm1, xmm0 # xmm15 = xmm1[0],xmm0[0]
	vpmuludq	xmm0, xmm11, xmm13
	vpmuludq	xmm1, xmm9, xmm10
	vpaddq	xmm1, xmm1, xmm1
	vpmuludq	xmm4, xmm15, xmm8
	vpaddq	xmm0, xmm0, xmm4
	vpaddq	xmm0, xmm1, xmm0
	vmovq	qword ptr [rdi + 112], xmm0
	vpextrq	qword ptr [rcx + 112], xmm0, 1
	mov	eax, dword ptr [rsi + 20]
	mov	ebx, dword ptr [r8 + 20]
	vmovq	xmm0, rbx
	vmovq	xmm1, rax
	vpunpcklqdq	xmm12, xmm1, xmm0 # xmm12 = xmm1[0],xmm0[0]
	mov	r13d, dword ptr [rdx + 20]
	mov	r12d, dword ptr [r9 + 20]
	vmovq	xmm0, r12
	vmovq	xmm1, r13
	vpunpcklqdq	xmm14, xmm1, xmm0 # xmm14 = xmm1[0],xmm0[0]
	vpmuludq	xmm0, xmm12, xmm13
	vpmuludq	xmm1, xmm11, xmm9
	vpmuludq	xmm4, xmm15, xmm10
	vpaddq	xmm1, xmm4, xmm1
	vpmuludq	xmm4, xmm14, xmm8
	vpaddq	xmm0, xmm0, xmm4
	vpaddq	xmm0, xmm1, xmm0
	vmovq	qword ptr [rdi + 104], xmm0
	vpextrq	qword ptr [rcx + 104], xmm0, 1
	mov	eax, dword ptr [rsi + 16]
	mov	ebx, dword ptr [r8 + 16]
	vmovq	xmm0, rbx
	vmovq	xmm1, rax
	vpunpcklqdq	xmm3, xmm1, xmm0 # xmm3 = xmm1[0],xmm0[0]
	mov	ebp, dword ptr [r9 + 16]
	vmovq	xmm0, rbp
	mov	ebp, dword ptr [rdx + 16]
	vmovq	xmm1, rbp
	vpunpcklqdq	xmm0, xmm1, xmm0 # xmm0 = xmm1[0],xmm0[0]
	vpmuludq	xmm1, xmm12, xmm9
	vpmuludq	xmm4, xmm14, xmm10
	vpaddq	xmm1, xmm4, xmm1
	vpaddq	xmm1, xmm1, xmm1
	vpmuludq	xmm4, xmm13, xmm3
	vpmuludq	xmm6, xmm15, xmm11
	vpaddq	xmm4, xmm6, xmm4
	vpmuludq	xmm6, xmm8, xmm0
	vpaddq	xmm4, xmm4, xmm6
	vpaddq	xmm1, xmm1, xmm4
	vmovq	qword ptr [rdi + 96], xmm1
	vpextrq	qword ptr [rcx + 96], xmm1, 1
	mov	ebp, dword ptr [r8 + 12]
	vmovq	xmm1, rbp
	mov	ebp, dword ptr [rsi + 12]
	vmovq	xmm4, rbp
	vpunpcklqdq	xmm1, xmm4, xmm1 # xmm1 = xmm4[0],xmm1[0]
	mov	ebp, dword ptr [r9 + 12]
	vmovq	xmm4, rbp
	mov	ebp, dword ptr [rdx + 12]
	vmovq	xmm6, rbp
	vpunpcklqdq	xmm4, xmm6, xmm4 # xmm4 = xmm6[0],xmm4[0]
	vpmuludq	xmm6, xmm13, xmm1
	vpmuludq	xmm2, xmm9, xmm3
	vpmuludq	xmm7, xmm12, xmm15
	vpmuludq	xmm5, xmm14, xmm11
	vpaddq	xmm5, xmm5, xmm7
	vpmuludq	xmm7, xmm10, xmm0
	vpaddq	xmm2, xmm2, xmm7
	vpaddq	xmm2, xmm5, xmm2
	vpmuludq	xmm5, xmm8, xmm4
	vpaddq	xmm5, xmm6, xmm5
	vpaddq	xmm2, xmm2, xmm5
	vmovq	qword ptr [rdi + 88], xmm2
	vpextrq	qword ptr [rcx + 88], xmm2, 1
	mov	ebp, dword ptr [r8 + 8]
	vmovq	xmm2, rbp
	mov	ebp, dword ptr [rsi + 8]
	vmovq	xmm5, rbp
	mov	ebp, dword ptr [r9 + 8]
	vmovq	xmm6, rbp
	mov	ebp, dword ptr [rdx + 8]
	vmovq	xmm7, rbp
	vpunpcklqdq	xmm2, xmm5, xmm2 # xmm2 = xmm5[0],xmm2[0]
	vpunpcklqdq	xmm5, xmm7, xmm6 # xmm5 = xmm7[0],xmm6[0]
	vpmuludq	xmm6, xmm9, xmm1
	vpmuludq	xmm7, xmm14, xmm12
	vpaddq	xmm6, xmm6, xmm7
	vpmuludq	xmm7, xmm10, xmm4
	vpaddq	xmm6, xmm6, xmm7
	vpaddq	xmm14, xmm6, xmm6
	vpmuludq	xmm7, xmm13, xmm2
	vpmuludq	xmm6, xmm15, xmm3
	vpmuludq	xmm3, xmm11, xmm0
	vpaddq	xmm3, xmm3, xmm6
	vpaddq	xmm3, xmm3, xmm7
	vpmuludq	xmm6, xmm8, xmm5
	vpaddq	xmm3, xmm3, xmm6
	vpaddq	xmm3, xmm14, xmm3
	vmovq	qword ptr [rdi + 80], xmm3
	vpextrq	qword ptr [rcx + 80], xmm3, 1
	vpmuludq	xmm2, xmm9, xmm2
	vpmuludq	xmm1, xmm15, xmm1
	vpmuludq	xmm0, xmm12, xmm0
	vpaddq	xmm0, xmm1, xmm0
	vpmuludq	xmm1, xmm11, xmm4
	vpaddq	xmm1, xmm1, xmm2
	vpaddq	xmm0, xmm0, xmm1
	vpmuludq	xmm1, xmm10, xmm5
	vpaddq	xmm0, xmm0, xmm1
	vmovq	qword ptr [rdi + 72], xmm0
	vpextrq	qword ptr [rcx + 72], xmm0, 1
	mov	esi, dword ptr [rsi + 4]
	imul	rsi, r15
	imul	rax, r13
	add	rax, rsi
	mov	edx, dword ptr [rdx + 4]
	imul	rdx, r14
	add	rax, qword ptr [rdi + 72]
	add	rax, rdx
	mov	qword ptr [rdi + 72], rax
	mov	eax, dword ptr [r8 + 4]
	imul	rax, r11
	imul	rbx, r12
	add	rbx, rax
	mov	eax, dword ptr [r9 + 4]
	imul	rax, r10
	add	rbx, qword ptr [rcx + 72]
	add	rbx, rax
	mov	qword ptr [rcx + 72], rbx
	pop	rbx
	pop	r12
	pop	r13
	pop	r14
	pop	r15
	pop	rbp
	ret
.Lfunc_end4:
	.size	sm2P256Mul2Way2, .Lfunc_end4-sm2P256Mul2Way2
                                        # -- End function
	.globl	sm2P256Mul2Way3         # -- Begin function sm2P256Mul2Way3
	.p2align	4, 0x90
	.type	sm2P256Mul2Way3,@function
sm2P256Mul2Way3:                        # @sm2P256Mul2Way3
# %bb.0:
	push	r15
	push	r14
	push	rbx
	sub	rsp, 144
	mov	r10d, dword ptr [rsi]
	mov	ebx, dword ptr [rsi + 4]
	mov	r14d, dword ptr [r8]
	mov	eax, dword ptr [r8 + 4]
	vmovq	xmm0, rax
	vmovq	xmm1, rbx
	vpunpcklqdq	xmm0, xmm1, xmm0 # xmm0 = xmm1[0],xmm0[0]
	vmovdqa	xmmword ptr [rsp + 128], xmm0 # 16-byte Spill
	mov	eax, dword ptr [rdx + 28]
	mov	ebx, dword ptr [r9 + 28]
	vmovq	xmm0, rbx
	vmovq	xmm1, rax
	vpunpcklqdq	xmm15, xmm1, xmm0 # xmm15 = xmm1[0],xmm0[0]
	mov	eax, dword ptr [rsi + 12]
	mov	ebx, dword ptr [r8 + 12]
	vmovq	xmm0, rbx
	vmovq	xmm1, rax
	vpunpcklqdq	xmm0, xmm1, xmm0 # xmm0 = xmm1[0],xmm0[0]
	vmovdqa	xmmword ptr [rsp + 112], xmm0 # 16-byte Spill
	mov	eax, dword ptr [rdx + 20]
	mov	ebx, dword ptr [r9 + 20]
	vmovq	xmm0, rbx
	vmovq	xmm1, rax
	vpunpcklqdq	xmm14, xmm1, xmm0 # xmm14 = xmm1[0],xmm0[0]
	mov	eax, dword ptr [rsi + 20]
	mov	ebx, dword ptr [r8 + 20]
	vmovq	xmm0, rbx
	vmovq	xmm1, rax
	vpunpcklqdq	xmm13, xmm1, xmm0 # xmm13 = xmm1[0],xmm0[0]
	mov	eax, dword ptr [rdx + 12]
	mov	ebx, dword ptr [r9 + 12]
	vmovq	xmm0, rbx
	vmovq	xmm1, rax
	vpunpcklqdq	xmm10, xmm1, xmm0 # xmm10 = xmm1[0],xmm0[0]
	mov	eax, dword ptr [rsi + 28]
	mov	ebx, dword ptr [r8 + 28]
	vmovq	xmm0, rbx
	vmovq	xmm1, rax
	vpunpcklqdq	xmm0, xmm1, xmm0 # xmm0 = xmm1[0],xmm0[0]
	vmovdqa	xmmword ptr [rsp + 96], xmm0 # 16-byte Spill
	mov	r11d, dword ptr [rdx]
	mov	eax, dword ptr [rdx + 4]
	mov	r15d, dword ptr [r9]
	mov	ebx, dword ptr [r9 + 4]
	vmovq	xmm0, rbx
	vmovq	xmm1, rax
	vpunpcklqdq	xmm7, xmm1, xmm0 # xmm7 = xmm1[0],xmm0[0]
	vmovq	xmm0, r14
	vmovq	xmm1, r10
	vpunpcklqdq	xmm9, xmm1, xmm0 # xmm9 = xmm1[0],xmm0[0]
	vmovdqa	xmmword ptr [rsp + 48], xmm9 # 16-byte Spill
	mov	eax, dword ptr [rsi + 8]
	mov	ebx, dword ptr [r8 + 8]
	vmovq	xmm0, rbx
	vmovq	xmm1, rax
	vpunpcklqdq	xmm0, xmm1, xmm0 # xmm0 = xmm1[0],xmm0[0]
	vmovdqa	xmmword ptr [rsp + 16], xmm0 # 16-byte Spill
	mov	eax, dword ptr [rdx + 24]
	mov	ebx, dword ptr [r9 + 24]
	vmovq	xmm1, rbx
	vmovq	xmm2, rax
	vpunpcklqdq	xmm12, xmm2, xmm1 # xmm12 = xmm2[0],xmm1[0]
	vmovdqa	xmmword ptr [rsp], xmm12 # 16-byte Spill
	mov	eax, dword ptr [rsi + 16]
	mov	ebx, dword ptr [r8 + 16]
	vmovq	xmm2, rbx
	vmovq	xmm3, rax
	vpunpcklqdq	xmm1, xmm3, xmm2 # xmm1 = xmm3[0],xmm2[0]
	mov	eax, dword ptr [rdx + 16]
	mov	ebx, dword ptr [r9 + 16]
	vmovq	xmm2, rbx
	vmovq	xmm4, rax
	vpunpcklqdq	xmm4, xmm4, xmm2 # xmm4 = xmm4[0],xmm2[0]
	mov	eax, dword ptr [rsi + 24]
	mov	ebx, dword ptr [r8 + 24]
	vmovq	xmm2, rbx
	vmovq	xmm5, rax
	vpunpcklqdq	xmm11, xmm5, xmm2 # xmm11 = xmm5[0],xmm2[0]
	mov	eax, dword ptr [rdx + 8]
	mov	ebx, dword ptr [r9 + 8]
	vmovq	xmm2, rbx
	vmovq	xmm6, rax
	vpunpcklqdq	xmm5, xmm6, xmm2 # xmm5 = xmm6[0],xmm2[0]
	vmovq	xmm6, r15
	vmovq	xmm8, r11
	vpunpcklqdq	xmm8, xmm8, xmm6 # xmm8 = xmm8[0],xmm6[0]
	vpmuludq	xmm6, xmm9, xmm15
	vmovdqa	xmm3, xmm14
	vmovdqa	xmmword ptr [rsp + 32], xmm14 # 16-byte Spill
	vpmuludq	xmm9, xmm14, xmm0
	vpaddq	xmm9, xmm9, xmm6
	vmovdqa	xmm0, xmmword ptr [rsp + 128] # 16-byte Reload
	vpmuludq	xmm6, xmm12, xmm0
	vmovdqa	xmm2, xmm10
	vmovdqa	xmmword ptr [rsp + 64], xmm10 # 16-byte Spill
	vpmuludq	xmm10, xmm10, xmm1
	vmovdqa	xmm12, xmm1
	vpaddq	xmm6, xmm10, xmm6
	vpaddq	xmm9, xmm9, xmm6
	vmovdqa	xmm10, xmmword ptr [rsp + 112] # 16-byte Reload
	vpmuludq	xmm6, xmm10, xmm4
	vmovdqa	xmm1, xmm7
	vmovdqa	xmmword ptr [rsp + 80], xmm7 # 16-byte Spill
	vpmuludq	xmm7, xmm11, xmm7
	vpaddq	xmm6, xmm6, xmm7
	vmovdqa	xmm14, xmm13
	vpmuludq	xmm7, xmm13, xmm5
	vpaddq	xmm6, xmm6, xmm7
	vpaddq	xmm6, xmm9, xmm6
	vmovdqa	xmm13, xmmword ptr [rsp + 96] # 16-byte Reload
	vpmuludq	xmm7, xmm8, xmm13
	vpaddq	xmm6, xmm6, xmm7
	vmovq	qword ptr [rdi + 56], xmm6
	vpextrq	qword ptr [rcx + 56], xmm6, 1
	vpmuludq	xmm6, xmm15, xmm0
	vmovdqa	xmm9, xmm0
	vpmuludq	xmm7, xmm10, xmm3
	vpaddq	xmm6, xmm7, xmm6
	vpmuludq	xmm7, xmm14, xmm2
	vpmuludq	xmm0, xmm13, xmm1
	vpaddq	xmm0, xmm7, xmm0
	vpaddq	xmm0, xmm6, xmm0
	mov	eax, dword ptr [rdx + 32]
	mov	edx, dword ptr [r9 + 32]
	vmovq	xmm6, rdx
	vmovq	xmm7, rax
	vpunpcklqdq	xmm6, xmm7, xmm6 # xmm6 = xmm7[0],xmm6[0]
	vmovdqa	xmm2, xmmword ptr [rsp + 16] # 16-byte Reload
	vmovdqa	xmm3, xmmword ptr [rsp] # 16-byte Reload
	vpmuludq	xmm7, xmm3, xmm2
	vpmuludq	xmm1, xmm12, xmm4
	vpaddq	xmm1, xmm7, xmm1
	vpmuludq	xmm7, xmm11, xmm5
	vpaddq	xmm1, xmm1, xmm7
	vpmuludq	xmm7, xmm6, xmmword ptr [rsp + 48] # 16-byte Folded Reload
	vpaddq	xmm1, xmm1, xmm7
	vpaddq	xmm0, xmm0, xmm0
	vpaddq	xmm0, xmm0, xmm1
	mov	eax, dword ptr [rsi + 32]
	mov	edx, dword ptr [r8 + 32]
	vmovq	xmm1, rdx
	vmovq	xmm7, rax
	vpunpcklqdq	xmm1, xmm7, xmm1 # xmm1 = xmm7[0],xmm1[0]
	vpmuludq	xmm7, xmm8, xmm1
	vpaddq	xmm0, xmm0, xmm7
	vmovq	qword ptr [rdi + 64], xmm0
	vpextrq	qword ptr [rcx + 64], xmm0, 1
	vpmuludq	xmm0, xmm9, xmm6
	vpmuludq	xmm6, xmm15, xmm2
	vpmuludq	xmm7, xmm10, xmm3
	vpaddq	xmm6, xmm7, xmm6
	vpmuludq	xmm3, xmm12, xmmword ptr [rsp + 32] # 16-byte Folded Reload
	vpmuludq	xmm4, xmm14, xmm4
	vpaddq	xmm3, xmm3, xmm4
	vpaddq	xmm3, xmm6, xmm3
	vpmuludq	xmm4, xmm11, xmmword ptr [rsp + 64] # 16-byte Folded Reload
	vpmuludq	xmm2, xmm13, xmm5
	vpaddq	xmm2, xmm4, xmm2
	vpaddq	xmm0, xmm2, xmm0
	vpaddq	xmm0, xmm3, xmm0
	vpmuludq	xmm1, xmm1, xmmword ptr [rsp + 80] # 16-byte Folded Reload
	vpaddq	xmm0, xmm0, xmm1
	vmovq	qword ptr [rdi + 72], xmm0
	vpextrq	qword ptr [rcx + 72], xmm0, 1
	add	rsp, 144
	pop	rbx
	pop	r14
	pop	r15
	ret
.Lfunc_end5:
	.size	sm2P256Mul2Way3, .Lfunc_end5-sm2P256Mul2Way3
                                        # -- End function
	.section	.rodata.cst16,"aM",@progbits,16
	.p2align	4               # -- Begin function reduceDegree_2way
.LCPI6_0:
	.quad	144115188075855744      # 0x1ffffffffffff80
	.quad	144115188075855744      # 0x1ffffffffffff80
.LCPI6_1:
	.quad	144114638320041984      # 0x1ffff8000000000
	.quad	144114638320041984      # 0x1ffff8000000000
.LCPI6_2:
	.quad	144115188075855872      # 0x200000000000000
	.quad	144115188075855872      # 0x200000000000000
.LCPI6_3:
	.quad	144115188075855871      # 0x1ffffffffffffff
	.quad	144115188075855871      # 0x1ffffffffffffff
.LCPI6_4:
	.quad	135107988821114880      # 0x1e0000000000000
	.quad	135107988821114880      # 0x1e0000000000000
.LCPI6_5:
	.quad	144115187807420416      # 0x1fffffff0000000
	.quad	144115187807420416      # 0x1fffffff0000000
.LCPI6_6:
	.quad	536870911               # 0x1fffffff
	.quad	536870911               # 0x1fffffff
	.text
	.globl	reduceDegree_2way
	.p2align	4, 0x90
	.type	reduceDegree_2way,@function
reduceDegree_2way:                      # @reduceDegree_2way
# %bb.0:
	vmovq	xmm0, rcx
	vmovq	xmm1, rsi
	vpunpcklqdq	xmm0, xmm1, xmm0 # xmm0 = xmm1[0],xmm0[0]
	vmovq	xmm1, qword ptr [rdx + 8] # xmm1 = mem[0],zero
	vmovq	xmm2, qword ptr [rdi + 8] # xmm2 = mem[0],zero
	vpsllq	xmm3, xmm0, 7
	vpand	xmm3, xmm3, xmmword ptr [rip + .LCPI6_0]
	vpsllq	xmm4, xmm0, 39
	vpand	xmm4, xmm4, xmmword ptr [rip + .LCPI6_1]
	vpor	xmm3, xmm3, xmmword ptr [rip + .LCPI6_2]
	vpunpcklqdq	xmm1, xmm2, xmm1 # xmm1 = xmm2[0],xmm1[0]
	vpsubq	xmm2, xmm3, xmm4
	vpaddq	xmm1, xmm2, xmm1
	vmovq	qword ptr [rdi + 8], xmm1
	vpextrq	qword ptr [rdx + 8], xmm1, 1
	vmovq	xmm1, qword ptr [rdx + 16] # xmm1 = mem[0],zero
	vmovq	xmm2, qword ptr [rdi + 16] # xmm2 = mem[0],zero
	vpunpcklqdq	xmm1, xmm2, xmm1 # xmm1 = xmm2[0],xmm1[0]
	vpsrad	xmm2, xmm0, 31
	vpsrad	xmm3, xmm0, 18
	vpshufd	xmm4, xmm3, 245         # xmm4 = xmm3[1,1,3,3]
	vpblendd	xmm2, xmm4, xmm2, 10 # xmm2 = xmm4[0],xmm2[1],xmm4[2],xmm2[3]
	vpsrlq	xmm4, xmm0, 18
	vpblendd	xmm3, xmm4, xmm3, 10 # xmm3 = xmm4[0],xmm3[1],xmm4[2],xmm3[3]
	vpsubq	xmm2, xmm2, xmm3
	vmovdqa	xmm3, xmmword ptr [rip + .LCPI6_3] # xmm3 = [144115188075855871,144115188075855871]
	vpaddq	xmm2, xmm2, xmm3
	vpaddq	xmm1, xmm1, xmm2
	vmovq	qword ptr [rdi + 16], xmm1
	vpextrq	qword ptr [rdx + 16], xmm1, 1
	vmovq	xmm1, qword ptr [rdx + 24] # xmm1 = mem[0],zero
	vmovq	xmm2, qword ptr [rdi + 24] # xmm2 = mem[0],zero
	vpsllq	xmm4, xmm0, 53
	vpand	xmm4, xmm4, xmmword ptr [rip + .LCPI6_4]
	vpunpcklqdq	xmm1, xmm2, xmm1 # xmm1 = xmm2[0],xmm1[0]
	vpxor	xmm2, xmm4, xmm3
	vpaddq	xmm1, xmm2, xmm1
	vmovq	qword ptr [rdi + 24], xmm1
	vpextrq	qword ptr [rdx + 24], xmm1, 1
	vmovq	xmm1, qword ptr [rdx + 32] # xmm1 = mem[0],zero
	vmovq	xmm2, qword ptr [rdi + 32] # xmm2 = mem[0],zero
	vpunpcklqdq	xmm1, xmm2, xmm1 # xmm1 = xmm2[0],xmm1[0]
	vpsrlq	xmm2, xmm0, 4
	vpsllq	xmm4, xmm0, 28
	vpand	xmm4, xmm4, xmmword ptr [rip + .LCPI6_5]
	vpandn	xmm2, xmm2, xmm3
	vpaddq	xmm1, xmm2, xmm1
	vpaddq	xmm1, xmm4, xmm1
	vmovq	qword ptr [rdi + 32], xmm1
	vpextrq	qword ptr [rdx + 32], xmm1, 1
	vmovq	xmm1, qword ptr [rdx + 40] # xmm1 = mem[0],zero
	vmovq	xmm2, qword ptr [rdi + 40] # xmm2 = mem[0],zero
	vpunpcklqdq	xmm1, xmm2, xmm1 # xmm1 = xmm2[0],xmm1[0]
	vpsrlq	xmm0, xmm0, 29
	vpand	xmm0, xmm0, xmmword ptr [rip + .LCPI6_6]
	vpcmpeqd	xmm2, xmm2, xmm2
	vpaddq	xmm1, xmm1, xmm2
	vpaddq	xmm0, xmm0, xmm1
	vmovq	qword ptr [rdi + 40], xmm0
	vpextrq	qword ptr [rdx + 40], xmm0, 1
	ret
.Lfunc_end6:
	.size	reduceDegree_2way, .Lfunc_end6-reduceDegree_2way
                                        # -- End function
	.section	.rodata.cst16,"aM",@progbits,16
	.p2align	4               # -- Begin function reduceDegree_2wayNew
.LCPI7_0:
	.quad	144115188075855744      # 0x1ffffffffffff80
	.quad	144115188075855744      # 0x1ffffffffffff80
.LCPI7_1:
	.quad	144114638320041984      # 0x1ffff8000000000
	.quad	144114638320041984      # 0x1ffff8000000000
.LCPI7_2:
	.quad	144115188075855872      # 0x200000000000000
	.quad	144115188075855872      # 0x200000000000000
.LCPI7_3:
	.quad	144115188075855871      # 0x1ffffffffffffff
	.quad	144115188075855871      # 0x1ffffffffffffff
.LCPI7_4:
	.quad	135107988821114880      # 0x1e0000000000000
	.quad	135107988821114880      # 0x1e0000000000000
.LCPI7_5:
	.quad	144115187807420416      # 0x1fffffff0000000
	.quad	144115187807420416      # 0x1fffffff0000000
	.text
	.globl	reduceDegree_2wayNew
	.p2align	4, 0x90
	.type	reduceDegree_2wayNew,@function
reduceDegree_2wayNew:                   # @reduceDegree_2wayNew
# %bb.0:
	push	r14
	push	rbx
	movabs	r10, 144115188075855871
	movabs	r8, 144114638320041984
	movabs	r9, 135107988821114880
	xor	ecx, ecx
	vmovdqa	xmm8, xmmword ptr [rip + .LCPI7_0] # xmm8 = [144115188075855744,144115188075855744]
	vmovdqa	xmm9, xmmword ptr [rip + .LCPI7_1] # xmm9 = [144114638320041984,144114638320041984]
	vmovdqa	xmm10, xmmword ptr [rip + .LCPI7_2] # xmm10 = [144115188075855872,144115188075855872]
	vmovdqa	xmm3, xmmword ptr [rip + .LCPI7_3] # xmm3 = [144115188075855871,144115188075855871]
	vmovdqa	xmm4, xmmword ptr [rip + .LCPI7_4] # xmm4 = [135107988821114880,135107988821114880]
	vmovdqa	xmm5, xmmword ptr [rip + .LCPI7_5] # xmm5 = [144115187807420416,144115187807420416]
	vpcmpeqd	xmm6, xmm6, xmm6
	jmp	.LBB7_1
	.p2align	4, 0x90
.LBB7_6:                                #   in Loop: Header=BB7_1 Depth=1
	vmovq	xmm7, r11
	vmovq	xmm0, rdx
	vpunpcklqdq	xmm7, xmm0, xmm7 # xmm7 = xmm0[0],xmm7[0]
	vmovq	xmm0, qword ptr [rsi + 8*rcx + 8] # xmm0 = mem[0],zero
	vmovq	xmm1, qword ptr [rdi + 8*rcx + 8] # xmm1 = mem[0],zero
	vpunpcklqdq	xmm0, xmm1, xmm0 # xmm0 = xmm1[0],xmm0[0]
	vpsllq	xmm1, xmm7, 7
	vpand	xmm1, xmm8, xmm1
	vpor	xmm1, xmm10, xmm1
	vpsllq	xmm2, xmm7, 39
	vpand	xmm2, xmm9, xmm2
	vpsubq	xmm1, xmm1, xmm2
	vpaddq	xmm0, xmm0, xmm1
	vmovq	qword ptr [rdi + 8*rcx + 8], xmm0
	vpextrq	qword ptr [rsi + 8*rcx + 8], xmm0, 1
	vmovq	xmm0, qword ptr [rsi + 8*rcx + 16] # xmm0 = mem[0],zero
	vmovq	xmm1, qword ptr [rdi + 8*rcx + 16] # xmm1 = mem[0],zero
	vpunpcklqdq	xmm0, xmm1, xmm0 # xmm0 = xmm1[0],xmm0[0]
	vpsrlq	xmm1, xmm7, 50
	vpsrlq	xmm2, xmm7, 18
	vpsubq	xmm1, xmm1, xmm2
	vpaddq	xmm0, xmm0, xmm3
	vpaddq	xmm0, xmm1, xmm0
	vmovq	qword ptr [rdi + 8*rcx + 16], xmm0
	vpextrq	qword ptr [rsi + 8*rcx + 16], xmm0, 1
	vmovq	xmm0, qword ptr [rsi + 8*rcx + 24] # xmm0 = mem[0],zero
	vmovq	xmm1, qword ptr [rdi + 8*rcx + 24] # xmm1 = mem[0],zero
	vpunpcklqdq	xmm0, xmm1, xmm0 # xmm0 = xmm1[0],xmm0[0]
	vpsllq	xmm1, xmm7, 53
	vpand	xmm1, xmm1, xmm4
	vpxor	xmm1, xmm1, xmm3
	vpaddq	xmm0, xmm0, xmm1
	vmovq	qword ptr [rdi + 8*rcx + 24], xmm0
	vpextrq	qword ptr [rsi + 8*rcx + 24], xmm0, 1
	vmovq	xmm0, qword ptr [rsi + 8*rcx + 32] # xmm0 = mem[0],zero
	vmovq	xmm1, qword ptr [rdi + 8*rcx + 32] # xmm1 = mem[0],zero
	vpunpcklqdq	xmm0, xmm1, xmm0 # xmm0 = xmm1[0],xmm0[0]
	vpsrlq	xmm1, xmm7, 4
	vpsllq	xmm2, xmm7, 28
	vpand	xmm2, xmm2, xmm5
	vpandn	xmm1, xmm1, xmm3
	vpaddq	xmm0, xmm1, xmm0
	vpaddq	xmm0, xmm2, xmm0
	vmovq	qword ptr [rdi + 8*rcx + 32], xmm0
	vpextrq	qword ptr [rsi + 8*rcx + 32], xmm0, 1
	vmovq	xmm0, qword ptr [rsi + 8*rcx + 40] # xmm0 = mem[0],zero
	vmovq	xmm1, qword ptr [rdi + 8*rcx + 40] # xmm1 = mem[0],zero
	vpunpcklqdq	xmm0, xmm1, xmm0 # xmm0 = xmm1[0],xmm0[0]
	vpaddq	xmm0, xmm0, xmm6
	vpsrlq	xmm1, xmm7, 29
	vpaddq	xmm0, xmm1, xmm0
	vmovq	qword ptr [rdi + 8*rcx + 40], xmm0
	vpextrq	qword ptr [rsi + 8*rcx + 40], xmm0, 1
.LBB7_11:                               #   in Loop: Header=BB7_1 Depth=1
	add	rcx, 1
	cmp	rcx, 5
	je	.LBB7_12
.LBB7_1:                                # =>This Inner Loop Header: Depth=1
	mov	rax, qword ptr [rdi + 8*rcx]
	cmp	rcx, 4
	jne	.LBB7_3
# %bb.2:                                #   in Loop: Header=BB7_1 Depth=1
	mov	edx, eax
	and	edx, 536870911
	and	rax, -536870912
	mov	qword ptr [rdi + 8*rcx], rax
	mov	rax, qword ptr [rsi + 8*rcx]
	mov	r11d, eax
	and	r11d, 536870911
	and	rax, -536870912
	mov	qword ptr [rsi + 8*rcx], rax
	test	rdx, rdx
	jne	.LBB7_5
	jmp	.LBB7_7
	.p2align	4, 0x90
.LBB7_3:                                #   in Loop: Header=BB7_1 Depth=1
	mov	rdx, rax
	shr	rdx, 57
	add	qword ptr [rdi + 8*rcx + 8], rdx
	and	rax, r10
	mov	r11, qword ptr [rsi + 8*rcx]
	mov	rdx, r11
	shr	rdx, 57
	add	qword ptr [rsi + 8*rcx + 8], rdx
	and	r11, r10
	mov	rdx, rax
	test	rdx, rdx
	je	.LBB7_7
.LBB7_5:                                #   in Loop: Header=BB7_1 Depth=1
	test	r11, r11
	jne	.LBB7_6
.LBB7_7:                                #   in Loop: Header=BB7_1 Depth=1
	test	rdx, rdx
	je	.LBB7_9
# %bb.8:                                #   in Loop: Header=BB7_1 Depth=1
	mov	rax, rdx
	shl	rax, 7
	lea	r14, [r10 - 127]
	and	r14, rax
	mov	rax, rdx
	shl	rax, 39
	and	rax, r8
	lea	rbx, [r10 + 1]
	or	rbx, r14
	sub	rbx, rax
	add	qword ptr [rdi + 8*rcx + 8], rbx
	mov	rax, rdx
	shr	rax, 50
	mov	rbx, rdx
	shr	rbx, 18
	add	rax, r10
	sub	rax, rbx
	add	qword ptr [rdi + 8*rcx + 16], rax
	mov	rax, rdx
	shl	rax, 53
	and	rax, r9
	xor	rax, r10
	add	qword ptr [rdi + 8*rcx + 24], rax
	mov	rax, rdx
	shr	rax, 4
	mov	rbx, rdx
	shl	rbx, 28
	lea	r14, [r10 - 268435455]
	and	r14, rbx
	xor	rax, r10
	add	rax, r14
	add	qword ptr [rdi + 8*rcx + 32], rax
	mov	rax, qword ptr [rdi + 8*rcx + 40]
	shr	rdx, 29
	add	rax, rdx
	add	rax, -1
	mov	qword ptr [rdi + 8*rcx + 40], rax
.LBB7_9:                                #   in Loop: Header=BB7_1 Depth=1
	test	r11, r11
	je	.LBB7_11
# %bb.10:                               #   in Loop: Header=BB7_1 Depth=1
	mov	rax, r11
	shl	rax, 7
	lea	rdx, [r10 - 127]
	and	rdx, rax
	mov	rax, r11
	shl	rax, 39
	and	rax, r8
	lea	rbx, [r10 + 1]
	or	rbx, rdx
	sub	rbx, rax
	add	qword ptr [rsi + 8*rcx + 8], rbx
	mov	rax, r11
	shr	rax, 50
	mov	rdx, r11
	shr	rdx, 18
	add	rax, r10
	sub	rax, rdx
	add	qword ptr [rsi + 8*rcx + 16], rax
	mov	rax, r11
	shl	rax, 53
	and	rax, r9
	xor	rax, r10
	add	qword ptr [rsi + 8*rcx + 24], rax
	mov	rax, r11
	shr	rax, 4
	mov	rdx, r11
	shl	rdx, 28
	lea	rbx, [r10 - 268435455]
	and	rbx, rdx
	xor	rax, r10
	add	rax, rbx
	add	qword ptr [rsi + 8*rcx + 32], rax
	mov	rax, qword ptr [rsi + 8*rcx + 40]
	shr	r11, 29
	add	rax, r11
	add	rax, -1
	mov	qword ptr [rsi + 8*rcx + 40], rax
	jmp	.LBB7_11
.LBB7_12:
	cmp	qword ptr [rdi + 72], -1
	je	.LBB7_13
# %bb.14:
	cmp	qword ptr [rsi + 72], -1
	je	.LBB7_15
.LBB7_16:
	pop	rbx
	pop	r14
	ret
.LBB7_13:
	mov	qword ptr [rdi + 72], 0
	movabs	rax, -144115188075855872
	add	qword ptr [rdi + 64], rax
	cmp	qword ptr [rsi + 72], -1
	jne	.LBB7_16
.LBB7_15:
	mov	qword ptr [rsi + 72], 0
	movabs	rax, -144115188075855872
	add	qword ptr [rsi + 64], rax
	jmp	.LBB7_16
.Lfunc_end7:
	.size	reduceDegree_2wayNew, .Lfunc_end7-reduceDegree_2wayNew
                                        # -- End function
	.section	.rodata.cst16,"aM",@progbits,16
	.p2align	4               # -- Begin function sm2P256DivideByR_2way
.LCPI8_0:
	.quad	268435456               # 0x10000000
	.quad	268435456               # 0x10000000
.LCPI8_1:
	.quad	536870911               # 0x1fffffff
	.quad	536870911               # 0x1fffffff
.LCPI8_2:
	.quad	268435455               # 0xfffffff
	.quad	268435455               # 0xfffffff
	.text
	.globl	sm2P256DivideByR_2way
	.p2align	4, 0x90
	.type	sm2P256DivideByR_2way,@function
sm2P256DivideByR_2way:                  # @sm2P256DivideByR_2way
# %bb.0:
	vmovq	xmm0, qword ptr [rcx + 32] # xmm0 = mem[0],zero
	vmovq	xmm1, qword ptr [rdx + 32] # xmm1 = mem[0],zero
	vpunpcklqdq	xmm0, xmm1, xmm0 # xmm0 = xmm1[0],xmm0[0]
	vmovq	xmm1, qword ptr [rcx + 40] # xmm1 = mem[0],zero
	vmovq	xmm2, qword ptr [rdx + 40] # xmm2 = mem[0],zero
	vpunpcklqdq	xmm4, xmm2, xmm1 # xmm4 = xmm2[0],xmm1[0]
	vpsrlq	xmm1, xmm0, 29
	vpxor	xmm0, xmm0, xmm0
	vpblendd	xmm2, xmm1, xmm0, 10 # xmm2 = xmm1[0],xmm0[1],xmm1[2],xmm0[3]
	vpsllq	xmm3, xmm4, 28
	vmovdqa	xmm1, xmmword ptr [rip + .LCPI8_0] # xmm1 = [268435456,268435456]
	vpand	xmm3, xmm3, xmm1
	vpaddq	xmm3, xmm3, xmm2
	vmovdqa	xmm2, xmmword ptr [rip + .LCPI8_1] # xmm2 = [536870911,536870911]
	vpand	xmm5, xmm3, xmm2
	vmovq	rax, xmm5
	mov	dword ptr [rdi], eax
	vpextrq	rax, xmm5, 1
	vpsrlq	xmm5, xmm3, 29
	mov	dword ptr [rsi], eax
	vpsrlq	xmm6, xmm4, 1
	vmovdqa	xmm3, xmmword ptr [rip + .LCPI8_2] # xmm3 = [268435455,268435455]
	vpand	xmm6, xmm6, xmm3
	vpaddq	xmm5, xmm5, xmm6
	vpsrlq	xmm6, xmm5, 28
	vpand	xmm5, xmm5, xmm3
	vmovq	rax, xmm5
	mov	dword ptr [rdi + 4], eax
	vpextrq	rax, xmm5, 1
	mov	dword ptr [rsi + 4], eax
	vmovq	xmm5, qword ptr [rcx + 48] # xmm5 = mem[0],zero
	vmovq	xmm7, qword ptr [rdx + 48] # xmm7 = mem[0],zero
	vpunpcklqdq	xmm5, xmm7, xmm5 # xmm5 = xmm7[0],xmm5[0]
	vpsrlq	xmm4, xmm4, 29
	vpblendd	xmm4, xmm4, xmm0, 10 # xmm4 = xmm4[0],xmm0[1],xmm4[2],xmm0[3]
	vpsllq	xmm7, xmm5, 28
	vpand	xmm7, xmm7, xmm1
	vpaddq	xmm4, xmm4, xmm7
	vpaddq	xmm4, xmm6, xmm4
	vpsrlq	xmm6, xmm4, 29
	vpand	xmm4, xmm4, xmm2
	vmovq	rax, xmm4
	mov	dword ptr [rdi + 8], eax
	vpextrq	rax, xmm4, 1
	mov	dword ptr [rsi + 8], eax
	vpsrlq	xmm4, xmm5, 1
	vpand	xmm4, xmm4, xmm3
	vpaddq	xmm4, xmm6, xmm4
	vpsrlq	xmm6, xmm4, 28
	vpand	xmm4, xmm4, xmm3
	vmovq	rax, xmm4
	mov	dword ptr [rdi + 12], eax
	vpextrq	rax, xmm4, 1
	mov	dword ptr [rsi + 12], eax
	vmovq	xmm4, qword ptr [rcx + 56] # xmm4 = mem[0],zero
	vmovq	xmm7, qword ptr [rdx + 56] # xmm7 = mem[0],zero
	vpunpcklqdq	xmm4, xmm7, xmm4 # xmm4 = xmm7[0],xmm4[0]
	vpsrlq	xmm5, xmm5, 29
	vpblendd	xmm5, xmm5, xmm0, 10 # xmm5 = xmm5[0],xmm0[1],xmm5[2],xmm0[3]
	vpsllq	xmm7, xmm4, 28
	vpand	xmm7, xmm7, xmm1
	vpaddq	xmm5, xmm5, xmm7
	vpaddq	xmm5, xmm6, xmm5
	vpsrlq	xmm6, xmm5, 29
	vpand	xmm5, xmm5, xmm2
	vmovq	rax, xmm5
	mov	dword ptr [rdi + 16], eax
	vpextrq	rax, xmm5, 1
	mov	dword ptr [rsi + 16], eax
	vpsrlq	xmm5, xmm4, 1
	vpand	xmm5, xmm5, xmm3
	vpaddq	xmm5, xmm6, xmm5
	vpsrlq	xmm6, xmm5, 28
	vpand	xmm5, xmm5, xmm3
	vmovq	rax, xmm5
	mov	dword ptr [rdi + 20], eax
	vpextrq	rax, xmm5, 1
	mov	dword ptr [rsi + 20], eax
	vmovq	xmm5, qword ptr [rcx + 64] # xmm5 = mem[0],zero
	vmovq	xmm7, qword ptr [rdx + 64] # xmm7 = mem[0],zero
	vpunpcklqdq	xmm5, xmm7, xmm5 # xmm5 = xmm7[0],xmm5[0]
	vpsrlq	xmm4, xmm4, 29
	vpblendd	xmm4, xmm4, xmm0, 10 # xmm4 = xmm4[0],xmm0[1],xmm4[2],xmm0[3]
	vpsllq	xmm7, xmm5, 28
	vpand	xmm7, xmm7, xmm1
	vpaddq	xmm4, xmm4, xmm7
	vpaddq	xmm4, xmm6, xmm4
	vpsrlq	xmm6, xmm4, 29
	vpand	xmm4, xmm4, xmm2
	vmovq	rax, xmm4
	mov	dword ptr [rdi + 24], eax
	vpextrq	rax, xmm4, 1
	mov	dword ptr [rsi + 24], eax
	vpsrlq	xmm4, xmm5, 1
	vpand	xmm4, xmm4, xmm3
	vpaddq	xmm4, xmm6, xmm4
	vpsrlq	xmm6, xmm4, 28
	vpand	xmm3, xmm4, xmm3
	vmovq	rax, xmm3
	mov	dword ptr [rdi + 28], eax
	vpextrq	rax, xmm3, 1
	mov	dword ptr [rsi + 28], eax
	vmovq	xmm3, qword ptr [rcx + 72] # xmm3 = mem[0],zero
	vmovq	xmm4, qword ptr [rdx + 72] # xmm4 = mem[0],zero
	vpunpcklqdq	xmm3, xmm4, xmm3 # xmm3 = xmm4[0],xmm3[0]
	vpsrlq	xmm4, xmm5, 29
	vpblendd	xmm0, xmm4, xmm0, 10 # xmm0 = xmm4[0],xmm0[1],xmm4[2],xmm0[3]
	vpsllq	xmm3, xmm3, 28
	vpand	xmm1, xmm3, xmm1
	vpaddq	xmm0, xmm0, xmm1
	vpaddq	xmm0, xmm6, xmm0
	vpsrlq	xmm1, xmm0, 29
	vpand	xmm0, xmm0, xmm2
	vmovq	rax, xmm0
	mov	dword ptr [rdi + 32], eax
	vpextrq	rax, xmm0, 1
	mov	dword ptr [rsi + 32], eax
	vmovq	rcx, xmm1
	vpextrq	rax, xmm1, 1
	shl	rax, 32
	or	rax, rcx
	ret
.Lfunc_end8:
	.size	sm2P256DivideByR_2way, .Lfunc_end8-sm2P256DivideByR_2way
                                        # -- End function
	.section	.rodata.cst16,"aM",@progbits,16
	.p2align	4               # -- Begin function sm2P256FromLargeElement_2Way
.LCPI9_0:
	.quad	144115187538984960      # 0x1ffffffe0000000
	.quad	144115187538984960      # 0x1ffffffe0000000
.LCPI9_1:
	.quad	144115188075855871      # 0x1ffffffffffffff
	.quad	144115188075855871      # 0x1ffffffffffffff
	.text
	.globl	sm2P256FromLargeElement_2Way
	.p2align	4, 0x90
	.type	sm2P256FromLargeElement_2Way,@function
sm2P256FromLargeElement_2Way:           # @sm2P256FromLargeElement_2Way
# %bb.0:
	vmovq	xmm0, qword ptr [rcx]   # xmm0 = mem[0],zero
	vmovq	xmm1, qword ptr [rsi]   # xmm1 = mem[0],zero
	vpunpcklqdq	xmm1, xmm1, xmm0 # xmm1 = xmm1[0],xmm0[0]
	vmovq	xmm0, qword ptr [rcx + 8] # xmm0 = mem[0],zero
	vmovq	xmm2, qword ptr [rsi + 8] # xmm2 = mem[0],zero
	vpunpcklqdq	xmm2, xmm2, xmm0 # xmm2 = xmm2[0],xmm0[0]
	vpsllq	xmm3, xmm2, 29
	vmovdqa	xmm0, xmmword ptr [rip + .LCPI9_0] # xmm0 = [144115187538984960,144115187538984960]
	vpand	xmm3, xmm3, xmm0
	vpaddq	xmm3, xmm3, xmm1
	vpsrad	xmm4, xmm3, 31
	vmovdqa	xmm1, xmmword ptr [rip + .LCPI9_1] # xmm1 = [144115188075855871,144115188075855871]
	vpand	xmm5, xmm3, xmm1
	vmovq	qword ptr [rdi], xmm5
	vpextrq	qword ptr [rdx], xmm5, 1
	vpsrad	xmm3, xmm3, 25
	vpshufd	xmm3, xmm3, 245         # xmm3 = xmm3[1,1,3,3]
	vpblendd	xmm3, xmm3, xmm4, 10 # xmm3 = xmm3[0],xmm4[1],xmm3[2],xmm4[3]
	vmovq	xmm4, qword ptr [rcx + 16] # xmm4 = mem[0],zero
	vmovq	xmm5, qword ptr [rsi + 16] # xmm5 = mem[0],zero
	vpunpcklqdq	xmm4, xmm5, xmm4 # xmm4 = xmm5[0],xmm4[0]
	vmovq	xmm5, qword ptr [rcx + 24] # xmm5 = mem[0],zero
	vmovq	xmm6, qword ptr [rsi + 24] # xmm6 = mem[0],zero
	vpunpcklqdq	xmm5, xmm6, xmm5 # xmm5 = xmm6[0],xmm5[0]
	vpsrad	xmm6, xmm2, 28
	vpsrlq	xmm2, xmm2, 28
	vpblendd	xmm2, xmm2, xmm6, 10 # xmm2 = xmm2[0],xmm6[1],xmm2[2],xmm6[3]
	vpaddq	xmm2, xmm2, xmm4
	vpsllq	xmm4, xmm5, 29
	vpand	xmm4, xmm4, xmm0
	vpaddq	xmm2, xmm2, xmm4
	vpaddq	xmm2, xmm3, xmm2
	vpsrad	xmm3, xmm2, 31
	vpand	xmm4, xmm2, xmm1
	vmovq	qword ptr [rdi + 8], xmm4
	vpextrq	qword ptr [rdx + 8], xmm4, 1
	vpsrad	xmm2, xmm2, 25
	vpshufd	xmm2, xmm2, 245         # xmm2 = xmm2[1,1,3,3]
	vpblendd	xmm2, xmm2, xmm3, 10 # xmm2 = xmm2[0],xmm3[1],xmm2[2],xmm3[3]
	vmovq	xmm3, qword ptr [rcx + 32] # xmm3 = mem[0],zero
	vmovq	xmm4, qword ptr [rsi + 32] # xmm4 = mem[0],zero
	vpunpcklqdq	xmm3, xmm4, xmm3 # xmm3 = xmm4[0],xmm3[0]
	vmovq	xmm4, qword ptr [rcx + 40] # xmm4 = mem[0],zero
	vmovq	xmm6, qword ptr [rsi + 40] # xmm6 = mem[0],zero
	vpunpcklqdq	xmm4, xmm6, xmm4 # xmm4 = xmm6[0],xmm4[0]
	vpsrad	xmm6, xmm5, 28
	vpsrlq	xmm5, xmm5, 28
	vpblendd	xmm5, xmm5, xmm6, 10 # xmm5 = xmm5[0],xmm6[1],xmm5[2],xmm6[3]
	vpaddq	xmm3, xmm5, xmm3
	vpsllq	xmm5, xmm4, 29
	vpand	xmm5, xmm5, xmm0
	vpaddq	xmm3, xmm3, xmm5
	vpaddq	xmm2, xmm2, xmm3
	vpsrad	xmm3, xmm2, 31
	vpand	xmm5, xmm2, xmm1
	vmovq	qword ptr [rdi + 16], xmm5
	vpextrq	qword ptr [rdx + 16], xmm5, 1
	vpsrad	xmm2, xmm2, 25
	vpshufd	xmm2, xmm2, 245         # xmm2 = xmm2[1,1,3,3]
	vpblendd	xmm2, xmm2, xmm3, 10 # xmm2 = xmm2[0],xmm3[1],xmm2[2],xmm3[3]
	vmovq	xmm3, qword ptr [rcx + 48] # xmm3 = mem[0],zero
	vmovq	xmm5, qword ptr [rsi + 48] # xmm5 = mem[0],zero
	vpunpcklqdq	xmm3, xmm5, xmm3 # xmm3 = xmm5[0],xmm3[0]
	vmovq	xmm5, qword ptr [rcx + 56] # xmm5 = mem[0],zero
	vmovq	xmm6, qword ptr [rsi + 56] # xmm6 = mem[0],zero
	vpunpcklqdq	xmm5, xmm6, xmm5 # xmm5 = xmm6[0],xmm5[0]
	vpsrad	xmm6, xmm4, 28
	vpsrlq	xmm4, xmm4, 28
	vpblendd	xmm4, xmm4, xmm6, 10 # xmm4 = xmm4[0],xmm6[1],xmm4[2],xmm6[3]
	vpaddq	xmm3, xmm4, xmm3
	vpsllq	xmm4, xmm5, 29
	vpand	xmm4, xmm4, xmm0
	vpaddq	xmm3, xmm3, xmm4
	vpaddq	xmm2, xmm2, xmm3
	vpsrad	xmm3, xmm2, 31
	vpand	xmm4, xmm2, xmm1
	vmovq	qword ptr [rdi + 24], xmm4
	vpextrq	qword ptr [rdx + 24], xmm4, 1
	vpsrad	xmm2, xmm2, 25
	vpshufd	xmm2, xmm2, 245         # xmm2 = xmm2[1,1,3,3]
	vpblendd	xmm2, xmm2, xmm3, 10 # xmm2 = xmm2[0],xmm3[1],xmm2[2],xmm3[3]
	vmovq	xmm3, qword ptr [rcx + 64] # xmm3 = mem[0],zero
	vmovq	xmm4, qword ptr [rsi + 64] # xmm4 = mem[0],zero
	vpunpcklqdq	xmm3, xmm4, xmm3 # xmm3 = xmm4[0],xmm3[0]
	vmovq	xmm4, qword ptr [rcx + 72] # xmm4 = mem[0],zero
	vmovq	xmm6, qword ptr [rsi + 72] # xmm6 = mem[0],zero
	vpunpcklqdq	xmm4, xmm6, xmm4 # xmm4 = xmm6[0],xmm4[0]
	vpsrad	xmm6, xmm5, 28
	vpsrlq	xmm5, xmm5, 28
	vpblendd	xmm5, xmm5, xmm6, 10 # xmm5 = xmm5[0],xmm6[1],xmm5[2],xmm6[3]
	vpaddq	xmm3, xmm5, xmm3
	vpsllq	xmm5, xmm4, 29
	vpand	xmm5, xmm5, xmm0
	vpaddq	xmm3, xmm3, xmm5
	vpaddq	xmm2, xmm2, xmm3
	vpsrad	xmm3, xmm2, 31
	vpand	xmm5, xmm2, xmm1
	vmovq	qword ptr [rdi + 32], xmm5
	vpextrq	qword ptr [rdx + 32], xmm5, 1
	vpsrad	xmm2, xmm2, 25
	vpshufd	xmm2, xmm2, 245         # xmm2 = xmm2[1,1,3,3]
	vpblendd	xmm2, xmm2, xmm3, 10 # xmm2 = xmm2[0],xmm3[1],xmm2[2],xmm3[3]
	vmovq	xmm3, qword ptr [rcx + 80] # xmm3 = mem[0],zero
	vmovq	xmm5, qword ptr [rsi + 80] # xmm5 = mem[0],zero
	vpunpcklqdq	xmm3, xmm5, xmm3 # xmm3 = xmm5[0],xmm3[0]
	vmovq	xmm5, qword ptr [rcx + 88] # xmm5 = mem[0],zero
	vmovq	xmm6, qword ptr [rsi + 88] # xmm6 = mem[0],zero
	vpunpcklqdq	xmm5, xmm6, xmm5 # xmm5 = xmm6[0],xmm5[0]
	vpsrad	xmm6, xmm4, 28
	vpsrlq	xmm4, xmm4, 28
	vpblendd	xmm4, xmm4, xmm6, 10 # xmm4 = xmm4[0],xmm6[1],xmm4[2],xmm6[3]
	vpaddq	xmm3, xmm4, xmm3
	vpsllq	xmm4, xmm5, 29
	vpand	xmm4, xmm4, xmm0
	vpaddq	xmm3, xmm3, xmm4
	vpaddq	xmm2, xmm2, xmm3
	vpsrad	xmm3, xmm2, 31
	vpand	xmm4, xmm2, xmm1
	vmovq	qword ptr [rdi + 40], xmm4
	vpextrq	qword ptr [rdx + 40], xmm4, 1
	vpsrad	xmm2, xmm2, 25
	vpshufd	xmm2, xmm2, 245         # xmm2 = xmm2[1,1,3,3]
	vpblendd	xmm2, xmm2, xmm3, 10 # xmm2 = xmm2[0],xmm3[1],xmm2[2],xmm3[3]
	vmovq	xmm3, qword ptr [rcx + 96] # xmm3 = mem[0],zero
	vmovq	xmm4, qword ptr [rsi + 96] # xmm4 = mem[0],zero
	vpunpcklqdq	xmm3, xmm4, xmm3 # xmm3 = xmm4[0],xmm3[0]
	vmovq	xmm4, qword ptr [rcx + 104] # xmm4 = mem[0],zero
	vmovq	xmm6, qword ptr [rsi + 104] # xmm6 = mem[0],zero
	vpunpcklqdq	xmm4, xmm6, xmm4 # xmm4 = xmm6[0],xmm4[0]
	vpsrad	xmm6, xmm5, 28
	vpsrlq	xmm5, xmm5, 28
	vpblendd	xmm5, xmm5, xmm6, 10 # xmm5 = xmm5[0],xmm6[1],xmm5[2],xmm6[3]
	vpaddq	xmm3, xmm5, xmm3
	vpsllq	xmm5, xmm4, 29
	vpand	xmm5, xmm5, xmm0
	vpaddq	xmm3, xmm3, xmm5
	vpaddq	xmm2, xmm2, xmm3
	vpsrad	xmm3, xmm2, 31
	vpand	xmm5, xmm2, xmm1
	vmovq	qword ptr [rdi + 48], xmm5
	vpextrq	qword ptr [rdx + 48], xmm5, 1
	vpsrad	xmm2, xmm2, 25
	vpshufd	xmm2, xmm2, 245         # xmm2 = xmm2[1,1,3,3]
	vpblendd	xmm2, xmm2, xmm3, 10 # xmm2 = xmm2[0],xmm3[1],xmm2[2],xmm3[3]
	vmovq	xmm3, qword ptr [rcx + 112] # xmm3 = mem[0],zero
	vmovq	xmm5, qword ptr [rsi + 112] # xmm5 = mem[0],zero
	vpunpcklqdq	xmm3, xmm5, xmm3 # xmm3 = xmm5[0],xmm3[0]
	vmovq	xmm5, qword ptr [rcx + 120] # xmm5 = mem[0],zero
	vmovq	xmm6, qword ptr [rsi + 120] # xmm6 = mem[0],zero
	vpunpcklqdq	xmm5, xmm6, xmm5 # xmm5 = xmm6[0],xmm5[0]
	vpsrad	xmm6, xmm4, 28
	vpsrlq	xmm4, xmm4, 28
	vpblendd	xmm4, xmm4, xmm6, 10 # xmm4 = xmm4[0],xmm6[1],xmm4[2],xmm6[3]
	vpaddq	xmm3, xmm4, xmm3
	vpsllq	xmm4, xmm5, 29
	vpand	xmm0, xmm4, xmm0
	vpaddq	xmm0, xmm3, xmm0
	vpaddq	xmm0, xmm2, xmm0
	vpsrad	xmm2, xmm0, 31
	vpand	xmm1, xmm0, xmm1
	vmovq	qword ptr [rdi + 56], xmm1
	vpextrq	qword ptr [rdx + 56], xmm1, 1
	vpsrad	xmm0, xmm0, 25
	vpshufd	xmm0, xmm0, 245         # xmm0 = xmm0[1,1,3,3]
	vpblendd	xmm0, xmm0, xmm2, 10 # xmm0 = xmm0[0],xmm2[1],xmm0[2],xmm2[3]
	vmovq	xmm1, qword ptr [rcx + 128] # xmm1 = mem[0],zero
	vmovq	xmm2, qword ptr [rsi + 128] # xmm2 = mem[0],zero
	vpunpcklqdq	xmm1, xmm2, xmm1 # xmm1 = xmm2[0],xmm1[0]
	vpsrad	xmm2, xmm5, 28
	vpsrlq	xmm3, xmm5, 28
	vpblendd	xmm2, xmm3, xmm2, 10 # xmm2 = xmm3[0],xmm2[1],xmm3[2],xmm2[3]
	vpaddq	xmm1, xmm2, xmm1
	vpaddq	xmm0, xmm0, xmm1
	vmovq	qword ptr [rdi + 64], xmm0
	vpextrq	qword ptr [rdx + 64], xmm0, 1
	mov	qword ptr [rdi + 72], 0
	mov	qword ptr [rdx + 72], 0
	ret
.Lfunc_end9:
	.size	sm2P256FromLargeElement_2Way, .Lfunc_end9-sm2P256FromLargeElement_2Way
                                        # -- End function
	.section	.rodata.cst16,"aM",@progbits,16
	.p2align	4               # -- Begin function sm2ReduceDegree_2way
.LCPI10_0:
	.quad	144115187538984960      # 0x1ffffffe0000000
	.quad	144115187538984960      # 0x1ffffffe0000000
.LCPI10_1:
	.quad	144115188075855871      # 0x1ffffffffffffff
	.quad	144115188075855871      # 0x1ffffffffffffff
.LCPI10_2:
	.quad	144115188075855744      # 0x1ffffffffffff80
	.quad	144115188075855744      # 0x1ffffffffffff80
.LCPI10_3:
	.quad	144114638320041984      # 0x1ffff8000000000
	.quad	144114638320041984      # 0x1ffff8000000000
.LCPI10_4:
	.quad	144115188075855872      # 0x200000000000000
	.quad	144115188075855872      # 0x200000000000000
.LCPI10_5:
	.quad	135107988821114880      # 0x1e0000000000000
	.quad	135107988821114880      # 0x1e0000000000000
.LCPI10_6:
	.quad	144115187807420416      # 0x1fffffff0000000
	.quad	144115187807420416      # 0x1fffffff0000000
.LCPI10_7:
	.quad	268435456               # 0x10000000
	.quad	268435456               # 0x10000000
.LCPI10_8:
	.quad	536870911               # 0x1fffffff
	.quad	536870911               # 0x1fffffff
.LCPI10_9:
	.quad	268435455               # 0xfffffff
	.quad	268435455               # 0xfffffff
	.text
	.globl	sm2ReduceDegree_2way
	.p2align	4, 0x90
	.type	sm2ReduceDegree_2way,@function
sm2ReduceDegree_2way:                   # @sm2ReduceDegree_2way
# %bb.0:
	push	r15
	push	r14
	push	r12
	push	rbx
	vmovq	xmm0, qword ptr [rcx]   # xmm0 = mem[0],zero
	vmovq	xmm1, qword ptr [rdx]   # xmm1 = mem[0],zero
	vpunpcklqdq	xmm0, xmm1, xmm0 # xmm0 = xmm1[0],xmm0[0]
	vmovq	xmm1, qword ptr [rcx + 8] # xmm1 = mem[0],zero
	vmovq	xmm2, qword ptr [rdx + 8] # xmm2 = mem[0],zero
	vpunpcklqdq	xmm2, xmm2, xmm1 # xmm2 = xmm2[0],xmm1[0]
	vpsllq	xmm3, xmm2, 29
	vmovdqa	xmm1, xmmword ptr [rip + .LCPI10_0] # xmm1 = [144115187538984960,144115187538984960]
	vpand	xmm3, xmm3, xmm1
	vpaddq	xmm3, xmm3, xmm0
	vpsrad	xmm4, xmm3, 31
	vmovdqa	xmm0, xmmword ptr [rip + .LCPI10_1] # xmm0 = [144115188075855871,144115188075855871]
	vpand	xmm5, xmm3, xmm0
	vmovq	qword ptr [r8], xmm5
	vpextrq	qword ptr [r9], xmm5, 1
	vpsrad	xmm3, xmm3, 25
	vpshufd	xmm3, xmm3, 245         # xmm3 = xmm3[1,1,3,3]
	vpblendd	xmm3, xmm3, xmm4, 10 # xmm3 = xmm3[0],xmm4[1],xmm3[2],xmm4[3]
	vmovq	xmm4, qword ptr [rcx + 16] # xmm4 = mem[0],zero
	vmovq	xmm5, qword ptr [rdx + 16] # xmm5 = mem[0],zero
	vpunpcklqdq	xmm4, xmm5, xmm4 # xmm4 = xmm5[0],xmm4[0]
	vmovq	xmm5, qword ptr [rcx + 24] # xmm5 = mem[0],zero
	vmovq	xmm6, qword ptr [rdx + 24] # xmm6 = mem[0],zero
	vpunpcklqdq	xmm5, xmm6, xmm5 # xmm5 = xmm6[0],xmm5[0]
	vpsrad	xmm6, xmm2, 28
	vpsrlq	xmm2, xmm2, 28
	vpblendd	xmm2, xmm2, xmm6, 10 # xmm2 = xmm2[0],xmm6[1],xmm2[2],xmm6[3]
	vpaddq	xmm2, xmm2, xmm4
	vpsllq	xmm4, xmm5, 29
	vpand	xmm4, xmm4, xmm1
	vpaddq	xmm2, xmm2, xmm4
	vpaddq	xmm2, xmm3, xmm2
	vpsrad	xmm3, xmm2, 31
	vpand	xmm4, xmm2, xmm0
	vmovq	qword ptr [r8 + 8], xmm4
	vpextrq	qword ptr [r9 + 8], xmm4, 1
	vpsrad	xmm2, xmm2, 25
	vpshufd	xmm2, xmm2, 245         # xmm2 = xmm2[1,1,3,3]
	vpblendd	xmm2, xmm2, xmm3, 10 # xmm2 = xmm2[0],xmm3[1],xmm2[2],xmm3[3]
	vmovq	xmm3, qword ptr [rcx + 32] # xmm3 = mem[0],zero
	vmovq	xmm4, qword ptr [rdx + 32] # xmm4 = mem[0],zero
	vpunpcklqdq	xmm3, xmm4, xmm3 # xmm3 = xmm4[0],xmm3[0]
	vmovq	xmm4, qword ptr [rcx + 40] # xmm4 = mem[0],zero
	vmovq	xmm6, qword ptr [rdx + 40] # xmm6 = mem[0],zero
	vpunpcklqdq	xmm4, xmm6, xmm4 # xmm4 = xmm6[0],xmm4[0]
	vpsrad	xmm6, xmm5, 28
	vpsrlq	xmm5, xmm5, 28
	vpblendd	xmm5, xmm5, xmm6, 10 # xmm5 = xmm5[0],xmm6[1],xmm5[2],xmm6[3]
	vpaddq	xmm3, xmm5, xmm3
	vpsllq	xmm5, xmm4, 29
	vpand	xmm5, xmm5, xmm1
	vpaddq	xmm3, xmm3, xmm5
	vpaddq	xmm2, xmm2, xmm3
	vpsrad	xmm3, xmm2, 31
	vpand	xmm5, xmm2, xmm0
	vmovq	qword ptr [r8 + 16], xmm5
	vpextrq	qword ptr [r9 + 16], xmm5, 1
	vpsrad	xmm2, xmm2, 25
	vpshufd	xmm2, xmm2, 245         # xmm2 = xmm2[1,1,3,3]
	vpblendd	xmm2, xmm2, xmm3, 10 # xmm2 = xmm2[0],xmm3[1],xmm2[2],xmm3[3]
	vmovq	xmm3, qword ptr [rcx + 48] # xmm3 = mem[0],zero
	vmovq	xmm5, qword ptr [rdx + 48] # xmm5 = mem[0],zero
	vpunpcklqdq	xmm3, xmm5, xmm3 # xmm3 = xmm5[0],xmm3[0]
	vmovq	xmm5, qword ptr [rcx + 56] # xmm5 = mem[0],zero
	vmovq	xmm6, qword ptr [rdx + 56] # xmm6 = mem[0],zero
	vpunpcklqdq	xmm5, xmm6, xmm5 # xmm5 = xmm6[0],xmm5[0]
	vpsrad	xmm6, xmm4, 28
	vpsrlq	xmm4, xmm4, 28
	vpblendd	xmm4, xmm4, xmm6, 10 # xmm4 = xmm4[0],xmm6[1],xmm4[2],xmm6[3]
	vpaddq	xmm3, xmm4, xmm3
	vpsllq	xmm4, xmm5, 29
	vpand	xmm4, xmm4, xmm1
	vpaddq	xmm3, xmm3, xmm4
	vpaddq	xmm2, xmm2, xmm3
	vpsrad	xmm3, xmm2, 31
	vpand	xmm4, xmm2, xmm0
	vmovq	qword ptr [r8 + 24], xmm4
	vpextrq	qword ptr [r9 + 24], xmm4, 1
	vpsrad	xmm2, xmm2, 25
	vpshufd	xmm2, xmm2, 245         # xmm2 = xmm2[1,1,3,3]
	vpblendd	xmm2, xmm2, xmm3, 10 # xmm2 = xmm2[0],xmm3[1],xmm2[2],xmm3[3]
	vmovq	xmm3, qword ptr [rcx + 64] # xmm3 = mem[0],zero
	vmovq	xmm4, qword ptr [rdx + 64] # xmm4 = mem[0],zero
	vpunpcklqdq	xmm3, xmm4, xmm3 # xmm3 = xmm4[0],xmm3[0]
	vmovq	xmm4, qword ptr [rcx + 72] # xmm4 = mem[0],zero
	vmovq	xmm6, qword ptr [rdx + 72] # xmm6 = mem[0],zero
	vpunpcklqdq	xmm4, xmm6, xmm4 # xmm4 = xmm6[0],xmm4[0]
	vpsrad	xmm6, xmm5, 28
	vpsrlq	xmm5, xmm5, 28
	vpblendd	xmm5, xmm5, xmm6, 10 # xmm5 = xmm5[0],xmm6[1],xmm5[2],xmm6[3]
	vpaddq	xmm3, xmm5, xmm3
	vpsllq	xmm5, xmm4, 29
	vpand	xmm5, xmm5, xmm1
	vpaddq	xmm3, xmm3, xmm5
	vpaddq	xmm2, xmm2, xmm3
	vpsrad	xmm3, xmm2, 31
	vpand	xmm5, xmm2, xmm0
	vmovq	qword ptr [r8 + 32], xmm5
	vpextrq	qword ptr [r9 + 32], xmm5, 1
	vpsrad	xmm2, xmm2, 25
	vpshufd	xmm2, xmm2, 245         # xmm2 = xmm2[1,1,3,3]
	vpblendd	xmm2, xmm2, xmm3, 10 # xmm2 = xmm2[0],xmm3[1],xmm2[2],xmm3[3]
	vmovq	xmm3, qword ptr [rcx + 80] # xmm3 = mem[0],zero
	vmovq	xmm5, qword ptr [rdx + 80] # xmm5 = mem[0],zero
	vpunpcklqdq	xmm3, xmm5, xmm3 # xmm3 = xmm5[0],xmm3[0]
	vmovq	xmm5, qword ptr [rcx + 88] # xmm5 = mem[0],zero
	vmovq	xmm6, qword ptr [rdx + 88] # xmm6 = mem[0],zero
	vpunpcklqdq	xmm5, xmm6, xmm5 # xmm5 = xmm6[0],xmm5[0]
	vpsrad	xmm6, xmm4, 28
	vpsrlq	xmm4, xmm4, 28
	vpblendd	xmm4, xmm4, xmm6, 10 # xmm4 = xmm4[0],xmm6[1],xmm4[2],xmm6[3]
	vpaddq	xmm3, xmm4, xmm3
	vpsllq	xmm4, xmm5, 29
	vpand	xmm4, xmm4, xmm1
	vpaddq	xmm3, xmm3, xmm4
	vpaddq	xmm2, xmm2, xmm3
	vpsrad	xmm3, xmm2, 31
	vpand	xmm4, xmm2, xmm0
	vmovq	qword ptr [r8 + 40], xmm4
	vpextrq	qword ptr [r9 + 40], xmm4, 1
	vpsrad	xmm2, xmm2, 25
	vpshufd	xmm2, xmm2, 245         # xmm2 = xmm2[1,1,3,3]
	vpblendd	xmm2, xmm2, xmm3, 10 # xmm2 = xmm2[0],xmm3[1],xmm2[2],xmm3[3]
	vmovq	xmm3, qword ptr [rcx + 96] # xmm3 = mem[0],zero
	vmovq	xmm4, qword ptr [rdx + 96] # xmm4 = mem[0],zero
	vpunpcklqdq	xmm3, xmm4, xmm3 # xmm3 = xmm4[0],xmm3[0]
	vmovq	xmm4, qword ptr [rcx + 104] # xmm4 = mem[0],zero
	vmovq	xmm6, qword ptr [rdx + 104] # xmm6 = mem[0],zero
	vpunpcklqdq	xmm4, xmm6, xmm4 # xmm4 = xmm6[0],xmm4[0]
	vpsrad	xmm6, xmm5, 28
	vpsrlq	xmm5, xmm5, 28
	vpblendd	xmm5, xmm5, xmm6, 10 # xmm5 = xmm5[0],xmm6[1],xmm5[2],xmm6[3]
	vpaddq	xmm3, xmm5, xmm3
	vpsllq	xmm5, xmm4, 29
	vpand	xmm5, xmm5, xmm1
	vpaddq	xmm3, xmm3, xmm5
	vpaddq	xmm2, xmm2, xmm3
	vpsrad	xmm3, xmm2, 31
	vpand	xmm5, xmm2, xmm0
	vmovq	qword ptr [r8 + 48], xmm5
	vpextrq	qword ptr [r9 + 48], xmm5, 1
	vpsrad	xmm2, xmm2, 25
	vpshufd	xmm2, xmm2, 245         # xmm2 = xmm2[1,1,3,3]
	vpblendd	xmm2, xmm2, xmm3, 10 # xmm2 = xmm2[0],xmm3[1],xmm2[2],xmm3[3]
	vmovq	xmm3, qword ptr [rcx + 112] # xmm3 = mem[0],zero
	vmovq	xmm5, qword ptr [rdx + 112] # xmm5 = mem[0],zero
	vpunpcklqdq	xmm3, xmm5, xmm3 # xmm3 = xmm5[0],xmm3[0]
	vmovq	xmm5, qword ptr [rcx + 120] # xmm5 = mem[0],zero
	vmovq	xmm6, qword ptr [rdx + 120] # xmm6 = mem[0],zero
	vpunpcklqdq	xmm5, xmm6, xmm5 # xmm5 = xmm6[0],xmm5[0]
	vpsrad	xmm6, xmm4, 28
	vpsrlq	xmm4, xmm4, 28
	vpblendd	xmm4, xmm4, xmm6, 10 # xmm4 = xmm4[0],xmm6[1],xmm4[2],xmm6[3]
	vpaddq	xmm3, xmm4, xmm3
	vpsllq	xmm4, xmm5, 29
	vpand	xmm1, xmm4, xmm1
	vpaddq	xmm1, xmm3, xmm1
	vpaddq	xmm1, xmm2, xmm1
	vpsrad	xmm2, xmm1, 31
	vpand	xmm3, xmm1, xmm0
	vmovq	qword ptr [r8 + 56], xmm3
	vpextrq	qword ptr [r9 + 56], xmm3, 1
	vpsrad	xmm1, xmm1, 25
	vpshufd	xmm1, xmm1, 245         # xmm1 = xmm1[1,1,3,3]
	vpblendd	xmm1, xmm1, xmm2, 10 # xmm1 = xmm1[0],xmm2[1],xmm1[2],xmm2[3]
	vmovq	xmm2, qword ptr [rcx + 128] # xmm2 = mem[0],zero
	vmovq	xmm3, qword ptr [rdx + 128] # xmm3 = mem[0],zero
	vpunpcklqdq	xmm2, xmm3, xmm2 # xmm2 = xmm3[0],xmm2[0]
	vpsrad	xmm3, xmm5, 28
	vpsrlq	xmm4, xmm5, 28
	vpblendd	xmm3, xmm4, xmm3, 10 # xmm3 = xmm4[0],xmm3[1],xmm4[2],xmm3[3]
	vpaddq	xmm2, xmm3, xmm2
	vpaddq	xmm1, xmm1, xmm2
	vmovq	qword ptr [r8 + 64], xmm1
	vpextrq	qword ptr [r9 + 64], xmm1, 1
	movabs	r14, 144115188075855871
	movabs	r10, 144114638320041984
	movabs	r11, 135107988821114880
	mov	qword ptr [r8 + 72], 0
	mov	qword ptr [r9 + 72], 0
	xor	ecx, ecx
	vmovdqa	xmm8, xmmword ptr [rip + .LCPI10_2] # xmm8 = [144115188075855744,144115188075855744]
	vmovdqa	xmm9, xmmword ptr [rip + .LCPI10_3] # xmm9 = [144114638320041984,144114638320041984]
	vmovdqa	xmm10, xmmword ptr [rip + .LCPI10_4] # xmm10 = [144115188075855872,144115188075855872]
	vmovdqa	xmm4, xmmword ptr [rip + .LCPI10_5] # xmm4 = [135107988821114880,135107988821114880]
	vmovdqa	xmm5, xmmword ptr [rip + .LCPI10_6] # xmm5 = [144115187807420416,144115187807420416]
	vpcmpeqd	xmm6, xmm6, xmm6
	jmp	.LBB10_1
	.p2align	4, 0x90
.LBB10_6:                               #   in Loop: Header=BB10_1 Depth=1
	vmovq	xmm7, r12
	vmovq	xmm1, rax
	vpunpcklqdq	xmm7, xmm1, xmm7 # xmm7 = xmm1[0],xmm7[0]
	vmovq	xmm1, qword ptr [r9 + 8*rcx + 8] # xmm1 = mem[0],zero
	vmovq	xmm2, qword ptr [r8 + 8*rcx + 8] # xmm2 = mem[0],zero
	vpunpcklqdq	xmm1, xmm2, xmm1 # xmm1 = xmm2[0],xmm1[0]
	vpsllq	xmm2, xmm7, 7
	vpand	xmm2, xmm8, xmm2
	vpor	xmm2, xmm10, xmm2
	vpsllq	xmm3, xmm7, 39
	vpand	xmm3, xmm9, xmm3
	vpsubq	xmm2, xmm2, xmm3
	vpaddq	xmm1, xmm1, xmm2
	vmovq	qword ptr [r8 + 8*rcx + 8], xmm1
	vpextrq	qword ptr [r9 + 8*rcx + 8], xmm1, 1
	vmovq	xmm1, qword ptr [r9 + 8*rcx + 16] # xmm1 = mem[0],zero
	vmovq	xmm2, qword ptr [r8 + 8*rcx + 16] # xmm2 = mem[0],zero
	vpunpcklqdq	xmm1, xmm2, xmm1 # xmm1 = xmm2[0],xmm1[0]
	vpsrlq	xmm2, xmm7, 50
	vpsrlq	xmm3, xmm7, 18
	vpsubq	xmm2, xmm2, xmm3
	vpaddq	xmm1, xmm1, xmm0
	vpaddq	xmm1, xmm2, xmm1
	vmovq	qword ptr [r8 + 8*rcx + 16], xmm1
	vpextrq	qword ptr [r9 + 8*rcx + 16], xmm1, 1
	vmovq	xmm1, qword ptr [r9 + 8*rcx + 24] # xmm1 = mem[0],zero
	vmovq	xmm2, qword ptr [r8 + 8*rcx + 24] # xmm2 = mem[0],zero
	vpunpcklqdq	xmm1, xmm2, xmm1 # xmm1 = xmm2[0],xmm1[0]
	vpsllq	xmm2, xmm7, 53
	vpand	xmm2, xmm2, xmm4
	vpxor	xmm2, xmm2, xmm0
	vpaddq	xmm1, xmm1, xmm2
	vmovq	qword ptr [r8 + 8*rcx + 24], xmm1
	vpextrq	qword ptr [r9 + 8*rcx + 24], xmm1, 1
	vmovq	xmm1, qword ptr [r9 + 8*rcx + 32] # xmm1 = mem[0],zero
	vmovq	xmm2, qword ptr [r8 + 8*rcx + 32] # xmm2 = mem[0],zero
	vpunpcklqdq	xmm1, xmm2, xmm1 # xmm1 = xmm2[0],xmm1[0]
	vpsrlq	xmm2, xmm7, 4
	vpsllq	xmm3, xmm7, 28
	vpand	xmm3, xmm3, xmm5
	vpandn	xmm2, xmm2, xmm0
	vpaddq	xmm1, xmm2, xmm1
	vpaddq	xmm1, xmm3, xmm1
	vmovq	qword ptr [r8 + 8*rcx + 32], xmm1
	vpextrq	qword ptr [r9 + 8*rcx + 32], xmm1, 1
	vmovq	xmm1, qword ptr [r9 + 8*rcx + 40] # xmm1 = mem[0],zero
	vmovq	xmm2, qword ptr [r8 + 8*rcx + 40] # xmm2 = mem[0],zero
	vpunpcklqdq	xmm1, xmm2, xmm1 # xmm1 = xmm2[0],xmm1[0]
	vpaddq	xmm1, xmm1, xmm6
	vpsrlq	xmm2, xmm7, 29
	vpaddq	xmm1, xmm2, xmm1
	vmovq	qword ptr [r8 + 8*rcx + 40], xmm1
	vpextrq	qword ptr [r9 + 8*rcx + 40], xmm1, 1
.LBB10_11:                              #   in Loop: Header=BB10_1 Depth=1
	add	rcx, 1
	cmp	rcx, 5
	je	.LBB10_12
.LBB10_1:                               # =>This Inner Loop Header: Depth=1
	mov	rbx, qword ptr [r8 + 8*rcx]
	cmp	rcx, 4
	jne	.LBB10_3
# %bb.2:                                #   in Loop: Header=BB10_1 Depth=1
	mov	eax, ebx
	and	eax, 536870911
	and	rbx, -536870912
	mov	qword ptr [r8 + 8*rcx], rbx
	mov	rbx, qword ptr [r9 + 8*rcx]
	mov	r12d, ebx
	and	r12d, 536870911
	and	rbx, -536870912
	mov	qword ptr [r9 + 8*rcx], rbx
	test	rax, rax
	jne	.LBB10_5
	jmp	.LBB10_7
	.p2align	4, 0x90
.LBB10_3:                               #   in Loop: Header=BB10_1 Depth=1
	mov	rax, rbx
	shr	rax, 57
	add	qword ptr [r8 + 8*rcx + 8], rax
	and	rbx, r14
	mov	r12, qword ptr [r9 + 8*rcx]
	mov	rax, r12
	shr	rax, 57
	add	qword ptr [r9 + 8*rcx + 8], rax
	and	r12, r14
	mov	rax, rbx
	test	rax, rax
	je	.LBB10_7
.LBB10_5:                               #   in Loop: Header=BB10_1 Depth=1
	test	r12, r12
	jne	.LBB10_6
.LBB10_7:                               #   in Loop: Header=BB10_1 Depth=1
	test	rax, rax
	je	.LBB10_9
# %bb.8:                                #   in Loop: Header=BB10_1 Depth=1
	mov	rbx, rax
	shl	rbx, 7
	lea	r15, [r14 - 127]
	and	r15, rbx
	mov	rbx, rax
	shl	rbx, 39
	and	rbx, r10
	lea	rdx, [r14 + 1]
	or	rdx, r15
	sub	rdx, rbx
	add	qword ptr [r8 + 8*rcx + 8], rdx
	mov	rdx, rax
	shr	rdx, 50
	mov	rbx, rax
	shr	rbx, 18
	add	rdx, r14
	sub	rdx, rbx
	add	qword ptr [r8 + 8*rcx + 16], rdx
	mov	rdx, rax
	shl	rdx, 53
	and	rdx, r11
	xor	rdx, r14
	add	qword ptr [r8 + 8*rcx + 24], rdx
	mov	rdx, rax
	shr	rdx, 4
	mov	rbx, rax
	shl	rbx, 28
	lea	r15, [r14 - 268435455]
	and	r15, rbx
	xor	rdx, r14
	add	rdx, r15
	add	qword ptr [r8 + 8*rcx + 32], rdx
	mov	rdx, qword ptr [r8 + 8*rcx + 40]
	shr	rax, 29
	add	rax, rdx
	add	rax, -1
	mov	qword ptr [r8 + 8*rcx + 40], rax
.LBB10_9:                               #   in Loop: Header=BB10_1 Depth=1
	test	r12, r12
	je	.LBB10_11
# %bb.10:                               #   in Loop: Header=BB10_1 Depth=1
	mov	rax, r12
	shl	rax, 7
	lea	rdx, [r14 - 127]
	and	rdx, rax
	mov	rax, r12
	shl	rax, 39
	and	rax, r10
	lea	rbx, [r14 + 1]
	or	rbx, rdx
	sub	rbx, rax
	add	qword ptr [r9 + 8*rcx + 8], rbx
	mov	rax, r12
	shr	rax, 50
	mov	rdx, r12
	shr	rdx, 18
	add	rax, r14
	sub	rax, rdx
	add	qword ptr [r9 + 8*rcx + 16], rax
	mov	rax, r12
	shl	rax, 53
	and	rax, r11
	xor	rax, r14
	add	qword ptr [r9 + 8*rcx + 24], rax
	mov	rax, r12
	shr	rax, 4
	mov	rdx, r12
	shl	rdx, 28
	lea	rbx, [r14 - 268435455]
	and	rbx, rdx
	xor	rax, r14
	add	rax, rbx
	add	qword ptr [r9 + 8*rcx + 32], rax
	mov	rax, qword ptr [r9 + 8*rcx + 40]
	shr	r12, 29
	add	rax, r12
	add	rax, -1
	mov	qword ptr [r9 + 8*rcx + 40], rax
	jmp	.LBB10_11
.LBB10_12:
	cmp	qword ptr [r8 + 72], -1
	je	.LBB10_13
# %bb.14:
	cmp	qword ptr [r9 + 72], -1
	je	.LBB10_15
.LBB10_16:
	vmovq	xmm0, qword ptr [r9 + 32] # xmm0 = mem[0],zero
	vmovq	xmm1, qword ptr [r8 + 32] # xmm1 = mem[0],zero
	vpunpcklqdq	xmm0, xmm1, xmm0 # xmm0 = xmm1[0],xmm0[0]
	vmovq	xmm1, qword ptr [r9 + 40] # xmm1 = mem[0],zero
	vmovq	xmm2, qword ptr [r8 + 40] # xmm2 = mem[0],zero
	vpunpcklqdq	xmm4, xmm2, xmm1 # xmm4 = xmm2[0],xmm1[0]
	vpsrlq	xmm1, xmm0, 29
	vpxor	xmm0, xmm0, xmm0
	vpblendd	xmm2, xmm1, xmm0, 10 # xmm2 = xmm1[0],xmm0[1],xmm1[2],xmm0[3]
	vpsllq	xmm3, xmm4, 28
	vmovdqa	xmm1, xmmword ptr [rip + .LCPI10_7] # xmm1 = [268435456,268435456]
	vpand	xmm3, xmm3, xmm1
	vpaddq	xmm3, xmm3, xmm2
	vmovdqa	xmm2, xmmword ptr [rip + .LCPI10_8] # xmm2 = [536870911,536870911]
	vpand	xmm5, xmm3, xmm2
	vmovq	rax, xmm5
	mov	dword ptr [rdi], eax
	vpextrq	rax, xmm5, 1
	vpsrlq	xmm5, xmm3, 29
	mov	dword ptr [rsi], eax
	vpsrlq	xmm6, xmm4, 1
	vmovdqa	xmm3, xmmword ptr [rip + .LCPI10_9] # xmm3 = [268435455,268435455]
	vpand	xmm6, xmm6, xmm3
	vpaddq	xmm5, xmm5, xmm6
	vpsrlq	xmm6, xmm5, 28
	vpand	xmm5, xmm5, xmm3
	vmovq	rax, xmm5
	mov	dword ptr [rdi + 4], eax
	vpextrq	rax, xmm5, 1
	mov	dword ptr [rsi + 4], eax
	vmovq	xmm5, qword ptr [r9 + 48] # xmm5 = mem[0],zero
	vmovq	xmm7, qword ptr [r8 + 48] # xmm7 = mem[0],zero
	vpunpcklqdq	xmm5, xmm7, xmm5 # xmm5 = xmm7[0],xmm5[0]
	vpsrlq	xmm4, xmm4, 29
	vpblendd	xmm4, xmm4, xmm0, 10 # xmm4 = xmm4[0],xmm0[1],xmm4[2],xmm0[3]
	vpsllq	xmm7, xmm5, 28
	vpand	xmm7, xmm7, xmm1
	vpaddq	xmm4, xmm4, xmm7
	vpaddq	xmm4, xmm6, xmm4
	vpsrlq	xmm6, xmm4, 29
	vpand	xmm4, xmm4, xmm2
	vmovq	rax, xmm4
	mov	dword ptr [rdi + 8], eax
	vpextrq	rax, xmm4, 1
	mov	dword ptr [rsi + 8], eax
	vpsrlq	xmm4, xmm5, 1
	vpand	xmm4, xmm4, xmm3
	vpaddq	xmm4, xmm6, xmm4
	vpsrlq	xmm6, xmm4, 28
	vpand	xmm4, xmm4, xmm3
	vmovq	rax, xmm4
	mov	dword ptr [rdi + 12], eax
	vpextrq	rax, xmm4, 1
	mov	dword ptr [rsi + 12], eax
	vmovq	xmm4, qword ptr [r9 + 56] # xmm4 = mem[0],zero
	vmovq	xmm7, qword ptr [r8 + 56] # xmm7 = mem[0],zero
	vpunpcklqdq	xmm4, xmm7, xmm4 # xmm4 = xmm7[0],xmm4[0]
	vpsrlq	xmm5, xmm5, 29
	vpblendd	xmm5, xmm5, xmm0, 10 # xmm5 = xmm5[0],xmm0[1],xmm5[2],xmm0[3]
	vpsllq	xmm7, xmm4, 28
	vpand	xmm7, xmm7, xmm1
	vpaddq	xmm5, xmm5, xmm7
	vpaddq	xmm5, xmm6, xmm5
	vpsrlq	xmm6, xmm5, 29
	vpand	xmm5, xmm5, xmm2
	vmovq	rax, xmm5
	mov	dword ptr [rdi + 16], eax
	vpextrq	rax, xmm5, 1
	mov	dword ptr [rsi + 16], eax
	vpsrlq	xmm5, xmm4, 1
	vpand	xmm5, xmm5, xmm3
	vpaddq	xmm5, xmm6, xmm5
	vpsrlq	xmm6, xmm5, 28
	vpand	xmm5, xmm5, xmm3
	vmovq	rax, xmm5
	mov	dword ptr [rdi + 20], eax
	vpextrq	rax, xmm5, 1
	mov	dword ptr [rsi + 20], eax
	vmovq	xmm5, qword ptr [r9 + 64] # xmm5 = mem[0],zero
	vmovq	xmm7, qword ptr [r8 + 64] # xmm7 = mem[0],zero
	vpunpcklqdq	xmm5, xmm7, xmm5 # xmm5 = xmm7[0],xmm5[0]
	vpsrlq	xmm4, xmm4, 29
	vpblendd	xmm4, xmm4, xmm0, 10 # xmm4 = xmm4[0],xmm0[1],xmm4[2],xmm0[3]
	vpsllq	xmm7, xmm5, 28
	vpand	xmm7, xmm7, xmm1
	vpaddq	xmm4, xmm4, xmm7
	vpaddq	xmm4, xmm6, xmm4
	vpsrlq	xmm6, xmm4, 29
	vpand	xmm4, xmm4, xmm2
	vmovq	rax, xmm4
	mov	dword ptr [rdi + 24], eax
	vpextrq	rax, xmm4, 1
	mov	dword ptr [rsi + 24], eax
	vpsrlq	xmm4, xmm5, 1
	vpand	xmm4, xmm4, xmm3
	vpaddq	xmm4, xmm6, xmm4
	vpsrlq	xmm6, xmm4, 28
	vpand	xmm3, xmm4, xmm3
	vmovq	rax, xmm3
	mov	dword ptr [rdi + 28], eax
	vpextrq	rax, xmm3, 1
	mov	dword ptr [rsi + 28], eax
	vmovq	xmm3, qword ptr [r9 + 72] # xmm3 = mem[0],zero
	vmovq	xmm4, qword ptr [r8 + 72] # xmm4 = mem[0],zero
	vpunpcklqdq	xmm3, xmm4, xmm3 # xmm3 = xmm4[0],xmm3[0]
	vpsrlq	xmm4, xmm5, 29
	vpblendd	xmm0, xmm4, xmm0, 10 # xmm0 = xmm4[0],xmm0[1],xmm4[2],xmm0[3]
	vpsllq	xmm3, xmm3, 28
	vpand	xmm1, xmm3, xmm1
	vpaddq	xmm0, xmm0, xmm1
	vpaddq	xmm0, xmm6, xmm0
	vpsrlq	xmm1, xmm0, 29
	vpand	xmm0, xmm0, xmm2
	vmovq	rax, xmm0
	mov	dword ptr [rdi + 32], eax
	vpextrq	rax, xmm0, 1
	mov	dword ptr [rsi + 32], eax
	vmovq	rcx, xmm1
	vpextrq	rax, xmm1, 1
	shl	rax, 32
	or	rax, rcx
	pop	rbx
	pop	r12
	pop	r14
	pop	r15
	ret
.LBB10_13:
	mov	qword ptr [r8 + 72], 0
	movabs	rax, -144115188075855872
	add	qword ptr [r8 + 64], rax
	cmp	qword ptr [r9 + 72], -1
	jne	.LBB10_16
.LBB10_15:
	mov	qword ptr [r9 + 72], 0
	movabs	rax, -144115188075855872
	add	qword ptr [r9 + 64], rax
	jmp	.LBB10_16
.Lfunc_end10:
	.size	sm2ReduceDegree_2way, .Lfunc_end10-sm2ReduceDegree_2way
                                        # -- End function
	.globl	sm2P256Square2Way       # -- Begin function sm2P256Square2Way
	.p2align	4, 0x90
	.type	sm2P256Square2Way,@function
sm2P256Square2Way:                      # @sm2P256Square2Way
# %bb.0:
	mov	r8d, dword ptr [rsi]
	mov	eax, dword ptr [rcx]
	vmovq	xmm0, rax
	vmovq	xmm1, r8
	vpunpcklqdq	xmm4, xmm1, xmm0 # xmm4 = xmm1[0],xmm0[0]
	vpmuludq	xmm0, xmm4, xmm4
	vmovq	qword ptr [rdi], xmm0
	vpextrq	qword ptr [rdx], xmm0, 1
	mov	r8d, dword ptr [rsi + 4]
	mov	eax, dword ptr [rcx + 4]
	vmovq	xmm0, rax
	vmovq	xmm1, r8
	vpunpcklqdq	xmm0, xmm1, xmm0 # xmm0 = xmm1[0],xmm0[0]
	vpmuludq	xmm1, xmm0, xmm4
	vpaddq	xmm1, xmm1, xmm1
	vmovq	qword ptr [rdi + 8], xmm1
	vpextrq	qword ptr [rdx + 8], xmm1, 1
	mov	r8d, dword ptr [rsi + 8]
	mov	eax, dword ptr [rcx + 8]
	vmovq	xmm1, rax
	vmovq	xmm2, r8
	vpunpcklqdq	xmm8, xmm2, xmm1 # xmm8 = xmm2[0],xmm1[0]
	vpmuludq	xmm1, xmm8, xmm4
	vpmuludq	xmm2, xmm0, xmm0
	vpaddq	xmm1, xmm1, xmm2
	vpaddq	xmm1, xmm1, xmm1
	vmovq	qword ptr [rdi + 16], xmm1
	vpextrq	qword ptr [rdx + 16], xmm1, 1
	mov	r8d, dword ptr [rsi + 12]
	mov	eax, dword ptr [rcx + 12]
	vmovq	xmm1, rax
	vmovq	xmm2, r8
	vpunpcklqdq	xmm2, xmm2, xmm1 # xmm2 = xmm2[0],xmm1[0]
	vpmuludq	xmm1, xmm2, xmm4
	vpmuludq	xmm3, xmm8, xmm0
	vpaddq	xmm1, xmm1, xmm3
	vpaddq	xmm1, xmm1, xmm1
	vmovq	qword ptr [rdi + 24], xmm1
	vpextrq	qword ptr [rdx + 24], xmm1, 1
	mov	r8d, dword ptr [rsi + 16]
	mov	eax, dword ptr [rcx + 16]
	vmovq	xmm1, rax
	vmovq	xmm3, r8
	vpunpcklqdq	xmm13, xmm3, xmm1 # xmm13 = xmm3[0],xmm1[0]
	vpmuludq	xmm1, xmm13, xmm4
	vpmuludq	xmm5, xmm2, xmm0
	vpaddq	xmm5, xmm5, xmm5
	vpaddq	xmm1, xmm1, xmm5
	vpaddq	xmm1, xmm1, xmm1
	vpmuludq	xmm5, xmm8, xmm8
	vpaddq	xmm1, xmm1, xmm5
	vmovq	qword ptr [rdi + 32], xmm1
	vpextrq	qword ptr [rdx + 32], xmm1, 1
	mov	r8d, dword ptr [rsi + 20]
	mov	eax, dword ptr [rcx + 20]
	vmovq	xmm1, rax
	vmovq	xmm5, r8
	vpunpcklqdq	xmm5, xmm5, xmm1 # xmm5 = xmm5[0],xmm1[0]
	vpmuludq	xmm1, xmm5, xmm4
	vpmuludq	xmm6, xmm13, xmm0
	vpmuludq	xmm7, xmm8, xmm2
	vpaddq	xmm6, xmm6, xmm7
	vpaddq	xmm1, xmm6, xmm1
	vpaddq	xmm1, xmm1, xmm1
	vmovq	qword ptr [rdi + 40], xmm1
	vpextrq	qword ptr [rdx + 40], xmm1, 1
	mov	r8d, dword ptr [rsi + 24]
	mov	eax, dword ptr [rcx + 24]
	vmovq	xmm1, rax
	vmovq	xmm6, r8
	vpunpcklqdq	xmm6, xmm6, xmm1 # xmm6 = xmm6[0],xmm1[0]
	vpmuludq	xmm9, xmm6, xmm4
	vpaddq	xmm10, xmm5, xmm5
	vpmuludq	xmm7, xmm10, xmm0
	vpsrlq	xmm11, xmm10, 32
	vpmuludq	xmm1, xmm11, xmm0
	vpsllq	xmm1, xmm1, 32
	vpaddq	xmm12, xmm7, xmm1
	vpmuludq	xmm7, xmm13, xmm8
	vpmuludq	xmm1, xmm2, xmm2
	vpaddq	xmm1, xmm7, xmm1
	vpaddq	xmm1, xmm9, xmm1
	vpaddq	xmm1, xmm12, xmm1
	vpaddq	xmm1, xmm1, xmm1
	vmovq	qword ptr [rdi + 48], xmm1
	vpextrq	qword ptr [rdx + 48], xmm1, 1
	mov	r8d, dword ptr [rsi + 28]
	mov	eax, dword ptr [rcx + 28]
	vmovq	xmm1, rax
	vmovq	xmm7, r8
	vpunpcklqdq	xmm12, xmm7, xmm1 # xmm12 = xmm7[0],xmm1[0]
	vpmuludq	xmm1, xmm12, xmm4
	vpmuludq	xmm3, xmm6, xmm0
	vpaddq	xmm1, xmm3, xmm1
	vpmuludq	xmm3, xmm8, xmm5
	vpmuludq	xmm7, xmm13, xmm2
	vpaddq	xmm3, xmm3, xmm7
	vpaddq	xmm1, xmm3, xmm1
	vpaddq	xmm1, xmm1, xmm1
	vmovq	qword ptr [rdi + 56], xmm1
	vpextrq	qword ptr [rdx + 56], xmm1, 1
	mov	eax, dword ptr [rsi + 32]
	mov	ecx, dword ptr [rcx + 32]
	vmovq	xmm1, rcx
	vmovq	xmm3, rax
	vpunpcklqdq	xmm1, xmm3, xmm1 # xmm1 = xmm3[0],xmm1[0]
	vpmuludq	xmm14, xmm1, xmm4
	vpaddq	xmm4, xmm12, xmm12
	vpmuludq	xmm7, xmm4, xmm0
	vpsrlq	xmm9, xmm4, 32
	vpmuludq	xmm3, xmm9, xmm0
	vpsllq	xmm3, xmm3, 32
	vpaddq	xmm15, xmm7, xmm3
	vpmuludq	xmm7, xmm8, xmm6
	vpmuludq	xmm3, xmm10, xmm2
	vpaddq	xmm3, xmm3, xmm7
	vpmuludq	xmm7, xmm11, xmm2
	vpsllq	xmm7, xmm7, 32
	vpaddq	xmm3, xmm7, xmm3
	vpaddq	xmm3, xmm15, xmm3
	vpaddq	xmm3, xmm14, xmm3
	vpaddq	xmm3, xmm3, xmm3
	vpmuludq	xmm7, xmm13, xmm13
	vpaddq	xmm3, xmm3, xmm7
	vmovq	qword ptr [rdi + 64], xmm3
	vpextrq	qword ptr [rdx + 64], xmm3, 1
	vpmuludq	xmm0, xmm1, xmm0
	vpmuludq	xmm3, xmm12, xmm8
	vpaddq	xmm0, xmm3, xmm0
	vpmuludq	xmm3, xmm6, xmm2
	vpmuludq	xmm7, xmm13, xmm5
	vpaddq	xmm3, xmm3, xmm7
	vpaddq	xmm0, xmm3, xmm0
	vpaddq	xmm0, xmm0, xmm0
	vmovq	qword ptr [rdi + 72], xmm0
	vpextrq	qword ptr [rdx + 72], xmm0, 1
	vpmuludq	xmm8, xmm8, xmm1
	vpmuludq	xmm3, xmm4, xmm2
	vpmuludq	xmm7, xmm9, xmm2
	vpsllq	xmm7, xmm7, 32
	vpaddq	xmm3, xmm3, xmm7
	vpmuludq	xmm7, xmm13, xmm6
	vpmuludq	xmm0, xmm5, xmm5
	vpaddq	xmm0, xmm7, xmm0
	vpaddq	xmm0, xmm8, xmm0
	vpaddq	xmm0, xmm3, xmm0
	vpaddq	xmm0, xmm0, xmm0
	vmovq	qword ptr [rdi + 80], xmm0
	vpextrq	qword ptr [rdx + 80], xmm0, 1
	vpmuludq	xmm0, xmm1, xmm2
	vpmuludq	xmm2, xmm12, xmm13
	vpmuludq	xmm3, xmm6, xmm5
	vpaddq	xmm2, xmm2, xmm3
	vpaddq	xmm0, xmm2, xmm0
	vpaddq	xmm0, xmm0, xmm0
	vmovq	qword ptr [rdi + 88], xmm0
	vpextrq	qword ptr [rdx + 88], xmm0, 1
	vpmuludq	xmm0, xmm13, xmm1
	vpmuludq	xmm2, xmm4, xmm5
	vpaddq	xmm0, xmm2, xmm0
	vpmuludq	xmm2, xmm9, xmm5
	vpsllq	xmm2, xmm2, 32
	vpaddq	xmm0, xmm2, xmm0
	vpaddq	xmm0, xmm0, xmm0
	vpmuludq	xmm2, xmm6, xmm6
	vpaddq	xmm0, xmm0, xmm2
	vmovq	qword ptr [rdi + 96], xmm0
	vpextrq	qword ptr [rdx + 96], xmm0, 1
	vpmuludq	xmm0, xmm1, xmm5
	vpmuludq	xmm2, xmm12, xmm6
	vpaddq	xmm0, xmm0, xmm2
	vpaddq	xmm0, xmm0, xmm0
	vmovq	qword ptr [rdi + 104], xmm0
	vpextrq	qword ptr [rdx + 104], xmm0, 1
	vpmuludq	xmm0, xmm1, xmm6
	vpmuludq	xmm2, xmm12, xmm12
	vpaddq	xmm0, xmm0, xmm2
	vpaddq	xmm0, xmm0, xmm0
	vmovq	qword ptr [rdi + 112], xmm0
	vpextrq	qword ptr [rdx + 112], xmm0, 1
	vpmuludq	xmm0, xmm12, xmm1
	vpaddq	xmm0, xmm0, xmm0
	vmovq	qword ptr [rdi + 120], xmm0
	vpextrq	qword ptr [rdx + 120], xmm0, 1
	vpmuludq	xmm0, xmm1, xmm1
	vmovq	qword ptr [rdi + 128], xmm0
	vpextrq	qword ptr [rdx + 128], xmm0, 1
	ret
.Lfunc_end11:
	.size	sm2P256Square2Way, .Lfunc_end11-sm2P256Square2Way
                                        # -- End function
	.type	bottom28BitsMask,@object # @bottom28BitsMask
	.section	.rodata,"a",@progbits
	.globl	bottom28BitsMask
	.p2align	3
bottom28BitsMask:
	.quad	268435455               # 0xfffffff
	.size	bottom28BitsMask, 8

	.type	bottom29BitsMask,@object # @bottom29BitsMask
	.globl	bottom29BitsMask
	.p2align	3
bottom29BitsMask:
	.quad	536870911               # 0x1fffffff
	.size	bottom29BitsMask, 8

	.type	bottom32BitsMask,@object # @bottom32BitsMask
	.globl	bottom32BitsMask
	.p2align	3
bottom32BitsMask:
	.quad	4294967295              # 0xffffffff
	.size	bottom32BitsMask, 8

	.type	top32BitsMask,@object   # @top32BitsMask
	.globl	top32BitsMask
	.p2align	3
top32BitsMask:
	.quad	-4294967296             # 0xffffffff00000000
	.size	top32BitsMask, 8

	.type	bottom57BitsMask,@object # @bottom57BitsMask
	.globl	bottom57BitsMask
	.p2align	3
bottom57BitsMask:
	.quad	144115188075855871      # 0x1ffffffffffffff
	.size	bottom57BitsMask, 8

	.type	twoPower57,@object      # @twoPower57
	.globl	twoPower57
	.p2align	3
twoPower57:
	.quad	144115188075855872      # 0x200000000000000
	.size	twoPower57, 8

	.type	sm2P256Zero31,@object   # @sm2P256Zero31
	.globl	sm2P256Zero31
	.p2align	4
sm2P256Zero31:
	.long	2147483640              # 0x7ffffff8
	.long	1073741820              # 0x3ffffffc
	.long	2147484668              # 0x800003fc
	.long	1073733628              # 0x3fffdffc
	.long	2147483644              # 0x7ffffffc
	.long	1073741820              # 0x3ffffffc
	.long	2147483644              # 0x7ffffffc
	.long	939524092               # 0x37fffffc
	.long	2147483644              # 0x7ffffffc
	.size	sm2P256Zero31, 36

	.ident	"Ubuntu clang version 10.0.1-++20200618012851+f5a9c661a35-1~exp1~20200617233447.176 "
	.section	".note.GNU-stack","",@progbits
	.addrsig
