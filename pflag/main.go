package main

// REF : https://www.cnblogs.com/sparkdev/p/10833186.html
import (
	"fmt"
	"strings"

	"github.com/spf13/pflag"
)

// 定义命令行参数对应的变量
var cliName = pflag.StringP("name", "n", "nick", "Input Your Name")
var cliAge = pflag.IntP("age", "a",22, "Input Your Age")
var cliGender = pflag.StringP("gender", "g","male", "Input Your Gender")
var cliOK = pflag.BoolP("ok", "o", false, "Input Are You OK")
var cliDes = pflag.StringP("des-detail", "d", "", "Input Description")
var cliOldFlag = pflag.StringP("badflag", "b", "just for test", "Input badflag")

func main() {
	// 设置标准化参数名称的函数 wordSepNormalizeFunc： -分隔符，_分隔符， .分隔符等效
	pflag.CommandLine.SetNormalizeFunc(wordSepNormalizeFunc)

	// 为 age 参数设置 NoOptDefVal。当--age是默认是25，当不传时,age默认是22，当--age=88是,age是88
	pflag.Lookup("age").NoOptDefVal = "25"

	// 把 badflag 参数标记为即将废弃的，请用户使用 des-detail 参数。当用户使用--badflag=abc时，会给出使用--des-detail的建议
	pflag.CommandLine.MarkDeprecated("badflag", "please use --des-detail instead")
	// 把 badflag 参数的 shorthand 标记为即将废弃的，请用户使用 des-detail 的 shorthand 参数。当用户使用-b=abc时，会给出使用--des-detail的shorthhand的建议即-d
	pflag.CommandLine.MarkShorthandDeprecated("badflag", "please use -d instead")

	// 在帮助文档中隐藏参数 gender。在命令行使用-h 或者 --help时，隐藏badflag这个flag的说明
	pflag.CommandLine.MarkHidden("badflag")

	// 把用户传递的命令行参数解析为对应变量的值
	pflag.Parse()

	fmt.Println("name=", *cliName)
	fmt.Println("age=", *cliAge)
	fmt.Println("gender=", *cliGender)
	fmt.Println("ok=", *cliOK)
	fmt.Println("des=", *cliDes)
}

// 保证 --abc_d=123 和 --abc-d=123 和 --abc.d=123是等价的，都可以解析的到
func wordSepNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.Replace(name, sep, to, -1)
	}
	return pflag.NormalizedName(name)
}