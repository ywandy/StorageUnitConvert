package cli

import (
	"StorageUnitConvert/util"
	"flag"
	"fmt"
)

var G_CmdVal *CmdStruct

type CmdStruct struct {
	ConvertType *string
	RawValue    *float64
	BlockSize   *int64
}

func InitCmd() {
	var ConvertType string
	var RawValue float64
	var BlockSize int64
	flag.StringVar(&ConvertType, "t", "auto", "auto : 自动转换")
	flag.Float64Var(&RawValue, "r", 1024, "原始数据")
	flag.Int64Var(&BlockSize, "b", 512, "块大小")
	cs := CmdStruct{ConvertType: &ConvertType, RawValue: &RawValue, BlockSize: &BlockSize}
	flag.Parse()
	G_CmdVal = &cs
}

func ProcessCmd() {
	switch *G_CmdVal.ConvertType {
	case "auto":
		raw_val := *G_CmdVal.RawValue
		block_size := *G_CmdVal.BlockSize
		val_res := util.UnitConvertAuto(raw_val * float64(block_size))
		fmt.Println(val_res)
	default:
		fmt.Println("参数错误")
	}
}
