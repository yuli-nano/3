package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"sync"
	"unsafe"
)

// CUDA handle for reducedot kernel
var reducedot_code cu.Function

// Stores the arguments for reducedot kernel invocation
type reducedot_args_t struct {
	arg_x1      unsafe.Pointer
	arg_x2      unsafe.Pointer
	arg_dst     unsafe.Pointer
	arg_initVal float32
	arg_n       int
	argptr      [5]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for reducedot kernel invocation
var reducedot_args reducedot_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	reducedot_args.argptr[0] = unsafe.Pointer(&reducedot_args.arg_x1)
	reducedot_args.argptr[1] = unsafe.Pointer(&reducedot_args.arg_x2)
	reducedot_args.argptr[2] = unsafe.Pointer(&reducedot_args.arg_dst)
	reducedot_args.argptr[3] = unsafe.Pointer(&reducedot_args.arg_initVal)
	reducedot_args.argptr[4] = unsafe.Pointer(&reducedot_args.arg_n)
}

// Wrapper for reducedot CUDA kernel, asynchronous.
func k_reducedot_async(x1 unsafe.Pointer, x2 unsafe.Pointer, dst unsafe.Pointer, initVal float32, n int, cfg *config) {
	if Synchronous { // debug
		Sync()
	}

	reducedot_args.Lock()
	defer reducedot_args.Unlock()

	if reducedot_code == 0 {
		reducedot_code = fatbinLoad(reducedot_map, "reducedot")
	}

	reducedot_args.arg_x1 = x1
	reducedot_args.arg_x2 = x2
	reducedot_args.arg_dst = dst
	reducedot_args.arg_initVal = initVal
	reducedot_args.arg_n = n

	args := reducedot_args.argptr[:]
	cu.LaunchKernel(reducedot_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
	}
}

// maps compute capability on PTX code for reducedot kernel.
var reducedot_map = map[int]string{0: "",
	20: reducedot_ptx_20,
	30: reducedot_ptx_30,
	35: reducedot_ptx_35}

