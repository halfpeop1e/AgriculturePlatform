import type { KnowledgeArticle } from '@/types/knowledgeArticle'
import { defineStore } from 'pinia'
export const ARTICLES: KnowledgeArticle[] = [
  {
    id: '1',
    title: '有机苹果栽培要点：从土壤到采摘',
    excerpt: '介绍有机苹果种植的关键环节：土壤改良、病虫害绿色防治与采摘时机。',
    content: '<p>有机苹果栽培强调生态平衡与土壤健康。首先进行土壤检测并改良，增加有机质；其次采用轮作和覆盖物保持水分与抑制杂草；病虫害以生物防治为主，例如释放天敌、使用诱捕和生物农药；采摘时机以果实糖度和着色为准。</p><p>实践要点：</p><ul><li>土壤pH维持在6.0-6.8之间。</li><li>早期修剪利于通风透光，降低病害。</li><li>采后冷链与分级包装能显著减少损耗。</li></ul>',
    image: '/images/know_apple.jpg',
    tags: ['种植', '有机'],
    author: '农业专家 张三',
    date: '2025-11-01'
  },
  {
    id: '2',
    title: '土鸡蛋的饲养管理与品质提升',
    excerpt: '土鸡蛋品质与饲养方式密切相关，本文总结日常管理要点与营养配比建议。',
    content: '<p>提高土鸡蛋品质首先要改善饲料与散养环境。合理蛋白质与钙磷配比能提高产蛋率与蛋壳强度。散养让鸡群获得更多运动与天然饲料，能提升风味。</p>',
    image: '/images/know_egg.jpg',
    tags: ['养殖', '营养'],
    author: '养殖顾问 李四',
    date: '2025-10-20'
  },
  {
    id: '3',
    title: '玉米高产栽培的三大技术',
    excerpt: '从品种选择、密植到水肥管理，提升玉米产量的实用技术汇总。',
    content: '<p>选择适应本地的优良品种，结合合理定植密度与精准施肥，可以在同等投入下显著提升产量。重视灌溉与病虫害监测。</p>',
    image: '/images/know_corn.jpg',
    tags: ['种植', '高产'],
    author: '田间技术员 王五',
    date: '2025-09-12'
  }
]//示例数据
export const useKnowledgeDataStore = defineStore('knowledgeDataStore', {
  state: () => ({
    articles: [] as KnowledgeArticle[]
  }),
  actions: {
    getArticleById(id: string): KnowledgeArticle | undefined {
      return this.articles.find(article => article.id === id)
    }
  }
}) 