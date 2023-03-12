import Taro from '@tarojs/taro'

export default {
    tokenKey: 'token',
    userProfiltKey: 'user',

    set(key, data) {
        console.log('store set', key, data)
        try {
            Taro.setStorageSync(key, JSON.stringify(data))
        } catch (e) {
            console.error(e)
        }
    },
    get(key) {
        const d = Taro.getStorageSync(key)
        return d ? JSON.parse(d ? d : '{}') : ''
    }
}