<template>
  <common-layout>
    <template #main>
      <div class="profile-container" v-loading="isLoading">
        <!-- 顶部个人信息区域 -->
        <div class="profile-header" :class="{ 'header-loaded': !isLoading }">
          <div class="profile-cover">
            <div class="profile-info">
              <div class="left-section">
                <el-avatar 
                  :size="90" 
                  :src="userInfo?.avatar"
                  class="profile-avatar"
                />
                <div class="user-meta">
                  <h1>{{ userInfo?.name || userInfo?.username }}</h1>
                  <p class="bio">{{ userInfo?.bio || '这个人很懒，什么都没写~' }}</p>
                  <el-button 
                    v-if="isCurrentUser" 
                    type="primary" 
                    class="edit-btn"
                    size="small"
                    @click="showEditDialog = true"
                  >
                    <el-icon><Edit /></el-icon>
                    编辑资料
                  </el-button>
                </div>
              </div>
              <div class="right-section">
                <div class="user-stats">
                  <div class="stat-item" v-for="(stat, index) in userStats" :key="index">
                    <h3>{{ stat.value }}</h3>
                    <p>{{ stat.label }}</p>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- 主要内容区域 -->
        <div class="profile-content">
          <el-tabs v-model="activeTab" class="profile-tabs">
            <el-tab-pane label="个人信息" name="info">
              <el-descriptions :column="2" border>
                <el-descriptions-item label="用户名">
                  {{ userInfo?.username }}
                </el-descriptions-item>
                <el-descriptions-item label="性别">
                  {{ formatGender(userInfo?.gender) }}
                </el-descriptions-item>
                <el-descriptions-item label="生日">
                  {{ formatDate(userInfo?.birthday) }}
                </el-descriptions-item>
                <el-descriptions-item label="地区">
                  {{ userInfo?.location || '未设置' }}
                </el-descriptions-item>
                <el-descriptions-item label="邮箱">
                  {{ userInfo?.email || '未设置' }}
                </el-descriptions-item>
                <el-descriptions-item label="手机">
                  {{ formatPhone(userInfo?.phone) }}
                </el-descriptions-item>
              </el-descriptions>
            </el-tab-pane>
          </el-tabs>
        </div>

        <!-- 编辑资料对话框 -->
        <el-dialog
          v-model="showEditDialog"
          title="编辑个人资料"
          width="500px"
          class="edit-dialog"
        >
          <el-form :model="editForm" label-width="80px">
            <el-form-item label="头像">
              <el-upload
                class="avatar-uploader"
                :action="uploadUrl"
                :show-file-list="false"
                :on-success="handleAvatarSuccess"
                :before-upload="beforeAvatarUpload"
              >
                <img v-if="editForm.avatar" :src="editForm.avatar" class="avatar" />
                <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
              </el-upload>
            </el-form-item>
            <el-form-item label="昵称">
              <el-input v-model="editForm.name" />
            </el-form-item>
            <el-form-item label="个人简介">
              <el-input v-model="editForm.bio" type="textarea" :rows="3" />
            </el-form-item>
            <el-form-item label="性别">
              <el-radio-group v-model="editForm.gender">
                <el-radio :label="1">男</el-radio>
                <el-radio :label="2">女</el-radio>
                <el-radio :label="0">保密</el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form-item label="生日">
              <el-date-picker
                v-model="editForm.birthday"
                type="date"
                placeholder="选择日期"
                format="YYYY-MM-DD"
              />
            </el-form-item>
            <el-form-item label="地区">
              <el-input v-model="editForm.location" />
            </el-form-item>
            <el-form-item label="邮箱">
              <el-input v-model="editForm.email" />
            </el-form-item>
            <el-form-item label="手机">
              <el-input v-model="editForm.phone" />
            </el-form-item>
          </el-form>
          <template #footer>
            <span class="dialog-footer">
              <el-button @click="showEditDialog = false">取消</el-button>
              <el-button type="primary" @click="handleUpdateProfile">
                保存
              </el-button>
            </span>
          </template>
        </el-dialog>
      </div>
    </template>

    <template #aside>
      <transition-group name="card-fade">
        <el-card class="side-card" key="activity">
          <template #header>
            <div class="card-header">
              <span>个人动态</span>
            </div>
          </template>
          <div class="activity-list">
            <el-empty description="暂无动态" />
          </div>
        </el-card>

        <el-card class="side-card" key="visitors">
          <template #header>
            <div class="card-header">
              <span>最近访客</span>
            </div>
          </template>
          <div class="visitors-list">
            <el-empty description="暂无访客" />
          </div>
        </el-card>
      </transition-group>
    </template>
  </common-layout>
</template>

<script setup>
import CommonLayout from '@/layouts/CommonLayout.vue'
import { ref, computed, onMounted } from 'vue'
import { useRoute } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { getUserInfo, updateUserInfo } from '@/api/axiosFile'
import { ElMessage } from 'element-plus'
import { Plus, Edit } from '@element-plus/icons-vue'

const route = useRoute()
const userStore = useUserStore()

// 状态变量
const isLoading = ref(false)
const userInfo = ref(null)
const activeTab = ref('info')
const showEditDialog = ref(false)
const editForm = ref({
  avatar: '',
  name: '',
  bio: '',
  gender: 0,
  birthday: '',
  location: '',
  email: '',
  phone: ''
})

