import type { Expert, ExpertDetail } from '~/types/expert'
import { defineStore } from 'pinia'
import { getExpertList, getExpertDetail } from '~/composables/getExpert'

export const EXPERTS: Expert[] = [
  {
    id: 'e1',
    name: '张教授',
    avatar: '/experts/zhang.jpg',
    field: '种植业',
    introduction: '从事水稻种植研究30年，擅长水稻病虫害防治和高产技术，曾获省级农业科技进步奖。',
    skills: ['水稻种植', '病虫害防治', '高产技术'],
    consultCount: 128,
    rating: 98,
    responseTime: '2小时内',
    recentCases: [
      {
        date: '2024-05-10',
        question: '水稻出现黄叶病怎么办？',
        answer: '黄叶病多由真菌感染引起，建议使用多菌灵喷雾，并注意田间排水...'
      }
    ]
  },
  {
    id: 'e2',
    name: '李专家',
    avatar: '/experts/li.jpg',
    field: '养殖业',
    introduction: '专注生猪养殖技术研究，在规模化养殖和疫病防控方面有丰富经验。',
    skills: ['生猪养殖', '疫病防控', '饲料配比'],
    consultCount: 95,
    rating: 96,
    responseTime: '4小时内'
  }
  // 可添加更多专家数据
]

export const useExpertDataStore = defineStore('expertDataStore', {
  state: () => ({
    experts: [] as Expert[],
    currentExpert: null as ExpertDetail | null,
    total: 0,
    page: 1,
    pageSize: 10,
    hasMore: false
  }),
  actions: {
    async fetchExperts(page: number = 1, pageSize: number = 10, field?: string, searchKey?: string) {
      const data = await getExpertList(page, pageSize, field, searchKey)
      if (data) {
        this.experts = data.list
        this.total = data.total
        this.page = data.page
        this.pageSize = data.pageSize
        this.hasMore = data.hasMore
      }
    },
    async fetchExpertDetail(id: string) {
      const data = await getExpertDetail(id)
      if (data) {
        this.currentExpert = data
      }
    },
    clearCurrentExpert() {
      this.currentExpert = null
    }
  }
})