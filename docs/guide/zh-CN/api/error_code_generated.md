# 错误码

！！IAM 系统错误码列表，由 `codegen -type=int -doc` 命令生成，不要对此文件做任何更改。

## 功能说明

如果返回结果中存在 `code` 字段，则表示调用 API 接口失败。例如：

```json
{
  "code": 100101,
  "message": "Database error"
}
```

上述返回中 `code` 表示错误码，`message` 表示该错误的具体信息。每个错误同时也对应一个 HTTP 状态码，比如上述错误码对应了 HTTP 状态码 500(Internal Server Error)。

## 错误码列表

IAM 系统支持的错误码列表如下：

| Identifier | Code | HTTP Code | Description |
| ---------- | ---- | --------- | ----------- |
| ErrDatabase | 100101 | 500 | Database error |
| ErrUnauthorized | 100102 | 401 | Unauthorized access |
| ErrNotFound | 100103 | 404 | Resource not found |
| ErrBadRequest | 100104 | 400 | Bad request |
| ErrForbidden | 100105 | 403 | Access forbidden |
| ErrInternalServerError | 100106 | 500 | Internal server error |

