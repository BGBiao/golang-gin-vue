# vue-test

> A Vue.js project

## Build Setup

``` bash
# install dependencies
cnpm install

# serve with hot reload at localhost:8080
cnpm run dev

# build for production with minification
cnpm run build

# build for production and view the bundle analyzer report
cnpm run build --report

# run unit tests
cnpm run unit

# run e2e tests
cnpm run e2e

# run all tests
cnpm test
```

For a detailed explanation on how things work, check out the [guide](http://vuejs-templates.github.io/webpack/) and [docs for vue-loader](http://vuejs.github.io/vue-loader).

## Run the golang-gin-vue

```
# 安装异步请求包
$ cnpm install --save axios
# 需要注意的是，如果跨域请求需要接口处增加请求的Header

# 运行前端代码
$ cd vue-test
$ cnpm run dev

# 运行后端接口服务
$ go run main.go
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /v1/hello                 --> main.HelloPage (3 handlers)
[GIN-debug] GET    /v1/hello/:name           --> main.main.func1 (3 handlers)
[GIN-debug] GET    /v1/line                  --> main.main.func2 (3 handlers)


```
