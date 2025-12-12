import type { Question, QuestionDetail } from '@/types/questionAnswer'
import { defineStore } from 'pinia'
import { getQuestionList, getQuestionDetail } from '~/composables/useQuestionAnswer'

export const useQuestionDataStore = defineStore('questionDataStore', {
  state: () => ({
    questions: [] as Question[],
    currentQuestion: null as QuestionDetail | null
  }),
  actions: {
    async fetchQuestions() {
      const data = await getQuestionList()
      if (data) {
        this.questions = data
      }
    },
    async fetchQuestionDetail(id: string) {
      const data = await getQuestionDetail(id)
      if (data) {
        this.currentQuestion = data
      }
    },
    clearCurrentQuestion() {
      this.currentQuestion = null
    }
  }
})