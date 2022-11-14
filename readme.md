## 微博发送到ipfsTweet
* 依赖爬虫 [weibo-crawler](https://github.com/dataabc/weibo-crawler)
* 爬虫配置和使用方法请阅读爬虫readme文档
* 微博的数据库和节点的数据库不要用同一个
* 需要本地运行节点 https://github.com/ethtweet/ethtweet 并且配置MySQL存储模式

### 步骤

1. 修改 `weibo-crawler/config.json` 配置增加需要爬取的 user_id
3. 爬取数据 `python3 weibo-crawler/weibo.py`
2. 同步到推特 `go run main.go`