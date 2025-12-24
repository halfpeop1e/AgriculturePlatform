import { useAxios } from "./useAxios";
import { ElMessage } from 'element-plus';
import type { Expert,ExpertDetail } from "~/types/expert";


const useAxiosInstance = useAxios();

// 获取专家列表（支持分页和筛选）
export async function getExpertList(page: number = 1, pageSize: number = 10, field?: string, searchKey?: string) {
  try {
    const response = await useAxiosInstance.get<{
      data:{list: Expert[];
      total: number;
      page: number;
      pageSize: number;
      hasMore: boolean;}
    }>('/expert/list', {
      params: {
        page,
        pageSize,
        field,
        searchKey
      }
    });
    if (response.status === 200) {
      console.log('获取专家列表成功', response.data);
      return response.data;
    }
    throw new Error('获取专家列表失败');
  } catch (err) {
    console.error('获取专家列表失败', err);
    ElMessage.error('获取专家列表失败');
  }
}

// 获取专家详情
export async function getExpertDetail(id: string) {
  try {
    const response = await useAxiosInstance.get<{data: ExpertDetail}>(`/expert/${id}`);
    if (response.status === 200) {
      console.log('获取专家详情成功', response.data);
      return response.data.data;
    }
    throw new Error('获取专家详情失败');
  } catch (err) {
    console.error('获取专家详情失败', err);
    ElMessage.error('获取专家详情失败');
  }
}

// 预约专家咨询
export async function bookExpertConsultation(expertId: string, data: {
  question: string;
  date: string;
  contactInfo: string;
}) {
  try {
    const response = await useAxiosInstance.post(`/expert/${expertId}/book`, data);
    if (response.status === 200) {
      ElMessage.success('预约咨询成功');
      return response.data;
    }
    throw new Error('预约咨询失败');
  } catch (err) {
    console.error('预约咨询失败', err);
    ElMessage.error('预约咨询失败');
  }
}
