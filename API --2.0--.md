# API文档 **2.0**

获得page内的的文章,通过query里的page来确定页数(若没有page，那么视为page=1)
**Get Articles(5/per request) :**
```javascript
url: '118.89.182.24/articles?page=(page)'
method: GET
query:
	page // 页数代表文章是第(page-1)*5-page*5
        
response: 
{
	Code //200 代表正确，其他代表相应错误，参照HTTP状态码
	     //实际没有必要，以后考虑删去
    Text //错误的话，其中保存错误原因
    Body // {
         //   “0”：              <article>  => 通用结构看最后的结构
         //   ...                ...
         //   "4":               <article>
         //   "total":           总页数
         //   “currentPage”：    当前页数，从1开始
            }
}
```
<br>
<br>

**Get Article :**
```javascript
{
    url: '118.89.182.24/articles/{id}'
    method: GET
    args: 
        id // 文章id，可以通过各种方法取得，例如上例中返回的Article结构

    response: 
    {
        Code //200 代表正确，其他代表相应错误，参照HTTP状态码
	         //实际没有必要，以后考虑删去
        Text //错误的话，其中保存错误原因
        Body //单个 Article 结构
    }
}
```
<br><br>
**Get Tags:**
```javascript
{
	url: '118.89.182.24/api/v1.0/tags'
	method: GET
	args: 
        
    response: 
    {
        Code //200 代表正确，其他代表相应错误，参照HTTP状态码
	         //实际没有必要，以后考虑删去
        Text //错误的话，其中保存错误原因
        Body //以数组存放的tag结构, tag[], 数据结构见最后
    }
}
```
<br><br>
**Get Tag:**
```javascript
{
	url: '118.89.182.24/tags/{id}'
	method: GET
	args: 
    	id   // Tag的id， 可以通过很多方法获得，但基本都是通过Tag结构
        
    response: 
    {
        Code //200 代表正确，其他代表相应错误，参照HTTP状态码
	         //实际没有必要，以后考虑删去
        Text //错误的话，其中保存错误原因
        Body //一个相应的Tag结构
    }
 	
    Tag struct:
    {
        Id   int,      // the sql id
        name string,   // the name of tags
    }
}
```
<br><br>
**Get Articles through Tags:**
```javascript
{
	url: '118.89.182.24/tags/{id}/articles?page=(page)'
	method: GET
	args: 
    	id   // Tag id, 获得方式同理
        page // 页数代表文章是第(page-1)*5-page*5
        
    response: 
    {
        Code //200 代表正确，其他代表相应错误，参照HTTP状态码
	         //实际没有必要，以后考虑删去
        Text //错误的话，其中保存错误原因
        Body // {
             //   “0”：              <article>  => 通用结构看最后的结构
             //   ...                ...
             //   "4":               <article>
             //   "total":           总页数
             //   “currentPage”：    当前页数，从1开始
             // }
    }
}
```
<br/><br>
**Get Tags through Articles:**
```javascript
{
	url: '118.89.182.24/articles/{id}/tags?'
	method: GET
	args: 
    	id   // Article id, 获得方式同理
        
    response: 
    {
        Code //200 代表正确，其他代表相应错误，参照HTTP状态码
	         //实际没有必要，以后考虑删去
        Text //错误的话，其中保存错误原因
        Body //以数组存放的tag结构, tag[], 数据结构见最后
    }
}
```

<br><br>
**各种数据结构**
```go
Article struct:
{
    名字    数据结构   注释
	Id      int        //id
	Title   string     //标题
	Content string     //文章内容
	Tag     Tag[]      //以数组存放的Tag结构
 	Time    time.Time  //yyyy-mm-dd hh:mm:ss
                       // 类似sql里面的datetime数据类型。
}
 	
Tag struct:
{
    名字    数据结构   注释
    Id      int        // id
    name    string     // 标签的中文名字，例如“操作系统”， “网络”等
}
```