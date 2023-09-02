
```bash
    goctl api go -api temp.api -dir /Users/hq/GeeVong/SimpleGo/testgoctl
```

```bash
goctl api ts --api temp.api --caller "webapi" --dir ./ts

```


curl --location --request POST -w "\nhttp code: %{http_code\n}" "'http://loccal:8888/ping' \ --header `Content-Type:application/json` \ --data-raw '{
    "time":0
}'