<template>
	<el-container>
		<el-header class="mine-el-header">
			<div class="panel-container">
				<div class="left-panel">

					<el-button
						icon="el-icon-plus"
						v-auth="['giftcard:save']"
						type="primary"
						@click="add"
					>创建</el-button>

					<el-button
						type="danger"
						plain
						icon="el-icon-delete"
						v-auth="['giftcard:delete']"
						:disabled="selection.length==0"
						@click="batchDel"
					>删除</el-button>

				</div>
				<div class="right-panel">
					<div class="right-panel-search">
						<el-input v-model="queryParams.card_no" placeholder="搜索礼品卡" clearable></el-input>

						<el-tooltip class="item" effect="dark" content="搜索" placement="top">
							<el-button type="primary" icon="el-icon-search" @click="handlerSearch"></el-button>
						</el-tooltip>

						<el-tooltip class="item" effect="dark" content="清空条件" placement="top">
							<el-button icon="el-icon-refresh" @click="resetSearch"></el-button>
						</el-tooltip>

						<el-button type="primary" link @click="toggleFilterPanel">
							{{ povpoerShow ? '关闭更多筛选' : '显示更多筛选'}}
							<el-icon><el-icon-arrow-down v-if="povpoerShow" /><el-icon-arrow-up v-else /></el-icon>
						</el-button>
					</div>
				</div>
			</div>
			<el-card class="filter-panel" shadow="never">
				<el-form label-width="80px" :inline="true">

					<el-form-item label="标识" prop="currency">
						<el-input v-model="queryParams.currency" placeholder="角色标识" clearable></el-input>
					</el-form-item>

					<el-form-item label="状态" prop="status">
						<el-select  v-model="queryParams.status" clearable placeholder="状态">
							<el-option label="启用" value="0">启用</el-option>
							<el-option label="停用" value="1">停用</el-option>
						</el-select>
					</el-form-item>

					<el-form-item label="创建时间">
						<el-date-picker
							clearable

							v-model="dateRange"
							type="daterange"
							range-separator="至"
							@change="handleDateChange"
							value-format="YYYY-MM-DD"
							start-placeholder="开始日期"
							end-placeholder="结束日期"
						></el-date-picker>
					</el-form-item>

				</el-form>
			</el-card>
		</el-header>
		<el-main class="nopadding">
			<meTable
				ref="table"
				:api="api"
				:showRecycle="true"
				@selection-change="selectionChange"
				@switch-data="switchData"
				stripe
				remoteSort
				remoteFilter
			>
				<el-table-column type="selection" width="50"></el-table-column>

				<el-table-column
					label="ID"
					prop="id"
					width="100"
				></el-table-column>


				<el-table-column
					label="礼品卡号"
					prop="card_no"
					sortable='custom'
					width="260"
				></el-table-column>

				<el-table-column
					label="币种"
					prop="currency"
					width="220"
				></el-table-column>

				<el-table-column
					label="状态"
					prop="status"
					width="200"
				>
					<template #default="scope">
						<el-switch
							v-model="scope.row.status"
							@change="handleStatus($event, scope.row)"
							:active-value="1"
							:inactive-value="2"
						></el-switch>
					</template>
				</el-table-column>

				<el-table-column
					label="最后登录时间"
					prop="login_time"
					width="260"
				></el-table-column>


				<el-table-column
					label="创建时间"
					prop="created_at"
					width="260"
					sortable='custom'
				></el-table-column>

				<!-- 正常数据操作按钮 -->
				<el-table-column label="操作" fixed="right" align="right" width="330" v-if="!isRecycle">
					<template #default="scope">

						<el-button
							type="primary" link
							@click="show(scope.row, scope.$index)"
						>查看</el-button>


						<el-button
							type="primary" link
							@click="edit(scope.row, scope.$index)"
							v-if="$AUTH('giftcard:update')"
						>编辑</el-button>

						<el-button
							type="primary" link
							v-if="$AUTH('giftcard:delete')"
							@click="deletes(scope.row.id)"
						>删除</el-button>

					</template>
				</el-table-column>

				<!-- 回收站操作按钮 -->
				<el-table-column label="操作" fixed="right" align="right" width="130" v-else>
					<template #default="scope">

						<el-button
							type="primary" link
							v-auth="['giftcard:recovery']"
							@click="recovery(scope.row.id)"
						>恢复</el-button>

						<el-button
							type="primary" link
							v-auth="['giftcard:realDelete']"
							@click="deletes(scope.row.id)"
						>删除</el-button>

					</template>
				</el-table-column>

			</meTable>
		</el-main>
	</el-container>

	<save-dialog v-if="dialog.save" ref="saveDialog" @success="handleSuccess" @closed="dialog.save=false" />

	<menu-dialog ref="menuDialog" />
	<data-dialog ref="dataDialog" @success="handleSuccess" />

