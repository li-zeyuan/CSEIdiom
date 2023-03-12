import Taro from '@tarojs/taro'

export const TokenKey = 'token'
export const  UserProfiltKey = 'user'

export function StoreSet(key, data) {
    console.log('store set', key, data)
}

export function StoreGet(key) {
    console.log('store get', key)
}

// export default {
//     tokenKey: 'token',
//     userProfiltKey: 'user',

//     set(key, data) {
//         console.log('store set', key, data)
//         try {
//             Taro.setStorageSync(key, JSON.stringify(data))
//         } catch (e) {
//             console.error(e)
//         }
//     },
//     get(key) {
//         console.log(key)
//         const d = Taro.getStorageSync(key)
//         return d ? JSON.parse(d ? d : '{}') : ''
//     }
// }