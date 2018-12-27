	.text
	.intel_syntax noprefix
	.file	"avx.c"
	.globl	avxFloat32Dot           # -- Begin function avxFloat32Dot
	.p2align	4, 0x90
	.type	avxFloat32Dot,@function
avxFloat32Dot:                          # @avxFloat32Dot
# %bb.0:
	push	rbp
	mov	rbp, rsp
	and	rsp, -8
	mov	r9, rdi
	and	r9, -8
	je	.LBB0_1
# %bb.2:
	lea	r10, [r9 - 1]
	mov	rax, r10
	shr	rax, 3
	lea	r8d, [rax + 1]
	and	r8d, 3
	cmp	r10, 24
	jae	.LBB0_13
# %bb.3:
	vxorps	xmm0, xmm0, xmm0
	xor	eax, eax
	test	r8, r8
	jne	.LBB0_5
	jmp	.LBB0_7
.LBB0_1:
	vxorps	xmm0, xmm0, xmm0
	jmp	.LBB0_7
.LBB0_13:
	lea	r10, [r8 - 1]
	sub	r10, rax
	vxorps	xmm1, xmm1, xmm1
	xor	eax, eax
	.p2align	4, 0x90
.LBB0_14:                               # =>This Inner Loop Header: Depth=1
	vmovups	ymm0, ymmword ptr [rsi + 4*rax]
	vmovups	ymm2, ymmword ptr [rsi + 4*rax + 32]
	vmovups	ymm3, ymmword ptr [rsi + 4*rax + 64]
	vmovups	ymm4, ymmword ptr [rsi + 4*rax + 96]
	vfmadd132ps	ymm0, ymm1, ymmword ptr [rdx + 4*rax]
	vfmadd231ps	ymm0, ymm2, ymmword ptr [rdx + 4*rax + 32]
	vfmadd231ps	ymm0, ymm3, ymmword ptr [rdx + 4*rax + 64]
	vfmadd231ps	ymm0, ymm4, ymmword ptr [rdx + 4*rax + 96]
	add	rax, 32
	vmovaps	ymm1, ymm0
	add	r10, 4
	jne	.LBB0_14
# %bb.4:
	test	r8, r8
	je	.LBB0_7
.LBB0_5:
	shl	rax, 2
	neg	r8
	.p2align	4, 0x90
.LBB0_6:                                # =>This Inner Loop Header: Depth=1
	vmovups	ymm1, ymmword ptr [rsi + rax]
	vfmadd231ps	ymm0, ymm1, ymmword ptr [rdx + rax]
	add	rax, 32
	add	r8, 1
	jne	.LBB0_6
.LBB0_7:
	vmovshdup	xmm1, xmm0      # xmm1 = xmm0[1,1,3,3]
	vaddss	xmm1, xmm0, xmm1
	vpermilpd	xmm2, xmm0, 1   # xmm2 = xmm0[1,0]
	vaddss	xmm1, xmm2, xmm1
	vpermilps	xmm2, xmm0, 231 # xmm2 = xmm0[3,1,2,3]
	vaddss	xmm1, xmm2, xmm1
	vextractf128	xmm0, ymm0, 1
	vaddss	xmm1, xmm0, xmm1
	vmovshdup	xmm2, xmm0      # xmm2 = xmm0[1,1,3,3]
	vaddss	xmm1, xmm2, xmm1
	vpermilpd	xmm2, xmm0, 1   # xmm2 = xmm0[1,0]
	vaddss	xmm1, xmm2, xmm1
	vpermilps	xmm0, xmm0, 231 # xmm0 = xmm0[3,1,2,3]
	vaddss	xmm0, xmm0, xmm1
	vmovss	dword ptr [rcx], xmm0
	cmp	r9, rdi
	jae	.LBB0_12
# %bb.8:
	lea	r8, [rdi - 1]
	sub	r8, r9
	mov	rax, rdi
	and	rax, 3
	je	.LBB0_11
# %bb.9:
	neg	rax
	.p2align	4, 0x90
