# Golang语言之JSON解码函数Unmarshal

## func Unmarshal

`func Unmarshal(data []byte, v interface{}) error`

unmarshal函数解析json编码的数据并将结果存入v指向的值。

Unmarshal和Marshal做相反的操作，必要时申请映射、切片或指针，有如下的附加规则：

要将json数据解码写入一个指针，Unmarshal函数首先处理json数据是json字面值null的情况。此时，函数将指针设为nil；否则，函数将json数据解码写入指针指向的值；如果指针本身是nil，函数会先申请一个值并使指针指向它。

要将json数据解码写入一个结构体，函数会匹配输入对象的键和Marshal使用的键（结构体字段名或者它的标签指定的键名），优先选择精确的匹配，但也接受大小写不敏感的匹配。

要将json数据解码写入一个接口类型值，函数会将数据解码为如下类型写入接口：

```javascript
bool, for JSON booleans
float64, for JSON numbers
string, for JSON strings
[]interface{}, for JSON arrays
map[string]interface{}, for JSON objects
nil for JSON null
```

如果一个JSON值不匹配给出的目标类型，或者如果一个json数字写入目标类型时溢出，Unmarshal函数会跳过该字段并尽量完成其余的解码操作。如果没有出现更加严重的错误，本函数会返回一个描述第一个此类错误的详细信息的UnmarshalTypeError。

JSON的null值解码为go的接口、指针、切片时会将它们设为nil，因为null在json里一般表示“不存在”。 解码json的null值到其他go类型时，不会造成任何改变，也不会产生错误。

当解码字符串时，不合法的utf-8或utf-16代理（字符）对不视为错误，而是将非法字符替换为unicode字符U+FFFD。

## 普通JSON

示例代码：

```javascript
package main

import (
   "encoding/json"
   "fmt"
)

// Actress 女演员
type Actress struct {
   Name       string
   Birthday   string
   BirthPlace string
   Opus       []string
}

func main() {

   // 普通JSON
   // 因为json.UnMarshal() 函数接收的参数是字节切片，   // 所以需要把JSON字符串转换成字节切片。   
   jsonData := []byte(`{
      "name":"迪丽热巴",
      "birthday":"1992-06-03",
      "birthPlace":"新疆乌鲁木齐市",
      "opus":[
         "《阿娜尔罕》",
         "《逆光之恋》",
         "《克拉恋人》"
      ]
   }`)

   var actress Actress
   err := json.Unmarshal(jsonData, &actress)
   if err != nil {
      fmt.Println("error:", err)
      return
   }
   fmt.Printf("姓名：%s\n", actress.Name)
   fmt.Printf("生日：%s\n", actress.Birthday)
   fmt.Printf("出生地：%s\n", actress.BirthPlace)
   fmt.Println("作品：")
   for _, val := range actress.Opus {
      fmt.Println("\t", val)
   }
}
```

运行结果：

姓名：迪丽热巴

生日：1992-06-03

出生地：新疆乌鲁木齐市

作品：

  《阿娜尔罕》

  《逆光之恋》

  《克拉恋人》

## JSON内嵌普通JSON

示例代码：

```javascript
package main

import (
   "encoding/json"
   "fmt"
)
// Opus 作品
type Opus struct {
   Date string
   Title string
}
// Actress 女演员
type Actress struct {
   Name       string
   Birthday   string
   BirthPlace string
   Opus       Opus
}

func main () {
   // JSON嵌套普通JSON
   jsonData := []byte(`{
      "name":"迪丽热巴",
      "birthday":"1992-06-03",
      "birthPlace":"新疆乌鲁木齐市",
      "opus": {
         "Date":"2013",
         "Title":"《阿娜尔罕》"
      }
   }`)
   var actress Actress
   err := json.Unmarshal(jsonData, &actress)
   if err != nil {
      fmt.Println("error:", err)
      return
   }
   fmt.Printf("姓名：%s\n", actress.Name)
   fmt.Printf("生日：%s\n", actress.Birthday)
   fmt.Printf("出生地：%s\n", actress.BirthPlace)
   fmt.Println("作品：")
fmt.Printf("\t%s:%s", actress.Opus.Date,   actress.Opus.Title)}
```

运行结果：

姓名：迪丽热巴

生日：1992-06-03

出生地：新疆乌鲁木齐市

作品：

 2013:《阿娜尔罕》

## JSON内嵌数组JSON

示例代码：

