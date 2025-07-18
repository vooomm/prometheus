# prometheus自定义监控指标-聚合

## 聚合演示示例代码
```bash
go mod init citytemp
go get github.com/prometheus/client_golang/prometheus
go run main.go
go build

http://localhost:2112/metrics
```

## windows 下编译成Linux平台

```bash
$env:GOOS = "linux"
$env:GOARCH = "amd64"
go build -o city-temp-exporter main.go
```

## systemd启动

```bash
vi /etc/systemd/system/city-temp-exporter.service

[Unit]
Description=City District Temperature Exporter
After=network.target

[Service]
ExecStart=/usr/local/bin/city-temp-exporter
Restart=on-failure
User=nobody

[Install]
WantedBy=multi-user.target
```

## 启动
```bash
systemctl daemon-reload
systemctl restart city-temp-exporter.service 
systemctl status  city-temp-exporter.service 

http://localhost:2112/metrics
```