.LBB0_10:                               # =>This Inner Loop Header: Depth=1
	vmovss	xmm1, dword ptr [rsi + 4*r9] # xmm1 = mem[0],zero,zero,zero
	vmulss	xmm1, xmm1, dword ptr [rdx + 4*r9]
	vaddss	xmm0, xmm0, xmm1
	vmovss	dword ptr [rcx], xmm0
	add	r9, 1
	add	rax, 1
	jne	.LBB0_10
.LBB0_11:
	cmp	r8, 3
	jb	.LBB0_12
	.p2align	4, 0x90
.LBB0_15:                               # =>This Inner Loop Header: Depth=1
	vmovss	xmm1, dword ptr [rsi + 4*r9] # xmm1 = mem[0],zero,zero,zero
	vmulss	xmm1, xmm1, dword ptr [rdx + 4*r9]
	vaddss	xmm0, xmm0, xmm1
	vmovss	dword ptr [rcx], xmm0
	vmovss	xmm1, dword ptr [rsi + 4*r9 + 4] # xmm1 = mem[0],zero,zero,zero
	vmulss	xmm1, xmm1, dword ptr [rdx + 4*r9 + 4]
	vaddss	xmm0, xmm0, xmm1
	vmovss	dword ptr [rcx], xmm0
	vmovss	xmm1, dword ptr [rsi + 4*r9 + 8] # xmm1 = mem[0],zero,zero,zero
	vmulss	xmm1, xmm1, dword ptr [rdx + 4*r9 + 8]
	vaddss	xmm0, xmm0, xmm1
	vmovss	dword ptr [rcx], xmm0
	vmovss	xmm1, dword ptr [rsi + 4*r9 + 12] # xmm1 = mem[0],zero,zero,zero
	vmulss	xmm1, xmm1, dword ptr [rdx + 4*r9 + 12]
	vaddss	xmm0, xmm0, xmm1
	vmovss	dword ptr [rcx], xmm0
	add	r9, 4
	cmp	rdi, r9
	jne	.LBB0_15
.LBB0_12:
	mov	rsp, rbp
	pop	rbp
	vzeroupper
	ret
.Lfunc_end0:
	.size	avxFloat32Dot, .Lfunc_end0-avxFloat32Dot
                                        # -- End function
	.globl	avxFloat64Dot           # -- Begin function avxFloat64Dot
	.p2align	4, 0x90
	.type	avxFloat64Dot,@function
avxFloat64Dot:                          # @avxFloat64Dot
# %bb.0:
	push	rbp
	mov	rbp, rsp
	and	rsp, -8
	mov	r9, rdi
	and	r9, -4
	je	.LBB1_1
# %bb.2:
	lea	r10, [r9 - 1]
	mov	rax, r10
	shr	rax, 2
	lea	r8d, [rax + 1]
	and	r8d, 3
	cmp	r10, 12
	jae	.LBB1_13
# %bb.3:
	vxorpd	xmm0, xmm0, xmm0
	xor	eax, eax
	test	r8, r8
	jne	.LBB1_5
	jmp	.LBB1_7
.LBB1_1:
	vxorpd	xmm0, xmm0, xmm0
	jmp	.LBB1_7
.LBB1_13:
	lea	r10, [r8 - 1]
	sub	r10, rax
	vxorpd	xmm1, xmm1, xmm1
	xor	eax, eax
	.p2align	4, 0x90
.LBB1_14:                               # =>This Inner Loop Header: Depth=1
	vmovupd	ymm0, ymmword ptr [rsi + 8*rax]
	vmovupd	ymm2, ymmword ptr [rsi + 8*rax + 32]
	vmovupd	ymm3, ymmword ptr [rsi + 8*rax + 64]
	vmovupd	ymm4, ymmword ptr [rsi + 8*rax + 96]
	vfmadd132pd	ymm0, ymm1, ymmword ptr [rdx + 8*rax]
	vfmadd231pd	ymm0, ymm2, ymmword ptr [rdx + 8*rax + 32]
	vfmadd231pd	ymm0, ymm3, ymmword ptr [rdx + 8*rax + 64]
	vfmadd231pd	ymm0, ymm4, ymmword ptr [rdx + 8*rax + 96]
	add	rax, 16
	vmovapd	ymm1, ymm0
	add	r10, 4
	jne	.LBB1_14
