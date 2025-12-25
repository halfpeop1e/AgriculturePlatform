interface LoanOrder {
  id: number;
  year: number;
  month: number;
  amount: number;
  loanName: string;
  loanStatus: string; // 例如: "Unpaid", "Paid"
}

interface CheckMyLoanRespond {
  loanedSum: number; // 所有贷款金额
  loanSum: number; // 所有待还金额
  loanList: LoanOrder[];
}

interface GiveMoney {
  loanId : number
}
interface LoanPlan {
  name: string;
  tag: string;
  rate: string;
  monthlyPayment: string;
  totalInterest: string;
  description: string;
  id:string;
}

interface AiRespond{
  aiSuggestion: string;
  loanPlans: LoanPlan[];
}

interface AiRequest{
  purpose: string;
  amount: number;
  term: number;
}