package qtgui

// /usr/include/qt/QtGui/qtextobject.h
// #include <qtextobject.h>
// #include <QtGui>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 1
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"

//  ext block end

//  body block begin

type QTextFrame struct {
	*QTextObject
}
type QTextFrame_ITF interface {
	QTextObject_ITF
	QTextFrame_PTR() *QTextFrame
}

func (ptr *QTextFrame) QTextFrame_PTR() *QTextFrame { return ptr }

func (this *QTextFrame) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QTextObject.GetCthis()
	}
}
func (this *QTextFrame) SetCthis(cthis unsafe.Pointer) {
	this.QTextObject = NewQTextObjectFromPointer(cthis)
}
func NewQTextFrameFromPointer(cthis unsafe.Pointer) *QTextFrame {
	bcthis0 := NewQTextObjectFromPointer(cthis)
	return &QTextFrame{bcthis0}
}
func (*QTextFrame) NewFromPointer(cthis unsafe.Pointer) *QTextFrame {
	return NewQTextFrameFromPointer(cthis)
}

// /usr/include/qt/QtGui/qtextobject.h:120
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QTextFrame) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextFrame10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextobject.h:123
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QTextFrame(QTextDocument *)
func NewQTextFrame(doc *QTextDocument /*777 QTextDocument **/) *QTextFrame {
	var convArg0 = doc.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextFrameC2EP13QTextDocument", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	gothis := NewQTextFrameFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtGui/qtextobject.h:124
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QTextFrame()
func DeleteQTextFrame(this *QTextFrame) {
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextFrameD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtGui/qtextobject.h:126
// index:0
// Public inline Visibility=Default Availability=Available
// [-2] void setFrameFormat(const QTextFrameFormat &)
func (this *QTextFrame) SetFrameFormat(format *QTextFrameFormat) {
	var convArg0 = format.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextFrame14setFrameFormatERK16QTextFrameFormat", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:127
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QTextFrameFormat frameFormat()
func (this *QTextFrame) FrameFormat() *QTextFrameFormat /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextFrame11frameFormatEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextFrameFormatFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextFrameFormat)
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:129
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCursor firstCursorPosition()
func (this *QTextFrame) FirstCursorPosition() *QTextCursor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextFrame19firstCursorPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:130
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextCursor lastCursorPosition()
func (this *QTextFrame) LastCursorPosition() *QTextCursor /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextFrame18lastCursorPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQTextCursorFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQTextCursor)
	return rv2
}

// /usr/include/qt/QtGui/qtextobject.h:131
// index:0
// Public Visibility=Default Availability=Available
// [4] int firstPosition()
func (this *QTextFrame) FirstPosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextFrame13firstPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:132
// index:0
// Public Visibility=Default Availability=Available
// [4] int lastPosition()
func (this *QTextFrame) LastPosition() int {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextFrame12lastPositionEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtrt.Cretval2go("int", rv).(int) // 1111
}

// /usr/include/qt/QtGui/qtextobject.h:134
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextFrameLayoutData * layoutData()
func (this *QTextFrame) LayoutData() *QTextFrameLayoutData /*777 QTextFrameLayoutData **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextFrame10layoutDataEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextFrameLayoutDataFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextobject.h:135
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setLayoutData(QTextFrameLayoutData *)
func (this *QTextFrame) SetLayoutData(data *QTextFrameLayoutData /*777 QTextFrameLayoutData **/) {
	var convArg0 = data.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN10QTextFrame13setLayoutDataEP20QTextFrameLayoutData", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtGui/qtextobject.h:138
// index:0
// Public Visibility=Default Availability=Available
// [8] QTextFrame * parentFrame()
func (this *QTextFrame) ParentFrame() *QTextFrame /*777 QTextFrame **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextFrame11parentFrameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQTextFrameFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtGui/qtextobject.h:181
// index:0
// Public Visibility=Default Availability=Available
// [32] QTextFrame::iterator begin()
func (this *QTextFrame) Begin() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextFrame5beginEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

// /usr/include/qt/QtGui/qtextobject.h:182
// index:0
// Public Visibility=Default Availability=Available
// [32] QTextFrame::iterator end()
func (this *QTextFrame) End() unsafe.Pointer /*444*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK10QTextFrame3endEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return unsafe.Pointer(uintptr(rv))
}

//  body block end

//  keep block begin

func init() {
	if false {
		reflect.TypeOf(123)
	}
	if false {
		reflect.TypeOf(unsafe.Sizeof(0))
	}
	if false {
		fmt.Println(123)
	}
	if false {
		qtrt.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  keep block end
