import {request} from '@/utils/request.js'

export default {

	attrOptions(params = {}) {
		return request({
			url: 'giftcard/attr_options',
			method: 'get',
			params
		})
	},
}
