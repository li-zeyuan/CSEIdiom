import Taro, { Component } from '@tarojs/taro'
import { View, Text } from '@tarojs/components'
import { AtButton } from 'taro-ui'
import { IButton } from './button'
import { Idiom } from './idiom'

import './index.scss'

export default class Index extends Component {

  componentWillMount () { }

  componentDidMount () { }

  componentWillUnmount () { }

  componentDidShow () { }

  componentDidHide () { }

  config = {
    navigationBarTitleText: '复习'
  }

  handleButClick() {
    console.log('no click')
  }

  render () {
    return (
      <View className='index'>
        <AtButton type='primary'>按钮文案</AtButton>
        <Idiom />
        <IButton onClick={this.handleButClick} />
      </View>
    )
  }
}
