export interface LoginRequest {
  username: string;
  password: string;
}
export interface LoginResponse {
  userId: string;
  tokens: string;
}
