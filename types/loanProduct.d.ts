export interface AgriculturalLoanProduct {
  // 产品基础信息
  productId: string;                    // 产品唯一标识
  productName: string;                  // 产品名称
  productAvatar: string;                
  
  financialInstitution:{
     id: string;                    // 机构ID
     name: string;                  // 机构全称
     customerService: string;     // 客服电话
  }
  // 融资条款
  loanAmountRange: {                   // 贷款金额范围
    min: number;                       // 最低贷款额
    max: number;                       // 最高贷款额
  };
  
  interestRate: {                      // 利率信息
    type: number;        // 0固定/1浮动利率
    finalRate: number;                 // 最终执行利率
    discountDescription?: string;      // 利率优惠说明
  };
  
  loanTerm: {                          // 贷款期限
    minMonths: number;                 // 最短期限(月)
    maxMonths: number;                 // 最长期限(月)
  };
  
  // 适用条件
  eligibility: {                       // 申请条件
    minOperatingYears: number;         // 最低经营年限
    creditRequirement: string;         // 征信要求
    collateralRequirements: string;   // 担保要求
  };
  
  // 支持的业务范围
  supportedPurposes: {                 // 贷款用途
    production: boolean;               // 农业生产
    equipment: boolean;                // 设备购置
    land: boolean;                     // 土地流转/租赁
    operating: boolean;                // 经营周转
    infrastructure: boolean;           // 设施建设
  };
  
  // 状态信息
  estimatedTime: string;             // 预计审批时间
  updateTime: Date;                    // 最后更新时间
  effectiveDate: Date;                 // 生效日期
  expiryDate?: Date;                   // 失效日期
}

export interface AgriculturalLoanProductList {
  hasmore: boolean;
  loanProductList: AgriculturalLoanProduct[];
  page: number;
  pageSize: number;
  total: number;
}