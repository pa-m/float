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
	mov	r8, rdx
	mov	r9, rdi
	movabs	rdx, -6148914691236517205
	mov	rax, rdi
	mul	rdx
	shr	rdx
	and	rdx, -16
	lea	rax, [rdx + 2*rdx]
	test	rax, rax
	je	.LBB0_1
# %bb.11:
	vxorps	xmm1, xmm1, xmm1
	xor	edx, edx
	vxorps	xmm5, xmm5, xmm5
	vxorps	xmm4, xmm4, xmm4
	vxorps	xmm3, xmm3, xmm3
	vxorps	xmm2, xmm2, xmm2
	vxorps	xmm0, xmm0, xmm0
	.p2align	4, 0x90
.LBB0_12:                               # =>This Inner Loop Header: Depth=1
	vmovups	ymm6, ymmword ptr [rsi]
	vmovups	ymm7, ymmword ptr [rsi + 32]
	vmovups	ymm8, ymmword ptr [rsi + 64]
	vmovups	ymm9, ymmword ptr [rsi + 96]
	vfmadd231ps	ymm0, ymm6, ymmword ptr [r8] # ymm0 = (ymm6 * mem) + ymm0
	vfmadd231ps	ymm2, ymm7, ymmword ptr [r8 + 32] # ymm2 = (ymm7 * mem) + ymm2
	vfmadd231ps	ymm3, ymm8, ymmword ptr [r8 + 64] # ymm3 = (ymm8 * mem) + ymm3
	vfmadd231ps	ymm4, ymm9, ymmword ptr [r8 + 96] # ymm4 = (ymm9 * mem) + ymm4
	vmovups	ymm6, ymmword ptr [rsi + 128]
	vfmadd231ps	ymm5, ymm6, ymmword ptr [r8 + 128] # ymm5 = (ymm6 * mem) + ymm5
	vmovups	ymm6, ymmword ptr [rsi + 160]
	vfmadd231ps	ymm1, ymm6, ymmword ptr [r8 + 160] # ymm1 = (ymm6 * mem) + ymm1
	add	rsi, 192
	add	r8, 192
	add	rdx, 48
	cmp	rdx, rax
	jb	.LBB0_12
	jmp	.LBB0_2
.LBB0_1:
	vxorps	xmm0, xmm0, xmm0
	vxorps	xmm2, xmm2, xmm2
	vxorps	xmm3, xmm3, xmm3
	vxorps	xmm4, xmm4, xmm4
	vxorps	xmm5, xmm5, xmm5
	vxorps	xmm1, xmm1, xmm1
.LBB0_2:
	vaddps	ymm1, ymm5, ymm1
	vaddps	ymm3, ymm3, ymm4
	vaddps	ymm0, ymm0, ymm2
	vaddps	ymm0, ymm0, ymm3
	vaddps	ymm0, ymm0, ymm1
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
	cmp	rax, r9
	jae	.LBB0_10
# %bb.3:
	mov	rdi, r9
	sub	rdi, rax
	mov	r10, rax
	not	r10
	add	r10, r9
	and	rdi, 3
	je	.LBB0_7
# %bb.4:
	neg	rdi
	xor	edx, edx
	.p2align	4, 0x90
.LBB0_5:                                # =>This Inner Loop Header: Depth=1
	vmovss	xmm1, dword ptr [rsi]   # xmm1 = mem[0],zero,zero,zero
	vmulss	xmm1, xmm1, dword ptr [r8]
	vaddss	xmm0, xmm0, xmm1
	vmovss	dword ptr [rcx], xmm0
	add	rsi, 4
	add	r8, 4
	add	rdx, -1
	cmp	rdi, rdx
	jne	.LBB0_5
# %bb.6:
	sub	rax, rdx
.LBB0_7:
	cmp	r10, 3
	jb	.LBB0_10
# %bb.8:
	xor	edx, edx
	.p2align	4, 0x90
