import {request} from '@/utils/request.js'

export default {

	attrOptions(params = {}) {
		return request({
			url: 'giftcard/attr_options',
			method: 'get',
			params
		})
	},
	pageList(params = {}) {
		return request({
			url: 'giftcard/index',
			method: 'get',
			params
		})
	},
	/**
	 * 添加用户
	 * @returns
	 */
	save(params = {}) {
		return request({
			url: 'giftcard/save',
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
			url: 'giftcard/save',
			method: 'post',
			data: params
		})
	},

	changeStatus (params = {}) {
		return request({
			url: 'giftcard/change_status',
			method: 'put',
			data: params
		})
	},

	detail (id) {
		return request({
			url: 'giftcard/detail/' + id,
			method: 'get'
		})
	},
}
