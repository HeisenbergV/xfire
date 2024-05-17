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
  primaryImage: 'https://s2.loli.net/2024/05/17/N6RGovQlj2p9eEJ.jpg',
  images: [
    'https://s2.loli.net/2024/05/17/N6RGovQlj2p9eEJ.jpg',
    'https://s2.loli.net/2024/05/17/N6RGovQlj2p9eEJ.jpg',
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
  primaryImage: 'https://s2.loli.net/2024/05/17/NVCKeMxdu2tgmyY.jpg',
  images: [
    'https://s2.loli.net/2024/05/17/NVCKeMxdu2tgmyY.jpg',
    'https://s2.loli.net/2024/05/17/NVCKeMxdu2tgmyY.jpg',
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
  primaryImage: 'https://s2.loli.net/2024/05/17/lDVf8KoqyUd7wJs.jpg',
  images: [
    'https://s2.loli.net/2024/05/17/lDVf8KoqyUd7wJs.jpg',
    'https://s2.loli.net/2024/05/17/lDVf8KoqyUd7wJs.jpg',
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
  primaryImage: 'https://s2.loli.net/2024/05/17/sMTkohBtlNnwviy.jpg',
  images: [
    'https://s2.loli.net/2024/05/17/sMTkohBtlNnwviy.jpg',
    'https://s2.loli.net/2024/05/17/sMTkohBtlNnwviy.jpg',
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