export interface productResponse {
    name:string;
    image:string;
    description:string;
    price:string;
    stock:string;
    saler:string;
}
export interface postProductRequest {
    name:string;
    images:string[];
    category:string;
    description:string;
    price:number;
    stock:string;
    saler:string;
}