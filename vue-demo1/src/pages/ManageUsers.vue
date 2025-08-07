<template>
  <CommonLayout :page-name="currentName">
	<template #main>
	  <div class="playlist-container">
		<!-- 搜索和筛选区域 -->
		<div class="filter-section">
		  <el-input
			  v-model="searchQuery"
			  placeholder="搜索用户"
			  class="search-input"
			  :prefix-icon="Search"
			  clearable
			  @clear="handleSearch"
			  @keyup.enter="handleSearch"
		  />
		  <el-select v-model="role" placeholder="用户角色" @change="handleSearch">
			<el-option label="全部" value=""/>
			<el-option label="普通用户" value="user"/>
			<el-option label="管理员" value="admin"/>
		  </el-select>
		</div>

		<!-- 加载状态 -->
		<el-loading
			:visible="isLoading"
			element-loading-text="加载中..."
			element-loading-background="rgba(255, 255, 255, 0.8)"
		/>

		<!-- 空状态 -->
		<el-empty
			v-if="!isLoading && users.length === 0"
			description="暂无用户"
		/>

		<!-- 用户列表 -->
		<el-table
			v-if="users.length > 0"
			:data="users"
			style="width: 100%"
			:header-cell-style="{ background: '#f5f7fa' }"
		>
		  <el-table-column prop="id" label="ID" width="80"/>

		  <el-table-column label="用户信息" min-width="200">
			<template #default="{ row }">
			  <div class="user-info">
				<el-avatar :src="row.avatar || '/assets/default-avatar.jpg'" :size="40"/>
				<div class="user-details">
				  <div class="username">{{ row.username }}</div>
				  <div class="email">{{ row.email }}</div>
				</div>
			  </div>
			</template>
		  </el-table-column>

		  <el-table-column prop="name" label="昵称" min-width="120"/>

		  <el-table-column label="角色" width="120">
			<template #default="{ row }">
			  <el-tag :type="row.role === 'admin' ? 'danger' : 'info'">
				{{ row.role === 'admin' ? '管理员' : '普通用户' }}
			  </el-tag>
			</template>
		  </el-table-column>

		  <el-table-column prop="created_at" label="注册时间" width="180"/>

		  <el-table-column label="状态" width="100">
			<template #default="{ row }">
			  <el-switch
				  v-model="row.is_active"
				  :active-value="1"
				  :inactive-value="0"
				  @change="handleStatusChange(row)"
			  />
			</template>
		  </el-table-column>

		  <el-table-column label="操作" width="200" fixed="right">
			<template #default="{ row }">
			  <el-button-group>
				<el-button
					type="primary"
					size="small"
					@click="editUser(row)"
				>
				  编辑
				</el-button>
				<el-button
					type="danger"
					size="small"
					@click="deleteUser(row)"
				>
				  删除
				</el-button>
			  </el-button-group>
			</template>
		  </el-table-column>
		</el-table>

		<!-- 分页器 -->
		<el-pagination
			v-if="users.length > 0"
			v-model:current-page="page"
			v-model:page-size="pageSize"
			:total="totalUsers"
			:page-sizes="[10, 20, 30, 50]"
			layout="total, sizes, prev, pager, next"
			class="pagination"
			@size-change="handlePageSizeChange"
			@current-change="handlePageChange"
		/>

		<!-- 编辑用户对话框 -->
		<el-dialog
			v-model="editDialogVisible"
			:title="editForm.id ? '编辑用户' : '新增用户'"
			width="500px"
		>
		  <el-form :model="editForm" label-width="80px">
			<el-form-item label="用户名">
			  <el-input v-model="editForm.username" :disabled="!!editForm.id"/>
			</el-form-item>
			<el-form-item label="昵称">
			  <el-input v-model="editForm.name"/>
			</el-form-item>
			<el-form-item label="邮箱">
			  <el-input v-model="editForm.email"/>
			</el-form-item>
			<el-form-item label="角色">
			  <el-select v-model="editForm.role">
				<el-option label="普通用户" value="user"/>
				<el-option label="管理员" value="admin"/>
			  </el-select>
			</el-form-item>
		  </el-form>
		  <template #footer>
			<el-button @click="editDialogVisible = false">取消</el-button>
			<el-button type="primary" @click="submitEdit">确定</el-button>
		  </template>
		</el-dialog>
	  </div>
	</template>
  </CommonLayout>
</template>

<script setup>
import {ref, onMounted} from 'vue'
import CommonLayout from 'src/layouts/CommonLayout.vue'
import {ElMessage, ElMessageBox} from 'element-plus'
import {Search} from '@element-plus/icons-vue'

const currentName = ref('用户管理')
const users = ref([])
const page = ref(1)
const pageSize = ref(10)
const totalUsers = ref(0)
const isLoading = ref(false)
const searchQuery = ref('')
const role = ref('')
const editDialogVisible = ref(false)
const editForm = ref({
  id: '',
  username: '',
  name: '',
  email: '',
  role: 'user'
})

const handleSearch = async () => {
  page.value = 1
  await loadData()
}

const handlePageChange = async (newPage) => {
  page.value = newPage
  await loadData()
}

const handlePageSizeChange = async (newSize) => {
  pageSize.value = newSize
  page.value = 1
  await loadData()
}

const loadData = async () => {
  isLoading.value = true
  try {
	// TODO: 实现获取用户列表的接口
	// const response = await getUsers(page.value, pageSize.value, role.value)
	isLoading.value = false
  } catch (error) {
	console.error('加载数据失败:', error)
	ElMessage.error('加载数据失败，请稍后重试')
	isLoading.value = false
  }
}

const editUser = (user) => {
  editForm.value = {...user}
  editDialogVisible.value = true
}

const submitEdit = async () => {
  // TODO: 实现编辑用户的接口
  ElMessage.success('用户信息已更新')
  editDialogVisible.value = false
  await loadData()
}

const handleStatusChange = async (user) => {
  // TODO: 实现更新用户状态的接口
  ElMessage.success(`用户状态已${user.is_active ? '启用' : '禁用'}`)
}

const deleteUser = async (user) => {
  try {
	await ElMessageBox.confirm(
		'确定要删除这个用户吗？此操作不可恢复！',
		'删除用户',
		{
		  confirmButtonText: '确定',
		  cancelButtonText: '取消',
		  type: 'warning'
		}
	)
	// TODO: 实现删除用户的接口
	ElMessage.success('用户已删除')
	await loadData()
  } catch {
	// 用户取消删除
  }
}

onMounted(() => {
  loadData()
})
</script>

<style scoped>
@import '../styles/common.css';

.user-info {
  display: flex;
  align-items: center;
  gap: 12px;
}

.user-details {
  display: flex;
  flex-direction: column;
}

.username {
  font-weight: 500;
}

.email {
  font-size: 12px;
  color: #666;
}

.el-table {
  margin: 20px 0;
}
</style>