package demo

import (
    "fmt"
    "unicode/utf8"
    "strings"
    "unicode"
)

const (
    a string = "hello aa"
    b string = "hello bb"
    r rune = 'a'
    s string = "a"
    t byte = 'b'
    f string = "asd fgh jkl zxcv"
    e string = "as,dfg,xvbg,rty,lkj"
    w string = "AADerrBHJ"
)

//判断两个字符串是否相等
func equal(str1 string, str2 string) bool {
    var str1EqualStr2 bool = strings.EqualFold(str1, str2)
    fmt.Println("字符串str1 ",str1, ", 字符串str2 ", str2, "是否相等", str1EqualStr2)
    if str1EqualStr2 {
        return true
    } else {
        return false
    }
}

func contains (str1 string, str2 string) bool {
    str1ContainsStr2 := strings.Contains(str1, str2)
    if str1ContainsStr2 {
        return true
    } else {
        return false
    }
}

func containsRune (str1 string, str2 rune) bool {
    str1ContainsRuneStr2 := strings.ContainsRune(str1, str2)
    if str1ContainsRuneStr2 {
        return true
    } else {
        return false
    }
}

func containsAny (str1 string, str2 string) bool {
    containsAnyStr := strings.ContainsAny(str1, str2)
    if containsAnyStr {
        return true
    } else {
        return false
    }
}

func count (str1 string, str2 string) int {
    if len(str1) == 0 || len(str2) == 0 {
        return 0
    }
    if str1 == "" || str2 == "" {
        return 0
    }
    countStr := strings.Count(str1, str2)
    return countStr
}

func hasPrefix (str1 string, str2 string) bool {
    hasPrefixStr := strings.HasPrefix(str1, str2)
    if hasPrefixStr {
        return true
    } else {
        return false
    }
}

func lastIndexAny (str1 string, str2 string) int {
    return strings.LastIndexAny(str1, str2)
}

func index (str1 string, str2 string) int {
    return strings.Index(str1, str2)
}

func indexAny (str1 string, str2 string) int {
    return strings.IndexAny(str1, str2)
}

func indexByte (str1 string, str2 byte) int {
    return strings.IndexByte(str1, str2)
}

func lastIndex(str1 string, str2 string) int  {
    lastIndex := strings.LastIndex(str1, str2)
    return lastIndex
}

//这个函数的作用是按照1：n个空格来分割字符串最后返回的是 []string的切片
func fields(str1 string) []string {
    return strings.Fields(str1)
}

func indexRune(str1 string, str2 rune) int {
    indexRune := strings.IndexRune(str1, str2)
    return indexRune
}

func lastIndexByte(str1 string, str2 byte) int {
    lastIndexByte := strings.LastIndexByte(str1, str2)
    return lastIndexByte
}

func split(str1 string, sep string) []string {
    return strings.Split(str1, sep)
}

//这个是切割字符串的时候自己定义长度，如果sep为空，那么每一个字符都分割
func splitN(str1 string, sep string, n int) []string {
    return strings.SplitN(str1, sep, n)
}

//该函数s根据sep分割，返回分割之后子字符串的slice,和split一样，
//只是返回的子字符串保留sep，如果sep为空，那么每一个字符都分割
func splitAfterN(str1 string, sep string, n int) []string {
    return strings.SplitAfterN(str1, sep, n)
}

//这个函数是在前边的切割完成之后再后边在加上sep分割符
func splitAfter(str1 string, sep string) []string {
    return strings.SplitAfter(str1, sep)
}

//这个跟php中的implode差不多，这个函数是将一个[]string的切片通过分隔符，分割成一个字符串
func join(str []string, sep string) string {
    return strings.Join(str, sep)
}

func repeat(str string, n int) string {
    return strings.Repeat(str, n)
}

func toUpper(str string) string {
    return strings.ToUpper(str)
}


func isSlash(r rune) bool {
    return strings.ContainsRune(",|/#", r)
}

func replace(s string, old string, new string, n int) string  {
    return strings.Replace(s, old, new, n)
}

//https://www.jianshu.com/p/f2e3f11bf6cb
//strings包
func main() {
    s := "asSASA ddd dsjkdsjs dk"
    fmt.Println("length of string s is %d", len([]byte(s)))
    fmt.Println("length of string s is %d", utf8.RuneCount([]byte(s)))

    //判断两个字符串是否相等
    fmt.Println("a equal b ", equal(a, b))

    //判断字符串url中是否包含字符串
    fmt.Println("a contains c ", contains(a, s))

    //判断字符串中是否包含字符a(rune('a')为字符的阿斯克码值)
    fmt.Println("containsRune ", containsRune(a, r))

    //判断字符串str中是否包含字符w(多个只要有一个符合就返回true)
    fmt.Println("containsAnyStr ", containsAny(a, s))

    //在一段字符串中有多少匹配到的字符或字符串
    fmt.Println("countStr ", count(a, s))

    //判断字符串str是否以He开头
    fmt.Println("hasPrefixStr ", hasPrefix(a, s))

    //判断字符串str是否以He结尾
    fmt.Println("LastIndexAny ", lastIndexAny(a, s))

    //判断字符或者字符串s第一次在字符串中的位置
    fmt.Println("index ", index(a, s))

    //字符串第一次出现的位置，如果不存在就返回-1
    fmt.Println("indexAny ", indexAny(a, s))

    //判断字符x第一次在字符串中的位置
    fmt.Println("indexByte ", indexByte(a, t))

    //判断字符或字符串最后一次出现的位置
    fmt.Println("lastIndex ", lastIndex(a, s))

    //将字符串按空格分割成数组
    fmt.Println("fields ", fields(f))

    fmt.Println("indexRune ", indexRune(a, r))

    fmt.Println("lastIndexByte ", lastIndexByte(a, t))

    splitStr := split(e, ",")
    fmt.Println("str split ", splitStr)
    fmt.Println("str split 1", splitStr[1])
    fmt.Println("str splitN", splitN(e, ",", 3))
    fmt.Println("str splitAfterN", splitAfterN(e, ",", 4))
    fmt.Println("str splitAfter", splitAfter(e, ","))
    //把切片转换成字符串。
    j  := []string{"qw", "zsd", "serg", "kiuyhb"}
    jStr := strings.Join(j, "**")
    fmt.Println("jStr ", jStr)

    repeatStr := repeat(a, 2)
    fmt.Println("repeat str ", repeatStr)

    fmt.Println("toUpper ", toUpper(a))
    fmt.Println("toLower ", strings.ToLower(w))
    fmt.Println("toTitle ", strings.ToTitle(a))
    var SC unicode.SpecialCase
    fmt.Println("ToUpperSpecial:", strings.ToUpperSpecial(SC, "Gopher"))
    //isSlash := isSlash('#')
    fmt.Println("Trim ", strings.Trim("##rtty#%", "#"))
    //利用string的修改操作，对str操作，用新字符串替代旧字符串。
    replaceStr := replace("qwerqwerqwer", "qwer", "asdf", 2)
    fmt.Println("replace ", replaceStr)
    //使用特定规则对字符串进行替换操作
    //截取字符串
    str := "XBodyContentX"
    content := str[1 : len(str)-1]
    fmt.Println("substr ", content)
}
