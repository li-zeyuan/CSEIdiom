import { Component } from '@tarojs/taro'
import { View } from '@tarojs/components'
import { StoreGet } from '../../utils/store'

export default class Index extends Component {
  config = {
    navigationBarTitleText: '我的'
  }

  componentWillMount() {
    debugger
    // const uProfile = StoreGet('user')
    // console.log(uProfile)
    console.log(111)
  }

  componentDidMount() {

  }

  componentWillUnmount() { }

  componentDidShow() { }

  componentDidHide() { }

  render() {
    return (
      <View className='index'>
        hello mine
      </View>
    )
  }
}
