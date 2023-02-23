export default defineAppConfig({
  pages: [
    'pages/index/index',
    'pages/mine/mine'
  ],
  window: {
    backgroundTextStyle: 'dark',
    navigationBarBackgroundColor: '#fff',
    navigationBarTitleText: 'WeChat',
    navigationBarTextStyle: 'black'
  },
  tabBar: {
    list: [
      {
        pagePath: 'pages/index/index',
        text: 'Studying',
        iconPath: './assets/images/tab_trend.png',
        selectedIconPath: './assets/images/tab_trend_s.png'
      },
      {
        pagePath: 'pages/mine/mine',
        text: 'Me',
        iconPath: './assets/images/tab_news.png',
        selectedIconPath: './assets/images/tab_news_s.png'
      }
    ],
    color: '#8a8a8a',
    selectedColor: '#2d8cf0',
    backgroundColor: '#ffffff',
    borderStyle: 'white',
    darkmode: false
  },

})
