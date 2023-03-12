import { Component } from '@tarojs/taro'
import { View } from '@tarojs/components'
import { IButton } from './button'
import { Idiom } from './idiom'
import { AtButton } from 'taro-ui'

import './index.scss'


export default class Index extends Component {

  componentWillMount() { }

  componentDidMount() { }

  componentWillUnmount() { }

  componentDidShow() { }

  componentDidHide() { }

  config = {
    navigationBarTitleText: '复习'
  }

  handleButClick() {
    console.log('no click')
  }

  handleLogin() {
    Taro.navigateTo({
      url: '/pages/login/index',
    })
  }

  render() {
    return (
      <View className='index'>
        <Idiom />
        <IButton onClick={this.handleButClick} />
        {/* // todo */}
        {/* <Login /> */}
        <AtButton type='primary' size='normal' onClick={this.handleLogin}>登陆</AtButton>
      </View>
    )
  }
}
