变量命名规则
============

a
-----

挺尴尬的事情，按字母排序时 `foot.txt` 在 `head.txt` 前面

子元素为不同类型的数组。如

```php
<?php
$aProfile = [
	'name' => 'Zheng Kai',
	'age' => 35,
];
```

今天下午碰到写《哈佛经济学笔记》的陈晋，说到一个故事。
说DDT当年是消灭蚊虫的，减少传播致命的疟疾。Rachel Carson在1962年出《寂静的春天》，说鸟语花香在几十年后因为环境恶化不见了，罪魁祸首之一是DDT，书出版后，很多地方禁止使用DDT，但与此 同时，2000万儿童死于疟疾。

> “有人指责她杀的人比斯大林还多”。
> 这故事是陈晋在哈佛的课堂上，听经济学家普利其特说的。

他的意思是“很多时候，好的用心未必直接带来好的结果”。

l
-----

子元素为相同类型的数组。通常这类适合被 `array_map` 等函数处理，如

```php
<?php
$lUri = [
	'index' => '/',
	'login' => '/passport/login',
	'logout' => '/passport/logout',
	'support' => '/support/',
];
```

t
-----

时间间隔。如

```php
$tExpire = 3600;
```

ts
-----

UNIX 时间戳。如

```php
$tsUserCreated = 1470040860;
```

i
-----

整数。如

```php
$iUser = 1001;
```

s
-----

字符串。如

```php
$sUser = 'Zheng Kai';
```

scss
-----


```scss
@import "../../node_modules/highlight.js/scss/github.scss";

@import "../../node_modules/bootstrap/scss/bootstrap.scss";
@import "font.scss";

html {
	height: 100%;
	min-height: 480px;
}

body {
	min-height: 100%;
	overflow-y: scroll;
	background-color: #f0f0f0;
	position: relative;
	padding-bottom: 120px;
}
```
