# API文档
#写得比较简陋不要在意<br>
#你可以先测试起来<br><br>
// 获得page内的的文章,通过query里的page来确定页数(若没有page，那么视为page=1)<br />
**Get Articles(20/per request) :**
```javascript
url: '118.89.182.24/articles?page=(page)'
method: GET
query:
	page // the articles between (page-1)*20 
         // to page*20 will be returned to you
         // this is a query argument
        
response: 
{
    // 注意首字母大写
	Code //200 is ok, else is error
    Text //if error, there will be some error message
    Body //an array of Artcile struct
    	 // 我没有给count, 想要的话跟我说
}

Article struct:
{
    // 注意首字母大写
	Id      int        // sql id
	Title   string     // 标题
	Content string     //文章内容
	Time    time.Time  //yyyy-mm-dd hh:mm:ss
    // 类似sql里面的datetime数据类型。
}
```
<br>
<br>

**Get Article() :**
```javascript
{
    url: '../api/v1.0/articles/{id}'
    method: GET
    args: 
        id // the article sql id

    response: 
    {
    	// 注意首字母大写
        Code //200 is ok, else is error
        Text //if error, there will be some error message
        Body //a Artcile struct if found
    }

    Article struct:
    {
        // 注意首字母大写
        Id      int        // sql id
        Title   string     // 标题
        Content string     //文章内容
        Time    time.Time  //yyyy-mm-dd hh:mm:ss
        // 类似sql里面的datetime数据类型。
    }
}
```
<br><br>
// Get all tags, with no arguements<br>
**Get Tags:**
```javascript
{
	url: '../api/v1.0/tags'
	method: GET
	args: 
        
    response: 
    {
        Code //200 is ok, else is error
        Text //if error, there will be some error message
        Body //an array of Tag struct
    }
 	
    Tag struct:
    {
        Id   int,      // the sql id
        name string,   // the name of tags
    }
}
```
<br><br>
**Get Tag:**
```javascript
{
	url: '../api/v1.0/tags/{id}'
	method: GET
	args: 
    	id   // the sql id of one tag
        
    response: 
    {
        Code //200 is ok, else is error
        Text //if error, there will be some error message
        Body //a Tag struct if found
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
	url: '../api/v1.0/tags/{id}/articles?page=(page)'
	method: GET
	args: 
    	id   // the sql id of one tag
        page // the articles of this tag between (page-1)*20 
             // to page*20 will be returned to you
             // this is a query argument

        
    response: 
    {
        Code //200 is ok, else is error
        Text //if error, there will be some error message
        Body //an array of Article struct
    }

    Article struct:
    {
        // 注意首字母大写
        Id      int        // sql id
        Title   string     // 标题
        Content string     //文章内容
        Time    time.Time  //yyyy-mm-dd hh:mm:ss
        // 类似sql里面的datetime数据类型。
    }
}
```
<br/><br>
**Get Tags through Articles:**
```javascript
{
	url: '../api/v1.0/articles/{id}/tags?'
	method: GET
	args: 
    	id   // the sql id of one article
        
    response: 
    {
        Code //200 is ok, else is error
        Text //if error, there will be some error message
        Body //an array of Tag struct
    }

    Tag struct:
    {
        // 注意首字母大写
        Id      int        // sql id
        Name    String     // name of tag
    }
}
```
备注:
POST方面如果你想写个管理界面，我就把api给你，如果你不想的写的话，就通过命令行管理，我来写好了.我哪里没写跟我说，如果有改动我也会跟你说的, 之后我把命令行程序给你<br>~(>-<)~