#!/bin/sh
# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

COMMAND="mksysnum_plan9.sh $@"

cat <<EOF
// $COMMAND
// MACHINE GENERATED BY THE ABOVE COMMAND; DO NOT EDIT

package plan9

const(
EOF

SP='[ 	]' # space or tab
sed "s/^#define${SP}\\([A-Z0-9_][A-Z0-9_]*\\)${SP}${SP}*\\([0-9][0-9]*\\)/SYS_\\1=\\2/g" \
	<$1 | grep -v SYS__

cat <<EOF
)
EOF
