export default {
  path: "/product",
  redirect: "/product",
  meta: {
    icon: "ri:flow-chart",
    // showLink: false,
    title: "产品",
    rank: 2
  },
  children: [
   {
      path: "/product/brand",
      name: "brand",
      component: () => import("@/views/product/brand/brand.vue"),
      meta: {
        title: "品牌"
      }
    },
    {
      path: "/product",
      name: "product",
      component: () => import("@/views/product/product.vue"),
      meta: {
        title: "产品"
      }
    },
    {
      path: "/product/material",
      name: "material",
      component: () => import("@/views/product/material/material.vue"),
      meta: {
        title: "原料"
      }
    }
  ]
} satisfies RouteConfigsTable;
