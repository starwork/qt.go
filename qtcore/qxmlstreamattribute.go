package qtcore

// /usr/include/qt/QtCore/qxmlstream.h
// #include <qxmlstream.h>
// #include <QtCore>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 10
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"

//  ext block end

//  body block begin

type QXmlStreamAttribute struct {
	*qtrt.CObject
}
type QXmlStreamAttribute_ITF interface {
	QXmlStreamAttribute_PTR() *QXmlStreamAttribute
}

func (ptr *QXmlStreamAttribute) QXmlStreamAttribute_PTR() *QXmlStreamAttribute { return ptr }

func (this *QXmlStreamAttribute) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.Cthis
	}
}
func (this *QXmlStreamAttribute) SetCthis(cthis unsafe.Pointer) {
	if this.CObject == nil {
		this.CObject = &qtrt.CObject{cthis}
	} else {
		this.CObject.Cthis = cthis
	}
}
func NewQXmlStreamAttributeFromPointer(cthis unsafe.Pointer) *QXmlStreamAttribute {
	return &QXmlStreamAttribute{&qtrt.CObject{cthis}}
}
func (*QXmlStreamAttribute) NewFromPointer(cthis unsafe.Pointer) *QXmlStreamAttribute {
	return NewQXmlStreamAttributeFromPointer(cthis)
}

// /usr/include/qt/QtCore/qxmlstream.h:109
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QXmlStreamAttribute()
func NewQXmlStreamAttribute() *QXmlStreamAttribute {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QXmlStreamAttributeC2Ev", qtrt.FFI_TYPE_POINTER)
	qtrt.ErrPrint(err, rv)
	gothis := NewQXmlStreamAttributeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamAttribute)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:110
// index:1
// Public Visibility=Default Availability=Available
// [-2] void QXmlStreamAttribute(const QString &, const QString &)
func NewQXmlStreamAttribute_1(qualifiedName string, value string) *QXmlStreamAttribute {
	var tmpArg0 = NewQString_5(qualifiedName)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(value)
	var convArg1 = tmpArg1.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QXmlStreamAttributeC2ERK7QStringS2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQXmlStreamAttributeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamAttribute)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:112
// index:2
// Public Visibility=Default Availability=Available
// [-2] void QXmlStreamAttribute(const QString &, const QString &, const QString &)
func NewQXmlStreamAttribute_2(namespaceUri string, name string, value string) *QXmlStreamAttribute {
	var tmpArg0 = NewQString_5(namespaceUri)
	var convArg0 = tmpArg0.GetCthis()
	var tmpArg1 = NewQString_5(name)
	var convArg1 = tmpArg1.GetCthis()
	var tmpArg2 = NewQString_5(value)
	var convArg2 = tmpArg2.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN19QXmlStreamAttributeC2ERK7QStringS2_S2_", qtrt.FFI_TYPE_POINTER, convArg0, convArg1, convArg2)
	qtrt.ErrPrint(err, rv)
	gothis := NewQXmlStreamAttributeFromPointer(unsafe.Pointer(uintptr(rv)))
	qtrt.SetFinalizer(gothis, DeleteQXmlStreamAttribute)
	return gothis
}

// /usr/include/qt/QtCore/qxmlstream.h:137
// index:0
// Public Visibility=Default Availability=Available
// [-2] void ~QXmlStreamAttribute()
func DeleteQXmlStreamAttribute(this *QXmlStreamAttribute) {
	rv, err := qtrt.InvokeQtFunc6("_ZN19QXmlStreamAttributeD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 80)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtCore/qxmlstream.h:140
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef namespaceUri()
func (this *QXmlStreamAttribute) NamespaceUri() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QXmlStreamAttribute12namespaceUriEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:141
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef name()
func (this *QXmlStreamAttribute) Name() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QXmlStreamAttribute4nameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:142
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef qualifiedName()
func (this *QXmlStreamAttribute) QualifiedName() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QXmlStreamAttribute13qualifiedNameEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:143
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef prefix()
func (this *QXmlStreamAttribute) Prefix() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QXmlStreamAttribute6prefixEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:148
// index:0
// Public inline Visibility=Default Availability=Available
// [16] QStringRef value()
func (this *QXmlStreamAttribute) Value() *QStringRef /*123*/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QXmlStreamAttribute5valueEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	rv2 := /*==*/ NewQStringRefFromPointer(unsafe.Pointer(uintptr(rv))) // 333
	qtrt.SetFinalizer(rv2 /*==*/, DeleteQStringRef)
	return rv2
}

// /usr/include/qt/QtCore/qxmlstream.h:149
// index:0
// Public inline Visibility=Default Availability=Available
// [1] bool isDefault()
func (this *QXmlStreamAttribute) IsDefault() bool {
	rv, err := qtrt.InvokeQtFunc6("_ZNK19QXmlStreamAttribute9isDefaultEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return rv != 0
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
}

//  keep block end