# %bb.4:
	test	r8, r8
	je	.LBB1_7
.LBB1_5:
	shl	rax, 3
	neg	r8
	.p2align	4, 0x90
.LBB1_6:                                # =>This Inner Loop Header: Depth=1
	vmovupd	ymm1, ymmword ptr [rsi + rax]
	vfmadd231pd	ymm0, ymm1, ymmword ptr [rdx + rax]
	add	rax, 32
	add	r8, 1
	jne	.LBB1_6
.LBB1_7:
	vpermilpd	xmm1, xmm0, 1   # xmm1 = xmm0[1,0]
	vaddsd	xmm1, xmm0, xmm1
	vextractf128	xmm0, ymm0, 1
	vaddsd	xmm1, xmm0, xmm1
	vpermilpd	xmm0, xmm0, 1   # xmm0 = xmm0[1,0]
	vaddsd	xmm0, xmm0, xmm1
	vmovsd	qword ptr [rcx], xmm0
	cmp	r9, rdi
	jae	.LBB1_12
# %bb.8:
	lea	r8, [rdi - 1]
	sub	r8, r9
	mov	rax, rdi
	and	rax, 3
	je	.LBB1_11
# %bb.9:
	neg	rax
	.p2align	4, 0x90
.LBB1_10:                               # =>This Inner Loop Header: Depth=1
	vmovsd	xmm1, qword ptr [rsi + 8*r9] # xmm1 = mem[0],zero
	vmulsd	xmm1, xmm1, qword ptr [rdx + 8*r9]
	vaddsd	xmm0, xmm0, xmm1
	vmovsd	qword ptr [rcx], xmm0
	add	r9, 1
	add	rax, 1
	jne	.LBB1_10
.LBB1_11:
	cmp	r8, 3
	jb	.LBB1_12
	.p2align	4, 0x90
.LBB1_15:                               # =>This Inner Loop Header: Depth=1
	vmovsd	xmm1, qword ptr [rsi + 8*r9] # xmm1 = mem[0],zero
	vmulsd	xmm1, xmm1, qword ptr [rdx + 8*r9]
	vaddsd	xmm0, xmm0, xmm1
	vmovsd	qword ptr [rcx], xmm0
	vmovsd	xmm1, qword ptr [rsi + 8*r9 + 8] # xmm1 = mem[0],zero
	vmulsd	xmm1, xmm1, qword ptr [rdx + 8*r9 + 8]
	vaddsd	xmm0, xmm0, xmm1
	vmovsd	qword ptr [rcx], xmm0
	vmovsd	xmm1, qword ptr [rsi + 8*r9 + 16] # xmm1 = mem[0],zero
	vmulsd	xmm1, xmm1, qword ptr [rdx + 8*r9 + 16]
	vaddsd	xmm0, xmm0, xmm1
	vmovsd	qword ptr [rcx], xmm0
	vmovsd	xmm1, qword ptr [rsi + 8*r9 + 24] # xmm1 = mem[0],zero
	vmulsd	xmm1, xmm1, qword ptr [rdx + 8*r9 + 24]
	vaddsd	xmm0, xmm0, xmm1
	vmovsd	qword ptr [rcx], xmm0
	add	r9, 4
	cmp	rdi, r9
	jne	.LBB1_15
.LBB1_12:
	mov	rsp, rbp
	pop	rbp
	vzeroupper
	ret
.Lfunc_end1:
	.size	avxFloat64Dot, .Lfunc_end1-avxFloat64Dot
                                        # -- End function

	.ident	"clang version 6.0.0-1ubuntu2 (tags/RELEASE_600/final)"
	.section	".note.GNU-stack","",@progbits
