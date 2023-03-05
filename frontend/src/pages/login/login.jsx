import Taro, { Component } from '@tarojs/taro'
import { AtButton, AtIcon } from 'taro-ui'
import { wxLogin } from '../../api/api'
import store from '../../utils/store'
import { View, Text } from '@tarojs/components'

export class Login extends Component {
  config = {
    navigationBarTitleText: '登陆'
  }

  componentWillMount() { }

  componentDidMount() { }

  componentWillUnmount() { }

  componentDidShow() { }

  componentDidHide() { }

  onMpLoginClick() {
    Taro.showLoading({
      title: "登陆中",
    })
    setTimeout(() => {
      Taro.hideLoading()
    }, 5000)

    Taro.login({
      success: function (res) {
        if (res.code) {
          wxLogin({ code: res.code })
            .then((res) => {
              if (res.code === 0) {
                store.set('token', res.data.token)

                // todo 获取个人信息
              } else {
                Taro.showToast({
                  title: res.errMsg,
                  duration: 500,
                })
              }
            })
            .catch((res) => {
              Taro.showToast({
                title: res.errMsg,
                duration: 500,
              })
            })
        }
      },
      fail: function (failRes) {
        Taro.showToast({
          title: failRes.errMsg,
          duration: 500,
        })
      }
    })
  }

  render() {
    return (
      <View className="container">
        <View className="logo">
          {/* <AtAvatar customStyle="background-color: rgba(0, 0, 0, 0);" image={require("../../static/images/sakurago_logo.png")}></AtAvatar> */}
          <Text>SakuraGo</Text>
        </View>
        <View className="loginbtn">
          <AtButton type='primary' size='normal' onClick={this.onMpLoginClick}>
            <View className="loginbtn-inner">
              <AtIcon value='star' size='20' color='#fff'></AtIcon>
              <Text className="loginbtn-text">微信授权登录</Text>
            </View>
          </AtButton>
        </View>
      </View>
    )
  }
}
