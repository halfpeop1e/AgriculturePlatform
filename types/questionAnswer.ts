export interface Question {
  id: string
  expertId: string
  expert: string
  title: string
  content: string
  author: string
  authorId: string
  date: string
  tags: string[]
  answerCount: number
  isAnswered: boolean
}

export interface Answer {
  id: string
  questionId: string
  content: string
  expert: string
  expertId: string
  date: string
  avatar?: string
  likes?: number
}

export interface QuestionDetail extends Question {
  answer: string[]
}
