import store from "./store"
import config from "../config/config"

const apiHost = config.apiHost

export default function request(method, url, data) {
    const token = store.get('token')
    if (!config.notNeedTokenUrl.includes(url) && !token) {
        Taro.reLaunch({
            url: '/pages/login/login',
        })
    }

    return new Promise((resolve, reject) => {
        let apiUrl = apiHost
        switch (method) {
            case 'get': {
                apiUrl += url + '?' + Object.keys(data).map(k => {
                    const v = data[k]
                    return `${k}=${v}`
                }).join('&')
                break
            }
            case 'post': {
                apiUrl += url
            }
        }

        Taro.request({
            method: method,
            url: apiHost,
            data: method == 'post' ? data : {},
            header: { 'Authorization': token },
            success: function (res) {
                if (config.skipToLoginCode.includes(res.code)) {
                    Taro.reLaunch({
                        url: '/pages/login/login',
                    })
                }

                resolve(res)
            },
            fail: function (res) {
                reject(res)
            }
        })
    })
}