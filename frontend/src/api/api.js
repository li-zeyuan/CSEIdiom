import request from "../utils/request";

export function wxLogin(data) {
    return request(
        'post',
        '/api/login/wechat',
        data
    )
}

export function userProfile(data) {
    return request(
        'get',
        '/api/user/detail',
        data
    )
}