export default {
  path: "/error",
  redirect: "/error/403",
  meta: {
    icon: "ri:flow-chart",
    // showLink: false,
    title: "工厂",
    rank: 0
  },
  children: [
    {
      path: "/factory/product",
      name: "product",
      component: () => import("@/views/factory/product.vue"),
      meta: {
        title: "产品"
      }
    },
    {
      path: "/factory/brand",
      name: "brand",
      component: () => import("@/views/factory/brand.vue"),
      meta: {
        title: "品牌"
      }
    }
  ]
} satisfies RouteConfigsTable;
