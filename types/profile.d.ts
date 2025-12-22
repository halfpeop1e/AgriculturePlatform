/*
  profile.d.ts
  说明：用户资料数据结构（后端返回）
  字段：nickname, email, avatar, bio, tags, location, joinedAt, phone, address, lastActive
  使用场景：getUserProfile 返回此类型；userStore 的 setUserProfile 接受此类型
*/
export interface profileResponse{
    code:number;
    data:{
    Nickname:string;
    Email:string ;
    Avatar: string;
    Bio:  string;
    Tags: string[];
    Location:string ;
    JoinedAt:string ;
    Phone: string;
    Address:string ;
    LastActive: string;
    role:string 
    }
    message:string;
}

export interface ExpertProfile {
  name: string
  avatar: string
  field: string
  introduction: string
  skills: string[]
  education?: string
  experience?: string
  certification?: string[]
  availableTime?: string[]
  serviceScope?: string[]
  price?: number
  completed?: boolean
}

export interface ExpertProfileResponse {
  code: number
  data: ExpertProfile
  message: string
}

