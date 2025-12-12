export interface Question {
  id: string;
  title: string;
  content: string;
  author: string;
  authorId: string;
  date: string;
  tags: string[];
  answerCount: number;
  isAnswered: boolean;
}

export interface Answer {
  id: string;
  questionId: string;
  content: string;
  expert: string;
  expertId: string;
  date: string;
  isAccepted: boolean;
}

export interface QuestionDetail extends Question {
  answers: Answer[];
}