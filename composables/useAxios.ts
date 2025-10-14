// composables/useAxios.ts
import axios from "axios";

export const useAxios = () => {
  const AxiosInstance = axios.create({
    baseURL: "http://localhost:3001/", // 对应你 Node.js 后端接口
    timeout: 5000,
  });

  // 拦截器示例（可选）
    AxiosInstance.interceptors.response.use(
    (response) => response, // 不再只返回 data
    (error) => {
        console.error("请求错误：", error);
        return Promise.reject(error);
    }
);

  return AxiosInstance;
};