## MVVM

model-view-viewmodel

名词解释：

view视图，给用户看的。写在html-body-div id="app"

model模型，比view深一层的。写在html-script的前面

viewModel,视图模型，Vue实例，new view;vue的核心，是前两个的桥。写在html-script的后面。

#### 发展历史：

1. 原生API
2. jQuery
3. MVC模式
4. MVVM

MVVM相比于jQuery，更加高内聚低耦合，简单。

从操作DOM变成了如何更新JS的状态。



它将DOM控制的的视图和JS控制的模型分离了，然后利用viewmodel将两者联系，视图的数据由[[DOM监听]]传给model，模型的数据由[[绑定]]传给视图。

#### 单项绑定

默认：vue会自动监听Model的任何变化，并实时反映在view上。

#### 双向绑定，关键字：v-model

在填写表单时，需要将view的数据传给model。

```html
<input type="text" v-model="msg"/>//用户在表单中改变msg时，model的msg也会发生改变。
```

不同表单v-model数据的类型要和model的数据类型相同，比如多个checkbox是数组，单个checkbox是bool，select下拉框是string

#### 处理事件，关键字v-on

例子

```html
<form id='vm' v-on:submit.prevent='register'>
```

```js
methods:{
	register:function(){
		...
	}
}
```

处理form提交事件时，采用了submit方法，其中.prevent的作用是阻止[[事件冒泡]]，register是事件处理函数的名字。



#### 操作DOM树的简单指令（条件渲染），关键字v-if，v-for等

1. 在标签内部v-if="*expression*"，true则显示这个标签，false则隐藏这个标签。

	```html
	<h1 v-if="...">标题</h1>
	```

2. v-show,元素一定会被渲染

3. 

	



