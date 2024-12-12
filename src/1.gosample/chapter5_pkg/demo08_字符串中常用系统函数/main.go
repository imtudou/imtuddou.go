package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//1.统计字符串的长度，按字节 len(str)
	//golang的编码统一为utf-8 （ascii的字符（字母和数字）占一个字节，汉字占用3个字节
	str := "上海shanghai"
	fmt.Println("str len=", len(str))

	//2.字符串遍历，同时处理有中文的问题 r ：= []rune(str)
	str2 := "上海shanghai"
	r := []rune(str2)
	for i := 0; i < len(r); i++ {
		fmt.Printf("字符=%c\n", r[i])
	}

	//3.字符串转整数：n, err := strconv.Atoi("12")
	n, err := strconv.Atoi("1bnm,")
	fmt.Printf("n=%v,err=%v \n", n, err)

	//4)整数转字符串 str = strconv.Itoa(12345)
	str = strconv.Itoa(12345)
	fmt.Printf("str=%v\n", str)

	//5字符串 转 []byte： var bytes = []byte("hello go")
	var bytes = []byte("hello go")
	fmt.Printf("bytes=%v\n", bytes)

	//6)[]byte 转 字符串： str = string([]byte{97, 98, 99})
	str = string([]byte{97, 98, 99})
	fmt.Println("str=", str)

	//7.查找子串是否在指定的字符串中：strings.Contains(“seafood"，"foo")//true
	fmt.Println(strings.Contains("seafood", "foo"))

	//8.统计一个字符串有几个指定的子串 ： strings.Count("ceheese"，"e")//4
	fmt.Println(strings.Count("ceheese", "e"))

	//10）不区分大小写的字符串比较(==是区分字母大小写的）：fmt.Println（strings.EqualFold（"abc"，Abc)
	fmt.Println(strings.EqualFold("abc", "Abc"))

	//11)返回子串在字符串第一次出现的index值，如果没有返回-1//strings.Index("NLT abc", "abc") // 4
	fmt.Println(strings.Index("NLT abc", "abc"))

	//12)返回子串在字符串最后一次出现的 index，如没有返回-1: strings.LastIndex("go golang"，"go"

	fmt.Println(strings.LastIndex("golang", "g"))

	//13)将指定的子串替换成另外一个子串: strings Replace("go go hello", "go", "go语言", n) n可以指定你希望替换几个，如果 n=-1 表示全部替换

	fmt.Println(strings.Replace("aaabbbccc", "b", "d", 1))
	sstr := "aaabbbccc"
	fmt.Println(strings.Replace(sstr, "b", "d", 1))

	//14)按照指定的某个字符,为分割标识,将一个字符串拆分成字符串数组,strings.Split("hello,wrold,ok", ".")
	arr := strings.Split("hello,wrold,ok", ",")
	fmt.Println(arr)

	//15)将字符串的字母进行大小写的转换: strings.ToLower(Go") // go strings.ToUpper("Go") // Go
	fmt.Println(strings.ToLower("strings.ToLower"))
	fmt.Println(strings.ToUpper("strings.ToUpper"))

	// 16) 将字符串左右两边的空格去掉： strings.TrimSpace(" tn a lone gopher ntmI
	fmt.Println(strings.TrimSpace(" tn a lone gopher ntmI "))
	//17) 将字符串左右两边指定的字符去掉：strings.Trim("! hello!"."！")/["hello"]//将左右两边 ！和""去掉
	//18)将字符串左边指定的字符去掉: strings.TrimLeft("! hello! ", "!") // ["hello"] //将左边!和 ""去掉
	//19) 将字符串右边指定的字符去掉 ：strings.TrimRight("! hello!"，"！")/["hello"]//将右边 ！和""去掉
	//20) 判断字符串是否以指定的字符串开头: strings.HasPrefix("ftp://192.168.10.1","ftp") // true
	//21)判断字符串是否以指定的字符串结束: strings.HasSuffixI"NLT_abc.jpg", "abc") //false
}
