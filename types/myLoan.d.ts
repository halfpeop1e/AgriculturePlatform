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