export default defineAppConfig({
  pages: [
    'pages/index/index',
    'pages/me/profile'
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
        pagePath: 'pages/me/profile',
        text: 'Me',
        iconPath: './assets/images/tab_me.png',
        selectedIconPath: './assets/images/tab_me_s.png'
      }
    ],
    color: '#8a8a8a',
    selectedColor: '#2d8cf0',
    backgroundColor: '#ffffff',
    borderStyle: 'white',
    darkmode: false
  },

})
