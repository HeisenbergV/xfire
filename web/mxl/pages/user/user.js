// pages/about/about.js
Page({
  data: {
    developer: {
      img: "/mm.jpg",
      items: [{
          title: "手机号/微信",
          text: "15222201847"
        }, {
          title: "地址",
          text: "静海区静海镇良王庄"
        },
        {
          title: "关于我们",
          text: `
          我们的面包一直秉承健康美味，以绿色文化、鱼水文化为产品核心价值。在天津我们拥有15000+超市，是天津老百姓的首选
          石家庄、廊坊、济南等地正在发展目前有600+超市选择我们的面包`
        }
      ],
      versions: "1.0.0"
    }
  },
  onShareAppMessage() {
    wx.showShareMenu({
      withShareTicket: true,
      menus: ['shareAppMessage', 'shareTimeline']
    })
  },
})