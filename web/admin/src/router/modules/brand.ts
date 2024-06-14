// 最简代码，也就是这些字段必须有
export default {
  path: "/brand/brand",
  meta: {
    title: "品牌",
    icon: "ri:xing-fill",
  },
  children: [
    {
      path: "/brand/brand",
      name: "brand",
      component: () => import("@/views/brand/brand.vue"),
      meta: {
    icon: "ri:xing-fill",

        title: "品牌"
      }
    }
  ]
};
