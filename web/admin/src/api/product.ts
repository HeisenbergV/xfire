

import { http } from "@/utils/http";
import { ref } from "vue";

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


export type BuildData = {
  name: string;
  remake: string;
  info: string;
  materialUnit: number;
  water: number;
  bakingtime: number;
  bakingtem: number;
  build_data:  { [key: string]: number };
}

export type GoodsCostResponse = {
  code: number;
  msg: string;
  data: {
    buildData: BuildData;
    yield: number;         
    cost:      number;    
  };
}

export type Product = {
  code: number;
  msg: string;
  data: {
    list: Array<Mianbao>;
    total: number;
  };
};

export type Response = {
  code: number;
  msg: string;
  data: object;
}

export const getProductList = (data?: object) => {
  let ee = http.request<Product>("post", baseUrlApi("/factory/getGoodsList"), { data });
  return ee;
};

export const getProductBuildInfo = (id?: number) => {
  const idinfo = { id:id };
  let ee = http.request<GoodsCostResponse>("post", baseUrlApi("/factory/getProductBuildInfo"), { data:idinfo });
  return ee;
};

export const delProduct = (id?: number) => {
  const idsToDelete = { ids: [id] };
  return http.request<Response>("delete", baseUrlApi("/factory/deleteGoodsByIds"), { data: idsToDelete });
};
