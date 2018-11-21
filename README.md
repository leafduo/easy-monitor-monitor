# easy-monitor-monitor

```bash
go build -o easy-monitor-monitor .
sudo mv easy-monitor-monitor /usr/local/bin
sudo cp com.leafduo.EasyMonitorMonitor.plist /Library/LaunchDaemons/com.leafduo.EasyMonitorMonitor.plist
sudo launchctl load com.leafduo.EasyMonitorMonitor.plist
```
