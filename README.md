# <img src="https://github.com/Aranxu/Nodepanels-probe/blob/main/favicon.ico" width = "40" height = "40" alt="" align=center/> Nodepanels-probe

<img src="https://img.shields.io/badge/Go-1.18-ff69b4"/> <img src="https://img.shields.io/badge/version-v1.1.0-orange"/> <img src="https://img.shields.io/badge/TG-@nodepanels-green?logo=telegram&style=plastic"/>

Nodepanels探针，采集服务器基础数据，服务器与本系统的连接通道

项目Github：[点击直达](https://github.com/Aranxu/Nodepanels)

项目网站：https://nodepanels.com

### 程序信息：

版本：`v1.1.0`

更新日期：`2022/06/01`

支持系统：`Linux 32/64位` | `Linux ARM 32/64位` | `Windows 32/64位`

### 采集的数据：

<details>
  <summary>系统信息</summary>
  <br>
  <ul>
    <li>主机名</li>
    <li>系统启动时间</li>
    <li>系统架构</li>
    <li>系统架构版本</li>
    <li>系统类型</li>
    <li>系统家族</li>
    <li>系统版本</li>
  </ul>
</details>

<details>
  <summary>CPU</summary>
  <br>
  <ul>
    <li>CPU数量</li>
    <li>CPU物理核心数</li>
    <li>CPU逻辑核心数</li>
    <li>CPU型号</li>
    <li>CPU供应商编号</li>
    <li>CPU赫兹</li>
    <li>CPU缓存</li>
    <li>CPU总使用率</li>
    <li>CPU各核心使用率</li>
  </ul>
</details>

<details>
  <summary>内存</summary>
  <br>
  <ul>
    <li>内存大小</li>
    <li>SWAP大小</li>
    <li>内存使用率</li>
    <li>SWAP使用率</li>
  </ul>
</details>

<details>
  <summary>磁盘</summary>
  <br>
  <ul>
    <li>分区信息（分区名，分区挂载点，分区格式，分区大小）</li>
    <li>分区使用率</li>
    <li>磁盘使用率</li>
    <li>磁盘IO</li>
  </ul>
</details>

<details>
  <summary>负载</summary>
  <br>
  <ul>
    <li>系统1分钟负载</li>
  </ul>
</details>

<details>
  <summary>进程</summary>
  <br>
  <ul>
    <li>进程列表</li>
    <li>进程数量</li>
    <li>监控的进程信息（Pid，进程名，进程命令行路径，进程CPU使用率，进程内存使用率，进程IO）</li>
  </ul>
</details>

<details>
  <summary>网络</summary>
  <br>
  <ul>
    <li>公网IP</li>
    <li>磁盘IO</li>
    <li>网络进出流量</li>
  </ul>
</details>

### 计划内容：

<table>
        <tr>
            <th>内容</th>
            <th>完成</th>
            <th>进度</th>
            <th>完成时间</th>
            <th>备注</th>
        </tr>
        <tr>
            <th>临时文件存放到temp目录</th>
            <th>✔</th>
            <th>100%</th>
            <th>2022.03.14</th>
            <th></th>
        </tr>
        <tr>
            <th>调用工具包后删除临时文件</th>
            <th>✔</th>
            <th>100%</th>
            <th>2022.03.14</th>
            <th></th>
        </tr>
        <tr>
            <th>Linux、Windows代码隔离</th>
            <th>✔</th>
            <th>100%</th>
            <th>2022.06.01</th>
            <th></th>
        </tr>
        <tr>
            <th>支持获取2秒粒度数据</th>
            <th>✔</th>
            <th>100%</th>
            <th>2022.06.01</th>
            <th></th>
        </tr>
        <tr>
            <th>每10分钟更新系统软硬件信息</th>
            <th>✔</th>
            <th>100%</th>
            <th>2022.06.01</th>
            <th></th>
        </tr>
    </table>
