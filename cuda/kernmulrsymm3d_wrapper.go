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

// CUDA handle for kernmulRSymm3D kernel
var kernmulRSymm3D_code cu.Function

// Stores the arguments for kernmulRSymm3D kernel invocation
type kernmulRSymm3D_args_t struct {
	arg_fftMx  unsafe.Pointer
	arg_fftMy  unsafe.Pointer
	arg_fftMz  unsafe.Pointer
	arg_fftKxx unsafe.Pointer
	arg_fftKyy unsafe.Pointer
	arg_fftKzz unsafe.Pointer
	arg_fftKyz unsafe.Pointer
	arg_fftKxz unsafe.Pointer
	arg_fftKxy unsafe.Pointer
	arg_Nx     int
	arg_Ny     int
	arg_Nz     int
	argptr     [12]unsafe.Pointer
	sync.Mutex
}

// Stores the arguments for kernmulRSymm3D kernel invocation
var kernmulRSymm3D_args kernmulRSymm3D_args_t

func init() {
	// CUDA driver kernel call wants pointers to arguments, set them up once.
	kernmulRSymm3D_args.argptr[0] = unsafe.Pointer(&kernmulRSymm3D_args.arg_fftMx)
	kernmulRSymm3D_args.argptr[1] = unsafe.Pointer(&kernmulRSymm3D_args.arg_fftMy)
	kernmulRSymm3D_args.argptr[2] = unsafe.Pointer(&kernmulRSymm3D_args.arg_fftMz)
	kernmulRSymm3D_args.argptr[3] = unsafe.Pointer(&kernmulRSymm3D_args.arg_fftKxx)
	kernmulRSymm3D_args.argptr[4] = unsafe.Pointer(&kernmulRSymm3D_args.arg_fftKyy)
	kernmulRSymm3D_args.argptr[5] = unsafe.Pointer(&kernmulRSymm3D_args.arg_fftKzz)
	kernmulRSymm3D_args.argptr[6] = unsafe.Pointer(&kernmulRSymm3D_args.arg_fftKyz)
	kernmulRSymm3D_args.argptr[7] = unsafe.Pointer(&kernmulRSymm3D_args.arg_fftKxz)
	kernmulRSymm3D_args.argptr[8] = unsafe.Pointer(&kernmulRSymm3D_args.arg_fftKxy)
	kernmulRSymm3D_args.argptr[9] = unsafe.Pointer(&kernmulRSymm3D_args.arg_Nx)
	kernmulRSymm3D_args.argptr[10] = unsafe.Pointer(&kernmulRSymm3D_args.arg_Ny)
	kernmulRSymm3D_args.argptr[11] = unsafe.Pointer(&kernmulRSymm3D_args.arg_Nz)
}

// Wrapper for kernmulRSymm3D CUDA kernel, asynchronous.
func k_kernmulRSymm3D_async(fftMx unsafe.Pointer, fftMy unsafe.Pointer, fftMz unsafe.Pointer, fftKxx unsafe.Pointer, fftKyy unsafe.Pointer, fftKzz unsafe.Pointer, fftKyz unsafe.Pointer, fftKxz unsafe.Pointer, fftKxy unsafe.Pointer, Nx int, Ny int, Nz int, cfg *config) {
	if Synchronous { // debug
		Sync()
	}

	kernmulRSymm3D_args.Lock()
	defer kernmulRSymm3D_args.Unlock()

	if kernmulRSymm3D_code == 0 {
		kernmulRSymm3D_code = fatbinLoad(kernmulRSymm3D_map, "kernmulRSymm3D")
	}

	kernmulRSymm3D_args.arg_fftMx = fftMx
	kernmulRSymm3D_args.arg_fftMy = fftMy
	kernmulRSymm3D_args.arg_fftMz = fftMz
	kernmulRSymm3D_args.arg_fftKxx = fftKxx
	kernmulRSymm3D_args.arg_fftKyy = fftKyy
	kernmulRSymm3D_args.arg_fftKzz = fftKzz
	kernmulRSymm3D_args.arg_fftKyz = fftKyz
	kernmulRSymm3D_args.arg_fftKxz = fftKxz
	kernmulRSymm3D_args.arg_fftKxy = fftKxy
	kernmulRSymm3D_args.arg_Nx = Nx
	kernmulRSymm3D_args.arg_Ny = Ny
	kernmulRSymm3D_args.arg_Nz = Nz

	args := kernmulRSymm3D_args.argptr[:]
	cu.LaunchKernel(kernmulRSymm3D_code, cfg.Grid.X, cfg.Grid.Y, cfg.Grid.Z, cfg.Block.X, cfg.Block.Y, cfg.Block.Z, 0, stream0, args)

	if Synchronous { // debug
		Sync()
	}
}

