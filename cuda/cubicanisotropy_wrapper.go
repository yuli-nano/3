package cuda

/*
 THIS FILE IS AUTO-GENERATED BY CUDA2GO.
 EDITING IS FUTILE.
*/

import (
	"github.com/barnex/cuda5/cu"
	"unsafe"
)

var addcubicanisotropy_code cu.Function

type addcubicanisotropy_args struct {
	arg_Bx      unsafe.Pointer
	arg_By      unsafe.Pointer
	arg_Bz      unsafe.Pointer
	arg_mx      unsafe.Pointer
	arg_my      unsafe.Pointer
	arg_mz      unsafe.Pointer
	arg_K1LUT   unsafe.Pointer
	arg_C1xLUT  unsafe.Pointer
	arg_C1yLUT  unsafe.Pointer
	arg_C1zLUT  unsafe.Pointer
	arg_C2xLUT  unsafe.Pointer
	arg_C2yLUT  unsafe.Pointer
	arg_C2zLUT  unsafe.Pointer
	arg_regions unsafe.Pointer
	arg_N       int
	argptr      [15]unsafe.Pointer
}

// Wrapper for addcubicanisotropy CUDA kernel, asynchronous.
func k_addcubicanisotropy_async(Bx unsafe.Pointer, By unsafe.Pointer, Bz unsafe.Pointer, mx unsafe.Pointer, my unsafe.Pointer, mz unsafe.Pointer, K1LUT unsafe.Pointer, C1xLUT unsafe.Pointer, C1yLUT unsafe.Pointer, C1zLUT unsafe.Pointer, C2xLUT unsafe.Pointer, C2yLUT unsafe.Pointer, C2zLUT unsafe.Pointer, regions unsafe.Pointer, N int, cfg *config, str int) {
	if addcubicanisotropy_code == 0 {
		addcubicanisotropy_code = fatbinLoad(addcubicanisotropy_map, "addcubicanisotropy")
	}

	var _a_ addcubicanisotropy_args

	_a_.arg_Bx = Bx
	_a_.argptr[0] = unsafe.Pointer(&_a_.arg_Bx)
	_a_.arg_By = By
	_a_.argptr[1] = unsafe.Pointer(&_a_.arg_By)
	_a_.arg_Bz = Bz
	_a_.argptr[2] = unsafe.Pointer(&_a_.arg_Bz)
	_a_.arg_mx = mx
	_a_.argptr[3] = unsafe.Pointer(&_a_.arg_mx)
	_a_.arg_my = my
	_a_.argptr[4] = unsafe.Pointer(&_a_.arg_my)
	_a_.arg_mz = mz
	_a_.argptr[5] = unsafe.Pointer(&_a_.arg_mz)
	_a_.arg_K1LUT = K1LUT
	_a_.argptr[6] = unsafe.Pointer(&_a_.arg_K1LUT)
	_a_.arg_C1xLUT = C1xLUT
	_a_.argptr[7] = unsafe.Pointer(&_a_.arg_C1xLUT)
	_a_.arg_C1yLUT = C1yLUT
	_a_.argptr[8] = unsafe.Pointer(&_a_.arg_C1yLUT)
	_a_.arg_C1zLUT = C1zLUT
	_a_.argptr[9] = unsafe.Pointer(&_a_.arg_C1zLUT)
	_a_.arg_C2xLUT = C2xLUT
	_a_.argptr[10] = unsafe.Pointer(&_a_.arg_C2xLUT)
	_a_.arg_C2yLUT = C2yLUT
	_a_.argptr[11] = unsafe.Pointer(&_a_.arg_C2yLUT)
	_a_.arg_C2zLUT = C2zLUT
	_a_.argptr[12] = unsafe.Pointer(&_a_.arg_C2zLUT)
	_a_.arg_regions = regions
	_a_.argptr[13] = unsafe.Pointer(&_a_.arg_regions)
	_a_.arg_N = N
	_a_.argptr[14] = unsafe.Pointer(&_a_.arg_N)

	args := _a_.argptr[:]
	cu.LaunchKernel(addcubicanisotropy_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream[str], args)
}

