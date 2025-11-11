// composables/useAxios.ts
import axios from "axios";

// 在 Axios 请求中自动注入 JWT Token（从 cookie 或 localStorage）
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
  const m = document.cookie.match(/(?:^|; )authToken=([^;]+)/);
  if (m && m[1]) return decodeURIComponent(m[1]);
      }
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
      if (token) {
        if (isJwtExpired(token)) {
          // token 已过期：清理并重定向到登录页
          try {
            // @ts-ignore
            if (typeof useCookie !== "undefined") {
              // @ts-ignore
              useCookie("authToken").value = null;
            }
          } catch (e) {
            // ignore
          }
          if (typeof window !== "undefined") {
            try {
              localStorage.removeItem("authToken");
            } catch (e) {}
            // 直接跳转到登录页（不带 Nuxt 全局 API，保证在任何环境都能工作）
            window.location.href = "/login";
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
          // 清理 token 并跳转登录
          try {
            // @ts-ignore
            if (typeof useCookie !== "undefined") {
              // @ts-ignore
              useCookie("authToken").value = null;
            }
          } catch (e) {}
          if (typeof window !== "undefined") {
            try {
              localStorage.removeItem("authToken");
            } catch (e) {}
            window.location.href = "/login";
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
