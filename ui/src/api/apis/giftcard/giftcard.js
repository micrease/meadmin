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


	/**
	 * 移到回收站
	 * @returns
	 */
	deletes (ids) {
		return request({
			url: 'giftcard/delete/' + ids,
			method: 'delete'
		})
	},

	/**
	 * 恢复数据
	 * @returns
	 */
	recoverys (ids) {
		return request({
			url: 'giftcard/recovery/' + ids,
			method: 'put'
		})
	},

	/**
	 * 真实删除
	 * @returns
	 */
	realDeletes (ids) {
		return request({
			url: 'giftcard/realDelete/' + ids,
			method: 'delete'
		})
	},

}
