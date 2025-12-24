import type { ApplyListRespond } from "~/types/loanApply";
import { useAxios } from "./useAxios";
import type { AgriculturalLoanProduct, AgriculturalLoanProductList } from "~/types/loanProduct";
import { all } from "axios";
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
    }>("/loan/apply/list", {
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

export async function useApplyLoan(productId: string, userId: string, amount: number,deadline: number) {
  try {
    const response = await useAxiosInstance.post<{
      data: ApplyListRespond;
    }>("/loan/apply", {
        user_id : userId,
        product_id:productId,
        amount:parseFloat(amount.toString()),
        deadline:parseInt(deadline.toString()),
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

export async function useApplyLoanResult(ApplyID: string,result : boolean) {
  try {
    const response = await useAxiosInstance.post<{
      data: ApplyListRespond;
    }>("/loan/allow", {
        apply_id : ApplyID,
        allow:result,
    });
    if (response.status === 200) {
      console.log("审批成功", response.data);
      return 
    } else {
      throw new Error("获取产品列表失败");
    }
  } catch (err) {
    console.error("获取产品列表失败", err);
  }
}
