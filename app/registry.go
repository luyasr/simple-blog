package app

import (
	// 注册Handler到ioc容器
	_ "github.com/luyasr/simple-blog/app/api"
	// 注册Controller到ioc容器
	_ "github.com/luyasr/simple-blog/app/blog"
	_ "github.com/luyasr/simple-blog/app/token"
	_ "github.com/luyasr/simple-blog/app/user"
)
