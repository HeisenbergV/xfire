

import { http } from "@/utils/http";

export const baseUrlApi = (url: string) =>
  process.env.NODE_ENV === "development"
    ? `/api/${url}`
    : `http://127.0.0.1:9999/${url}`;




export type Mianbao = {
      id: number;
      name: string;
      brand: string;
      remake: string;
      barcode: string;
      created_at: Date;
}
export type Product = {
  /** 是否请求成功 */
  code: number;
  msg: string;
  data: {
    list: Array<Mianbao>;
    total: number;
  };
};


export const getProductList = (data?: object) => {
  let ee = http.request<Product>("post", baseUrlApi("/factory/getGoodsList"), { data });
  return ee;
};