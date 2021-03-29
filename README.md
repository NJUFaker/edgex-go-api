# Edgex-go-api

### 环境
go1.15

### 运行命令
```./build.sh && ./output/bootstrap.sh ```

#### Redis 配置验证
```redis-cli -h 127.0.0.1 -p 6379```

```config set requirepass edgex_go```

验证

http://localhost:6789/ping

http://localhost:6789/test_redis