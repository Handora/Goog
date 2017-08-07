# API文档
#写得比较简陋不要在意
#你可以先测试起来<br>

**Get Articles(20/per request) :**
```javascript
url: '../api/v1.0/articles?page=(page)'
method: GET
args: 
	page // the articles between (page-1)*20 
        // to page*20 will be returned to you
        
response: 
{
	Code //200 is ok, else is error
    Text //if error, there will be some error message
    Body //an array of Artcile struct
    	 // 我没有给count, 想要的话跟我说
}

Article struct:
{
	id      int        // sql id
	title   string     // 标题
	content string     //文章内容
	time    time.Time  //yyyy-mm-dd hh:mm:ss
    // 类似sql里面的datetime数据类型。
}
```
<br>
<br>

**Get Article() :**
```javascript
{
    url: '../api/v1.0/articles/(id)'
    method: GET
    args: 
        id // the article sql id

    response: 
    {
        Code //200 is ok, else is error
        Text //if error, there will be some error message
        Body //a Artcile struct
    }

    Article struct:
    {
        id      int        // sql id
        title   string     // 标题
        content string     //文章内容
        time    time.Time  //yyyy-mm-dd hh:mm:ss
        // 类似sql里面的datetime数据类型。
    }
}
```
<br>
**Get Tags:**
```javascript
{
	url: '../api/v1.0/tags'
	method: GET
	args: 
		page // the articles between (page-1)*20 
        	 // to page*20 will be returned to you
      
        
    response: 
    {
        Code //200 is ok, else is error
        Text //if error, there will be some error message
        Body //an array of Article struct
    }
 	
    Article struct:
    {
        id      int        // sql id，可能无用
        title   string     // 标题
        content string     //文章内容
        time    time.Time  //yyyy-mm-dd hh:mm:ss
        // 类似sql里面的datetime数据类型。
    }
    
}
```

<br>
**Get Articles through Tag id:**
```javascript
{
	url: '../api/v1.0/tags/(id)/articles?page=(page)'
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
        id      int        // sql id，通过id找文章
        name    string     // tags' name
    }
}
```
<br/>
备注:
POST方面如果你想写个管理界面，我就把api给你，如果你不想的写的话，就通过命令行管理，我来写好了.我哪里没写跟我说，如果有改动我也会跟你说的，我写了大致一半，休息2天在写<br>~(>-<)~