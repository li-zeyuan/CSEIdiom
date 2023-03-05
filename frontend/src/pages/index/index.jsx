import  { Component } from '@tarojs/taro'
import { View } from '@tarojs/components'
import { IButton } from './button'
import { Idiom } from './idiom'
import { Login } from '../login/login'

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

  render() {
    return (
      <View className='index'>
        <Idiom />
        <IButton onClick={this.handleButClick} />
        <Login />
      </View>
    )
  }
}
