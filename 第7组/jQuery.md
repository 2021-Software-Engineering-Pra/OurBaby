# jQuery

#### 介绍

##### 名词解释：

​	selector选择器，action动作

###### 基本语法

​	$(selector).action(function(){});//有时候action()括号内部不需要参数或参数不为函数。

$指代jQuery

分别负责选择，事件，动作。

jQuery的函数会在DOM树加载完成之后才会开始对其进行修改。（$(document).ready()）

是一个JS函数库，写得少做的多。

这种以过程为导向的框架真的很没效率。

2020年已经被淘汰，是一种落后的技术，但是仍然有学习价值。

#### 选择器

基于CSS选择器，常用的选择器有：

- 元素选择器,$("element")

- id选择器,$("#id")

- 类选择器,$(".test")

- 通配符$("*")

- 复合选择器：

	| ∪              | $("xx,xx")                 |
	| :------------- | -------------------------- |
	| ∩              | $("xx.xx")                 |
	| 所有后代       | $("xx xx")                 |
	| 第一代，子元素 | $("xx>xx")                 |
	| 相邻兄弟       | $("prevSibing+nextSibing") |
	| 后面的所有兄弟 | $("prev~sibing")           |
	| 属性           | $("xx[xx]")                |

- 属性选择器[]

- 过滤选择器:

- 筛选选择器.

	这三个个人认为没必要记，需要再去学。

#### 事件

初始化DOM树

```js
	$(document).ready(function(){
        继续添加jquery语句。
    });
```

去掉了on-的DOM事件

- 鼠标
	- 点击click
	- 双击dbclick
	- 移入mouseenter
	- 移出mouseleave
	- 在上面hover
	- 按下mousedown
	- 弹起mouseup
- 键盘
	- 按下弹起keypress
	- 按下keydown
	- 弹起keyup
- 表单
	- submit提交
	- 输入字段改变，change
	- 聚焦focus
	- 失焦blur
- 文档|窗口
	- resize调整浏览器窗口大小
	- scroll滚动元素

#### 动作

*就是在function(){里面的代码}。*

##### 效果

###### 隐藏和显示

$("p").hide(time,way,func);	$("p").show(time);

```js
$("p").hide(1000,"linear",function(){alert("111")});//线性隐藏1000ms，之后执行反馈函数，提示111。
```

.toggle();

显示和隐藏的开关，相当于show()|hide()



$(selector).fadeIn( |slow|2000);

将display:none的元素逐渐呈现

$(selector).fadeIOut( |slow|2000);

将display:none的元素逐渐呈现

.fadeToggle()

上面两个的开关

渐隐fade：不透明度减到0。

隐藏hide：变小直至消失，块级元素的隐藏和滑上去很像。

.fadeTo(speed,opacity,callback);

渐变到指定的不透明度。

###### 滑动

slideDown滑下来

slideUp滑上去

slideToggle

###### 动画

animate({},speed,callback)

第一个是必需的。

里面是CSS代码，可以写的有

1. 绝对值
2. 相对值
3. 预定义的值
4. 队列，依次实现各类效果。

stop()停止jQuery效果

stop(stopAll,gotoEnd)这两个参数默认是false,第一个参数会停止所有的任务队列，第二个参数会直接完成动画。

当没有callback的时候，会先执行下一条语句，有的话会先执行callback。

###### 链

$("#p1").css("color","red").slideUp(2000).slideDown(2000);

形如以上的链条可以队列执行这些效果。

##### DOM树

##### 获取（查

###### 获取内容

text()返回会所选元素的文本内容

html()返回所选元素的html内容

val()返回表单字段的值，修改value属性值

在括号内添加字符串参数可以修改内容

###### 获取属性

attr("xx")获取元素的属性为xx的值，找不到就返回undefined

attr("xx","xxx")将xx属性的值修改为xxx

prop("xx")和attr的区别是，它用于本身就带有的固有属性，找不到会返回""空字符串。

##### 修改（改

text(),html(),val()

里面可以填值也可以填函数，这个函数名称是回调函数，呈现出返回值。

attr("xx","xxx")修改xx属性的值为xxx。

xxx也可以是回调函数。

##### 添加（增

append()和prepend()

在元素**内部**的结尾和开头插入内容

after()和before()

在元素的后面和前面插入dom树

##### 删除（删

remove()

删除被选元素和子元素

empty()

删除子元素

括号内参数代表过滤出该元素。要删只删它一个，但是如果同级没有而子元素有，这个子元素不会被删除。

##### 操作CSS

addClass("xx")给元素添加xx类，使其带有那个类的css属性。

removeClass("xx")删除xx类

toggleCLass()

###### css()

​	css("*propertyname*");返回指定的属性值

​	css("*propertyname*","*value*");更改指定属性值

​	css({"*propertyname*":"*value*","*propertyname*":"*value*",...});更改多个指定属性值。

##### 尺寸

width()height()获取或修改元素宽高

innerWidth()和innerHeight()获取或修改元素element和内边距padding的宽高

outerWidth()和outerHeight()获取或修改元素element和内边距padding和边框border的宽高

##### 遍历

###### 树遍历

parent()第一任父元素

parents()所有父元素直到根节点

$("ss").parentsUntil("zz")从ss向父节点遍历直到zz节点。

children()第一任子元素，children("p.1")类名为1的p元素。

find("xx")后代的所有xx元素

siblings()所有兄弟元素

next()下一个兄弟元素

nextAll()之后所有的兄弟元素

nextUntil()直到xx的所有兄弟元素

将next换为prev，就是更换了遍历方向。

##### 过滤

3个节点,div p p p

$("div p").first()=eq(0)

$("div p").last()=eq(2)

filter()过滤

not()除此之外

#### Ajax（异步的js，xhtml）

通过HTTP请求加载远程数据。

$(selector).load(URL,data,callback);

url必需，希望加载的URL。

$.get(URI,callback);

url必需，请求数据

$.post(*URL,data,callback*);

向服务器提交数据

详情[[Ajax]]



$.noConflict();

使得其他脚本也能使用$符号



JSONP

由于同源策略（一种安全协议）使得只有script这个元素可以访问其他网站，所以用JSONP解决跨域数据访问问题。

*这个为什么要用jquery实现？js不香么？*

[[json]]

