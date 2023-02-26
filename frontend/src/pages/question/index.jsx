import Taro, { Component } from '@tarojs/taro'
import { View } from '@tarojs/components'


export default class Index extends Component {
  config = {
    navigationBarTitleText: '题库'
  }

  componentWillMount() { }

  componentDidMount() { }

  componentWillUnmount() { }

  componentDidShow() { }

  componentDidHide() { }

  render() {
    return (
      <View className='index'>
        题库
      </View>
    )
  }
}
