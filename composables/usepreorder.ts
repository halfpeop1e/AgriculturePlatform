import { useAxios } from "./useAxios";
import { ElMessage } from 'element-plus';
export interface PreOrder {
  expertId: string;
  expertName: string;
  authorName: string;
  authorId: string;
  time: string;
  preorderId: string;
}
export interface PreOrderRequest {
  expertId: string;
  authorId: string;
  time: string;
}
export async function PreOrder(expertId: string, scheduleTime?: string | Date) {
  const useAxiosInstance = useAxios();
  try {
    let normalizedTime: Date;
    if (scheduleTime instanceof Date) {
      normalizedTime = scheduleTime;
    } else if (typeof scheduleTime === 'string' && scheduleTime) {
      const parsed = new Date(scheduleTime);
      normalizedTime = isNaN(parsed.getTime()) ? new Date() : parsed;
    } else {
      normalizedTime = new Date();
    }
    const userStore = useUserStore();
    const payload: PreOrderRequest = {
      expertId,
      authorId: userStore.userId,
      time: normalizedTime.toISOString()
    };
    const response = await useAxiosInstance.post('/book/preorder', payload);
    if (response.status === 200) {
      console.log('预约成功', response.data);
      ElMessage.success('预约成功');
      return response.data;
    } else {
      throw new Error('预约失败');
    }
  } catch (err) {
    console.error('预约失败', err);
    ElMessage.error('预约失败');
  }
}
export async function GetPreOrder() {
  const useAxiosInstance = useAxios();
  try {
    const userStore= useUserStore()
    const response = await useAxiosInstance.get<{data:PreOrder[]}>(`/book/preorder/list`,
      {
        params: { 
          userId: userStore.userId,
          role: userStore.role
        }
      }
    );
    if (response.status === 200) {
      console.log('获取预下单信息成功', response.data);
      return response.data.data;
    } else {
      throw new Error('获取预下单信息失败');
    }
  } catch (err) {
    console.error('获取预下单信息失败', err);
    ElMessage.error('获取预下单信息失败');
  }
}