```javascript
package main

import (
   "encoding/json"
   "fmt"
)

type Opus struct {
   Date string
   Title string
}
type Actress struct {
   Name string
   Birthday string
   BirthPlace string
   Opus []Opus
}

func main () {
   // JSON嵌套数组JSON
   jsonData := []byte(`{
      "name":"迪丽热巴",
      "birthday":"1992-06-03",
      "birthPlace":"新疆乌鲁木齐市",
      "opus":[
         {
            "date":"2013",
            "title":"《阿娜尔罕》"
         },
         {
            "date":"2014",
            "title":"《逆光之恋》"
         },
         {
            "date":"2015",
            "title":"《克拉恋人》"
         }
      ]
   }`)
   var actress Actress
   err := json.Unmarshal(jsonData, &actress)
   if err != nil {
      fmt.Println("error:", err)
      return
   }
   fmt.Printf("姓名：%s\n", actress.Name)
   fmt.Printf("生日：%s\n", actress.Birthday)
   fmt.Printf("出生地：%s\n", actress.BirthPlace)
   fmt.Println("作品：")
   for _, val := range actress.Opus {
      fmt.Printf("\t%s - %s\n", val.Date, val.Title)
   }
}
```

运行结果：

姓名：迪丽热巴

生日：1992-06-03

出生地：新疆乌鲁木齐市

作品：

 2013 - 《阿娜尔罕》

 2014 - 《逆光之恋》

 2015 - 《克拉恋人》

## JSON内嵌具有动态Key的JSON

示例代码：

```javascript
package main

import (
   "encoding/json"
   "fmt"
)

// Opus 作品
type Opus struct {
   Type string
   Title string
}
// Actress 女演员
type Actress struct {
   Name string
   Birthday string
   BirthPlace string
   Opus map[string]Opus
}

func main () {
   jsonData := []byte(`{
      "name":"迪丽热巴",
      "birthday":"1992-06-03",
      "birthPlace":"新疆乌鲁木齐市",
      "opus":{
         "2013":{
            "Type":"近代革命剧",
            "Title":"《阿娜尔罕》"
         },
         "2014":{
            "Type":"奇幻剧",
            "Title":"《逆光之恋》"
         },
         "2015":{
            "Type":"爱情剧",
            "Title":"《克拉恋人》"
         }
      }
   }`)
   var actress Actress
   err := json.Unmarshal(jsonData, &actress)
   if err != nil {
      fmt.Println("error:", err)
      return
   }
   fmt.Printf("姓名：%s\n", actress.Name)
   fmt.Printf("生日：%s\n", actress.Birthday)
   fmt.Printf("出生地：%s\n", actress.BirthPlace)
   fmt.Println("作品：")
   for index, value := range actress.Opus {
      fmt.Printf("\t日期：%s\n", index)
      fmt.Printf("\t\t分类：%s\n", value.Type)
      fmt.Printf("\t\t标题：%s\n", value.Title)
   }
}
```

运行结果：

姓名：迪丽热巴

生日：1992-06-03

出生地：新疆乌鲁木齐市

作品：

 日期：2013

 分类：近代革命剧

 标题：《阿娜尔罕》

 日期：2014

 分类：奇幻剧

 标题：《逆光之恋》

 日期：2015

 分类：爱情剧

 标题：《克拉恋人》

## 总结

我们先是介绍了Golang标准库的encoding/json包中的Unmarshal函数，然后通过上面4个示例代码，分别介绍了如何解码以下4种JSON格式数据：

```javascript
JSON格式1：
{
    "name":"迪丽热巴",
    "birthday":"1992-06-03",
    "birthPlace":"新疆乌鲁木齐市",
    "opus":[
        "《阿娜尔罕》",
        "《逆光之恋》",
        "《克拉恋人》"
    ]
}
JSON格式2：
{
    "name":"迪丽热巴",
    "birthday":"1992-06-03",
    "birthPlace":"新疆乌鲁木齐市",
    "opus":{
        "Date":"2013",
        "Title":"《阿娜尔罕》"
    }
}
JSON格式3：
{
    "name":"迪丽热巴",
    "birthday":"1992-06-03",
    "birthPlace":"新疆乌鲁木齐市",
    "opus":[
        {
            "date":"2013",
            "title":"《阿娜尔罕》"
        },
        {
            "date":"2014",
            "title":"《逆光之恋》"
        },
        {
            "date":"2015",
            "title":"《克拉恋人》"
        }
    ]
}
JSON格式4：
{
    "name":"迪丽热巴",
    "birthday":"1992-06-03",
    "birthPlace":"新疆乌鲁木齐市",
    "opus":{
        "2013":{
            "Type":"近代革命剧",
            "Title":"《阿娜尔罕》"
        },
        "2014":{
            "Type":"奇幻剧",
            "Title":"《逆光之恋》"
        },
        "2015":{
            "Type":"爱情剧",
            "Title":"《克拉恋人》"
        }
    }
}
```