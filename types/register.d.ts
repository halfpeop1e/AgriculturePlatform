/*
  register.d.ts
  说明：注册请求数据结构
  字段：
    - username: 用户名
    - email: 邮箱
    - password: 密码
    - compassword: 确认密码（注：拼写应为 confirmPassword）
    - isverified: 邮箱验证码是否已验证
*/
export interface RegisterRequest { 
    username: string; 
    email:string; 
    password: string; 
    compassword:string;
}
