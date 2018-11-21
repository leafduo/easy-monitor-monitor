# easy-monitor-monitor

使得 [EasyConnect](http://www.sangfor.com.cn/product/net-safe-mobile-security-easyconnect.html) 在不连接时不再驻留内存，相反地，我驻留内存 XD

```bash
go build -o easy-monitor-monitor .
sudo mv easy-monitor-monitor /usr/local/bin
sudo cp com.leafduo.EasyMonitorMonitor.plist /Library/LaunchDaemons/com.leafduo.EasyMonitorMonitor.plist
sudo launchctl load com.leafduo.EasyMonitorMonitor.plist
```