</template>

<script>
import saveDialog from './save'
import menuDialog from './menuForm'
import dataDialog from './dataForm'

export default {
	name: 'giftcard:index',
	components: {
		saveDialog,
		menuDialog,
		dataDialog,
	},

	data() {
		return {
			dialog: {
				save: false
			},
			povpoerShow: false,
			dateRange:'',
			showDeptloading: false,
			deptFilterText: '',
			dept: [],
			api: {
				list: this.$API.giftcard.pageList,
				recycleList: this.$API.giftcard.getRecyclePageList,
			},
			selection: [],
			queryParams: {
				card_no: undefined,
				currency: undefined,
				maxDate: undefined,
				minDate: undefined,
				status: undefined
			},
			isRecycle: false,
		}
	},
	methods: {
		//添加
		add(){
			this.dialog.save = true
			this.$nextTick(() => {
				this.$refs.saveDialog.open()
			})
		},
		//编辑
		edit(row){
			this.dialog.save = true
			this.$nextTick(() => {
				this.$refs.saveDialog.open('edit').setData(row)
			})
		},
		//查看
		show(row){
			this.dialog.save = true
			this.$nextTick(() => {
				this.$refs.saveDialog.open('show').setData(row)
			})
		},
		//批量删除
		async batchDel(){
			await this.$confirm(`确定删除选中的 ${this.selection.length} 项吗？`, '提示', {
				confirmButtonText: '确定',
				cancelButtonText: '取消',
				type: 'warning'
			}).then(() => {
				const loading = this.$loading();
				let ids = []
				this.selection.map(item => ids.push(item.id))
				if (this.isRecycle) {
					this.$API.giftcard.realDeletes(ids.join(',')).then(() => {
						this.$refs.table.upData(this.queryParams)
					})
				} else {
					this.$API.giftcard.deletes(ids.join(',')).then(() => {
						this.$refs.table.upData(this.queryParams)
					})
				}
				loading.close();
				this.$message.success("操作成功")
			})
		},

		// 单个删除
		async deletes(id) {
			await this.$confirm(`确定删除该数据吗？`, '提示', {
				confirmButtonText: '确定',
				cancelButtonText: '取消',
				type: 'warning'
			}).then(() => {
				const loading = this.$loading();
				if (this.isRecycle) {
					this.$API.giftcard.realDeletes(id).then(() => {
						this.$refs.table.upData(this.queryParams)
					})
				} else {
					this.$API.giftcard.deletes(id).then(() => {
						this.$refs.table.upData(this.queryParams)
					})
				}
				loading.close();
				this.$message.success("操作成功")
			}).catch(()=>{})
		},

		// 恢复数据
		async recovery (id) {
			await this.$API.giftcard.recoverys(id).then(res => {
				this.$message.success(res.message)
				this.$refs.table.upData(this.queryParams)
			})
		},

		//表格选择后回调事件
		selectionChange(selection){
			this.selection = selection;
		},

		// 选择时间事件
		handleDateChange (values) {
			if (values !== null) {
				this.queryParams.minDate = values[0]
				this.queryParams.maxDate = values[1]
			}
		},

		toggleFilterPanel() {
			this.povpoerShow = ! this.povpoerShow
			document.querySelector('.filter-panel').style.display = this.povpoerShow ? 'block' : 'none'
		},

		//搜索
		handlerSearch(){
			this.$refs.table.upData(this.queryParams)
		},

		// 切换数据类型回调
		switchData(isRecycle) {
			this.isRecycle = isRecycle
		},

		// 状态更改
		handleStatus (val, row) {
			const status = row.status === 1 ? 1 : 2
			const text = row.status === 1 ? '启用' : '停用'
			this.$confirm(`确认要${text} ${row.username} 商户吗？`, '提示', {
				type: 'warning',
				confirmButtonText: '确定',
				cancelButtonText: '取消'
			}).then(() => {
				this.$API.merchant.changeStatus({ id: row.id, status }).then(() => {
					this.$message.success(text + '成功')
				})
			}).catch(() => {
				row.status = (row.status === 1) ? 2 : 1
			})
		},

		resetSearch() {
			this.queryParams = {
				name: undefined,
				code: undefined,
				maxDate: undefined,
				minDate: undefined,
				status: undefined
			}
			this.$refs.table.upData(this.queryParams)
		},

		//本地更新数据
		handleSuccess(){
			this.$refs.table.upData(this.queryParams)
		}
	}
}
</script>
