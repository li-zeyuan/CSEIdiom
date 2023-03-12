import Taro, { Component, uploadFile } from '@tarojs/taro'
import { View } from '@tarojs/components'
import store from '../../utils/store'

export default class Index extends Component {
  config = {
    navigationBarTitleText: '我的'
  }

  componentWillMount() {
    const uProfile = store.get(store.userProfiltKey)
    console.log(uProfile)
  }

  componentDidMount() {}

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