.LBB0_9:                                # =>This Inner Loop Header: Depth=1
	vmovss	xmm1, dword ptr [rsi + 4*rdx] # xmm1 = mem[0],zero,zero,zero
	vmulss	xmm1, xmm1, dword ptr [r8 + 4*rdx]
	vaddss	xmm0, xmm0, xmm1
	vmovss	dword ptr [rcx], xmm0
	vmovss	xmm1, dword ptr [rsi + 4*rdx + 4] # xmm1 = mem[0],zero,zero,zero
	vmulss	xmm1, xmm1, dword ptr [r8 + 4*rdx + 4]
	vaddss	xmm0, xmm0, xmm1
	vmovss	dword ptr [rcx], xmm0
	vmovss	xmm1, dword ptr [rsi + 4*rdx + 8] # xmm1 = mem[0],zero,zero,zero
	vmulss	xmm1, xmm1, dword ptr [r8 + 4*rdx + 8]
	vaddss	xmm0, xmm0, xmm1
	vmovss	dword ptr [rcx], xmm0
	vmovss	xmm1, dword ptr [rsi + 4*rdx + 12] # xmm1 = mem[0],zero,zero,zero
	vmulss	xmm1, xmm1, dword ptr [r8 + 4*rdx + 12]
	vaddss	xmm0, xmm0, xmm1
	vmovss	dword ptr [rcx], xmm0
	lea	rdi, [rax + rdx]
	add	rdi, 4
	add	rdx, 4
	cmp	rdi, r9
	jb	.LBB0_9
.LBB0_10:
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
	mov	r8, rdx
	mov	r9, rdi
	movabs	rdx, -6148914691236517205
	mov	rax, rdi
	mul	rdx
	shr	rdx
	and	rdx, -8
	lea	rax, [rdx + 2*rdx]
	test	sil, 31
	jne	.LBB1_6
# %bb.1:
	test	r8b, 31
	je	.LBB1_2
.LBB1_6:
	test	rax, rax
	je	.LBB1_3
# %bb.7:
	vxorpd	xmm1, xmm1, xmm1
	xor	edx, edx
	vxorpd	xmm5, xmm5, xmm5
	vxorpd	xmm4, xmm4, xmm4
	vxorpd	xmm3, xmm3, xmm3
	vxorpd	xmm2, xmm2, xmm2
	vxorpd	xmm0, xmm0, xmm0
	.p2align	4, 0x90
.LBB1_8:                                # =>This Inner Loop Header: Depth=1
	vmovupd	ymm6, ymmword ptr [rsi]
	vmovupd	ymm7, ymmword ptr [rsi + 32]
	vmovupd	ymm8, ymmword ptr [rsi + 64]
	vmovupd	ymm9, ymmword ptr [rsi + 96]
	vfmadd231pd	ymm0, ymm6, ymmword ptr [r8] # ymm0 = (ymm6 * mem) + ymm0
	vfmadd231pd	ymm2, ymm7, ymmword ptr [r8 + 32] # ymm2 = (ymm7 * mem) + ymm2
	vfmadd231pd	ymm3, ymm8, ymmword ptr [r8 + 64] # ymm3 = (ymm8 * mem) + ymm3
	vfmadd231pd	ymm4, ymm9, ymmword ptr [r8 + 96] # ymm4 = (ymm9 * mem) + ymm4
	vmovupd	ymm6, ymmword ptr [rsi + 128]
	vfmadd231pd	ymm5, ymm6, ymmword ptr [r8 + 128] # ymm5 = (ymm6 * mem) + ymm5
	vmovupd	ymm6, ymmword ptr [rsi + 160]
	vfmadd231pd	ymm1, ymm6, ymmword ptr [r8 + 160] # ymm1 = (ymm6 * mem) + ymm1
	add	rsi, 192
	add	r8, 192
	add	rdx, 24
	cmp	rdx, rax
	jb	.LBB1_8
	jmp	.LBB1_9
.LBB1_2:
	test	rax, rax
	je	.LBB1_3
# %bb.4:
	vxorpd	xmm1, xmm1, xmm1
	xor	edx, edx
	vxorpd	xmm5, xmm5, xmm5
	vxorpd	xmm4, xmm4, xmm4
	vxorpd	xmm3, xmm3, xmm3
	vxorpd	xmm2, xmm2, xmm2
	vxorpd	xmm0, xmm0, xmm0
	.p2align	4, 0x90
