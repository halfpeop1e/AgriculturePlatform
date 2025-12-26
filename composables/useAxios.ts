/*
  useAxios.ts
  说明：创建并返回一个 Axios 实例，封装常用的请求拦截与响应拦截逻辑：
    - 从 cookie / localStorage / Nuxt 的 useCookie 中读取 token 并注入 Authorization header
    - 简单检查 JWT 的 exp 字段，若过期则清理 token 并在客户端跳转到 /login
    - 全局捕获 401 响应并清理 token 后跳转登录
  注意：
    - 该实现为轻量客户端处理，不会校验签名或自动刷新 token；若后端支持 refresh token，可扩展此处
    - 在 SSR 环境使用时请注意跳转逻辑（window 不存在），已做简化兼容处理
*/
// composables/useAxios.ts
import axios from "axios";
import { useUserStore } from "~/utils/userStore";

export const useAxios = () => {
  const AxiosInstance = axios.create({
    baseURL: "http://47.94.106.89:8080/", // 对应你 Node.js 后端接口
    timeout: 5000,
  });

  // Helper: 从可用位置读取 token（优先 useCookie -> localStorage -> document.cookie）
  const getToken = (): string | null => {
    try {
      // 在 Nuxt composable 环境中优先使用 useCookie（若可用）
      // @ts-ignore
      if (typeof useCookie !== "undefined") {
        // @ts-ignore
        const c = useCookie("AuthToken");
        if (c && c.value) return c.value as string;
      }
    } catch (e) {
      // ignore
    }

    try {
      if (typeof window !== "undefined") {
        const local = localStorage.getItem("AuthToken");
        if (local) return local;
        // 兼容性读取：document.cookie 中可能用小写或大小写不一致，使用 case-insensitive 查找
        const m = document.cookie.match(/(?:^|; )(?:(?:AuthToken|authToken))=([^;]+)/i);
        if (m && m[1]) return decodeURIComponent(m[1]);
      }
    } catch (e) {
      // ignore
    }

    try {
      const store = useUserStore()
      if (store?.tokens) return store.tokens
    } catch (e) {
      // ignore
    }

    return null;
  };

  // Helper: base64url 解码
  const base64UrlDecode = (str: string) => {
    try {
      const s = str.replace(/-/g, "+").replace(/_/g, "/");
      if (typeof window === "undefined") {
        // server / node 环境
        // @ts-ignore
        return Buffer.from(s, "base64").toString("utf8");
      } else {
        // 浏览器
        return decodeURIComponent(
          Array.prototype.map
            .call(window.atob(s), (c: string) => "%" + ("00" + c.charCodeAt(0).toString(16)).slice(-2))
            .join("")
        );
      }
    } catch (e) {
      return null;
    }
  };

  // 检查 JWT 是否过期（如果无法解析则当作已过期）
  const isJwtExpired = (token: string) => {
    try {
  const parts = token.split(".");
  if (parts.length < 2) return true;
  if (!parts[1]) return true;
  const payload = base64UrlDecode(parts[1]);
  if (!payload) return true;
      const obj = JSON.parse(payload);
      if (obj && obj.exp) {
        const now = Date.now() / 1000;
        return now > obj.exp;
      }
      // 无 exp 字段则认为未过期（可根据需求调整）
      return false;
    } catch (e) {
      return true;
    }
  };

  // 请求拦截器：注入 Authorization header，并检查过期
  AxiosInstance.interceptors.request.use(
    (config) => {
      const token = getToken();
      if (token == null){console.log("no token found");}
      if (token) {
        if (isJwtExpired(token)) {
          // token 已过期：清理并重定向到登录页
          try {
            // @ts-ignore
            if (typeof useCookie !== "undefined") {
              // @ts-ignore
              useCookie("AuthToken").value = null;
            }
          } catch (e) {
            // ignore
          }
          if (typeof window !== "undefined") {
            try {
              localStorage.removeItem("AuthToken");
            } catch (e) {}
            // 在客户端尽量使用 Nuxt 的导航以避免硬刷新；若不可用，才回退到 location
            try {
              // @ts-ignore
              if (typeof navigateTo !== 'undefined') {
                // @ts-ignore
                navigateTo('/login')
              } else {
                window.location.href = '/login'
              }
            } catch (e) {
              window.location.href = '/login'
            }
          }
          return Promise.reject(new Error("JWT expired"));
        }

        config.headers = config.headers || {};
        // 确保使用 Bearer 模式
        // @ts-ignore
        config.headers.Authorization = `Bearer ${token}`;
      }
      return config;
    },
    (error) => Promise.reject(error)
  );

  // 响应拦截器：统一处理 401（未授权）等错误
  AxiosInstance.interceptors.response.use(
    (response) => response,
    (error) => {
      try {
        const status = error?.response?.status;
        if (status === 401) {
          // 清理 token 并跳转登录（尽量使用 Nuxt 导航）
          try {
            // @ts-ignore
            if (typeof useCookie !== "undefined") {
              // @ts-ignore
              useCookie("AuthToken").value = null;
            }
          } catch (e) {}
          if (typeof window !== "undefined") {
            try {
              localStorage.removeItem("AuthToken");
            } catch (e) {}
            try {
              // @ts-ignore
              if (typeof navigateTo !== 'undefined') {
                // @ts-ignore
                navigateTo('/login')
              } else {
                window.location.href = '/login'
              }
            } catch (e) {
              window.location.href = '/login'
            }
          }
        }
      } catch (e) {
        // ignore
      }
      console.error("请求错误：", error);
      return Promise.reject(error);
    }
  );

  return AxiosInstance;
};
