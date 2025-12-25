import { useAxios } from "./useAxios";

const useAxiosInstance = useAxios();

export async function useAiSuggestion(requestData: AiRequest) {
  try {
    const response = await useAxiosInstance.post<{ data: AiRespond }>(
      "/loan/ai",
      requestData
    );

    if (response.status === 200) {
      console.log("获取数据成功", response.data);

      return response.data.data;
    } else {
      throw new Error("获取数据失败");
    }
  } catch (err) {
    console.error("获取数据失败", err);
  }
}