.LBB1_5:                                # =>This Inner Loop Header: Depth=1
	vmovapd	ymm6, ymmword ptr [rsi]
	vmovapd	ymm7, ymmword ptr [rsi + 32]
	vmovapd	ymm8, ymmword ptr [rsi + 64]
	vmovapd	ymm9, ymmword ptr [rsi + 96]
	vfmadd231pd	ymm0, ymm6, ymmword ptr [r8] # ymm0 = (ymm6 * mem) + ymm0
	vfmadd231pd	ymm2, ymm7, ymmword ptr [r8 + 32] # ymm2 = (ymm7 * mem) + ymm2
	vfmadd231pd	ymm3, ymm8, ymmword ptr [r8 + 64] # ymm3 = (ymm8 * mem) + ymm3
	vfmadd231pd	ymm4, ymm9, ymmword ptr [r8 + 96] # ymm4 = (ymm9 * mem) + ymm4
	vmovapd	ymm6, ymmword ptr [rsi + 128]
	vfmadd231pd	ymm5, ymm6, ymmword ptr [r8 + 128] # ymm5 = (ymm6 * mem) + ymm5
	vmovapd	ymm6, ymmword ptr [rsi + 160]
	vfmadd231pd	ymm1, ymm6, ymmword ptr [r8 + 160] # ymm1 = (ymm6 * mem) + ymm1
	add	rsi, 192
	add	r8, 192
	add	rdx, 24
	cmp	rdx, rax
	jb	.LBB1_5
	jmp	.LBB1_9
.LBB1_3:
	vxorpd	xmm0, xmm0, xmm0
	vxorpd	xmm2, xmm2, xmm2
	vxorpd	xmm3, xmm3, xmm3
	vxorpd	xmm4, xmm4, xmm4
	vxorpd	xmm5, xmm5, xmm5
	vxorpd	xmm1, xmm1, xmm1
.LBB1_9:
	vaddpd	ymm1, ymm5, ymm1
	vaddpd	ymm3, ymm3, ymm4
	vaddpd	ymm0, ymm0, ymm2
	vaddpd	ymm0, ymm0, ymm3
	vaddpd	ymm0, ymm0, ymm1
	vpermilpd	xmm1, xmm0, 1   # xmm1 = xmm0[1,0]
	vaddsd	xmm1, xmm0, xmm1
	vextractf128	xmm0, ymm0, 1
	vaddsd	xmm1, xmm0, xmm1
	vpermilpd	xmm0, xmm0, 1   # xmm0 = xmm0[1,0]
	vaddsd	xmm0, xmm0, xmm1
	vmovsd	qword ptr [rcx], xmm0
	cmp	rax, r9
	jae	.LBB1_17
# %bb.10:
	mov	rdi, r9
	sub	rdi, rax
	mov	r10, rax
	not	r10
	add	r10, r9
	and	rdi, 3
	je	.LBB1_14
# %bb.11:
	neg	rdi
	xor	edx, edx
	.p2align	4, 0x90
.LBB1_12:                               # =>This Inner Loop Header: Depth=1
	vmovsd	xmm1, qword ptr [rsi]   # xmm1 = mem[0],zero
	vmulsd	xmm1, xmm1, qword ptr [r8]
	vaddsd	xmm0, xmm0, xmm1
	vmovsd	qword ptr [rcx], xmm0
	add	rsi, 8
	add	r8, 8
	add	rdx, -1
	cmp	rdi, rdx
	jne	.LBB1_12
# %bb.13:
	sub	rax, rdx
.LBB1_14:
	cmp	r10, 3
	jb	.LBB1_17
# %bb.15:
	xor	edx, edx
	.p2align	4, 0x90
