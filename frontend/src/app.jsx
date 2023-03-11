import Taro, { Component } from '@tarojs/taro'
import Index from './pages/index'

import './app.scss'

// 如果需要在 h5 环境中开启 React Devtools
// 取消以下注释：
// if (process.env.NODE_ENV !== 'production' && process.env.TARO_ENV === 'h5')  {
//   require('nerv-devtools')
// }

class App extends Component {

  componentDidMount() { }

  componentDidShow() { }

  componentDidHide() { }

  componentDidCatchError() { }

  config = {
    pages: [
      'pages/index/index',
      'pages/question/index',
      'pages/mine/index',
      'pages/login/index'
    ],
    window: {
      backgroundTextStyle: 'light',
      navigationBarBackgroundColor: '#fff',
      navigationBarTitleText: 'WeChat',
      navigationBarTextStyle: 'black'
    },
    tabBar: {
      list: [
        {
          pagePath: 'pages/index/index',
          text: '复习',
          iconPath: './assets/images/tab_trend.png',
          selectedIconPath: './assets/images/tab_trend_s.png'
        },
        {
          pagePath: 'pages/question/index',
          text: '题库',
          iconPath: './assets/images/tab_news.png',
          selectedIconPath: './assets/images/tab_news_s.png'
        },
        {
          pagePath: 'pages/mine/index',
          text: '我的',
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
  }

  // 在 App 类中的 render() 函数没有实际作用
  // 请勿修改此函数
  render() {
    return (
      <Index />
    )
  }
}

Taro.render(<App />, document.getElementById('app'))
