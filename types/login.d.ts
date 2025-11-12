/*
  login.d.ts
  说明：登录相关的请求和响应类型
  接口：
    - LoginRequest: 登录请求（username 可以是用户名或邮箱，password 为密码）
    - LoginResponse: 登录成功的响应（包含 userId 和 tokens（JWT 或 session））
*/
export interface LoginRequest {
  username: string;
  password: string;
}
export interface LoginResponse {
  userId: string;
  tokens: string;
}