.LBB1_16:                               # =>This Inner Loop Header: Depth=1
	vmovsd	xmm1, qword ptr [rsi + 8*rdx] # xmm1 = mem[0],zero
	vmulsd	xmm1, xmm1, qword ptr [r8 + 8*rdx]
	vaddsd	xmm0, xmm0, xmm1
	vmovsd	qword ptr [rcx], xmm0
	vmovsd	xmm1, qword ptr [rsi + 8*rdx + 8] # xmm1 = mem[0],zero
	vmulsd	xmm1, xmm1, qword ptr [r8 + 8*rdx + 8]
	vaddsd	xmm0, xmm0, xmm1
	vmovsd	qword ptr [rcx], xmm0
	vmovsd	xmm1, qword ptr [rsi + 8*rdx + 16] # xmm1 = mem[0],zero
	vmulsd	xmm1, xmm1, qword ptr [r8 + 8*rdx + 16]
	vaddsd	xmm0, xmm0, xmm1
	vmovsd	qword ptr [rcx], xmm0
	vmovsd	xmm1, qword ptr [rsi + 8*rdx + 24] # xmm1 = mem[0],zero
	vmulsd	xmm1, xmm1, qword ptr [r8 + 8*rdx + 24]
	vaddsd	xmm0, xmm0, xmm1
	vmovsd	qword ptr [rcx], xmm0
	lea	rdi, [rax + rdx]
	add	rdi, 4
	add	rdx, 4
	cmp	rdi, r9
	jb	.LBB1_16
.LBB1_17:
	mov	rsp, rbp
	pop	rbp
	vzeroupper
	ret
.Lfunc_end1:
	.size	avxFloat64Dot, .Lfunc_end1-avxFloat64Dot
                                        # -- End function
	.globl	avxFloat32Sum           # -- Begin function avxFloat32Sum
	.p2align	4, 0x90
	.type	avxFloat32Sum,@function
avxFloat32Sum:                          # @avxFloat32Sum
# %bb.0:
	push	rbp
	mov	rbp, rsp
	and	rsp, -8
	mov	rcx, rdx
	mov	r8, rdi
	movabs	rdx, -6148914691236517205
	mov	rax, rdi
	mul	rdx
	shr	rdx
	and	rdx, -16
	lea	rax, [rdx + 2*rdx]
	test	rax, rax
	je	.LBB2_1
# %bb.9:
	vxorps	xmm1, xmm1, xmm1
	xor	edx, edx
	vxorps	xmm5, xmm5, xmm5
	vxorps	xmm4, xmm4, xmm4
	vxorps	xmm3, xmm3, xmm3
	vxorps	xmm2, xmm2, xmm2
	vxorps	xmm0, xmm0, xmm0
	.p2align	4, 0x90
.LBB2_10:                               # =>This Inner Loop Header: Depth=1
	vaddps	ymm0, ymm0, ymmword ptr [rsi]
	vaddps	ymm2, ymm2, ymmword ptr [rsi + 32]
	vaddps	ymm3, ymm3, ymmword ptr [rsi + 64]
	vaddps	ymm4, ymm4, ymmword ptr [rsi + 96]
	vaddps	ymm5, ymm5, ymmword ptr [rsi + 128]
	vaddps	ymm1, ymm1, ymmword ptr [rsi + 160]
	add	rsi, 192
	add	rdx, 48
	cmp	rdx, rax
	jb	.LBB2_10
	jmp	.LBB2_2
.LBB2_1:
	vxorps	xmm0, xmm0, xmm0
	vxorps	xmm2, xmm2, xmm2
	vxorps	xmm3, xmm3, xmm3
	vxorps	xmm4, xmm4, xmm4
	vxorps	xmm5, xmm5, xmm5
	vxorps	xmm1, xmm1, xmm1
.LBB2_2:
	vaddps	ymm1, ymm5, ymm1
	vaddps	ymm3, ymm3, ymm4
	vaddps	ymm0, ymm0, ymm2
	vaddps	ymm0, ymm0, ymm3
	vaddps	ymm0, ymm0, ymm1
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
	cmp	rax, r8
	jae	.LBB2_8
# %bb.3:
	mov	rdi, r8
	sub	rdi, rax
	mov	r9, rax
	not	r9
	add	r9, r8
	and	rdi, 3
	je	.LBB2_7
