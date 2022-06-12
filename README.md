# redrock-test
# redrock-exam
> 红岩大一下学期期末考核
> 
> 部署：8.130.103.141:9090
> 
> 客户端用的postman
# api
## 一、输入的数据
> 操作棋子 x y tx ty 
> 
> x,y表示当前的坐标，tx，ty表示你想要下的地方的横纵坐标，均以空格分割

## 二、用户登录
| 返回参数 | 说明                                                  |
| -------- | ----------------------------------------------------- |
| id        |个人id，用于url中做参数                                                |
| code      | 状态码 |
| message     | 提示消息                   |
| data     | 自定义返回的数据                   |
### 1.注册api
> 无需token
> 
> 访问：localhost:9090/user/register
> 
- 请求参数

| 请求参数 | 类型                  | 说明                                    |
| -------- | ----------------------------------- | --------------------------------------- |
| password | string,必选 | 密码，要求长度8到16位即可 |
| username    | string,必选 | 用户名，要求1-20位之间长度的字符串即可    |
| rePassword | string,必选 | 重复输入密码，要求与password相同|

- 返回参数

| 返回参数 | 类型                  | 说明                                    |
|-------- | ----------------------------------- | --------------------------------------- |
| status| int ,必选 | 错误码 |
| message| string,必选 |提示信息   |

- 返回实例
```json
{
    "message": "注册成功！请登录！",
    "status": 0
}
```


### 2.密码登录

> 不需要token
> 
> 访问:localhost:9090/user/login

- 请求参数

| 请求参数 | 类型                                | 说明                                    |
| -------- | ----------------------------------- | --------------------------------------- |
| username    | string,必选 | 注册时使用的用户名           |
| password |string,必选 | 密码，要求长度8到16位即可 |

- 返回参数

| 返回参数 | 类型                  | 说明                                    |
|-------- | ----------------------------------- | --------------------------------------- |
| code| int ,必选 | 错误码 |
| refreshToken| string,必选 |签发的refreshToken   |
| token| string,必选 |  签发的token |
| uid| int,必选 |用户id   |
| message| string,必选 |提示信息   |

- 返回实例
```json
    {
  "code": 0,
  "data": {
    "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Imh5cW10ZXN0IiwiZXhwIjoxNjU1MDE5ODIxLCJpc3MiOiJBbHNhY2UiLCJuYmYiOjE2NTUwMTYyMjF9.b727OLT3jCJ3Tl4AOCKCNdz8T6u_yNsQBvlyjg6-2Xk",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Imh5cW10ZXN0IiwiZXhwIjoxNjU1MDI3MDIxLCJpc3MiOiJBbHNhY2UiLCJuYmYiOjE2NTUwMTYxNjF9.w_578fS_ChRLx4KwVhIUB04zfMwjzDdm-mS3sbJwQYw",
    "uid": 1
  },
  "message": "登录成功，欢迎进入！"
}
```







# 其他
时间真的很紧（六级加专业选修课报告），只能完成基础项了QAQ
（待后续补充）
