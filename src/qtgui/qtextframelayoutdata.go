//  header block begin
// /usr/include/qt/QtGui/qtextobject.h
// #include <qtextobject.h>
// #include <QtGui>
package qtgui

//  header block end

//  ext block begin

/*
#include <stdlib.h>
// extern C begin: 7
*/
// import "C"
import "unsafe"
import "reflect"
import "fmt"
import "qtrt"
import "mkuse/cffiqt"
import "gopp"
import "qtcore"

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
		ffiqt.KeepMe()
	}
	if false {
		gopp.KeepMe()
	}
	if false {
		qtcore.KeepMe()
	}
}

//  ext block end

//  body block begin
type QTextFrameLayoutData struct {
	*qtrt.CObject
}

func (this *QTextFrameLayoutData) GetCthis() unsafe.Pointer {
	return this.Cthis
}
func NewQTextFrameLayoutDataFromPointer(cthis unsafe.Pointer) *QTextFrameLayoutData {
	return &QTextFrameLayoutData{&qtrt.CObject{cthis}}
}

// /usr/include/qt/QtGui/qtextobject.h:114
// index:0
// Public virtual
// void ~QTextFrameLayoutData()
func DeleteQTextFrameLayoutData(*QTextFrameLayoutData) {
	rv, err := ffiqt.InvokeQtFunc6("_ZN20QTextFrameLayoutDataD2Ev", ffiqt.FFI_TYPE_VOID)
	gopp.ErrPrint(err, rv)
}

//  body block end