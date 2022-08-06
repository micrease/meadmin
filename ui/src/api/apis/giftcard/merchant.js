import { request } from '@/utils/request.js'

export default {


    /**
     * 添加用户
     * @returns
     */
    save (params = {}) {
        return request({
            url: 'giftcard/merchant/save',
            method: 'post',
            data: params
        })
    },
}
