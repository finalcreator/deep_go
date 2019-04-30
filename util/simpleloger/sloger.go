package simpleloger

import (
	"log"
)

//type LogChecker interface {
//	chk(err error)
//}

//type Simpleloger struct{}
//
//var Sloger *Simpleloger

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	//Sloger = &Simpleloger{}
}

//func (sl *Simpleloger) chkError(err error) {
//	log.Printf("%s\n", err)
//}

func PanicError(err error) {
	if err != nil {
		panic(err)
	}

}

//func ChkError(err error) {
//	if err != nil {
//		panic(err)
//	}
//
//}

//func ChkError(err error)  {
//	if err != nil {
//		Sloger.chkError(err)
//
//	}
//	return nil
//}

