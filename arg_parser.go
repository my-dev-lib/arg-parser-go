package arg_parser_go

import (
	"flag"
	"fmt"
)

const VERSION = "1.0.1"

var TypeInt = "int"
var TypeBool = "bool"
var TypeUint = "uint"
var TypeString = "string"
var TypeUint64 = "uint64"
var TypeFloat64 = "float64"

type argInfo struct {
	typo string
	val  any
}

type ArgParser struct {
	args [][]any
}

func NewArgParser(args [][]any) *ArgParser {
	return &ArgParser{args: args}
}

func parseTypeArgValue(typeArg *argInfo) any {
	switch typeArg.typo {
	case TypeInt:
		return *(typeArg.val.(*int))
	case TypeBool:
		return *(typeArg.val.(*bool))
	case TypeUint:
		return *(typeArg.val.(*uint))
	case TypeString:
		return *(typeArg.val.(*string))
	case TypeUint64:
		return *(typeArg.val.(*uint64))
	case TypeFloat64:
		return *(typeArg.val.(*float64))
	default:
		return 0
	}
}

func (ap *ArgParser) Parse() (map[string]any, error) {
	args := map[string]*argInfo{}
	for _, arg := range ap.args {
		name := arg[0].(string)
		typo := arg[1].(string)
		help := arg[2].(string)
		def := arg[3]

		switch typo {
		case TypeInt:
			defVal, ok := def.(int)
			if !ok {
				return nil, defValError(name, def)
			}

			val := flag.Int(name, defVal, help)
			args[name] = &argInfo{typo: typo, val: val}
		case TypeBool:
			defVal, ok := def.(bool)
			if !ok {
				return nil, defValError(name, def)
			}

			val := flag.Bool(name, defVal, help)
			args[name] = &argInfo{typo: typo, val: val}
		case TypeUint:
			defVal, ok := def.(uint)
			if !ok {
				return nil, defValError(name, def)
			}

			val := flag.Uint(name, defVal, help)
			args[name] = &argInfo{typo: typo, val: val}
		case TypeString:
			defVal, ok := def.(string)
			if !ok {
				return nil, defValError(name, def)
			}

			val := flag.String(name, defVal, help)
			args[name] = &argInfo{typo: typo, val: val}
		case TypeUint64:
			defVal, ok := def.(uint64)
			if !ok {
				return nil, defValError(name, def)
			}

			val := flag.Uint64(name, defVal, help)
			args[name] = &argInfo{typo: typo, val: val}
		case TypeFloat64:
			defVal, ok := def.(float64)
			if !ok {
				return nil, defValError(name, def)
			}

			val := flag.Float64(name, defVal, help)
			args[name] = &argInfo{typo: typo, val: val}
		default:
			return nil, fmt.Errorf("不支持此参数类型，%s, type: %v", name, typo)
		}
	}

	flag.Parse()

	var ret = map[string]any{}
	for key, typeArg := range args {
		value := parseTypeArgValue(typeArg)
		ret[key] = value
	}

	return ret, nil
}

func (ap *ArgParser) PrintHelp() {
	flag.Usage()
}

func defValError(name string, def any) error {
	return fmt.Errorf("参数默认值类型错误, %s, def: %v", name, def)
}

/*
func main() {
	// 声明参数
	argParser := NewArgParser([][]any{
		// 参数名、参数类型、帮助文本、参数默认值
		{"testInt", TypeInt, "TypeInt help", 320},
		{"testBool", TypeBool, "TypeBool help", false},
		{"testUint", TypeUint, "TypeUint help", uint(32)},
		{"testString", TypeString, "TypeString help", "stringV"},
		{"testString2", TypeString, "TypeString help", "stringArg2"},
		{"testString3", TypeString, "TypeString help", "stringArg3"},
		{"testUint64", TypeUint64, "TypeUint64 help", uint64(64)},
		{"testFloat64", TypeFloat64, "TypeFloat64 help", 1234.0},
	})

	// 解析参数
	ret, err := argParser.Parse()
	if err != nil {
		fmt.Printf("err: %v", err)
		return
	}

	// 获取参数值
	fmt.Println("testInt: " + fmt.Sprint(ret["testInt"].(int)))
	fmt.Println("testBool: " + fmt.Sprint(ret["testBool"].(bool)))
	fmt.Println("testUint: " + fmt.Sprint(ret["testUint"].(uint)))
	fmt.Println("testString: " + ret["testString"].(string))
	fmt.Println("testString2: " + ret["testString2"].(string))
	fmt.Println("testString3: " + ret["testString3"].(string))
	fmt.Println("testUint64: " + fmt.Sprint(ret["testUint64"].(uint64)))
	fmt.Println("testFloat64: " + fmt.Sprint(ret["testFloat64"].(float64)))
}

// */
