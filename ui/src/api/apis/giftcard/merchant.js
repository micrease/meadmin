import {request} from '@/utils/request.js'

export default {


	/**
	 * 添加用户
	 * @returns
	 */
	save(params = {}) {
		return request({
			url: 'giftcard/merchant/save',
			method: 'post',
			data: params
		})
	},

	/**
	 * 更新数据,和添加一样
	 * @returns
	 */
	update (id, params = {}) {
		return request({
			url: 'giftcard/merchant/save',
			method: 'post',
			data: params
		})
	},

	pageList(params = {}) {
		return request({
			url: 'giftcard/merchant/index',
			method: 'get',
			params
		})
	},

	changeStatus (params = {}) {
		return request({
			url: 'giftcard/merchant/change_status',
			method: 'put',
			data: params
		})
	},

	detail (id) {
		return request({
			url: 'giftcard/merchant/detail/' + id,
			method: 'get'
		})
	},
}
