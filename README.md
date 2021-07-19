# NewBell
Go练习项目
# Setup
跟目录新建 'setting.yaml' 配置文件,粘贴下面配置
```yaml
app:
  name: 'ginapp'
  env: 'dev'
  port: 3000
  version: '1.0.0'

mysql:
  host: '127.0.0.1'
  port: 3306
  user: 'root'
  passwd: 'root'
  dbname: 'test'
  max_open_conns: 20
  max_idle_conns: 10

redis:
  host: '127.0.0.1'
  port: 3306
  passwd: 'redis123'
  poolsize: 100

log:
  level: 'debug'
  filename: './logs/ginapp.log'
  maxsize: 1
  maxbackups: 5
  maxage: 30
```
将MySQL和Redis配置改成自己的

```bash
$ go mod tidy
$ go run main.go
```