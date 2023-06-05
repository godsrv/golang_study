### string结构

![](https://store.babyzoo.club/study/string.png)

**编码知识点：**

  ASCII: 规定了128个字符的编码 占用一个字节

  ASCII码中，一个英文字母(不区分大小写) 占用一个字节的空间，C语音中而一个汉字占用2个字节的空间。

GBK编码：GBK是GB2312的扩展，它能够表示更多的汉字和符号。GBK同样采用双字节编码，每个汉字占用两个字节，而一些罕用汉字则使用四个字节来表示。对于非汉字字符，GBK与GB2312兼容，仍然占用一个字节。

Unicode：国际标准码，融合了目前人类使用的所有字符，为每个字符分配唯一的字符码，所有的文字都用两个字节来表示；

UTF-8编码：UTF-8是一种变长编码方式，它能够表示Unicode字符集中的所有字符，包括汉字、拉丁字母、数字、符号等等。在UTF-8中，英文字母和数字占用一个字节，汉字占用三个字节，一些特殊的符号可能占用更多的字节。

**string结构体知识点**

1. 两部分组成： 指向底层[]byte数组的指针 和 长度

   2.string常量会在编译期分配到只读段，对应数据地址不可写入，故string不支持修改。

   3. 要修改必须转[]byte，string和[]byte转换，会将这段只读内存的数据复制到堆/栈上。

**string使用**

```
1.长度计算
s1 := "s时间"
len(s1) = 7
len 返回的是字节数，如果要返回字符串长度
s3 := utf8.RuneCountInString(s2) = 3

2.类型转换
string转成int：int, err := strconv.Atoi(string)
string转成int64：int64, err := strconv.ParseInt(string, 10, 64)
int转成string： string := strconv.Itoa(int)：
int64转成string： string := strconv.FormatInt(int64,10)
字符串到float32/float64： 
    float32, err = ParseFloat(string, 32)     
    float64,err = ParseFloat(string,64)

interface{}类型转换成string类型
    value, ok := a.(string)
```
