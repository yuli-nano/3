package ptx

//This file is auto-generated. Editing is futile.

func init() { Code["normalize"] = NORMALIZE }

const NORMALIZE = `
//
// Generated by NVIDIA NVVM Compiler
// Compiler built on Sat Sep 22 02:35:14 2012 (1348274114)
// Cuda compilation tools, release 5.0, V0.2.1221
//

.version 3.1
.target sm_30
.address_size 64

	.file	1 "/tmp/tmpxft_00004d21_00000000-9_normalize.cpp3.i"
	.file	2 "/home/arne/src/code.google.com/p/nimble-cube/gpu/ptx/common_func.h"
	.file	3 "/usr/local/cuda-5.0/nvvm/ci_include.h"
	.file	4 "/home/arne/src/code.google.com/p/nimble-cube/gpu/ptx/normalize.cu"

.visible .func  (.param .align 4 .b8 func_retval0[12]) _Z9normalize6float3(
	.param .align 4 .b8 _Z9normalize6float3_param_0[12]
)
{
	.reg .pred 	%p<2>;
	.reg .f32 	%f<15>;


	ld.param.f32 	%f3, [_Z9normalize6float3_param_0+8];
	ld.param.f32 	%f1, [_Z9normalize6float3_param_0];
	ld.param.f32 	%f2, [_Z9normalize6float3_param_0+4];
	.loc 2 65 1
	mul.f32 	%f8, %f2, %f2;
	fma.rn.f32 	%f9, %f1, %f1, %f8;
	fma.rn.f32 	%f10, %f3, %f3, %f9;
	.loc 3 991 5
	sqrt.rn.f32 	%f4, %f10;
	mov.f32 	%f14, 0f00000000;
	.loc 2 65 1
	setp.eq.f32 	%p1, %f4, 0f00000000;
	@%p1 bra 	BB0_2;

	rcp.rn.f32 	%f14, %f4;

BB0_2:
	.loc 2 66 1
	mul.f32 	%f11, %f14, %f1;
	mul.f32 	%f12, %f14, %f2;
	mul.f32 	%f13, %f14, %f3;
	st.param.f32	[func_retval0+0], %f11;
	st.param.f32	[func_retval0+4], %f12;
	st.param.f32	[func_retval0+8], %f13;
	ret;
}

.visible .entry normalize(
	.param .u64 normalize_param_0,
	.param .u64 normalize_param_1,
	.param .u64 normalize_param_2,
	.param .u32 normalize_param_3
)
{



	.loc 4 15 2
	ret;
}


`
