# rpc-logger

 Simple logging middleware for rpc services.

 View the [docs](http://godoc.org/github.com/gohttp/rpc-logger).

```
2015-06-15 19:59:38 [INFO] >> POST /rpc {"method":"Coupons.GetById","params":[{"id":"trial"}], "id":1}
2015-06-15 19:59:38 [WARNING] << POST /rpc {"method":"Coupons.GetById","params":[{"id":"trial"}], "id":1} 404 (19B) in 142.103µs
```

# License

 MIT
