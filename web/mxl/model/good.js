import {
  cdnBase
} from '../config/index';
const imgPrefix = cdnBase;

const defaultDesc = [`${imgPrefix}/goods/details-1.png`];

const allGoods = [{
  saasId: '1',
  storeId: '1000',
  spuId: '0',
  title: '麦小六果子面包',
  primaryImage: 'https://i.ibb.co/HFjR2ZR/image.jpg',
  images: [
    'https://i.ibb.co/HFjR2ZR/image.jpg',
    'https://i.ibb.co/HFjR2ZR/image.jpg',
  ],
  spuTagList: [{
    id: '13001',
    title: '爆款',
    image: null
  }],
  desc: [
    'https://cdn-we-retail.ym.tencent.com/tsr/goods/nz-09c.png',
    'https://cdn-we-retail.ym.tencent.com/tsr/goods/nz-09d.png',
  ],
  etitle: '',
}, {
  saasId: '2',
  storeId: '1000',
  spuId: '0',
  title: '俄罗斯大列巴',
  primaryImage: 'https://i.ibb.co/DbjtncN/image.jpg',
  images: [
    'https://i.ibb.co/DbjtncN/image.jpg',
    'https://i.ibb.co/DbjtncN/image.jpg',
  ],
  spuTagList: [{
    id: '13001',
    title: '爆款',
    image: null
  }],
  desc: [
    'https://cdn-we-retail.ym.tencent.com/tsr/goods/nz-09c.png',
    'https://cdn-we-retail.ym.tencent.com/tsr/goods/nz-09d.png',
  ],
  etitle: '',
}, {
  saasId: '3',
  storeId: '1000',
  spuId: '0',
  title: '无蔗糖吐司面包',
  primaryImage: 'https://i.ibb.co/51Ncbn3/image.jpg',
  images: [
    'https://i.ibb.co/51Ncbn3/image.jpg',
    'https://i.ibb.co/51Ncbn3/image.jpg',
  ],
  spuTagList: [{
    id: '13001',
    title: '爆款',
    image: null
  }],
  desc: [
    'https://cdn-we-retail.ym.tencent.com/tsr/goods/nz-09c.png',
    'https://cdn-we-retail.ym.tencent.com/tsr/goods/nz-09d.png',
  ],
  etitle: '',
}, {
  saasId: '4',
  storeId: '1000',
  spuId: '0',
  title: '俄罗斯小列巴',
  primaryImage: 'https://i.ibb.co/wrcrzPZ/image.jpg',
  images: [
    'https://i.ibb.co/wrcrzPZ/image.jpg',
    'https://i.ibb.co/wrcrzPZ/image.jpg',
  ],
  spuTagList: [{
    id: '13001',
    title: '爆款',
    image: null
  }],
  desc: [
    'https://cdn-we-retail.ym.tencent.com/tsr/goods/nz-09c.png',
    'https://cdn-we-retail.ym.tencent.com/tsr/goods/nz-09d.png',
  ],
  etitle: '',
}];

/**
 * @param {string} id
 * @param {number} [available] 库存, 默认1
 */
export function genGood(id, available = 1) {
  const specID = ['135681624', '135681628'];
  if (specID.indexOf(id) > -1) {
    return allGoods.filter((good) => good.spuId === id)[0];
  }
  const item = allGoods[id % allGoods.length];
  return {
    ...item,
    spuId: `${id}`,
    available: available,
    desc: item?.desc || defaultDesc,
    images: item?.images || [item?.primaryImage],
  };
}