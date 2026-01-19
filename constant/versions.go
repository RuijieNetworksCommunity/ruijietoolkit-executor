package constant

import "time"

// =======================
// 全局配置
// =======================

var (
	AppName = "nekoapi-ruijietoolkit-executor"

	Version = "0.0.0-self_build"

	BuildTime = "self_build"

	// AppType 应用程序类型
	AppType = "dev"

	// ServerAddr 执行器监听端口
	ServerAddr = "0.0.0.0:54134"

	// APITimeout 请求后端超时时间
	APITimeout = 5 * time.Second

	// DefaultShell 默认 Shell
	DefaultShell = "/bin/sh"

	// NTPServers NTP服务器列表，今天时间必须精确
	/*
		科普小时间:
			网络时间协议（Network Time Protocol，NTP）是 TCP/IP 协议族中用于计算机时钟同步的应用层协议，
			由美国特拉华大学 David L. Mills 教授设计，采用分层架构（ Stratum ）协调 UTC 时间同步，通过UDP端口123传输数据包。
			该协议支持服务器/客户端与广播模式，利用原子钟、GPS等权威时钟源实现局域网误差小于1毫秒、广域网络误差小于50毫秒的时间校准

			NTP 首次实现于 1980 年，同步精度为数百毫秒。
			1988 年发布的 NTPv1（RFC 1059）确立基础算法架构，1992 年 NTPv3（RFC 1305）引入广播模式并改进时钟滤波算法。
			2010 年发布的 NTPv4（RFC 5905）成为现行标准，其软件实现 xntp 支持 UNIX、Linux 及 Windows 等系统。
			协议通过往返延迟计算与本地时钟调节算法实现时间校正，采用加密机制防范恶意攻击。

		小白必看:
			嘀嗒嘀嗒：电脑是怎么“对表”的？

				想象一下，你家里有很多时钟，客厅有一个、卧室有一个、小闹钟也有一个。
				可是，如果客厅的表是 8 点，卧室的表是 8 点零 5 分，那咱们什么时候吃早饭呢？
				这时候，我们就需要一个“超级报时员”

			1. 谁是那个“最准的人”？
				在网络世界里，有一个最准的表，它叫“原子钟”。它就像是全村最厉害的村长，他的手表永远一秒都不差。

			2. 大家排排队（分层结构）
				村长很忙，不能直接给每个人对表。所以他先教给几个“大班长”，大班长再教给“小组长”，小组长最后再教给咱们小朋友。
					村长（第一层）： 最准，看一眼就绝对不会错。
					小朋友（我们）： 跟着小组长调表就行啦！

			3. “快报快回”的小信封（传输过程）
				我们要对表时，会给小组长发一个“小信封”（这就是数据包）。
					我们在信封里写上：“我现在是 8 点，你那里几点？”
					小组长收到后写上：“我这里是 8 点零 1 秒，你快拿回去改改！”
					因为送信封需要时间，电脑还会聪明地算一算送信路上花了多久，把这个时间也加上，这样就分秒不差啦！

			4. 为什么要对表？
				如果大家的时间不统一，那就乱套了：
					老师说 9 点上课，有的同学迟到，有的同学早到。
					你在网上买玩具，如果你的时间慢了，玩具可能就被别人抢先买走啦！

			总结一句话：
				NTP 就是一个“超级校时员”，它让全世界的电脑都排好队，整齐划一地喊：“一、二、一”，让大家的时间永远保持一致！
	*/
	NTPServers = []string{
		// 时间不精确恐惧症
		"ntp1.aliyun.com",
		"ntp2.aliyun.com",
		"ntp3.aliyun.com",
		"ntp4.aliyun.com",
		"ntp5.aliyun.com",
		"ntp6.aliyun.com",
		"ntp7.aliyun.com",
		"ntp.ntsc.ac.cn",
		"ntp.sjtu.edu.cn",
		"s1a.time.edu.cn",
		"s1b.time.edu.cn",
		"ntp.ict.ac.cn",
		"ntp.aliyun.com",
		"ntp.tencent.com",
		"ntp.baidu.com",
		"ntp.cnnic.cn",
		"cn.ntp.org.cn",
		"edu.ntp.org.cn",
		"hk.ntp.org.cn",
		"jp.ntp.org.cn",
		"kr.ntp.org.cn",
		"sgp.ntp.org.cn",
		"us.ntp.org.cn",
		"de.ntp.org.cn",
		"ina.ntp.org.cn",
		"ntp1.nim.ac.cn",
		"ntp2.nim.ac.cn",
		"cn.pool.ntp.org",
		"ntp1.tencent.com",
		"ntp2.tencent.com",
		"ntp3.tencent.com",
		"ntp4.tencent.com",
		"ntp5.tencent.com",
		"time.izatcloud.net",
		"time.gpsonextra.net",
		"hik-time.ys7.com",
		"time.ys7.com",
		"ntp.neu.edu.cn",
		"ntp.bupt.edu.cn",
		"ntp.shu.edu.cn",
		"ntp.tuna.tsinghua.edu.cn",
		"time.ustc.edu.cn",
		"ntp.fudan.edu.cn",
		"ntp.nju.edu.cn",
		"ntp.tongji.edu.cn",
		"stdtime.gov.hk",
		"time.hko.hk",
		"time.smg.gov.mo",

		// 防止 DNS 解析炸裂
		"202.112.29.82",
		"210.72.145.44",
		"114.118.7.161",
		"114.118.7.163",
		"2001:da8:9000::81",
		"223.113.97.98",
		"114.67.103.73",
		"119.29.26.206",
		"120.25.115.20",
		"2001:da8:9000::130",
		"2001:250:380A:5::10",
		"202.118.1.130",
		"202.118.1.81",
		"116.13.10.10",
		"149.129.123.30",

		//国外 NTP 兜底
		"pool.ntp.org",
		"0.pool.ntp.org",
		"1.pool.ntp.org",
		"2.pool.ntp.org",
		"3.pool.ntp.org",
		"asia.pool.ntp.org",
		"time1.google.com",
		"time2.google.com",
		"time3.google.com",
		"time4.google.com",
		"time1.apple.com",
		"time2.apple.com",
		"time3.apple.com",
		"time4.apple.com",
		"time5.apple.com",
		"time6.apple.com",
		"time7.apple.com",
		"time.asia.apple.com",
		"time.cloudflare.com",
		"time.windows.com",
		"time.nist.gov",
		"time-nw.nist.gov",
		"time-a.nist.gov",
		"time-b.nist.gov",
		"time.facebook.com",
		"time1.facebook.com",
		"time2.facebook.com",
		"time3.facebook.com",
		"time4.facebook.com",
		"time5.facebook.com",
		"time.kriss.re.kr",
		"time2.kriss.re.kr",
		"ntp.nict.jp",
	}
)