# %bb.4:
	neg	rdi
	xor	edx, edx
	.p2align	4, 0x90
.LBB2_5:                                # =>This Inner Loop Header: Depth=1
	vaddss	xmm0, xmm0, dword ptr [rsi]
	vmovss	dword ptr [rcx], xmm0
	add	rsi, 4
	add	rdx, -1
	cmp	rdi, rdx
	jne	.LBB2_5
# %bb.6:
	sub	rax, rdx
.LBB2_7:
	cmp	r9, 3
	jb	.LBB2_8
	.p2align	4, 0x90
.LBB2_11:                               # =>This Inner Loop Header: Depth=1
	vaddss	xmm0, xmm0, dword ptr [rsi]
	vmovss	dword ptr [rcx], xmm0
	vaddss	xmm0, xmm0, dword ptr [rsi + 4]
	vmovss	dword ptr [rcx], xmm0
	vaddss	xmm0, xmm0, dword ptr [rsi + 8]
	vmovss	dword ptr [rcx], xmm0
	vaddss	xmm0, xmm0, dword ptr [rsi + 12]
	vmovss	dword ptr [rcx], xmm0
	add	rax, 4
	add	rsi, 16
	cmp	rax, r8
	jb	.LBB2_11
.LBB2_8:
	mov	rsp, rbp
	pop	rbp
	vzeroupper
	ret
.Lfunc_end2:
	.size	avxFloat32Sum, .Lfunc_end2-avxFloat32Sum
                                        # -- End function
	.globl	avxFloat64Sum           # -- Begin function avxFloat64Sum
	.p2align	4, 0x90
	.type	avxFloat64Sum,@function
avxFloat64Sum:                          # @avxFloat64Sum
# %bb.0:
	push	rbp
	mov	rbp, rsp
	and	rsp, -8
	mov	r10, rdi
	and	r10, -4
	je	.LBB3_1
# %bb.2:
	lea	rcx, [r10 - 1]
	shr	rcx, 2
	lea	r8, [rcx + 2*rcx]
	lea	r9d, [rcx + 1]
	and	r9d, 1
	test	rcx, rcx
	je	.LBB3_3
# %bb.16:
	lea	rax, [r9 - 1]
	sub	rax, rcx
	vxorpd	xmm0, xmm0, xmm0
	vxorpd	xmm3, xmm3, xmm3
	vxorpd	xmm2, xmm2, xmm2
	vxorpd	xmm5, xmm5, xmm5
	vxorpd	xmm4, xmm4, xmm4
	vxorpd	xmm1, xmm1, xmm1
	mov	rcx, rsi
	.p2align	4, 0x90
.LBB3_17:                               # =>This Inner Loop Header: Depth=1
	vaddpd	ymm1, ymm1, ymmword ptr [rcx]
	vaddpd	ymm4, ymm4, ymmword ptr [rcx + 32]
	vaddpd	ymm5, ymm5, ymmword ptr [rcx + 64]
	vaddpd	ymm2, ymm2, ymmword ptr [rcx + 96]
	vaddpd	ymm3, ymm3, ymmword ptr [rcx + 128]
	vaddpd	ymm0, ymm0, ymmword ptr [rcx + 160]
	vaddpd	ymm1, ymm1, ymmword ptr [rcx + 192]
	vaddpd	ymm4, ymm4, ymmword ptr [rcx + 224]
	vaddpd	ymm5, ymm5, ymmword ptr [rcx + 256]
	vaddpd	ymm2, ymm2, ymmword ptr [rcx + 288]
	vaddpd	ymm3, ymm3, ymmword ptr [rcx + 320]
	vaddpd	ymm0, ymm0, ymmword ptr [rcx + 352]
	add	rcx, 384
	add	rax, 2
	jne	.LBB3_17
# %bb.4:
	lea	rax, [8*r8 + 24]
	test	r9, r9
	je	.LBB3_6