// Wrapper for addcubicanisotropy CUDA kernel, synchronized.
func k_addcubicanisotropy(Bx unsafe.Pointer, By unsafe.Pointer, Bz unsafe.Pointer, mx unsafe.Pointer, my unsafe.Pointer, mz unsafe.Pointer, K1LUT unsafe.Pointer, C1xLUT unsafe.Pointer, C1yLUT unsafe.Pointer, C1zLUT unsafe.Pointer, C2xLUT unsafe.Pointer, C2yLUT unsafe.Pointer, C2zLUT unsafe.Pointer, regions unsafe.Pointer, N int, cfg *config) {
	const stream = 0
	k_addcubicanisotropy_async(Bx, By, Bz, mx, my, mz, K1LUT, C1xLUT, C1yLUT, C1zLUT, C2xLUT, C2yLUT, C2zLUT, regions, N, cfg, stream)
	Sync(stream)
}

var addcubicanisotropy_map = map[int]string{0: "",
	20: addcubicanisotropy_ptx_20,
	30: addcubicanisotropy_ptx_30,
	35: addcubicanisotropy_ptx_35}

const (
	addcubicanisotropy_ptx_20 = `
.version 3.2
.target sm_20
.address_size 64


.visible .entry addcubicanisotropy(
	.param .u64 addcubicanisotropy_param_0,
	.param .u64 addcubicanisotropy_param_1,
	.param .u64 addcubicanisotropy_param_2,
	.param .u64 addcubicanisotropy_param_3,
	.param .u64 addcubicanisotropy_param_4,
	.param .u64 addcubicanisotropy_param_5,
	.param .u64 addcubicanisotropy_param_6,
	.param .u64 addcubicanisotropy_param_7,
	.param .u64 addcubicanisotropy_param_8,
	.param .u64 addcubicanisotropy_param_9,
	.param .u64 addcubicanisotropy_param_10,
	.param .u64 addcubicanisotropy_param_11,
	.param .u64 addcubicanisotropy_param_12,
	.param .u64 addcubicanisotropy_param_13,
	.param .u32 addcubicanisotropy_param_14
)
{
	.reg .pred 	%p<4>;
	.reg .s16 	%rs<2>;
	.reg .s32 	%r<11>;
	.reg .f32 	%f<78>;
	.reg .s64 	%rd<49>;


	ld.param.u64 	%rd3, [addcubicanisotropy_param_0];
	ld.param.u64 	%rd4, [addcubicanisotropy_param_1];
	ld.param.u64 	%rd15, [addcubicanisotropy_param_2];
	ld.param.u64 	%rd5, [addcubicanisotropy_param_3];
	ld.param.u64 	%rd6, [addcubicanisotropy_param_4];
	ld.param.u64 	%rd7, [addcubicanisotropy_param_5];
	ld.param.u64 	%rd8, [addcubicanisotropy_param_6];
	ld.param.u64 	%rd9, [addcubicanisotropy_param_7];
	ld.param.u64 	%rd10, [addcubicanisotropy_param_8];
	ld.param.u64 	%rd11, [addcubicanisotropy_param_9];
	ld.param.u64 	%rd12, [addcubicanisotropy_param_10];
	ld.param.u64 	%rd13, [addcubicanisotropy_param_11];
	ld.param.u64 	%rd14, [addcubicanisotropy_param_12];
	ld.param.u64 	%rd16, [addcubicanisotropy_param_13];
	ld.param.u32 	%r3, [addcubicanisotropy_param_14];
	cvta.to.global.u64 	%rd1, %rd15;
	cvta.to.global.u64 	%rd2, %rd16;
	.loc 1 18 1
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.y;
	mov.u32 	%r6, %ctaid.x;
	mad.lo.s32 	%r7, %r4, %r5, %r6;
	mov.u32 	%r8, %ntid.x;
	mov.u32 	%r9, %tid.x;
	mad.lo.s32 	%r1, %r7, %r8, %r9;
	.loc 1 19 1
	setp.ge.s32	%p1, %r1, %r3;
	@%p1 bra 	BB0_8;

	cvt.s64.s32	%rd17, %r1;
	add.s64 	%rd18, %rd2, %rd17;
	.loc 1 21 1
	ld.global.u8 	%rs1, [%rd18];
	cvt.u32.u16	%r10, %rs1;
	cvt.s32.s8 	%r2, %r10;
	cvta.to.global.u64 	%rd19, %rd8;
	cvt.u64.u16	%rd20, %rs1;
	cvt.s64.s8 	%rd21, %rd20;
	shl.b64 	%rd22, %rd21, 2;
	add.s64 	%rd23, %rd19, %rd22;
	.loc 1 22 1
	ld.global.f32 	%f1, [%rd23];
	cvta.to.global.u64 	%rd24, %rd9;
	add.s64 	%rd25, %rd24, %rd22;
	cvta.to.global.u64 	%rd26, %rd10;
	add.s64 	%rd27, %rd26, %rd22;
	cvta.to.global.u64 	%rd28, %rd11;
	add.s64 	%rd29, %rd28, %rd22;
	.loc 1 23 1
	ld.global.f32 	%f2, [%rd25];
	ld.global.f32 	%f3, [%rd27];
	mul.f32 	%f17, %f3, %f3;
	fma.rn.f32 	%f18, %f2, %f2, %f17;
	ld.global.f32 	%f4, [%rd29];
	fma.rn.f32 	%f19, %f4, %f4, %f18;
	.loc 2 3055 10
	sqrt.rn.f32 	%f5, %f19;
	setp.neu.f32	%p2, %f5, 0f00000000;
	@%p2 bra 	BB0_3;

	mov.f32 	%f76, 0f00000000;
	bra.uni 	BB0_4;

BB0_3:
	rcp.rn.f32 	%f76, %f5;

BB0_4:
	mul.f32 	%f8, %f76, %f2;
	mul.f32 	%f9, %f76, %f3;
	mul.f32 	%f10, %f76, %f4;
	cvta.to.global.u64 	%rd30, %rd12;
	mul.wide.s32 	%rd31, %r2, 4;
	add.s64 	%rd32, %rd30, %rd31;
	cvta.to.global.u64 	%rd33, %rd13;
	add.s64 	%rd34, %rd33, %rd31;
	cvta.to.global.u64 	%rd35, %rd14;
	add.s64 	%rd36, %rd35, %rd31;
	.loc 1 24 1
	ld.global.f32 	%f11, [%rd32];
	ld.global.f32 	%f12, [%rd34];
	mul.f32 	%f21, %f12, %f12;
	fma.rn.f32 	%f22, %f11, %f11, %f21;
	ld.global.f32 	%f13, [%rd36];
	fma.rn.f32 	%f23, %f13, %f13, %f22;
	.loc 2 3055 10
	sqrt.rn.f32 	%f14, %f23;
	setp.neu.f32	%p3, %f14, 0f00000000;
	@%p3 bra 	BB0_6;

	mov.f32 	%f77, 0f00000000;
	bra.uni 	BB0_7;

BB0_6:
	rcp.rn.f32 	%f77, %f14;

BB0_7:
	mul.f32 	%f25, %f77, %f13;
	mul.f32 	%f26, %f77, %f12;
	.loc 1 25 1
	mul.f32 	%f27, %f10, %f26;
	mul.f32 	%f28, %f25, %f9;
	sub.f32 	%f29, %f27, %f28;
	mul.f32 	%f30, %f77, %f11;
	.loc 1 25 1
	mul.f32 	%f31, %f8, %f25;
	mul.f32 	%f32, %f30, %f10;
	sub.f32 	%f33, %f31, %f32;
	mul.f32 	%f34, %f9, %f30;
	mul.f32 	%f35, %f26, %f8;
	sub.f32 	%f36, %f34, %f35;
	cvta.to.global.u64 	%rd37, %rd5;
	mul.wide.s32 	%rd38, %r1, 4;
	add.s64 	%rd39, %rd37, %rd38;
	cvta.to.global.u64 	%rd40, %rd6;
	add.s64 	%rd41, %rd40, %rd38;
	cvta.to.global.u64 	%rd42, %rd7;
	add.s64 	%rd43, %rd42, %rd38;
	.loc 1 26 1
	ld.global.f32 	%f37, [%rd39];
	ld.global.f32 	%f38, [%rd41];
	.loc 1 28 1
	mul.f32 	%f39, %f9, %f38;
	fma.rn.f32 	%f40, %f8, %f37, %f39;
	.loc 1 26 1
	ld.global.f32 	%f41, [%rd43];
	.loc 1 28 1
	fma.rn.f32 	%f42, %f10, %f41, %f40;
	.loc 1 29 1
	mul.f32 	%f43, %f26, %f38;
	fma.rn.f32 	%f44, %f30, %f37, %f43;
	fma.rn.f32 	%f45, %f25, %f41, %f44;
	.loc 1 30 1
	mul.f32 	%f46, %f33, %f38;
	fma.rn.f32 	%f47, %f29, %f37, %f46;
	fma.rn.f32 	%f48, %f36, %f41, %f47;
	.loc 1 32 1
	mul.f32 	%f49, %f45, %f45;
	mul.f32 	%f50, %f48, %f48;
	add.f32 	%f51, %f49, %f50;
	mul.f32 	%f52, %f42, %f51;
	fma.rn.f32 	%f53, %f42, %f42, %f50;
	mul.f32 	%f54, %f45, %f53;
	fma.rn.f32 	%f55, %f42, %f42, %f49;
	mul.f32 	%f56, %f48, %f55;
	mul.f32 	%f57, %f1, 0fC0000000;
	mul.f32 	%f58, %f57, %f52;
	mul.f32 	%f59, %f57, %f54;
	mul.f32 	%f60, %f57, %f56;
	.loc 1 34 1
	mul.f32 	%f61, %f59, %f30;
	fma.rn.f32 	%f62, %f58, %f8, %f61;
	fma.rn.f32 	%f63, %f60, %f29, %f62;
	cvta.to.global.u64 	%rd44, %rd3;
	add.s64 	%rd45, %rd44, %rd38;
	.loc 1 34 1
	ld.global.f32 	%f64, [%rd45];
	add.f32 	%f65, %f64, %f63;
	st.global.f32 	[%rd45], %f65;
	.loc 1 35 1
	mul.f32 	%f66, %f59, %f26;
	fma.rn.f32 	%f67, %f58, %f9, %f66;
	fma.rn.f32 	%f68, %f60, %f33, %f67;
	cvta.to.global.u64 	%rd46, %rd4;
	add.s64 	%rd47, %rd46, %rd38;
	.loc 1 35 1
	ld.global.f32 	%f69, [%rd47];
	add.f32 	%f70, %f69, %f68;
	st.global.f32 	[%rd47], %f70;
	.loc 1 36 1
	mul.f32 	%f71, %f59, %f25;
	fma.rn.f32 	%f72, %f58, %f10, %f71;
	fma.rn.f32 	%f73, %f60, %f36, %f72;
	add.s64 	%rd48, %rd1, %rd38;
	.loc 1 36 1
	ld.global.f32 	%f74, [%rd48];
	add.f32 	%f75, %f74, %f73;
	st.global.f32 	[%rd48], %f75;

BB0_8:
	.loc 1 38 2
	ret;
}


`
	addcubicanisotropy_ptx_30 = `
.version 3.2
.target sm_30
.address_size 64


.visible .entry addcubicanisotropy(
	.param .u64 addcubicanisotropy_param_0,
	.param .u64 addcubicanisotropy_param_1,
	.param .u64 addcubicanisotropy_param_2,
	.param .u64 addcubicanisotropy_param_3,
	.param .u64 addcubicanisotropy_param_4,
	.param .u64 addcubicanisotropy_param_5,
	.param .u64 addcubicanisotropy_param_6,
	.param .u64 addcubicanisotropy_param_7,
	.param .u64 addcubicanisotropy_param_8,
	.param .u64 addcubicanisotropy_param_9,
	.param .u64 addcubicanisotropy_param_10,
	.param .u64 addcubicanisotropy_param_11,
	.param .u64 addcubicanisotropy_param_12,
	.param .u64 addcubicanisotropy_param_13,
	.param .u32 addcubicanisotropy_param_14
)
{
	.reg .pred 	%p<4>;
	.reg .s16 	%rs<2>;
	.reg .s32 	%r<11>;
	.reg .f32 	%f<78>;
	.reg .s64 	%rd<49>;


	ld.param.u64 	%rd3, [addcubicanisotropy_param_0];
	ld.param.u64 	%rd4, [addcubicanisotropy_param_1];
	ld.param.u64 	%rd15, [addcubicanisotropy_param_2];
	ld.param.u64 	%rd5, [addcubicanisotropy_param_3];
	ld.param.u64 	%rd6, [addcubicanisotropy_param_4];
	ld.param.u64 	%rd7, [addcubicanisotropy_param_5];
	ld.param.u64 	%rd8, [addcubicanisotropy_param_6];
	ld.param.u64 	%rd9, [addcubicanisotropy_param_7];
	ld.param.u64 	%rd10, [addcubicanisotropy_param_8];
	ld.param.u64 	%rd11, [addcubicanisotropy_param_9];
	ld.param.u64 	%rd12, [addcubicanisotropy_param_10];
	ld.param.u64 	%rd13, [addcubicanisotropy_param_11];
	ld.param.u64 	%rd14, [addcubicanisotropy_param_12];
	ld.param.u64 	%rd16, [addcubicanisotropy_param_13];
	ld.param.u32 	%r3, [addcubicanisotropy_param_14];
	cvta.to.global.u64 	%rd1, %rd15;
	cvta.to.global.u64 	%rd2, %rd16;
	.loc 1 18 1
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.y;
	mov.u32 	%r6, %ctaid.x;
	mad.lo.s32 	%r7, %r4, %r5, %r6;
	mov.u32 	%r8, %ntid.x;
	mov.u32 	%r9, %tid.x;
	mad.lo.s32 	%r1, %r7, %r8, %r9;
	.loc 1 19 1
	setp.ge.s32	%p1, %r1, %r3;
	@%p1 bra 	BB0_8;

	cvt.s64.s32	%rd17, %r1;
	add.s64 	%rd18, %rd2, %rd17;
	.loc 1 21 1
	ld.global.u8 	%rs1, [%rd18];
	cvt.u32.u16	%r10, %rs1;
	cvt.s32.s8 	%r2, %r10;
	cvta.to.global.u64 	%rd19, %rd8;
	cvt.u64.u16	%rd20, %rs1;
	cvt.s64.s8 	%rd21, %rd20;
	shl.b64 	%rd22, %rd21, 2;
	add.s64 	%rd23, %rd19, %rd22;
	.loc 1 22 1
	ld.global.f32 	%f1, [%rd23];
	cvta.to.global.u64 	%rd24, %rd9;
	add.s64 	%rd25, %rd24, %rd22;
	cvta.to.global.u64 	%rd26, %rd10;
	add.s64 	%rd27, %rd26, %rd22;
	cvta.to.global.u64 	%rd28, %rd11;
	add.s64 	%rd29, %rd28, %rd22;
	.loc 1 23 1
	ld.global.f32 	%f2, [%rd25];
	ld.global.f32 	%f3, [%rd27];
	mul.f32 	%f17, %f3, %f3;
	fma.rn.f32 	%f18, %f2, %f2, %f17;
	ld.global.f32 	%f4, [%rd29];
	fma.rn.f32 	%f19, %f4, %f4, %f18;
	.loc 2 3055 10
	sqrt.rn.f32 	%f5, %f19;
	setp.neu.f32	%p2, %f5, 0f00000000;
	@%p2 bra 	BB0_3;

	mov.f32 	%f76, 0f00000000;
	bra.uni 	BB0_4;

BB0_3:
	rcp.rn.f32 	%f76, %f5;

BB0_4:
	mul.f32 	%f8, %f76, %f2;
	mul.f32 	%f9, %f76, %f3;
	mul.f32 	%f10, %f76, %f4;
	cvta.to.global.u64 	%rd30, %rd12;
	mul.wide.s32 	%rd31, %r2, 4;
	add.s64 	%rd32, %rd30, %rd31;
	cvta.to.global.u64 	%rd33, %rd13;
	add.s64 	%rd34, %rd33, %rd31;
	cvta.to.global.u64 	%rd35, %rd14;
	add.s64 	%rd36, %rd35, %rd31;
	.loc 1 24 1
	ld.global.f32 	%f11, [%rd32];
	ld.global.f32 	%f12, [%rd34];
	mul.f32 	%f21, %f12, %f12;
	fma.rn.f32 	%f22, %f11, %f11, %f21;
	ld.global.f32 	%f13, [%rd36];
	fma.rn.f32 	%f23, %f13, %f13, %f22;
	.loc 2 3055 10
	sqrt.rn.f32 	%f14, %f23;
	setp.neu.f32	%p3, %f14, 0f00000000;
	@%p3 bra 	BB0_6;

	mov.f32 	%f77, 0f00000000;
	bra.uni 	BB0_7;

BB0_6:
	rcp.rn.f32 	%f77, %f14;

BB0_7:
	mul.f32 	%f25, %f77, %f13;
	mul.f32 	%f26, %f77, %f12;
	.loc 1 25 1
	mul.f32 	%f27, %f10, %f26;
	mul.f32 	%f28, %f25, %f9;
	sub.f32 	%f29, %f27, %f28;
	mul.f32 	%f30, %f77, %f11;
	.loc 1 25 1
	mul.f32 	%f31, %f8, %f25;
	mul.f32 	%f32, %f30, %f10;
	sub.f32 	%f33, %f31, %f32;
	mul.f32 	%f34, %f9, %f30;
	mul.f32 	%f35, %f26, %f8;
	sub.f32 	%f36, %f34, %f35;
	cvta.to.global.u64 	%rd37, %rd5;
	mul.wide.s32 	%rd38, %r1, 4;
	add.s64 	%rd39, %rd37, %rd38;
	cvta.to.global.u64 	%rd40, %rd6;
	add.s64 	%rd41, %rd40, %rd38;
	cvta.to.global.u64 	%rd42, %rd7;
	add.s64 	%rd43, %rd42, %rd38;
	.loc 1 26 1
	ld.global.f32 	%f37, [%rd39];
	ld.global.f32 	%f38, [%rd41];
	.loc 1 28 1
	mul.f32 	%f39, %f9, %f38;
	fma.rn.f32 	%f40, %f8, %f37, %f39;
	.loc 1 26 1
	ld.global.f32 	%f41, [%rd43];
	.loc 1 28 1
	fma.rn.f32 	%f42, %f10, %f41, %f40;
	.loc 1 29 1
	mul.f32 	%f43, %f26, %f38;
	fma.rn.f32 	%f44, %f30, %f37, %f43;
	fma.rn.f32 	%f45, %f25, %f41, %f44;
	.loc 1 30 1
	mul.f32 	%f46, %f33, %f38;
	fma.rn.f32 	%f47, %f29, %f37, %f46;
	fma.rn.f32 	%f48, %f36, %f41, %f47;
	.loc 1 32 1
	mul.f32 	%f49, %f45, %f45;
	mul.f32 	%f50, %f48, %f48;
	add.f32 	%f51, %f49, %f50;
	mul.f32 	%f52, %f42, %f51;
	fma.rn.f32 	%f53, %f42, %f42, %f50;
	mul.f32 	%f54, %f45, %f53;
	fma.rn.f32 	%f55, %f42, %f42, %f49;
	mul.f32 	%f56, %f48, %f55;
	mul.f32 	%f57, %f1, 0fC0000000;
	mul.f32 	%f58, %f57, %f52;
	mul.f32 	%f59, %f57, %f54;
	mul.f32 	%f60, %f57, %f56;
	.loc 1 34 1
	mul.f32 	%f61, %f59, %f30;
	fma.rn.f32 	%f62, %f58, %f8, %f61;
	fma.rn.f32 	%f63, %f60, %f29, %f62;
	cvta.to.global.u64 	%rd44, %rd3;
	add.s64 	%rd45, %rd44, %rd38;
	.loc 1 34 1
	ld.global.f32 	%f64, [%rd45];
	add.f32 	%f65, %f64, %f63;
	st.global.f32 	[%rd45], %f65;
	.loc 1 35 1
	mul.f32 	%f66, %f59, %f26;
	fma.rn.f32 	%f67, %f58, %f9, %f66;
	fma.rn.f32 	%f68, %f60, %f33, %f67;
	cvta.to.global.u64 	%rd46, %rd4;
	add.s64 	%rd47, %rd46, %rd38;
	.loc 1 35 1
	ld.global.f32 	%f69, [%rd47];
	add.f32 	%f70, %f69, %f68;
	st.global.f32 	[%rd47], %f70;
	.loc 1 36 1
	mul.f32 	%f71, %f59, %f25;
	fma.rn.f32 	%f72, %f58, %f10, %f71;
	fma.rn.f32 	%f73, %f60, %f36, %f72;
	add.s64 	%rd48, %rd1, %rd38;
	.loc 1 36 1
	ld.global.f32 	%f74, [%rd48];
	add.f32 	%f75, %f74, %f73;
	st.global.f32 	[%rd48], %f75;

BB0_8:
	.loc 1 38 2
	ret;
}


`
	addcubicanisotropy_ptx_35 = `
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

.visible .entry addcubicanisotropy(
	.param .u64 addcubicanisotropy_param_0,
	.param .u64 addcubicanisotropy_param_1,
	.param .u64 addcubicanisotropy_param_2,
	.param .u64 addcubicanisotropy_param_3,
	.param .u64 addcubicanisotropy_param_4,
	.param .u64 addcubicanisotropy_param_5,
	.param .u64 addcubicanisotropy_param_6,
	.param .u64 addcubicanisotropy_param_7,
	.param .u64 addcubicanisotropy_param_8,
	.param .u64 addcubicanisotropy_param_9,
	.param .u64 addcubicanisotropy_param_10,
	.param .u64 addcubicanisotropy_param_11,
	.param .u64 addcubicanisotropy_param_12,
	.param .u64 addcubicanisotropy_param_13,
	.param .u32 addcubicanisotropy_param_14
)
{
	.reg .pred 	%p<4>;
	.reg .s16 	%rs<2>;
	.reg .s32 	%r<11>;
	.reg .f32 	%f<78>;
	.reg .s64 	%rd<49>;


	ld.param.u64 	%rd3, [addcubicanisotropy_param_0];
	ld.param.u64 	%rd4, [addcubicanisotropy_param_1];
	ld.param.u64 	%rd15, [addcubicanisotropy_param_2];
	ld.param.u64 	%rd5, [addcubicanisotropy_param_3];
	ld.param.u64 	%rd6, [addcubicanisotropy_param_4];
	ld.param.u64 	%rd7, [addcubicanisotropy_param_5];
	ld.param.u64 	%rd8, [addcubicanisotropy_param_6];
	ld.param.u64 	%rd9, [addcubicanisotropy_param_7];
	ld.param.u64 	%rd10, [addcubicanisotropy_param_8];
	ld.param.u64 	%rd11, [addcubicanisotropy_param_9];
	ld.param.u64 	%rd12, [addcubicanisotropy_param_10];
	ld.param.u64 	%rd13, [addcubicanisotropy_param_11];
	ld.param.u64 	%rd14, [addcubicanisotropy_param_12];
	ld.param.u64 	%rd16, [addcubicanisotropy_param_13];
	ld.param.u32 	%r3, [addcubicanisotropy_param_14];
	cvta.to.global.u64 	%rd1, %rd15;
	cvta.to.global.u64 	%rd2, %rd16;
	.loc 1 18 1
	mov.u32 	%r4, %nctaid.x;
	mov.u32 	%r5, %ctaid.y;
	mov.u32 	%r6, %ctaid.x;
	mad.lo.s32 	%r7, %r4, %r5, %r6;
	mov.u32 	%r8, %ntid.x;
	mov.u32 	%r9, %tid.x;
	mad.lo.s32 	%r1, %r7, %r8, %r9;
	.loc 1 19 1
	setp.ge.s32	%p1, %r1, %r3;
	@%p1 bra 	BB2_8;

	cvt.s64.s32	%rd17, %r1;
	add.s64 	%rd18, %rd2, %rd17;
	.loc 1 21 1
	ld.global.nc.u8 	%rs1, [%rd18];
	cvt.u32.u16	%r10, %rs1;
	cvt.s32.s8 	%r2, %r10;
	cvta.to.global.u64 	%rd19, %rd8;
	cvt.u64.u16	%rd20, %rs1;
	cvt.s64.s8 	%rd21, %rd20;
	shl.b64 	%rd22, %rd21, 2;
	add.s64 	%rd23, %rd19, %rd22;
	.loc 1 22 1
	ld.global.nc.f32 	%f1, [%rd23];
	cvta.to.global.u64 	%rd24, %rd9;
	add.s64 	%rd25, %rd24, %rd22;
	cvta.to.global.u64 	%rd26, %rd10;
	add.s64 	%rd27, %rd26, %rd22;
	cvta.to.global.u64 	%rd28, %rd11;
	add.s64 	%rd29, %rd28, %rd22;
	.loc 1 23 1
	ld.global.nc.f32 	%f2, [%rd25];
	ld.global.nc.f32 	%f3, [%rd27];
	mul.f32 	%f17, %f3, %f3;
	fma.rn.f32 	%f18, %f2, %f2, %f17;
	ld.global.nc.f32 	%f4, [%rd29];
	fma.rn.f32 	%f19, %f4, %f4, %f18;
	.loc 3 3055 10
	sqrt.rn.f32 	%f5, %f19;
	setp.neu.f32	%p2, %f5, 0f00000000;
	@%p2 bra 	BB2_3;

	mov.f32 	%f76, 0f00000000;
	bra.uni 	BB2_4;

BB2_3:
	rcp.rn.f32 	%f76, %f5;

BB2_4:
	mul.f32 	%f8, %f76, %f2;
	mul.f32 	%f9, %f76, %f3;
	mul.f32 	%f10, %f76, %f4;
	cvta.to.global.u64 	%rd30, %rd12;
	mul.wide.s32 	%rd31, %r2, 4;
	add.s64 	%rd32, %rd30, %rd31;
	cvta.to.global.u64 	%rd33, %rd13;
	add.s64 	%rd34, %rd33, %rd31;
	cvta.to.global.u64 	%rd35, %rd14;
	add.s64 	%rd36, %rd35, %rd31;
	.loc 1 24 1
	ld.global.nc.f32 	%f11, [%rd32];
	ld.global.nc.f32 	%f12, [%rd34];
	mul.f32 	%f21, %f12, %f12;
	fma.rn.f32 	%f22, %f11, %f11, %f21;
	ld.global.nc.f32 	%f13, [%rd36];
	fma.rn.f32 	%f23, %f13, %f13, %f22;
	.loc 3 3055 10
	sqrt.rn.f32 	%f14, %f23;
	setp.neu.f32	%p3, %f14, 0f00000000;
	@%p3 bra 	BB2_6;

	mov.f32 	%f77, 0f00000000;
	bra.uni 	BB2_7;

BB2_6:
	rcp.rn.f32 	%f77, %f14;

BB2_7:
	mul.f32 	%f25, %f77, %f13;
	mul.f32 	%f26, %f77, %f12;
	.loc 1 25 1
	mul.f32 	%f27, %f10, %f26;
	mul.f32 	%f28, %f25, %f9;
	sub.f32 	%f29, %f27, %f28;
	mul.f32 	%f30, %f77, %f11;
	.loc 1 25 1
	mul.f32 	%f31, %f8, %f25;
	mul.f32 	%f32, %f30, %f10;
	sub.f32 	%f33, %f31, %f32;
	mul.f32 	%f34, %f9, %f30;
	mul.f32 	%f35, %f26, %f8;
	sub.f32 	%f36, %f34, %f35;
	cvta.to.global.u64 	%rd37, %rd5;
	mul.wide.s32 	%rd38, %r1, 4;
	add.s64 	%rd39, %rd37, %rd38;
	cvta.to.global.u64 	%rd40, %rd6;
	add.s64 	%rd41, %rd40, %rd38;
	cvta.to.global.u64 	%rd42, %rd7;
	add.s64 	%rd43, %rd42, %rd38;
	.loc 1 26 1
	ld.global.nc.f32 	%f37, [%rd39];
	ld.global.nc.f32 	%f38, [%rd41];
	.loc 1 28 1
	mul.f32 	%f39, %f9, %f38;
	fma.rn.f32 	%f40, %f8, %f37, %f39;
	.loc 1 26 1
	ld.global.nc.f32 	%f41, [%rd43];
	.loc 1 28 1
	fma.rn.f32 	%f42, %f10, %f41, %f40;
	.loc 1 29 1
	mul.f32 	%f43, %f26, %f38;
	fma.rn.f32 	%f44, %f30, %f37, %f43;
	fma.rn.f32 	%f45, %f25, %f41, %f44;
	.loc 1 30 1
	mul.f32 	%f46, %f33, %f38;
	fma.rn.f32 	%f47, %f29, %f37, %f46;
	fma.rn.f32 	%f48, %f36, %f41, %f47;
	.loc 1 32 1
	mul.f32 	%f49, %f45, %f45;
	mul.f32 	%f50, %f48, %f48;
	add.f32 	%f51, %f49, %f50;
	mul.f32 	%f52, %f42, %f51;
	fma.rn.f32 	%f53, %f42, %f42, %f50;
	mul.f32 	%f54, %f45, %f53;
	fma.rn.f32 	%f55, %f42, %f42, %f49;
	mul.f32 	%f56, %f48, %f55;
	mul.f32 	%f57, %f1, 0fC0000000;
	mul.f32 	%f58, %f57, %f52;
	mul.f32 	%f59, %f57, %f54;
	mul.f32 	%f60, %f57, %f56;
	.loc 1 34 1
	mul.f32 	%f61, %f59, %f30;
	fma.rn.f32 	%f62, %f58, %f8, %f61;
	fma.rn.f32 	%f63, %f60, %f29, %f62;
	cvta.to.global.u64 	%rd44, %rd3;
	add.s64 	%rd45, %rd44, %rd38;
	.loc 1 34 1
	ld.global.f32 	%f64, [%rd45];
	add.f32 	%f65, %f64, %f63;
	st.global.f32 	[%rd45], %f65;
	.loc 1 35 1
	mul.f32 	%f66, %f59, %f26;
	fma.rn.f32 	%f67, %f58, %f9, %f66;
	fma.rn.f32 	%f68, %f60, %f33, %f67;
	cvta.to.global.u64 	%rd46, %rd4;
	add.s64 	%rd47, %rd46, %rd38;
	.loc 1 35 1
	ld.global.f32 	%f69, [%rd47];
	add.f32 	%f70, %f69, %f68;
	st.global.f32 	[%rd47], %f70;
	.loc 1 36 1
	mul.f32 	%f71, %f59, %f25;
	fma.rn.f32 	%f72, %f58, %f10, %f71;
	fma.rn.f32 	%f73, %f60, %f36, %f72;
	add.s64 	%rd48, %rd1, %rd38;
	.loc 1 36 1
	ld.global.f32 	%f74, [%rd48];
	add.f32 	%f75, %f74, %f73;
	st.global.f32 	[%rd48], %f75;

BB2_8:
	.loc 1 38 2
	ret;
}


`
)
