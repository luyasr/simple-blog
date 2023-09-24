package pkg

import (
	// 注册ApiHandler到ioc容器
	_ "github.com/luyasr/simple-blog/pkg/api"
	// 注册Controller到ioc容器
	_ "github.com/luyasr/simple-blog/pkg/token"
	_ "github.com/luyasr/simple-blog/pkg/user"
)
