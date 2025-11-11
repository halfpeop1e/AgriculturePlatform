import Cookies from "js-cookie";
export const setCookie = (label: string, value: any, cookieExpires: any) => {        
    Cookies.set(label, value, { expires: cookieExpires })
}
export const getTokenFromCookie = (label: string) => {
    return Cookies.get(label)
}
