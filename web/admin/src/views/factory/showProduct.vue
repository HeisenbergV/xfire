<script setup lang="ts">
import { ref, onMounted } from "vue";
import { Mianbao, getProductBuildInfo, GoodsCostResponse } from "@/api/product";

// 声明 props 类型
export interface Form2Props {
  formInline: {
    user: string;
    region: string;
    M: Mianbao;
    BuildInfo: GoodsCostResponse;
  };
}

// 声明 props 默认值
// 推荐阅读：https://cn.vuejs.org/guide/typescript/composition-api.html#typing-component-props
const props = withDefaults(defineProps<Form2Props>(), {
  formInline: () => ({ user: "", region: "", M: null, BuildInfo: null })
});

// vue 规定所有的 prop 都遵循着单向绑定原则，直接修改 prop 时，Vue 会抛出警告。此处的写法仅仅是为了消除警告。
// 因为对一个 reactive 对象执行 ref，返回 Ref 对象的 value 值仍为传入的 reactive 对象，
// 即 newFormInline === props.formInline 为 true，所以此处代码的实际效果，仍是直接修改 props.formInline。
// 但该写法仅适用于 props.formInline 是一个对象类型的情况，原始类型需抛出事件
// 推荐阅读：https://cn.vuejs.org/guide/components/props.html#one-way-data-flow
const newFormInline = ref(props.formInline);
// const getCardListData = async id => {
//   try {
//     let data = await getProductBuildInfo(id);
//     return data;
//   } finally {
//     setTimeout(() => {}, 500);
//   }
// };
console.log(newFormInline);
const data = newFormInline.value.BuildInfo.data;
</script>

<template>
  <div class="popup-wrapper">
    <div class="popup-container">
      <div class="section cost">
        <h3 class="subtitle">介绍</h3>
        <div class="cost-info">
          <p>一袋面产出 {{ data.yield }} 个面包</p>
          <p>材料总成本 {{ data.cost }} 元</p>
          <p>大约用水 {{ data.build_info.water / 500 }} 斤</p>
          <p>
            风炉温度 {{ data.build_info.bakingtem }}°， 烤
            {{ data.build_info.bakingtime }}分钟
          </p>
        </div>
      </div>
      <div class="section">
        <h3 class="subtitle">配方</h3>
        <ul class="ingredients-list">
          <li
            v-for="(ingredient, index) in data.build_info.build_data"
            :key="index"
            class="ingredient-item"
          >
            <span class="ingredient-name">{{ index }}:</span>
            <span class="ingredient-amount">{{ ingredient }}g</span>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<style scoped>
.ingredients-list {
  list-style-type: none;
  padding: 0;
}

.ingredient-item {
  display: inline-block;
  width: 100%;
  background-color: #d0e7ff;
  margin-bottom: 5px;
  border-radius: 5px;
  padding: 5px;
}

.ingredient-name {
  display: inline-block;
  font-weight: bold;
  color: #000;
  margin-right: 10px; /* 调整这个值可以控制名字和用量之间的距离 */
}

.ingredient-amount {
  display: inline-block;
  color: #555;
}

.popup-wrapper {
  max-height: 400px;
  overflow-y: auto;
}

.popup-container {
  background-color: #f0f8ff;
  border: 2px solid #ccc;
  border-radius: 10px;
  padding: 20px;
}

.section {
  margin-bottom: 30px;
  width: 100%;
}

.subtitle {
  margin-bottom: 10px;
  color: #333;
}

.cost {
  background-color: #add8e6;
  padding: 15px;
  border-radius: 10px;
  margin-bottom: 20px;
}

.cost-info p {
  margin: 5px 0;
}
.list {
  padding-left: 20px;
}
</style>
