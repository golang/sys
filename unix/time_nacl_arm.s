// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

#include "../../cmd/ld/textflag.h" // TODO: how to refer to this?

TEXT ·startTimer(SB),NOSPLIT,$0
	B time·startTimer(SB)

TEXT ·stopTimer(SB),NOSPLIT,$0
	B time·stopTimer(SB)
