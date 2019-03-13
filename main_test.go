package main

import (
	"StorageUnitConvert/util"
	"testing"
)

//func Test_UnitConvert(t *testing.T) {
//	var raw = float64(1024.0)
//	var srcGrade = uint(util.FLAG_Byte)
//	var destGrade = uint(util.FLAG_Kilobyte)
//	var rawbyte, res = util.UnitConvert(srcGrade, destGrade, raw)
//	t.Log(rawbyte, res)
//}

func Test_UnitConvertAuto(t *testing.T) {
	testconvert := util.UnitConvertAuto(2048)
	t.Log(testconvert)
}
