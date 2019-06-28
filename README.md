# ratelimiter
a simple go ratelimiter

使用 go 实现的限流算法，参考令牌桶算法实现

令牌桶算法
令牌桶算法的原理是系统会以一个恒定的速度往桶里放入令牌，而如果请求需要被处理，则需要先从桶里获取一个令牌，当桶里没有令牌可取时，则拒绝服务。

token-bucket

令牌桶算法示意图

https://github.com/yangwenmai/ratelimit/blob/master/doc/token-bucket.jpg

令牌桶算法通过发放令牌，根据令牌的rate频率做请求频率限制，容量限制等。
