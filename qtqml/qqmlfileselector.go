package qtqml

// /usr/include/qt/QtQml/qqmlfileselector.h
// #include <qqmlfileselector.h>
// #include <QtQml>

//  header block end

//  main block begin

//  main block end

//  use block begin

//  use block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 28
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "github.com/kitech/qt.go/qtrt"
import "github.com/kitech/qt.go/qtcore"
import "github.com/kitech/qt.go/qtnetwork"

//  ext block end

//  body block begin

type QQmlFileSelector struct {
	*qtcore.QObject
}
type QQmlFileSelector_ITF interface {
	qtcore.QObject_ITF
	QQmlFileSelector_PTR() *QQmlFileSelector
}

func (ptr *QQmlFileSelector) QQmlFileSelector_PTR() *QQmlFileSelector { return ptr }

func (this *QQmlFileSelector) GetCthis() unsafe.Pointer {
	if this == nil {
		return nil
	} else {
		return this.QObject.GetCthis()
	}
}
func (this *QQmlFileSelector) SetCthis(cthis unsafe.Pointer) {
	this.QObject = qtcore.NewQObjectFromPointer(cthis)
}
func NewQQmlFileSelectorFromPointer(cthis unsafe.Pointer) *QQmlFileSelector {
	bcthis0 := qtcore.NewQObjectFromPointer(cthis)
	return &QQmlFileSelector{bcthis0}
}
func (*QQmlFileSelector) NewFromPointer(cthis unsafe.Pointer) *QQmlFileSelector {
	return NewQQmlFileSelectorFromPointer(cthis)
}

// /usr/include/qt/QtQml/qqmlfileselector.h:54
// index:0
// Public virtual Visibility=Default Availability=Available
// [8] const QMetaObject * metaObject()
func (this *QQmlFileSelector) MetaObject() *qtcore.QMetaObject /*777 const QMetaObject **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QQmlFileSelector10metaObjectEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQMetaObjectFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlfileselector.h:57
// index:0
// Public Visibility=Default Availability=Available
// [-2] void QQmlFileSelector(QQmlEngine *, QObject *)
func NewQQmlFileSelector(engine *QQmlEngine /*777 QQmlEngine **/, parent *qtcore.QObject /*777 QObject **/) *QQmlFileSelector {
	var convArg0 = engine.GetCthis()
	var convArg1 = parent.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QQmlFileSelectorC2EP10QQmlEngineP7QObject", qtrt.FFI_TYPE_POINTER, convArg0, convArg1)
	qtrt.ErrPrint(err, rv)
	gothis := NewQQmlFileSelectorFromPointer(unsafe.Pointer(uintptr(rv)))
	return gothis
}

// /usr/include/qt/QtQml/qqmlfileselector.h:58
// index:0
// Public virtual Visibility=Default Availability=Available
// [-2] void ~QQmlFileSelector()
func DeleteQQmlFileSelector(this *QQmlFileSelector) {
	rv, err := qtrt.InvokeQtFunc6("_ZN16QQmlFileSelectorD2Ev", qtrt.FFI_TYPE_VOID, this.GetCthis())
	qtrt.Cmemset(this.GetCthis(), 9, 16)
	qtrt.ErrPrint(err, rv)
	this.SetCthis(nil)
}

// /usr/include/qt/QtQml/qqmlfileselector.h:59
// index:0
// Public Visibility=Default Availability=Available
// [8] QFileSelector * selector()
func (this *QQmlFileSelector) Selector() *qtcore.QFileSelector /*777 QFileSelector **/ {
	rv, err := qtrt.InvokeQtFunc6("_ZNK16QQmlFileSelector8selectorEv", qtrt.FFI_TYPE_POINTER, this.GetCthis())
	qtrt.ErrPrint(err, rv)
	return qtcore.NewQFileSelectorFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}

// /usr/include/qt/QtQml/qqmlfileselector.h:60
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setSelector(QFileSelector *)
func (this *QQmlFileSelector) SetSelector(selector *qtcore.QFileSelector /*777 QFileSelector **/) {
	var convArg0 = selector.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QQmlFileSelector11setSelectorEP13QFileSelector", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlfileselector.h:61
// index:0
// Public Visibility=Default Availability=Available
// [-2] void setExtraSelectors(QStringList &)
func (this *QQmlFileSelector) SetExtraSelectors(strings *qtcore.QStringList) {
	var convArg0 = strings.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QQmlFileSelector17setExtraSelectorsER11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlfileselector.h:62
// index:1
// Public Visibility=Default Availability=Available
// [-2] void setExtraSelectors(const QStringList &)
func (this *QQmlFileSelector) SetExtraSelectors_1(strings *qtcore.QStringList) {
	var convArg0 = strings.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QQmlFileSelector17setExtraSelectorsERK11QStringList", qtrt.FFI_TYPE_POINTER, this.GetCthis(), convArg0)
	qtrt.ErrPrint(err, rv)
}

// /usr/include/qt/QtQml/qqmlfileselector.h:63
// index:0
// Public static Visibility=Default Availability=Available
// [8] QQmlFileSelector * get(QQmlEngine *)
func (this *QQmlFileSelector) Get(arg0 *QQmlEngine /*777 QQmlEngine **/) *QQmlFileSelector /*777 QQmlFileSelector **/ {
	var convArg0 = arg0.GetCthis()
	rv, err := qtrt.InvokeQtFunc6("_ZN16QQmlFileSelector3getEP10QQmlEngine", qtrt.FFI_TYPE_POINTER, convArg0)
	qtrt.ErrPrint(err, rv)
	return /*==*/ NewQQmlFileSelectorFromPointer(unsafe.Pointer(uintptr(rv))) // 444
}
func QQmlFileSelector_Get(arg0 *QQmlEngine /*777 QQmlEngine **/) *QQmlFileSelector /*777 QQmlFileSelector **/ {
	var nilthis *QQmlFileSelector
	rv := nilthis.Get(arg0)
	return rv
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
	if false {
		qtnetwork.KeepMe()
	}
}

//  keep block end