// maps compute capability on PTX code for kernmulRSymm3D kernel.
var kernmulRSymm3D_map = map[int]string{0: "",
	20: kernmulRSymm3D_ptx_20,
	30: kernmulRSymm3D_ptx_30,
	35: kernmulRSymm3D_ptx_35}

// kernmulRSymm3D PTX code for various compute capabilities.
const (
	kernmulRSymm3D_ptx_20 = `
.version 3.2
.target sm_20
.address_size 64


.visible .entry kernmulRSymm3D(
	.param .u64 kernmulRSymm3D_param_0,
	.param .u64 kernmulRSymm3D_param_1,
	.param .u64 kernmulRSymm3D_param_2,
	.param .u64 kernmulRSymm3D_param_3,
	.param .u64 kernmulRSymm3D_param_4,
	.param .u64 kernmulRSymm3D_param_5,
	.param .u64 kernmulRSymm3D_param_6,
	.param .u64 kernmulRSymm3D_param_7,
	.param .u64 kernmulRSymm3D_param_8,
	.param .u32 kernmulRSymm3D_param_9,
	.param .u32 kernmulRSymm3D_param_10,
	.param .u32 kernmulRSymm3D_param_11
)
{
	.reg .pred 	%p<7>;
	.reg .s32 	%r<66>;
	.reg .f32 	%f<39>;
	.reg .s64 	%rd<43>;


	ld.param.u64 	%rd2, [kernmulRSymm3D_param_0];
	ld.param.u64 	%rd3, [kernmulRSymm3D_param_1];
	ld.param.u64 	%rd4, [kernmulRSymm3D_param_2];
	ld.param.u64 	%rd5, [kernmulRSymm3D_param_3];
	ld.param.u64 	%rd6, [kernmulRSymm3D_param_4];
	ld.param.u64 	%rd7, [kernmulRSymm3D_param_5];
	ld.param.u64 	%rd8, [kernmulRSymm3D_param_6];
	ld.param.u64 	%rd9, [kernmulRSymm3D_param_7];
	ld.param.u64 	%rd10, [kernmulRSymm3D_param_8];
	ld.param.u32 	%r8, [kernmulRSymm3D_param_9];
	ld.param.u32 	%r9, [kernmulRSymm3D_param_10];
	ld.param.u32 	%r10, [kernmulRSymm3D_param_11];
	cvta.to.global.u64 	%rd1, %rd4;
	.loc 1 35 1
	mov.u32 	%r11, %ntid.x;
	mov.u32 	%r12, %ctaid.x;
	mov.u32 	%r13, %tid.x;
	mad.lo.s32 	%r1, %r11, %r12, %r13;
	.loc 1 36 1
	mov.u32 	%r14, %ntid.y;
	mov.u32 	%r15, %ctaid.y;
	mov.u32 	%r16, %tid.y;
	mad.lo.s32 	%r2, %r14, %r15, %r16;
	.loc 1 37 1
	mov.u32 	%r17, %ntid.z;
	mov.u32 	%r18, %ctaid.z;
	mov.u32 	%r19, %tid.z;
	mad.lo.s32 	%r3, %r17, %r18, %r19;
	.loc 1 39 1
	setp.ge.s32	%p1, %r2, %r9;
	setp.ge.s32	%p2, %r1, %r8;
	or.pred  	%p3, %p2, %p1;
	.loc 1 39 1
	setp.ge.s32	%p4, %r3, %r10;
	or.pred  	%p5, %p3, %p4;
	.loc 1 39 1
	@%p5 bra 	BB0_5;

	.loc 1 45 1
	mul.lo.s32 	%r4, %r3, %r9;
	add.s32 	%r20, %r4, %r2;
	mad.lo.s32 	%r21, %r20, %r8, %r1;
	.loc 1 46 1
	shl.b32 	%r22, %r21, 1;
	cvta.to.global.u64 	%rd11, %rd2;
	mul.wide.s32 	%rd12, %r22, 4;
	add.s64 	%rd13, %rd11, %rd12;
	.loc 1 47 1
	ld.global.f32 	%f1, [%rd13];
	.loc 1 48 1
	ld.global.f32 	%f2, [%rd13+4];
	cvta.to.global.u64 	%rd14, %rd3;
	add.s64 	%rd15, %rd14, %rd12;
	.loc 1 49 1
	ld.global.f32 	%f3, [%rd15];
	.loc 1 50 1
	ld.global.f32 	%f4, [%rd15+4];
	add.s64 	%rd16, %rd1, %rd12;
	.loc 1 51 1
	ld.global.f32 	%f5, [%rd16];
	.loc 1 52 1
	ld.global.f32 	%f6, [%rd16+4];
	.loc 1 54 1
	shr.u32 	%r23, %r9, 31;
	add.s32 	%r24, %r9, %r23;
	shr.s32 	%r25, %r24, 1;
	add.s32 	%r26, %r25, 1;
	setp.lt.s32	%p6, %r2, %r26;
	@%p6 bra 	BB0_3;

	.loc 1 62 1
	sub.s32 	%r31, %r9, %r2;
	add.s32 	%r32, %r31, %r4;
	mad.lo.s32 	%r65, %r32, %r8, %r1;
	cvta.to.global.u64 	%rd17, %rd8;
	mul.wide.s32 	%rd18, %r65, 4;
	add.s64 	%rd19, %rd17, %rd18;
	.loc 1 66 1
	ld.global.f32 	%f13, [%rd19];
	neg.f32 	%f38, %f13;
	cvta.to.global.u64 	%rd20, %rd10;
	add.s64 	%rd21, %rd20, %rd18;
	.loc 1 68 1
	ld.global.f32 	%f14, [%rd21];
	neg.f32 	%f37, %f14;
	bra.uni 	BB0_4;

BB0_3:
	.loc 1 45 1
	mad.lo.s32 	%r45, %r3, %r9, %r2;
	mad.lo.s32 	%r65, %r45, %r8, %r1;
	cvta.to.global.u64 	%rd22, %rd8;
	mul.wide.s32 	%rd23, %r65, 4;
	add.s64 	%rd24, %rd22, %rd23;
	.loc 1 58 1
	ld.global.f32 	%f38, [%rd24];
	cvta.to.global.u64 	%rd25, %rd10;
	add.s64 	%rd26, %rd25, %rd23;
	.loc 1 60 1
	ld.global.f32 	%f37, [%rd26];

BB0_4:
	cvta.to.global.u64 	%rd27, %rd9;
	mul.wide.s32 	%rd28, %r65, 4;
	add.s64 	%rd29, %rd27, %rd28;
	cvta.to.global.u64 	%rd30, %rd7;
	add.s64 	%rd31, %rd30, %rd28;
	cvta.to.global.u64 	%rd32, %rd6;
	add.s64 	%rd33, %rd32, %rd28;
	cvta.to.global.u64 	%rd34, %rd5;
	add.s64 	%rd35, %rd34, %rd28;
	.loc 1 71 1
	ld.global.f32 	%f15, [%rd31];
	ld.global.f32 	%f16, [%rd33];
	ld.global.f32 	%f17, [%rd35];
	mul.f32 	%f18, %f3, %f37;
	fma.rn.f32 	%f19, %f1, %f17, %f18;
	ld.global.f32 	%f20, [%rd29];
	fma.rn.f32 	%f21, %f5, %f20, %f19;
	.loc 1 45 1
	mad.lo.s32 	%r58, %r3, %r9, %r2;
	mad.lo.s32 	%r63, %r58, %r8, %r1;
	.loc 1 46 1
	shl.b32 	%r64, %r63, 1;
	mul.wide.s32 	%rd37, %r64, 4;
	add.s64 	%rd38, %rd11, %rd37;
	.loc 1 71 1
	st.global.f32 	[%rd38], %f21;
	.loc 1 72 1
	mul.f32 	%f22, %f4, %f37;
	fma.rn.f32 	%f23, %f2, %f17, %f22;
	fma.rn.f32 	%f24, %f6, %f20, %f23;
	st.global.f32 	[%rd38+4], %f24;
	.loc 1 73 1
	mul.f32 	%f25, %f3, %f16;
	fma.rn.f32 	%f26, %f1, %f37, %f25;
	fma.rn.f32 	%f27, %f5, %f38, %f26;
	add.s64 	%rd40, %rd14, %rd37;
	.loc 1 73 1
	st.global.f32 	[%rd40], %f27;
	.loc 1 74 1
	mul.f32 	%f28, %f4, %f16;
	fma.rn.f32 	%f29, %f2, %f37, %f28;
	fma.rn.f32 	%f30, %f6, %f38, %f29;
	st.global.f32 	[%rd40+4], %f30;
	.loc 1 75 1
	mul.f32 	%f31, %f3, %f38;
	fma.rn.f32 	%f32, %f1, %f20, %f31;
	fma.rn.f32 	%f33, %f5, %f15, %f32;
	add.s64 	%rd42, %rd1, %rd37;
	.loc 1 75 1
	st.global.f32 	[%rd42], %f33;
	.loc 1 76 1
	mul.f32 	%f34, %f4, %f38;
	fma.rn.f32 	%f35, %f2, %f20, %f34;
	fma.rn.f32 	%f36, %f6, %f15, %f35;
	st.global.f32 	[%rd42+4], %f36;

BB0_5:
	.loc 1 77 2
	ret;
}


`
	kernmulRSymm3D_ptx_30 = `
.version 3.2
.target sm_30
.address_size 64


.visible .entry kernmulRSymm3D(
	.param .u64 kernmulRSymm3D_param_0,
	.param .u64 kernmulRSymm3D_param_1,
	.param .u64 kernmulRSymm3D_param_2,
	.param .u64 kernmulRSymm3D_param_3,
	.param .u64 kernmulRSymm3D_param_4,
	.param .u64 kernmulRSymm3D_param_5,
	.param .u64 kernmulRSymm3D_param_6,
	.param .u64 kernmulRSymm3D_param_7,
	.param .u64 kernmulRSymm3D_param_8,
	.param .u32 kernmulRSymm3D_param_9,
	.param .u32 kernmulRSymm3D_param_10,
	.param .u32 kernmulRSymm3D_param_11
)
{
	.reg .pred 	%p<7>;
	.reg .s32 	%r<66>;
	.reg .f32 	%f<39>;
	.reg .s64 	%rd<43>;


	ld.param.u64 	%rd2, [kernmulRSymm3D_param_0];
	ld.param.u64 	%rd3, [kernmulRSymm3D_param_1];
	ld.param.u64 	%rd4, [kernmulRSymm3D_param_2];
	ld.param.u64 	%rd5, [kernmulRSymm3D_param_3];
	ld.param.u64 	%rd6, [kernmulRSymm3D_param_4];
	ld.param.u64 	%rd7, [kernmulRSymm3D_param_5];
	ld.param.u64 	%rd8, [kernmulRSymm3D_param_6];
	ld.param.u64 	%rd9, [kernmulRSymm3D_param_7];
	ld.param.u64 	%rd10, [kernmulRSymm3D_param_8];
	ld.param.u32 	%r8, [kernmulRSymm3D_param_9];
	ld.param.u32 	%r9, [kernmulRSymm3D_param_10];
	ld.param.u32 	%r10, [kernmulRSymm3D_param_11];
	cvta.to.global.u64 	%rd1, %rd4;
	.loc 1 35 1
	mov.u32 	%r11, %ntid.x;
	mov.u32 	%r12, %ctaid.x;
	mov.u32 	%r13, %tid.x;
	mad.lo.s32 	%r1, %r11, %r12, %r13;
	.loc 1 36 1
	mov.u32 	%r14, %ntid.y;
	mov.u32 	%r15, %ctaid.y;
	mov.u32 	%r16, %tid.y;
	mad.lo.s32 	%r2, %r14, %r15, %r16;
	.loc 1 37 1
	mov.u32 	%r17, %ntid.z;
	mov.u32 	%r18, %ctaid.z;
	mov.u32 	%r19, %tid.z;
	mad.lo.s32 	%r3, %r17, %r18, %r19;
	.loc 1 39 1
	setp.ge.s32	%p1, %r2, %r9;
	setp.ge.s32	%p2, %r1, %r8;
	or.pred  	%p3, %p2, %p1;
	.loc 1 39 1
	setp.ge.s32	%p4, %r3, %r10;
	or.pred  	%p5, %p3, %p4;
	.loc 1 39 1
	@%p5 bra 	BB0_5;

	.loc 1 45 1
	mul.lo.s32 	%r4, %r3, %r9;
	add.s32 	%r20, %r4, %r2;
	mad.lo.s32 	%r21, %r20, %r8, %r1;
	.loc 1 46 1
	shl.b32 	%r22, %r21, 1;
	cvta.to.global.u64 	%rd11, %rd2;
	mul.wide.s32 	%rd12, %r22, 4;
	add.s64 	%rd13, %rd11, %rd12;
	.loc 1 47 1
	ld.global.f32 	%f1, [%rd13];
	.loc 1 48 1
	ld.global.f32 	%f2, [%rd13+4];
	cvta.to.global.u64 	%rd14, %rd3;
	add.s64 	%rd15, %rd14, %rd12;
	.loc 1 49 1
	ld.global.f32 	%f3, [%rd15];
	.loc 1 50 1
	ld.global.f32 	%f4, [%rd15+4];
	add.s64 	%rd16, %rd1, %rd12;
	.loc 1 51 1
	ld.global.f32 	%f5, [%rd16];
	.loc 1 52 1
	ld.global.f32 	%f6, [%rd16+4];
	.loc 1 54 1
	shr.u32 	%r23, %r9, 31;
	add.s32 	%r24, %r9, %r23;
	shr.s32 	%r25, %r24, 1;
	add.s32 	%r26, %r25, 1;
	setp.lt.s32	%p6, %r2, %r26;
	@%p6 bra 	BB0_3;

	.loc 1 62 1
	sub.s32 	%r31, %r9, %r2;
	add.s32 	%r32, %r31, %r4;
	mad.lo.s32 	%r65, %r32, %r8, %r1;
	cvta.to.global.u64 	%rd17, %rd8;
	mul.wide.s32 	%rd18, %r65, 4;
	add.s64 	%rd19, %rd17, %rd18;
	.loc 1 66 1
	ld.global.f32 	%f13, [%rd19];
	neg.f32 	%f38, %f13;
	cvta.to.global.u64 	%rd20, %rd10;
	add.s64 	%rd21, %rd20, %rd18;
	.loc 1 68 1
	ld.global.f32 	%f14, [%rd21];
	neg.f32 	%f37, %f14;
	bra.uni 	BB0_4;

BB0_3:
	.loc 1 45 1
	mad.lo.s32 	%r45, %r3, %r9, %r2;
	mad.lo.s32 	%r65, %r45, %r8, %r1;
	cvta.to.global.u64 	%rd22, %rd8;
	mul.wide.s32 	%rd23, %r65, 4;
	add.s64 	%rd24, %rd22, %rd23;
	.loc 1 58 1
	ld.global.f32 	%f38, [%rd24];
	cvta.to.global.u64 	%rd25, %rd10;
	add.s64 	%rd26, %rd25, %rd23;
	.loc 1 60 1
	ld.global.f32 	%f37, [%rd26];

BB0_4:
	cvta.to.global.u64 	%rd27, %rd9;
	mul.wide.s32 	%rd28, %r65, 4;
	add.s64 	%rd29, %rd27, %rd28;
	cvta.to.global.u64 	%rd30, %rd7;
	add.s64 	%rd31, %rd30, %rd28;
	cvta.to.global.u64 	%rd32, %rd6;
	add.s64 	%rd33, %rd32, %rd28;
	cvta.to.global.u64 	%rd34, %rd5;
	add.s64 	%rd35, %rd34, %rd28;
	.loc 1 71 1
	ld.global.f32 	%f15, [%rd31];
	ld.global.f32 	%f16, [%rd33];
	ld.global.f32 	%f17, [%rd35];
	mul.f32 	%f18, %f3, %f37;
	fma.rn.f32 	%f19, %f1, %f17, %f18;
	ld.global.f32 	%f20, [%rd29];
	fma.rn.f32 	%f21, %f5, %f20, %f19;
	.loc 1 45 1
	mad.lo.s32 	%r58, %r3, %r9, %r2;
	mad.lo.s32 	%r63, %r58, %r8, %r1;
	.loc 1 46 1
	shl.b32 	%r64, %r63, 1;
	mul.wide.s32 	%rd37, %r64, 4;
	add.s64 	%rd38, %rd11, %rd37;
	.loc 1 71 1
	st.global.f32 	[%rd38], %f21;
	.loc 1 72 1
	mul.f32 	%f22, %f4, %f37;
	fma.rn.f32 	%f23, %f2, %f17, %f22;
	fma.rn.f32 	%f24, %f6, %f20, %f23;
	st.global.f32 	[%rd38+4], %f24;
	.loc 1 73 1
	mul.f32 	%f25, %f3, %f16;
	fma.rn.f32 	%f26, %f1, %f37, %f25;
	fma.rn.f32 	%f27, %f5, %f38, %f26;
	add.s64 	%rd40, %rd14, %rd37;
	.loc 1 73 1
	st.global.f32 	[%rd40], %f27;
	.loc 1 74 1
	mul.f32 	%f28, %f4, %f16;
	fma.rn.f32 	%f29, %f2, %f37, %f28;
	fma.rn.f32 	%f30, %f6, %f38, %f29;
	st.global.f32 	[%rd40+4], %f30;
	.loc 1 75 1
	mul.f32 	%f31, %f3, %f38;
	fma.rn.f32 	%f32, %f1, %f20, %f31;
	fma.rn.f32 	%f33, %f5, %f15, %f32;
	add.s64 	%rd42, %rd1, %rd37;
	.loc 1 75 1
	st.global.f32 	[%rd42], %f33;
	.loc 1 76 1
	mul.f32 	%f34, %f4, %f38;
	fma.rn.f32 	%f35, %f2, %f20, %f34;
	fma.rn.f32 	%f36, %f6, %f15, %f35;
	st.global.f32 	[%rd42+4], %f36;

BB0_5:
	.loc 1 77 2
	ret;
}


`
	kernmulRSymm3D_ptx_35 = `
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

.visible .entry kernmulRSymm3D(
	.param .u64 kernmulRSymm3D_param_0,
	.param .u64 kernmulRSymm3D_param_1,
	.param .u64 kernmulRSymm3D_param_2,
	.param .u64 kernmulRSymm3D_param_3,
	.param .u64 kernmulRSymm3D_param_4,
	.param .u64 kernmulRSymm3D_param_5,
	.param .u64 kernmulRSymm3D_param_6,
	.param .u64 kernmulRSymm3D_param_7,
	.param .u64 kernmulRSymm3D_param_8,
	.param .u32 kernmulRSymm3D_param_9,
	.param .u32 kernmulRSymm3D_param_10,
	.param .u32 kernmulRSymm3D_param_11
)
{
	.reg .pred 	%p<7>;
	.reg .s32 	%r<66>;
	.reg .f32 	%f<39>;
	.reg .s64 	%rd<43>;


	ld.param.u64 	%rd2, [kernmulRSymm3D_param_0];
	ld.param.u64 	%rd3, [kernmulRSymm3D_param_1];
	ld.param.u64 	%rd4, [kernmulRSymm3D_param_2];
	ld.param.u64 	%rd5, [kernmulRSymm3D_param_3];
	ld.param.u64 	%rd6, [kernmulRSymm3D_param_4];
	ld.param.u64 	%rd7, [kernmulRSymm3D_param_5];
	ld.param.u64 	%rd8, [kernmulRSymm3D_param_6];
	ld.param.u64 	%rd9, [kernmulRSymm3D_param_7];
	ld.param.u64 	%rd10, [kernmulRSymm3D_param_8];
	ld.param.u32 	%r8, [kernmulRSymm3D_param_9];
	ld.param.u32 	%r9, [kernmulRSymm3D_param_10];
	ld.param.u32 	%r10, [kernmulRSymm3D_param_11];
	cvta.to.global.u64 	%rd1, %rd4;
	.loc 1 35 1
	mov.u32 	%r11, %ntid.x;
	mov.u32 	%r12, %ctaid.x;
	mov.u32 	%r13, %tid.x;
	mad.lo.s32 	%r1, %r11, %r12, %r13;
	.loc 1 36 1
	mov.u32 	%r14, %ntid.y;
	mov.u32 	%r15, %ctaid.y;
	mov.u32 	%r16, %tid.y;
	mad.lo.s32 	%r2, %r14, %r15, %r16;
	.loc 1 37 1
	mov.u32 	%r17, %ntid.z;
	mov.u32 	%r18, %ctaid.z;
	mov.u32 	%r19, %tid.z;
	mad.lo.s32 	%r3, %r17, %r18, %r19;
	.loc 1 39 1
	setp.ge.s32	%p1, %r2, %r9;
	setp.ge.s32	%p2, %r1, %r8;
	or.pred  	%p3, %p2, %p1;
	.loc 1 39 1
	setp.ge.s32	%p4, %r3, %r10;
	or.pred  	%p5, %p3, %p4;
	.loc 1 39 1
	@%p5 bra 	BB2_5;

	.loc 1 45 1
	mul.lo.s32 	%r4, %r3, %r9;
	add.s32 	%r20, %r4, %r2;
	mad.lo.s32 	%r21, %r20, %r8, %r1;
	.loc 1 46 1
	shl.b32 	%r22, %r21, 1;
	cvta.to.global.u64 	%rd11, %rd2;
	mul.wide.s32 	%rd12, %r22, 4;
	add.s64 	%rd13, %rd11, %rd12;
	.loc 1 47 1
	ld.global.f32 	%f1, [%rd13];
	.loc 1 48 1
	ld.global.f32 	%f2, [%rd13+4];
	cvta.to.global.u64 	%rd14, %rd3;
	add.s64 	%rd15, %rd14, %rd12;
	.loc 1 49 1
	ld.global.f32 	%f3, [%rd15];
	.loc 1 50 1
	ld.global.f32 	%f4, [%rd15+4];
	add.s64 	%rd16, %rd1, %rd12;
	.loc 1 51 1
	ld.global.f32 	%f5, [%rd16];
	.loc 1 52 1
	ld.global.f32 	%f6, [%rd16+4];
	.loc 1 54 1
	shr.u32 	%r23, %r9, 31;
	add.s32 	%r24, %r9, %r23;
	shr.s32 	%r25, %r24, 1;
	add.s32 	%r26, %r25, 1;
	setp.lt.s32	%p6, %r2, %r26;
	@%p6 bra 	BB2_3;

	.loc 1 62 1
	sub.s32 	%r31, %r9, %r2;
	add.s32 	%r32, %r31, %r4;
	mad.lo.s32 	%r65, %r32, %r8, %r1;
	cvta.to.global.u64 	%rd17, %rd8;
	mul.wide.s32 	%rd18, %r65, 4;
	add.s64 	%rd19, %rd17, %rd18;
	.loc 1 66 1
	ld.global.nc.f32 	%f13, [%rd19];
	neg.f32 	%f38, %f13;
	cvta.to.global.u64 	%rd20, %rd10;
	add.s64 	%rd21, %rd20, %rd18;
	.loc 1 68 1
	ld.global.nc.f32 	%f14, [%rd21];
	neg.f32 	%f37, %f14;
	bra.uni 	BB2_4;

BB2_3:
	.loc 1 45 1
	mad.lo.s32 	%r45, %r3, %r9, %r2;
	mad.lo.s32 	%r65, %r45, %r8, %r1;
	cvta.to.global.u64 	%rd22, %rd8;
	mul.wide.s32 	%rd23, %r65, 4;
	add.s64 	%rd24, %rd22, %rd23;
	.loc 1 58 1
	ld.global.nc.f32 	%f38, [%rd24];
	cvta.to.global.u64 	%rd25, %rd10;
	add.s64 	%rd26, %rd25, %rd23;
	.loc 1 60 1
	ld.global.nc.f32 	%f37, [%rd26];

BB2_4:
	cvta.to.global.u64 	%rd27, %rd9;
	mul.wide.s32 	%rd28, %r65, 4;
	add.s64 	%rd29, %rd27, %rd28;
	cvta.to.global.u64 	%rd30, %rd7;
	add.s64 	%rd31, %rd30, %rd28;
	cvta.to.global.u64 	%rd32, %rd6;
	add.s64 	%rd33, %rd32, %rd28;
	cvta.to.global.u64 	%rd34, %rd5;
	add.s64 	%rd35, %rd34, %rd28;
	.loc 1 71 1
	ld.global.nc.f32 	%f15, [%rd31];
	ld.global.nc.f32 	%f16, [%rd33];
	ld.global.nc.f32 	%f17, [%rd35];
	mul.f32 	%f18, %f3, %f37;
	fma.rn.f32 	%f19, %f1, %f17, %f18;
	ld.global.nc.f32 	%f20, [%rd29];
	fma.rn.f32 	%f21, %f5, %f20, %f19;
	.loc 1 45 1
	mad.lo.s32 	%r58, %r3, %r9, %r2;
	mad.lo.s32 	%r63, %r58, %r8, %r1;
	.loc 1 46 1
	shl.b32 	%r64, %r63, 1;
	mul.wide.s32 	%rd37, %r64, 4;
	add.s64 	%rd38, %rd11, %rd37;
	.loc 1 71 1
	st.global.f32 	[%rd38], %f21;
	.loc 1 72 1
	mul.f32 	%f22, %f4, %f37;
	fma.rn.f32 	%f23, %f2, %f17, %f22;
	fma.rn.f32 	%f24, %f6, %f20, %f23;
	st.global.f32 	[%rd38+4], %f24;
	.loc 1 73 1
	mul.f32 	%f25, %f3, %f16;
	fma.rn.f32 	%f26, %f1, %f37, %f25;
	fma.rn.f32 	%f27, %f5, %f38, %f26;
	add.s64 	%rd40, %rd14, %rd37;
	.loc 1 73 1
	st.global.f32 	[%rd40], %f27;
	.loc 1 74 1
	mul.f32 	%f28, %f4, %f16;
	fma.rn.f32 	%f29, %f2, %f37, %f28;
	fma.rn.f32 	%f30, %f6, %f38, %f29;
	st.global.f32 	[%rd40+4], %f30;
	.loc 1 75 1
	mul.f32 	%f31, %f3, %f38;
	fma.rn.f32 	%f32, %f1, %f20, %f31;
	fma.rn.f32 	%f33, %f5, %f15, %f32;
	add.s64 	%rd42, %rd1, %rd37;
	.loc 1 75 1
	st.global.f32 	[%rd42], %f33;
	.loc 1 76 1
	mul.f32 	%f34, %f4, %f38;
	fma.rn.f32 	%f35, %f2, %f20, %f34;
	fma.rn.f32 	%f36, %f6, %f15, %f35;
	st.global.f32 	[%rd42+4], %f36;

BB2_5:
	.loc 1 77 2
	ret;
}


`
)
