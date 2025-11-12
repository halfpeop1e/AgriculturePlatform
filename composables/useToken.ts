/*
    useToken.ts
    说明：简易的 cookie 操作封装，基于 js-cookie
    导出：
        - setCookie(label, value, expires)：设置 cookie
        - getTokenFromCookie(label)：读取指定名称的 cookie
    注意：该模块为轻量包装，未处理 JSON 序列化/反序列化或 SameSite 等高级选项。
*/
import Cookies from "js-cookie";
export const setCookie = (label: string, value: any, cookieExpires: any) => {        
        Cookies.set(label, value, { expires: cookieExpires })
}
export const getTokenFromCookie = (label: string) => {
        return Cookies.get(label)
}