// 计算属性
const isCurrentUser = computed(() => {
  return String(localStorage.getItem('userId')) === String(route.params.id)
})

// 格式化函数
const formatGender = (gender) => {
  const genderMap = {
    0: '保密',
    1: '男',
    2: '女'
  }
  return genderMap[gender] || '保密'
}

const formatDate = (date) => {
  if (!date) return '未设置'
  return new Date(date).toLocaleDateString()
}

const formatPhone = (phone) => {
  if (!phone) return '未设置'
  return phone.replace(/(\d{3})\d{4}(\d{4})/, '$1****$2')
}

// 获取用户信息
const fetchUserInfo = async () => {
  try {
    isLoading.value = true
    console.log(route)
    const resp = await getUserInfo(route.params.id)
    if (resp.data.message) {
      userInfo.value = resp.data.data.data
      // 如果是当前用户，初始化编辑表单
      if (isCurrentUser.value) {
        initEditForm()
      }
    }
  } catch (error) {
    console.error('获取用户信息失败:', error)
    ElMessage.error('获取用户信息失败')
  } finally {
    isLoading.value = false
  }
}

// 初始化编辑表单
const initEditForm = () => {
  const info = userInfo.value
  editForm.value = {
    avatar: info.avatar,
    name: info.name,
    bio: info.bio,
    gender: info.gender,
    birthday: info.birthday,
    location: info.location,
    email: info.email,
    phone: info.phone
  }
}

// 更新用户信息
const handleUpdateProfile = async () => {
  try {
    const resp = await updateUserInfo(route.params.id, editForm.value)
    if (resp.data.message) {
      ElMessage.success('更新成功')
      showEditDialog.value = false
      await fetchUserInfo()
    }
  } catch (error) {
    console.error('更新用户信息失败:', error)
    ElMessage.error('更新用户信息失败')
  }
}

// 头像上传相关
const uploadUrl = import.meta.env.VITE_API_URL + '/upload'
const handleAvatarSuccess = (res) => {
  editForm.value.avatar = res.url
}
const beforeAvatarUpload = (file) => {
  const isImage = file.type.startsWith('image/')
  const isLt2M = file.size / 1024 / 1024 < 2

  if (!isImage) {
    ElMessage.error('只能上传图片文件!')
    return false
  }
  if (!isLt2M) {
    ElMessage.error('图片大小不能超过 2MB!')
    return false
  }
  return true
}

// 用户统计数据
const userStats = computed(() => [
  { label: '粉丝', value: userInfo.value?.followers || 0 },
  { label: '关注', value: userInfo.value?.followings || 0 }
])

// 生命周期钩子
onMounted(() => {
  fetchUserInfo()
})
</script>

<style scoped>
.profile-container {
  width: 100%;
}

.profile-header {
  margin-bottom: 20px;
  background: var(--el-bg-color);
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.profile-cover {
  background: linear-gradient(135deg, var(--el-color-primary-light-3), var(--el-color-primary));
  padding: 30px;
  color: white;
}

.profile-info {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.left-section {
  display: flex;
  align-items: center;
  gap: 24px;
}

.profile-avatar {
  border: 3px solid rgba(255, 255, 255, 0.3);
}

.user-meta {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.user-meta h1 {
  margin: 0;
  font-size: 24px;
  font-weight: 600;
}

.bio {
  margin: 0;
  font-size: 15px;
  opacity: 0.9;
}

.edit-btn {
  margin-top: 8px;
  width: fit-content;
}

.right-section {
  padding-left: 40px;
}

.user-stats {
  display: flex;
  gap: 32px;
}

.stat-item {
  text-align: center;
}

.stat-item h3 {
  margin: 0;
  font-size: 22px;
  font-weight: 600;
}

.stat-item p {
  margin: 4px 0 0;
  font-size: 14px;
}

.profile-content {
  background: var(--el-bg-color);
  border-radius: 12px;
  padding: 24px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.profile-tabs :deep(.el-tabs__header) {
  margin-bottom: 20px;
}

/* 描述列表样式 */
.el-descriptions :deep(.el-descriptions__label) {
  font-size: 15px;
  padding: 16px 20px;
}

.el-descriptions :deep(.el-descriptions__content) {
  font-size: 15px;
  padding: 16px 20px;
}

/* 编辑对话框样式 */
.edit-dialog {
  :deep(.el-dialog__body) {
    padding: 30px;
  }
}

.avatar-uploader {
  text-align: center;
  
  .avatar {
    width: 60px;
    height: 60px;
    border-radius: 50%;
  }

  .avatar-uploader-icon {
    font-size: 20px;
    color: #8c939d;
    width: 60px;
    height: 60px;
    line-height: 60px;
    text-align: center;
    border: 1px dashed var(--el-border-color);
    border-radius: 50%;
  }
}

/* 响应式调整 */
@media (max-width: 768px) {
  .profile-info {
    flex-direction: column;
    gap: 24px;
  }

  .right-section {
    padding-left: 0;
    width: 100%;
  }

  .user-stats {
    justify-content: center;
  }

  .profile-cover {
    padding: 20px;
  }

  .left-section {
    flex-direction: column;
    text-align: center;
    gap: 16px;
  }

  .user-meta {
    align-items: center;
  }

  .edit-btn {
    margin-top: 12px;
  }
}
</style> 