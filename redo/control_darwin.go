// 30 july 2014

package ui

// #include "objc_darwin.h"
import "C"

type controlbase struct {
	*controldefs
	id	C.id
}

type controlParent struct {
	id	C.id
}

func newControl(id C.id) *controlbase {
	c := new(controlbase)
	c.id = id
	c.controldefs = new(controldefs)
	c.fsetParent = func(p *controlParent) {
		// redrawing the new window handled by C.parent()
		C.parent(c.id, p.id)
	}
	c.fcontainerShow = func() {
		C.controlSetHidden(c.id, C.NO)
	}
	c.fcontainerHide = func() {
		C.controlSetHidden(c.id, C.YES)
	}
	c.fallocate = baseallocate(c)
	c.fpreferredSize = func(d *sizing) (int, int) {
		s := C.controlPrefSize(c.id)
		return int(s.width), int(s.height)
	}
	c.fcommitResize = func(a *allocation, d *sizing) {
//TODO
/*
	if s.ctype == c_label && !s.alternate && c.neighbor != nil {
		c.neighbor.getAuxResizeInfo(d)
		if d.neighborAlign.baseline != 0 {		// no adjustment needed if the given control has no baseline
			// in order for the baseline value to be correct, the label MUST BE AT THE HEIGHT THAT OS X WANTS IT TO BE!
			// otherwise, the baseline calculation will be relative to the bottom of the control, and everything will be wrong
			origsize := C.controlPrefSize(s.id)
			c.height = int(origsize.height)
			newrect := C.struct_xrect{
				x:		C.intptr_t(c.x),
				y:		C.intptr_t(c.y),
				width:	C.intptr_t(c.width),
				height:	C.intptr_t(c.height),
			}
			ourAlign := C.alignmentInfo(s.id, newrect)
			// we need to find the exact Y positions of the baselines
			// fortunately, this is easy now that (x,y) is the bottom-left corner
			thisbasey := ourAlign.alignmentRect.y + ourAlign.baseline
			neighborbasey := d.neighborAlign.alignmentRect.y + d.neighborAlign.baseline
			// now the amount we have to move the label down by is easy to find
			yoff := neighborbasey - thisbasey
			// and we just add that
			c.y += int(yoff)
		}
		// TODO if there's no baseline, the alignment should be to the top /of the alignment rect/, not the frame
	}
*/
		C.moveControl(c.id, C.intptr_t(a.x), C.intptr_t(a.y), C.intptr_t(a.width), C.intptr_t(a.height))
	}
	c.fgetAuxResizeInfo = func(d *sizing) {
// TODO
//		d.neighborAlign = C.alignmentInfo(s.id, C.frame(s.id))
	}
	return c
}

type scrolledcontrol struct {
	*controlbase
	scroller			*controlbase
}

func newScrolledControl(id C.id) *scrolledcontrol {
	scroller := C.newScrollView(id)
	s := &scrolledcontrol{
		controlbase:		newControl(id),
		scroller:			newControl(scroller),
	}
	s.fsetParent = s.scroller.fsetParent
	s.fcommitResize = s.scroller.fcommitResize
	return s
}