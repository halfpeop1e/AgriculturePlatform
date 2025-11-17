import axios from "axios";

export const useEmailAxios = () => {
  const AxiosInstance = axios.create({
    baseURL: "http://localhost:3001", // 对应你 Node.js 后端接口
    timeout: 5000,
  });

  return AxiosInstance;
}
