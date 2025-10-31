export interface productResponse {
    id:string;
    name:string;
    image:string;
    description:string;
    price:number;
    stock:number;
    saler:string;
    salerId:string;
}
export interface postProductRequest {
    name:string;
    images:string[];
    category:string;
    description:string;
    price:number;
    stock:number;
    saler:string;
    salerId:string;
}