.LBB3_5:
	vaddpd	ymm0, ymm0, ymmword ptr [rcx + 160]
	vaddpd	ymm3, ymm3, ymmword ptr [rcx + 128]
	vaddpd	ymm2, ymm2, ymmword ptr [rcx + 96]
	vaddpd	ymm5, ymm5, ymmword ptr [rcx + 64]
	vaddpd	ymm4, ymm4, ymmword ptr [rcx + 32]
	vaddpd	ymm1, ymm1, ymmword ptr [rcx]
.LBB3_6:
	lea	rsi, [rsi + 8*rax]
	jmp	.LBB3_7
.LBB3_1:
	vxorpd	xmm1, xmm1, xmm1
	vxorpd	xmm4, xmm4, xmm4
	vxorpd	xmm5, xmm5, xmm5
	vxorpd	xmm2, xmm2, xmm2
	vxorpd	xmm3, xmm3, xmm3
	vxorpd	xmm0, xmm0, xmm0
.LBB3_7:
	vaddpd	ymm0, ymm3, ymm0
	vaddpd	ymm2, ymm5, ymm2
	vaddpd	ymm1, ymm1, ymm4
	vaddpd	ymm1, ymm1, ymm2
	vaddpd	ymm0, ymm1, ymm0
	vpermilpd	xmm1, xmm0, 1   # xmm1 = xmm0[1,0]
	vaddsd	xmm1, xmm0, xmm1
	vextractf128	xmm0, ymm0, 1
	vaddsd	xmm1, xmm0, xmm1
	vpermilpd	xmm0, xmm0, 1   # xmm0 = xmm0[1,0]
	vaddsd	xmm0, xmm0, xmm1
	vmovsd	qword ptr [rdx], xmm0
	cmp	r10, rdi
	jae	.LBB3_15
# %bb.8:
	mov	r8, r10
	not	r8
	add	r8, rdi
	mov	rcx, rdi
	and	rcx, 3
	je	.LBB3_12
# %bb.9:
	neg	rcx
	xor	eax, eax
	.p2align	4, 0x90
.LBB3_10:                               # =>This Inner Loop Header: Depth=1
	vaddsd	xmm0, xmm0, qword ptr [rsi]
	vmovsd	qword ptr [rdx], xmm0
	add	rsi, 8
	add	rax, -1
	cmp	rcx, rax
	jne	.LBB3_10
# %bb.11:
	sub	r10, rax
.LBB3_12:
	cmp	r8, 3
	jb	.LBB3_15
# %bb.13:
	sub	rdi, r10
	xor	eax, eax
	.p2align	4, 0x90
.LBB3_14:                               # =>This Inner Loop Header: Depth=1
	vaddsd	xmm0, xmm0, qword ptr [rsi + 8*rax]
	vmovsd	qword ptr [rdx], xmm0
	vaddsd	xmm0, xmm0, qword ptr [rsi + 8*rax + 8]
	vmovsd	qword ptr [rdx], xmm0
	vaddsd	xmm0, xmm0, qword ptr [rsi + 8*rax + 16]
	vmovsd	qword ptr [rdx], xmm0
	vaddsd	xmm0, xmm0, qword ptr [rsi + 8*rax + 24]
	vmovsd	qword ptr [rdx], xmm0
	add	rax, 4
	cmp	rdi, rax
	jne	.LBB3_14
.LBB3_15:
	mov	rsp, rbp
	pop	rbp
	vzeroupper
	ret
.LBB3_3:
	vxorpd	xmm0, xmm0, xmm0
	vxorpd	xmm3, xmm3, xmm3
	vxorpd	xmm2, xmm2, xmm2
	vxorpd	xmm5, xmm5, xmm5
	vxorpd	xmm4, xmm4, xmm4
	vxorpd	xmm1, xmm1, xmm1
	mov	rcx, rsi
	lea	rax, [8*r8 + 24]
	test	r9, r9
	jne	.LBB3_5
	jmp	.LBB3_6
.Lfunc_end3:
	.size	avxFloat64Sum, .Lfunc_end3-avxFloat64Sum
                                        # -- End function

	.ident	"clang version 8.0.0-3 (tags/RELEASE_800/final)"
	.section	".note.GNU-stack","",@progbits
	.addrsig
