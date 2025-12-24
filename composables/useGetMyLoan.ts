import { useAxios } from "./useAxios"


const useAxiosInstance = useAxios()

export async function addLoanProduct(){
    try {
    const response = await useAxiosInstance.get<{data:CheckMyLoanRespond}>("/loan/detail");
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

export async function giveMoney(loanId:number){

    try {

    const response = await useAxiosInstance.post<{data:GiveMoney}>("/loan/give",{loanId});

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