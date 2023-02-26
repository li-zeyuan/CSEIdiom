import Taro, { Component } from '@tarojs/taro'
import { View, Button } from '@tarojs/components'

export class IButton extends Component {
  state = {
    btn: [
      {
        key: 1,
        className: 'cognitive',
        text: '认识',
        type: 'primary'
      },
      {
        key: 2,
        className: 'forgotten',
        text: '忘记',
        type: 'warn'
      },
    ]
  }

  componentWillMount() { }

  componentDidMount() { }

  componentWillUnmount() { }

  componentDidShow() { }

  componentDidHide() { }

  render() {
    return (
      <View className='iButton'>
        {this.state.btn.map(item => {
          return (
            <Button
              key={item.key}
              className={item.className}
              type={item.type}
              onClick={() => this.props.onClick()}
            >
              {item.text}
            </Button>
          )
        })}
      </View>
    )
  }
}
