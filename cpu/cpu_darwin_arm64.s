//go:build arm64 && gc
// +build arm64
// +build gc

#include "textflag.h"

// func darwinCheckFeatureEnabled(feature_vec uint64) bool
TEXT ·darwinCheckFeatureEnabled(SB), NOSPLIT, $0-8
    MOVD    feature_vec+0(FP), R0
    MOVD    $0, ret+0(FP) // default to false
    MOVD    $1, R2        // set R2 as true boolean constan

#ifdef GOOS_darwin   // return if not darwin
#ifdef GOARCH_arm64  // return if not arm64
// These values from:
// https://github.com/apple/darwin-xnu/blob/main/osfmk/arm/cpu_capabilities.h
#define arm_commpage64_base_address         0x0000000fffffc000
#define arm_commpage64_cpu_capabilities64   (arm_commpage64_base_address+0x010)
    MOVD    $0xffffc000, R1
    MOVK    $(0xf<<32), R1
    MOVD    (R1), R1
    AND   R1, R0
    CBZ      R0, no_feature
    MOVD    R2, ret+8(FP)

no_feature:
#endif
#endif
    RET