// reducedot PTX code for various compute capabilities.
const (
	reducedot_ptx_20 = `
.version 3.2
.target sm_20
.address_size 64


.visible .entry reducedot(
	.param .u64 reducedot_param_0,
	.param .u64 reducedot_param_1,
	.param .u64 reducedot_param_2,
	.param .f32 reducedot_param_3,
	.param .u32 reducedot_param_4
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<15>;
	.reg .f32 	%f<31>;
	.reg .s64 	%rd<16>;
	// demoted variable
	.shared .align 4 .b8 reducedot$__cuda_local_var_33857_35_non_const_sdata[2048];

	ld.param.u64 	%rd5, [reducedot_param_0];
	ld.param.u64 	%rd6, [reducedot_param_1];
	ld.param.u64 	%rd7, [reducedot_param_2];
	ld.param.f32 	%f30, [reducedot_param_3];
	ld.param.u32 	%r9, [reducedot_param_4];
	cvta.to.global.u64 	%rd1, %rd7;
	cvta.to.global.u64 	%rd2, %rd6;
	cvta.to.global.u64 	%rd3, %rd5;
	.loc 1 9 1
	mov.u32 	%r14, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r13, %r14, %r10, %r2;
	mov.u32 	%r11, %nctaid.x;
	mul.lo.s32 	%r4, %r14, %r11;
	.loc 1 9 1
	setp.ge.s32	%p1, %r13, %r9;
	@%p1 bra 	BB0_2;

BB0_1:
	.loc 1 9 1
	mul.wide.s32 	%rd8, %r13, 4;
	add.s64 	%rd9, %rd3, %rd8;
	add.s64 	%rd10, %rd2, %rd8;
	ld.global.f32 	%f5, [%rd10];
	ld.global.f32 	%f6, [%rd9];
	fma.rn.f32 	%f30, %f6, %f5, %f30;
	add.s32 	%r13, %r13, %r4;
	.loc 1 9 1
	setp.lt.s32	%p2, %r13, %r9;
	@%p2 bra 	BB0_1;

BB0_2:
	.loc 1 9 1
	mul.wide.s32 	%rd11, %r2, 4;
	mov.u64 	%rd12, reducedot$__cuda_local_var_33857_35_non_const_sdata;
	add.s64 	%rd4, %rd12, %rd11;
	st.shared.f32 	[%rd4], %f30;
	bar.sync 	0;
	.loc 1 9 1
	setp.lt.u32	%p3, %r14, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	.loc 1 9 1
	mov.u32 	%r7, %r14;
	shr.u32 	%r14, %r7, 1;
	.loc 1 9 1
	setp.ge.u32	%p4, %r2, %r14;
	@%p4 bra 	BB0_5;

	.loc 1 9 1
	ld.shared.f32 	%f7, [%rd4];
	add.s32 	%r12, %r14, %r2;
	mul.wide.u32 	%rd13, %r12, 4;
	add.s64 	%rd15, %rd12, %rd13;
	ld.shared.f32 	%f8, [%rd15];
	add.f32 	%f9, %f7, %f8;
	st.shared.f32 	[%rd4], %f9;

BB0_5:
	.loc 1 9 1
	bar.sync 	0;
	.loc 1 9 1
	setp.gt.u32	%p5, %r7, 131;
	@%p5 bra 	BB0_3;

BB0_6:
	.loc 1 9 1
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	.loc 1 9 1
	ld.volatile.shared.f32 	%f10, [%rd4];
	ld.volatile.shared.f32 	%f11, [%rd4+128];
	add.f32 	%f12, %f10, %f11;
	st.volatile.shared.f32 	[%rd4], %f12;
	ld.volatile.shared.f32 	%f13, [%rd4+64];
	ld.volatile.shared.f32 	%f14, [%rd4];
	add.f32 	%f15, %f14, %f13;
	st.volatile.shared.f32 	[%rd4], %f15;
	ld.volatile.shared.f32 	%f16, [%rd4+32];
	ld.volatile.shared.f32 	%f17, [%rd4];
	add.f32 	%f18, %f17, %f16;
	st.volatile.shared.f32 	[%rd4], %f18;
	ld.volatile.shared.f32 	%f19, [%rd4+16];
	ld.volatile.shared.f32 	%f20, [%rd4];
	add.f32 	%f21, %f20, %f19;
	st.volatile.shared.f32 	[%rd4], %f21;
	ld.volatile.shared.f32 	%f22, [%rd4+8];
	ld.volatile.shared.f32 	%f23, [%rd4];
	add.f32 	%f24, %f23, %f22;
	st.volatile.shared.f32 	[%rd4], %f24;
	ld.volatile.shared.f32 	%f25, [%rd4+4];
	ld.volatile.shared.f32 	%f26, [%rd4];
	add.f32 	%f27, %f26, %f25;
	st.volatile.shared.f32 	[%rd4], %f27;

BB0_8:
	.loc 1 9 1
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	.loc 1 9 1
	ld.shared.f32 	%f28, [reducedot$__cuda_local_var_33857_35_non_const_sdata];
	.loc 2 3725 3
	atom.global.add.f32 	%f29, [%rd1], %f28;

BB0_10:
	.loc 1 10 2
	ret;
}


`
	reducedot_ptx_30 = `
.version 3.2
.target sm_30
.address_size 64


.visible .entry reducedot(
	.param .u64 reducedot_param_0,
	.param .u64 reducedot_param_1,
	.param .u64 reducedot_param_2,
	.param .f32 reducedot_param_3,
	.param .u32 reducedot_param_4
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<15>;
	.reg .f32 	%f<31>;
	.reg .s64 	%rd<16>;
	// demoted variable
	.shared .align 4 .b8 reducedot$__cuda_local_var_33930_35_non_const_sdata[2048];

	ld.param.u64 	%rd5, [reducedot_param_0];
	ld.param.u64 	%rd6, [reducedot_param_1];
	ld.param.u64 	%rd7, [reducedot_param_2];
	ld.param.f32 	%f30, [reducedot_param_3];
	ld.param.u32 	%r9, [reducedot_param_4];
	cvta.to.global.u64 	%rd1, %rd7;
	cvta.to.global.u64 	%rd2, %rd6;
	cvta.to.global.u64 	%rd3, %rd5;
	.loc 1 9 1
	mov.u32 	%r14, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r13, %r14, %r10, %r2;
	mov.u32 	%r11, %nctaid.x;
	mul.lo.s32 	%r4, %r14, %r11;
	.loc 1 9 1
	setp.ge.s32	%p1, %r13, %r9;
	@%p1 bra 	BB0_2;

BB0_1:
	.loc 1 9 1
	mul.wide.s32 	%rd8, %r13, 4;
	add.s64 	%rd9, %rd3, %rd8;
	add.s64 	%rd10, %rd2, %rd8;
	ld.global.f32 	%f5, [%rd10];
	ld.global.f32 	%f6, [%rd9];
	fma.rn.f32 	%f30, %f6, %f5, %f30;
	add.s32 	%r13, %r13, %r4;
	.loc 1 9 1
	setp.lt.s32	%p2, %r13, %r9;
	@%p2 bra 	BB0_1;

BB0_2:
	.loc 1 9 1
	mul.wide.s32 	%rd11, %r2, 4;
	mov.u64 	%rd12, reducedot$__cuda_local_var_33930_35_non_const_sdata;
	add.s64 	%rd4, %rd12, %rd11;
	st.shared.f32 	[%rd4], %f30;
	bar.sync 	0;
	.loc 1 9 1
	setp.lt.u32	%p3, %r14, 66;
	@%p3 bra 	BB0_6;

BB0_3:
	.loc 1 9 1
	mov.u32 	%r7, %r14;
	shr.u32 	%r14, %r7, 1;
	.loc 1 9 1
	setp.ge.u32	%p4, %r2, %r14;
	@%p4 bra 	BB0_5;

	.loc 1 9 1
	ld.shared.f32 	%f7, [%rd4];
	add.s32 	%r12, %r14, %r2;
	mul.wide.u32 	%rd13, %r12, 4;
	add.s64 	%rd15, %rd12, %rd13;
	ld.shared.f32 	%f8, [%rd15];
	add.f32 	%f9, %f7, %f8;
	st.shared.f32 	[%rd4], %f9;

BB0_5:
	.loc 1 9 1
	bar.sync 	0;
	.loc 1 9 1
	setp.gt.u32	%p5, %r7, 131;
	@%p5 bra 	BB0_3;

BB0_6:
	.loc 1 9 1
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB0_8;

	.loc 1 9 1
	ld.volatile.shared.f32 	%f10, [%rd4];
	ld.volatile.shared.f32 	%f11, [%rd4+128];
	add.f32 	%f12, %f10, %f11;
	st.volatile.shared.f32 	[%rd4], %f12;
	ld.volatile.shared.f32 	%f13, [%rd4+64];
	ld.volatile.shared.f32 	%f14, [%rd4];
	add.f32 	%f15, %f14, %f13;
	st.volatile.shared.f32 	[%rd4], %f15;
	ld.volatile.shared.f32 	%f16, [%rd4+32];
	ld.volatile.shared.f32 	%f17, [%rd4];
	add.f32 	%f18, %f17, %f16;
	st.volatile.shared.f32 	[%rd4], %f18;
	ld.volatile.shared.f32 	%f19, [%rd4+16];
	ld.volatile.shared.f32 	%f20, [%rd4];
	add.f32 	%f21, %f20, %f19;
	st.volatile.shared.f32 	[%rd4], %f21;
	ld.volatile.shared.f32 	%f22, [%rd4+8];
	ld.volatile.shared.f32 	%f23, [%rd4];
	add.f32 	%f24, %f23, %f22;
	st.volatile.shared.f32 	[%rd4], %f24;
	ld.volatile.shared.f32 	%f25, [%rd4+4];
	ld.volatile.shared.f32 	%f26, [%rd4];
	add.f32 	%f27, %f26, %f25;
	st.volatile.shared.f32 	[%rd4], %f27;

BB0_8:
	.loc 1 9 1
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB0_10;

	.loc 1 9 1
	ld.shared.f32 	%f28, [reducedot$__cuda_local_var_33930_35_non_const_sdata];
	.loc 2 3725 3
	atom.global.add.f32 	%f29, [%rd1], %f28;

BB0_10:
	.loc 1 10 2
	ret;
}


`
	reducedot_ptx_35 = `
.version 3.2
.target sm_35
.address_size 64


.weak .func  (.param .b32 func_retval0) cudaMalloc(
	.param .b64 cudaMalloc_param_0,
	.param .b64 cudaMalloc_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 66 3
	ret;
}

.weak .func  (.param .b32 func_retval0) cudaFuncGetAttributes(
	.param .b64 cudaFuncGetAttributes_param_0,
	.param .b64 cudaFuncGetAttributes_param_1
)
{
	.reg .s32 	%r<2>;


	mov.u32 	%r1, 30;
	st.param.b32	[func_retval0+0], %r1;
	.loc 2 71 3
	ret;
}

.visible .entry reducedot(
	.param .u64 reducedot_param_0,
	.param .u64 reducedot_param_1,
	.param .u64 reducedot_param_2,
	.param .f32 reducedot_param_3,
	.param .u32 reducedot_param_4
)
{
	.reg .pred 	%p<8>;
	.reg .s32 	%r<15>;
	.reg .f32 	%f<31>;
	.reg .s64 	%rd<16>;
	// demoted variable
	.shared .align 4 .b8 reducedot$__cuda_local_var_34093_35_non_const_sdata[2048];

	ld.param.u64 	%rd5, [reducedot_param_0];
	ld.param.u64 	%rd6, [reducedot_param_1];
	ld.param.u64 	%rd7, [reducedot_param_2];
	ld.param.f32 	%f30, [reducedot_param_3];
	ld.param.u32 	%r9, [reducedot_param_4];
	cvta.to.global.u64 	%rd1, %rd7;
	cvta.to.global.u64 	%rd2, %rd6;
	cvta.to.global.u64 	%rd3, %rd5;
	.loc 1 9 1
	mov.u32 	%r14, %ntid.x;
	mov.u32 	%r10, %ctaid.x;
	mov.u32 	%r2, %tid.x;
	mad.lo.s32 	%r13, %r14, %r10, %r2;
	mov.u32 	%r11, %nctaid.x;
	mul.lo.s32 	%r4, %r14, %r11;
	.loc 1 9 1
	setp.ge.s32	%p1, %r13, %r9;
	@%p1 bra 	BB2_2;

BB2_1:
	.loc 1 9 1
	mul.wide.s32 	%rd8, %r13, 4;
	add.s64 	%rd9, %rd3, %rd8;
	add.s64 	%rd10, %rd2, %rd8;
	ld.global.nc.f32 	%f5, [%rd10];
	ld.global.nc.f32 	%f6, [%rd9];
	fma.rn.f32 	%f30, %f6, %f5, %f30;
	add.s32 	%r13, %r13, %r4;
	.loc 1 9 1
	setp.lt.s32	%p2, %r13, %r9;
	@%p2 bra 	BB2_1;

BB2_2:
	.loc 1 9 1
	mul.wide.s32 	%rd11, %r2, 4;
	mov.u64 	%rd12, reducedot$__cuda_local_var_34093_35_non_const_sdata;
	add.s64 	%rd4, %rd12, %rd11;
	st.shared.f32 	[%rd4], %f30;
	bar.sync 	0;
	.loc 1 9 1
	setp.lt.u32	%p3, %r14, 66;
	@%p3 bra 	BB2_6;

BB2_3:
	.loc 1 9 1
	mov.u32 	%r7, %r14;
	shr.u32 	%r14, %r7, 1;
	.loc 1 9 1
	setp.ge.u32	%p4, %r2, %r14;
	@%p4 bra 	BB2_5;

	.loc 1 9 1
	ld.shared.f32 	%f7, [%rd4];
	add.s32 	%r12, %r14, %r2;
	mul.wide.u32 	%rd13, %r12, 4;
	add.s64 	%rd15, %rd12, %rd13;
	ld.shared.f32 	%f8, [%rd15];
	add.f32 	%f9, %f7, %f8;
	st.shared.f32 	[%rd4], %f9;

BB2_5:
	.loc 1 9 1
	bar.sync 	0;
	.loc 1 9 1
	setp.gt.u32	%p5, %r7, 131;
	@%p5 bra 	BB2_3;

BB2_6:
	.loc 1 9 1
	setp.gt.s32	%p6, %r2, 31;
	@%p6 bra 	BB2_8;

	.loc 1 9 1
	ld.volatile.shared.f32 	%f10, [%rd4];
	ld.volatile.shared.f32 	%f11, [%rd4+128];
	add.f32 	%f12, %f10, %f11;
	st.volatile.shared.f32 	[%rd4], %f12;
	ld.volatile.shared.f32 	%f13, [%rd4+64];
	ld.volatile.shared.f32 	%f14, [%rd4];
	add.f32 	%f15, %f14, %f13;
	st.volatile.shared.f32 	[%rd4], %f15;
	ld.volatile.shared.f32 	%f16, [%rd4+32];
	ld.volatile.shared.f32 	%f17, [%rd4];
	add.f32 	%f18, %f17, %f16;
	st.volatile.shared.f32 	[%rd4], %f18;
	ld.volatile.shared.f32 	%f19, [%rd4+16];
	ld.volatile.shared.f32 	%f20, [%rd4];
	add.f32 	%f21, %f20, %f19;
	st.volatile.shared.f32 	[%rd4], %f21;
	ld.volatile.shared.f32 	%f22, [%rd4+8];
	ld.volatile.shared.f32 	%f23, [%rd4];
	add.f32 	%f24, %f23, %f22;
	st.volatile.shared.f32 	[%rd4], %f24;
	ld.volatile.shared.f32 	%f25, [%rd4+4];
	ld.volatile.shared.f32 	%f26, [%rd4];
	add.f32 	%f27, %f26, %f25;
	st.volatile.shared.f32 	[%rd4], %f27;

BB2_8:
	.loc 1 9 1
	setp.ne.s32	%p7, %r2, 0;
	@%p7 bra 	BB2_10;

	.loc 1 9 1
	ld.shared.f32 	%f28, [reducedot$__cuda_local_var_34093_35_non_const_sdata];
	.loc 3 3725 3
	atom.global.add.f32 	%f29, [%rd1], %f28;

BB2_10:
	.loc 1 10 2
	ret;
}


`
)
