## JavaScript的模块实现

#### 	require          import,export

模块

​	相当于游戏里的mod，开发者只要写好自己的业务逻辑，就可以加载别人已经写好的模块。

​	模块的一大特点就是，不同模块之间的数据是不共享的。也称数据封装。



### require和以前的方法

​	为了实现模块的功能，先后用了以下几种方法：

- 函数封装：将一个模块的函数写在一起。缺点显而易见

- 对象封装：将函数写在一个对象内，但是它的变量不是private的

- [[立即执行函数]]：用立即执行函数实现封装。以下的都是以它为基础。

  - [立即执行函数]:立即执行函数.md

    



### 10年前的主流模块规范

- Common（node.js）一个文件就是一个模块

  - 始于2009年，用于服务器，如果用于客户端，会半天加载不出来。

  - 调用和使用模块

    - ```js
      var math = require('math');
      math.add(2.3);add是模块math的函数
      ```

- require.js推广的异步模块定义（AMD），利用函数创建和加载模块

  - 用于客户端浏览器，不必等模块完全加载就可以运行后面的语句。

  - 创建模块函数：define('modName',['依赖项'],function()){函数代码}

    - 前两个参数是可选的，第一个是字符串，第二个是数组，这个函数会加载“依赖项.js”文件

  - 加载模块函数：require(['modName'],function(obj){obj.methods()}

  - require()函数

    - AMD的require函数

      - ```
        require([mod],callback);
        ```

        

- CMD(SeaJS)通用模块定义

  - 创建模块函数：和AMD一样，用define函数定义，用require函数引用。

    - ```js
      define('modName',['依赖项'],function(require,exports,module)){函数代码}
      require:一个接受参数的方法
      exports:向外提供模块接口
      module:与当前模块相关联的属性和方法
      ```

  - CMD和AMD的异同有

    - 都是异步加载模块
    - CMD是懒加载，在require的时候才会加载模块，AMD是预加载，在定义模块的时候就加载好所有的依赖项。
    - 保留了CommonJS的风格

- UMD通用模块定义

  - CMD，AMD和CommonJS的缝合怪。





## 现在的标准——ES6

模块功能主要由两个命令构成：export指令导出接口，import引入模块。

ES6是自动采用严格模式的，特别有：this是指向undefined而不是全局。

#### export命令

​	导出接口的意思是：将对象从私有变成公共，可以被外部读取，解除封装性。

使用时的代码是：

```js
变量
export var name = 'Joe';
or
var name = 'Joe';
export{name};
函数和类
export function add(x,y){
    return x+y;
}
or
function add
(x,y){return x+y;}
export{add};

其中可以重命名对外的接口
export{
	name as JoeName;
	add as numberAdd;
}
```

export导出的接口是实时动态的，模块里面的值改变了，接口导出的值也会相应发生变化。（CommonJS是静态的）

尽管导出的值是实时的，但是导入的值一般情况下是静态的，不会跟着导出的值一起变化。

export不能位于{}内部，最好还是写在代码顶部。



export default function(){...}可以设置这个js文件的默认函数。

别的文件import就可以

```js
import anyName from './A';//anyName不需要用花括号包裹
```



#### import命令

A.js文件用export定义了对外接口之后（黑箱A出现了一个插座，开始供电），B.js文件就可以通过import命令加载这个模块了（黑箱B长出了一个插头接通了A的插座）。

```
import{JoeName as yourName,add} from './A.js';//之后B.js就可以使用yourName变量和add函数了
as关键字也可以将JoeName重命名为其他名字。

import * from './A.js';//导入A的所有变量，函数和对象
```

尽量不要修改引入的数据。



import命令会自动提升到顶部，所以不用担心调用的先后顺序。（好的代码风格是自己将import提升到顶部）



#### export和import的复合写法

不实用，没必要

#### 模块的继承

```
export * from './A';//继承A的全部内容。
```

其他。。。