##  README
中间件是一个`Handler` (`func(ctx iris.Context)`)

###   执行下一个`handler` 
* 调用函数
```go
ctx.Next
```

* 配置
```go
app.SetExecutionRules(iris.ExecutionRules{
	// Begin: ...
	// Main: ...
	Done: iris.ExecutionOptions{Force: true},
})
```

