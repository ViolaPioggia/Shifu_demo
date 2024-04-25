# Shifu_demo
在Kubernetes中运行Shifu并编写一个应用的实践过程



## 部署并运行 Shifu

```sh
sudo docker ps
```

![image-20240425113905126](./images/image-20240425113905126.png)

Shifu 安装已完成

```
sudo kubectl get pods -A
```

![image-20240425114934734](./images/image-20240425114934734.png)





## 运行一个酶标仪的数字孪生

### 准备

![image-20240425115503308](./images/image-20240425115503308.png)

### 运行

```sh
sudo kubectl apply -f run_dir/shifu/demo_device/edgedevice-plate-reader
```

![image-20240425115552852](./images/image-20240425115552852.png)

```sh
sudo kubectl get pods -A | grep plate
```

成功启动

![image-20240425115613590](./images/image-20240425115613590.png)



### 交互

进入 nginx

![image-20240425115757817](./images/image-20240425115757817.png)

```
curl "deviceshifu-plate-reader.deviceshifu.svc.cluster.local/get_measurement"
```

![image-20240425115907998](./images/image-20240425115907998.png)

## 编写一个 Go 程序

编写代码

```go
package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	targetUrl := "http://deviceshifu-plate-reader.deviceshifu.svc.cluster.local/get_measurement"
	req, _ := http.NewRequest("GET", targetUrl, nil)
	for {
		res, _ := http.DefaultClient.Do(req)
		body, _ := ioutil.ReadAll(res.Body)
		average := calculateAverage(body)
		log.Println("Average:", average)
		time.Sleep(2 * time.Second)
	}
}

func calculateAverage(data []byte) float64 {
	sum := 0
	count := 0
	for _, value := range data {
		sum += int(value)
		count++
	}
	if count > 0 {
		return float64(sum) / float64(count)
	}
	return 0
}

```

docker 打包

![image-20240425161303225](./images/image-20240425161303225.png)

将应用镜像加载到 `kind` 中

![image-20240425161347384](./images/image-20240425161347384.png)

运行容器 Pod

![image-20240425161445390](./images/image-20240425161445390.png)

成功运行

![image-20240425162057630](./images/image-20240425162057630.png)

### 检查应用输出

每两秒打印一次切片的值，符合程序预期

![image-20240425161838534](./images/image-20240425161838534.png)
