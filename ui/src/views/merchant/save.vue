<template>
	<el-dialog
		:title="titleMap[mode]"
		v-model="visible"
		append-to-body
		:width="800"
		destroy-on-close
		@closed="$emit('closed')"
	>
		<el-form
			:model="form"
			:rules="rules"
			:disabled="mode=='show'"
			ref="dialogForm"
			label-width="80px"
			label-position="right"
			v-loading="loading"
			element-loading-background="rgba(255, 255, 255, 0.01)"
			element-loading-text="数据加载中..."
		>

			<el-row :gutter="20">
				<el-col :span="12">
					<el-form-item label="商户名称" prop="merchant_name">
						<!--	<el-input v-model="form.merchant_name" placeholder="商户名称" clearable :disabled="mode!='add'" />-->
						<el-input v-model="form.merchant_name" placeholder="商户名称" clearable  />
					</el-form-item>
				</el-col>

				<el-col :span="12">
					<el-form-item label="登录帐号" prop="username">
						<el-input v-model="form.username" placeholder="用于登录系统" clearable  />
					</el-form-item>
				</el-col>


			</el-row>

			<el-row :gutter="20">

				<el-col :span="12">
					<el-form-item label="登录密码" prop="password" >
						<el-input type="password" v-model="form.password" clearable show-password />
						<!--
							<el-input type="password" v-model="form.password" v-if="mode=='add'" clearable show-password />
						<el-input v-else placeholder="密码不可修改" :disabled="true" />
						-->
					</el-form-item>
				</el-col>


				<el-col :span="12">
					<el-form-item label="资金密码" prop="fund_password" >
						<el-input type="fund_password" v-model="form.fund_password" clearable show-password />
						<!--
						<el-input type="fund_password" v-model="form.fund_password" v-if="mode=='add'" clearable show-password />
						<el-input v-else placeholder="密码不可修改" :disabled="true" />
						-->
					</el-form-item>
				</el-col>

			</el-row>

			<el-row :gutter="20">
				<el-col :span="12">
					<el-form-item label="角色" prop="role_ids">
						<el-select v-model="form.role_ids" clearable style="width:100%" multiple placeholder="请选择用户角色">
							<el-option v-for="item in roleData" :key="item.id" :label="item.name" :value="item.id">
							</el-option>
						</el-select>
					</el-form-item>
				</el-col>
			</el-row>

			<el-row :gutter="20">
				<el-col :span="12">
					<el-form-item label="邮箱" prop="email">
						<el-input v-model="form.email" placeholder="请输入电子邮箱" clearable />
					</el-form-item>
				</el-col>
				<el-col :span="12">
					<el-form-item label="手机 " prop="phone">
						<el-input v-model="form.phone" placeholder="请输入手机" />
					</el-form-item>
				</el-col>
			</el-row>

			<el-row :gutter="20">
				<el-col :span="24">
					<el-form-item label="备注" prop="remark">
						<el-input type="textarea" :rows="3" placeholder="用户备注信息" v-model="form.remark" maxlength="255" show-word-limit />
					</el-form-item>
				</el-col>
			</el-row>
		</el-form>
		<template #footer>
			<el-button @click="visible=false" >取 消</el-button>
			<el-button v-if="mode!='show'" type="primary" :loading="isSaveing" @click="submit()">保 存</el-button>
		</template>
	</el-dialog>
</template>

<script>
export default {
	emits: ['success', 'closed'],
	data() {
		return {
			mode: "add",
			loading: false,
			titleMap: {
				add: '新增商户',
				edit: '编辑商户',
				show: '查看'
			},
			// 用户选择器数据
			roleData: [],
			visible: false,
			isSaveing: false,
			//表单数据
			form: {
				id: 0,
				merchant_name: '',
				username: '',
				phone: '',
				password: '123456',
				fund_password: '123456',
				role_ids: '',
				email: '',
				remark: ''
			},
			//验证规则,merchant_name
			rules: {
				merchant_name: [{ required: true, message: '请输入商户名称', trigger: 'blur' }],
				username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
				password: [{ required: true, message: '请输入用户密码', trigger: 'blur' }],
				fund_password: [{ required: true, message: '请输入资金密码', trigger: 'blur' }],
				role_ids: [{ required: true, message: '请选择角色', trigger: 'blur' }],
				email: [{ type: 'email', message: '请输入正确的邮箱地址', trigger: ['blur', 'change'] }],
				phone: [{ pattern: /^1[3|4|5|6|7|8|9][0-9]\d{8}$/, message: '请输入正确的手机号码', trigger: ['blur'] }]
			}
		}
	},
	mounted() {
		this.initData()
	},
	methods: {

		//显示
		open(mode='add'){
			this.mode = mode;
			this.visible = true;
			return this
		},

		//表单提交方法
		submit(){
			this.$refs.dialogForm.validate(async (valid) => {
				if (valid) {
					this.isSaveing = true;
					let res = null
					if (this.mode == 'add') {
						res = await this.$API.merchant.save(this.form)
					} else {
						res = await this.$API.merchant.update(this.form.id, this.form)
					}
					this.isSaveing = false;
					if(res.success){
						this.$emit('success', this.form, this.mode)
						this.visible = false;
						this.$message.success(res.message)
					}else{
						this.$alert(res.message, "提示", { type: 'error' })
					}
				}else{
					return false;
				}
			})

		},

		handleResource(data) {
			this.form.avatar = this.viewImage(data)
		},

		// 请求部门、角色、岗位数据
		async initData () {
			await this.$API.role.getList().then(res => {
				this.roleData = res.data
			})
		},

		//表单注入数据
		async setData(data){
			this.loading = true
			this.form.id = data.id
			this.form.merchant_name = data.merchant_name
			this.form.username = data.username
			this.form.phone = data.phone
			this.form.role_ids = data.role_ids
			this.form.email = data.email
			this.form.remark = data.remark

			await this.$API.merchant.detail(data.id).then(res => {
				/*
				this.form.role_ids = res.data.roleList.map(item => {
					return item.id
				})
				this.form.post_ids = res.data.postList.map(item => {
					return item.id
				})
				 */
			})

			this.loading = false

			//可以和上面一样单个注入，也可以像下面一样直接合并进去
			//Object.assign(this.form, data)
		}
	}
}
</script>

<style scoped lang="scss">
:deep(.el-avatar--square) {
	display:flex;
	margin-bottom: 8px;
	justify-content: center;
	align-items:center
}
</style>
