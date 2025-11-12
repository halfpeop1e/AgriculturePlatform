/*
  profile.d.ts
  说明：用户资料数据结构（后端返回）
  字段：nickname, email, avatar, bio, tags, location, joinedAt, phone, address, lastActive
  使用场景：getUserProfile 返回此类型；userStore 的 setUserProfile 接受此类型
*/
export interface profileResponse{
    nickname:string;
    email:string ;
    avatar: string;
    bio:  string;
    tags: string[];
    location:string ;
    joinedAt:string ;
    phone: string;
    address:string ;
    lastActive: string; 
}

