import request from "../utils/request";

export function wxLogin(data) {
    return request(
        'post',
        '/api/login/wechat_login',
        data
    )
}