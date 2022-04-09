## LINE Notifyで自動誕生日通知

Go + mysql
開発:docker
本番:heroku

# memo
・herokuSchedulerで日次実行
・FixieにてOutboundIpを固定
　Proxy URLをつけてRequstすると固定IPでリクエストできる
```shell
curl https:://hoge.app --proxy $FIXIE_URL
```
