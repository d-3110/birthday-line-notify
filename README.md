# memo
herokuアドオン「Fixie」にてOutboundIpを固定
Proxy URLをつけてRequstすると固定IPでリクエストできる
```shell
curl https:://hoge.app --proxy $FIXIE_URL
```
