
# 常用的 HTTP 状态码如下，分为不同的类别：

## 1xx (信息性状态码)
这些状态码表示请求已被接收，服务器正在继续处理。

100 Continue：表示服务器已收到请求头，客户端可以继续发送请求体。
101 Switching Protocols：服务器根据客户端的请求切换协议。


## 2xx (成功状态码)
表示请求已成功处理。
200 OK：请求成功，通常用于 GET、POST 请求。
201 Created：请求成功并导致新资源的创建，通常用于 POST 请求。
202 Accepted：请求已接受，但处理尚未完成。
204 No Content：请求成功，但没有返回内容（常用于 DELETE 请求）。

## 3xx (重定向状态码)
表示需要客户端进一步操作才能完成请求。

301 Moved Permanently：请求的资源已被永久移动到新的 URL。
302 Found：请求的资源临时被移动到不同的 URL。
303 See Other：请求应当访问另一个 URL，通常用于响应 POST 请求时重定向到 GET 请求。
304 Not Modified：客户端缓存的资源是最新的，无需重新下载。
307 Temporary Redirect：请求的资源临时被移动到其他地方。


## 4xx (客户端错误状态码)
表示请求包含语法错误或无法完成请求。

400 Bad Request：请求无效，通常是请求格式或参数错误。
401 Unauthorized：未授权，客户端必须提供身份验证。
403 Forbidden：服务器理解请求，但拒绝执行（如权限不足）。
404 Not Found：请求的资源未找到。
405 Method Not Allowed：请求方法不被允许。
406 Not Acceptable：请求的资源不可接受（如不支持的媒体类型）。
408 Request Timeout：请求超时。
409 Conflict：请求与服务器资源的当前状态冲突。
410 Gone：请求的资源不再可用。
413 Payload Too Large：请求体过大。
415 Unsupported Media Type：请求的媒体类型不受支持。
422 Unprocessable Entity：请求格式正确，但语义有误，通常用于验证失败。
429 Too Many Requests：请求次数过多，达到限制。

## 5xx (服务器错误状态码)
表示服务器未能完成有效的请求。

##500 Internal Server Error：服务器内部错误，无法完成请求。
501 Not Implemented：服务器不支持请求的功能。
502 Bad Gateway：服务器作为网关或代理时，接收到无效的响应。
503 Service Unavailable：服务器暂时无法处理请求，可能是服务器过载或维护中。
504 Gateway Timeout：作为网关或代理的服务器未及时从上游服务器获取响应。
505 HTTP Version Not Supported：服务器不支持请求中使用的 HTTP 版本。
其他常见状态码：
418 I'm a teapot：这是一个愚人节笑话，由 IETF RFC 2324 定义，表示服务器是一个茶壶，无法泡茶。