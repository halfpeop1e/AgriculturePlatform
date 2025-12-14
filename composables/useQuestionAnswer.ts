import { useAxios } from "./useAxios"
import { ElMessage } from 'element-plus'  // 补充导入ElMessage
import type { Question, QuestionDetail, Answer } from "~/types/questionAnswer";

const useAxiosInstance = useAxios()

// 获取问答列表
export async function getQuestionList() {
  try {
    const response = await useAxiosInstance.get<Question[]>('/question/list')
    if (response.status === 200) {
      console.log('获取问答列表成功', response.data)
      return response.data
    }
    throw new Error('获取问答列表失败')
  } catch (err) {
    console.error('获取问答列表失败', err)
    ElMessage.error('获取问答列表失败')
  }
}

// 获取问答详情 - 修正语法错误和类型使用错误
export async function getQuestionDetail(id: string) {
  try {
    // 修正1：修复模板字符串语法错误（缺少反引号闭合）
    // 修正2：正确使用泛型参数（QuestionDetail是类型，应放在<>中）
    const response = await useAxiosInstance.get<QuestionDetail>(`/question/${id}`)
    if (response.status === 200) {
      console.log('获取问答详情成功', response.data)
      return response.data
    }
    throw new Error('获取问答详情失败')
  } catch (err) {
    console.error('获取问答详情失败', err)
    ElMessage.error('获取问答详情失败')
  }
}

// 提交问题
export async function submitQuestion(data: Omit<Question, 'id' | 'date' | 'author' | 'answerCount' | 'isAnswered'>) {
  try {
    const response = await useAxiosInstance.post('/question/create', data)
    if (response.status === 200) {
      ElMessage.success('问题提交成功')
      return response.data
    }
    throw new Error('问题提交失败')
  } catch (err) {
    console.error('问题提交失败', err)
    ElMessage.error('问题提交失败')
  }
}

// 回答问题
export async function submitAnswer(questionId: string, content: string) {
  try {
    const response = await useAxiosInstance.post('/answer/create', {
      questionId,
      content
    })
    if (response.status === 200) {
      ElMessage.success('回答提交成功')
      return response.data
    }
    throw new Error('回答提交失败')
  } catch (err) {
    console.error('回答提交失败', err)
    ElMessage.error('回答提交失败')
  }
}