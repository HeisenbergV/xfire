export default {
  path: "/factory",
  redirect: "/factory/product",
  meta: {
    icon: "ri:flow-chart",
    // showLink: false,
    title: "工厂",
    rank: 2
  },
  children: [
    {
      path: "/factory/product",
      name: "product",
      component: () => import("@/views/factory/product/product.vue"),
      meta: {
        title: "产品"
      }
    },
    {
      path: "/factory/material",
      name: "material",
      component: () => import("@/views/factory/material/material.vue"),
      meta: {
        title: "原料"
      }
    }
  ]
} satisfies RouteConfigsTable;
