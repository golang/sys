// Copyright 2019 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package consolidate

type (
	// FileStats holds statistics for a go source code file.
	FileStats struct {
		Name   string // File name
		Consts int    // Number of constants
		Types  int    // Number of types
		Funcs  int    // Number of functions
	}
	// AggStats holds statistics for an aggregated file.
	AggStats struct {
		FileStats             // Aggregated objects file
		Arch      []FileStats // Arch dependent files for this aggregate
	}
	// Stats holds all the statistics of a consolidated package.
	Stats struct {
		Agg []AggStats
	}
)

func (s *Stats) clear() {
	s.Agg = s.Agg[:0]
}

func (s Stats) String() string {
	var buf builder
	for _, as := range s.Agg {
		_, _ = buf.WriteString(as.String())
		buf.Newline()
	}
	return buf.String()
}

func (s *AggStats) clear() {
	s.FileStats.clear()
	s.Arch = s.Arch[:0]
}

func (s AggStats) String() string {
	var buf builder
	_, _ = buf.WriteString(s.FileStats.String())
	buf.Newline()
	for _, fs := range s.Arch {
		_, _ = buf.WriteString("  ")
		_, _ = buf.WriteString(fs.String())
		buf.Newline()
	}
	return buf.String()
}

func (s *FileStats) clear() {
	s.Name = ""
	s.Consts = 0
	s.Types = 0
	s.Funcs = 0
}

func (s FileStats) String() string {
	var buf builder
	buf.Printf("file=%q ", s.Name)
	buf.Printf("\t\tconsts=%d", s.Consts)
	buf.Printf("\t\ttypes=%d", s.Types)
	buf.Printf("\t\tfuncs=%d", s.Funcs)
	return buf.String()
}

func (s *FileStats) set(name string, k *kinds) {
	s.Name = name
	if k.consts != nil {
		for _, v := range k.consts {
			s.Consts += len(v.Names)
		}
	}
	s.Types = len(k.types)
	s.Funcs = len(k.funcs)
}
