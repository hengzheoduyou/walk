// Copyright 2010 The Walk Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package walk

import "github.com/lxn/go-winapi"

type Rectangle struct {
	X, Y, Width, Height int
}

func (r Rectangle) Left() int {
	return r.X
}

func (r Rectangle) Top() int {
	return r.Y
}

func (r Rectangle) Right() int {
	return r.X + r.Width - 1
}

func (r Rectangle) Bottom() int {
	return r.Y + r.Height - 1
}

func (r Rectangle) Location() Point {
	return Point{r.X, r.Y}
}

func (r *Rectangle) SetLocation(p Point) Rectangle {
	r.X = p.X
	r.Y = p.Y

	return *r
}

func (r Rectangle) Size() Size {
	return Size{r.Width, r.Height}
}

func (r *Rectangle) SetSize(s Size) Rectangle {
	r.Width = s.Width
	r.Height = s.Height

	return *r
}

func (r Rectangle) toRECT() winapi.RECT {
	return winapi.RECT{
		int32(r.X),
		int32(r.Y),
		int32(r.X + r.Width),
		int32(r.Y + r.Height),
	}
}