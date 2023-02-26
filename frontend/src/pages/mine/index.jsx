import Taro, { Component } from '@tarojs/taro'
import { View } from '@tarojs/components'

export default class Index extends Component {
  config = {
    navigationBarTitleText: '我的'
  }

  componentWillMount () { }

  componentDidMount () { }

  componentWillUnmount () { }

  componentDidShow () { }

  componentDidHide () { }

  render () {
    return (
      <View className='index'>
        hello mine
      </View>
    )
  }
}
