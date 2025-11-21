import type { ApplyListRespond } from "~/types/loanApply";
import { useAxios } from "./useAxios";
import type { AgriculturalLoanProduct, AgriculturalLoanProductList } from "~/types/loanProduct";
const useAxiosInstance = useAxios();
export async function getLoanProductList(page: number, pageSize: number) {
  try {
    const response = await useAxiosInstance.get<{
      data: AgriculturalLoanProductList;
    }>("/loan/list", {
      params: {
        page,
        pageSize,
      },
    });
    if (response.status === 200) {
      console.log("获取产品列表成功", response.data);
      return response.data.data;
    } else {
      throw new Error("获取产品列表失败");
    }
  } catch (err) {
    console.error("获取产品列表失败", err);
  }
}

export async function getLoanApplyList(userId: string) {
  try {
    const response = await useAxiosInstance.get<{
      data: ApplyListRespond[];
    }>("/loan/list", {
      params: {
        user_id : userId,
      },
    });
    if (response.status === 200) {
      console.log("获取产品列表成功", response.data);
      return response.data.data;
    } else {
      throw new Error("获取产品列表失败");
    }
  } catch (err) {
    console.error("获取产品列表失败", err);
  }
}
