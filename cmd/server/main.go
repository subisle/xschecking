package main

import (
	"auth-system/internal/config"
	"auth-system/internal/pkg"
	"auth-system/internal/router"
	"fmt"
	"log"
)

func main() {
	// 1. 加载配置
	if err := config.LoadConfig("config.yaml"); err != nil {
		log.Fatalf("加载配置文件失败: %v", err)
	}
	cfg := config.GetConfig()

	// 2. 初始化日志
	if err := pkg.InitLogger(cfg.Server.Mode); err != nil {
		log.Fatalf("初始化日志失败: %v", err)
	}
	defer pkg.GetLogger().Sync()

	pkg.GetLogger().Info("正在启动授权系统...")

	// 3. 初始化数据库
	if err := pkg.InitDatabase(&cfg.Database); err != nil {
		pkg.GetLogger().Fatal(fmt.Sprintf("初始化数据库失败: %v", err))
	}
	pkg.GetLogger().Info("数据库连接成功")

	// 4. 初始化Redis
	if err := pkg.InitRedis(&cfg.Redis); err != nil {
		pkg.GetLogger().Fatal(fmt.Sprintf("初始化Redis失败: %v", err))
	}
	pkg.GetLogger().Info("Redis连接成功")

	// 5. 设置路由
	r := router.SetupRouter(cfg.Server.Mode)

	// 6. 启动服务器
	addr := fmt.Sprintf(":%d", cfg.Server.Port)
	pkg.GetLogger().Info(fmt.Sprintf("服务器启动在 %s", addr))
	
	if err := r.Run(addr); err != nil {
		pkg.GetLogger().Fatal(fmt.Sprintf("服务器启动失败: %v", err))
	}
}
