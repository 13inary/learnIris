##  API
###   通常写法
```go
app := iris.New()

app.Handle("大写的方法", "路径", 可变数量的iris.Handler按执行顺序写入 )
												|
											函数名/匿名函数
```

###   简洁写法
```go
app := iris.New()

app.方法("路径", 可变数量的iris.Handler按执行顺序写入 )
								|
							函数名/匿名函数
```


###   iris.Handler
```go
func(ctx iris.Context){

}
```

