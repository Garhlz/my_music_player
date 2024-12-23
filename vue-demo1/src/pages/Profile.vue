<template>
  <CommonLayout :page-name="currentName">
    <template #main>
      <div class="profile-container">
        <!-- 个人信息卡片 -->
        <div class="profile-header">
          <div class="profile-cover"></div>
          <div class="profile-info">
            <el-avatar 
              :src="userInfo.avatar || '/assets/default-avatar.jpg'" 
              :size="100" 
              class="profile-avatar"
            />
            <div class="profile-details">
              <h1 class="username">{{ userInfo.name }}</h1>
              <p class="user-bio">{{ userInfo.bio || '这个人很懒，什么都没写~' }}</p>
              <div class="user-stats">
                <div class="stat-item">
                  <span class="stat-value">{{ userInfo.followings || 0 }}</span>
                  <span class="stat-label">关注</span>
                </div>
                <div class="stat-item">
                  <span class="stat-value">{{ userInfo.followers || 0 }}</span>
                  <span class="stat-label">粉丝</span>
                </div>
                <div class="stat-item">
                  <span class="stat-value">{{ userInfo.likes || 0 }}</span>
                  <span class="stat-label">获赞</span>
                </div>
              </div>
            </div>
            <el-button type="primary" @click="showEditDialog">编辑资料</el-button>
          </div>
        </div>

        <!-- 主要内容区域 -->
        <div class="profile-content">
          <el-tabs v-model="activeTab" class="profile-tabs">
            <!-- 基本信息标签页 -->
            <el-tab-pane label="基本信息" name="info">
              <div class="info-list">
                <div class="info-item">
                  <span class="info-label">用户名</span>
                  <span class="info-value">{{ userInfo.username }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">昵称</span>
                  <span class="info-value">{{ userInfo.name }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">邮箱</span>
                  <span class="info-value">{{ userInfo.email }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">手机</span>
                  <span class="info-value">{{ userInfo.phone || '未设置' }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">性别</span>
                  <span class="info-value">{{ userInfo.gender || '未设置' }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">生日</span>
                  <span class="info-value">{{ userInfo.birthday || '未设置' }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">地区</span>
                  <span class="info-value">{{ userInfo.location || '未设置' }}</span>
                </div>
                <div class="info-item">
                  <span class="info-label">注册时间</span>
                  <span class="info-value">{{ formatDate(userInfo.created_at) }}</span>
                </div>
              </div>
            </el-tab-pane>

            <!-- 我的音乐标签页 -->
            <el-tab-pane label="我的音乐" name="music">
              <div class="music-stats">
                <el-card class="stat-card">
                  <template #header>
                    <div class="card-header">
                      <span>收藏的音乐</span>
                      <el-button text>查看全部</el-button>
                    </div>
                  </template>
                  <div class="stat-number">{{ userInfo.likedSongs || 0 }}</div>
                </el-card>

                <el-card class="stat-card">
                  <template #header>
                    <div class="card-header">
                      <span>创建的歌单</span>
                      <el-button text>查看全部</el-button>
                    </div>
                  </template>
                  <div class="stat-number">{{ userInfo.playlists || 0 }}</div>
                </el-card>

                <el-card class="stat-card">
                  <template #header>
                    <div class="card-header">
                      <span>上传的音乐</span>
                      <el-button text>查看全部</el-button>
                    </div>
                  </template>
                  <div class="stat-number">{{ userInfo.uploadedSongs || 0 }}</div>
                </el-card>
              </div>
            </el-tab-pane>

            <!-- 修改密码标签页 -->
            <el-tab-pane label="修改密码" name="password">
              <el-form 
                ref="passwordForm"
                :model="passwordForm"
                :rules="passwordRules"
                label-width="100px"
                class="password-form"
              >
                <el-form-item label="原密码" prop="oldPassword">
                  <el-input 
                    v-model="passwordForm.oldPassword"
                    type="password"
                    show-password
                  />
                </el-form-item>
                <el-form-item label="新密码" prop="newPassword">
                  <el-input 
                    v-model="passwordForm.newPassword"
                    type="password"
                    show-password
                  />
                </el-form-item>
                <el-form-item label="确认密码" prop="confirmPassword">
                  <el-input 
                    v-model="passwordForm.confirmPassword"
                    type="password"
                    show-password
                  />
                </el-form-item>
                <el-form-item>
                  <el-button type="primary" @click="updatePassword">修改密码</el-button>
                </el-form-item>
              </el-form>
            </el-tab-pane>

            <!-- 账号设置标签页 -->
            <el-tab-pane label="账号设置" name="settings">
              <div class="settings-section">
                <h3>隐私设置</h3>
                <el-form :model="privacySettings" label-width="120px">
                  <el-form-item label="个人主页可见">
                    <el-switch v-model="privacySettings.profileVisible" />
                  </el-form-item>
                  <el-form-item label="播放记录可见">
                    <el-switch v-model="privacySettings.historyVisible" />
                  </el-form-item>
                </el-form>
              </div>
            </el-tab-pane>
          </el-tabs>
        </div>

        <!-- 编辑资料对话框 -->
        <el-dialog
          v-model="editDialogVisible"
          title="编辑个人资料"
          width="500px"
        >
          <el-form :model="editForm" :rules="editRules" ref="editForm" label-width="80px">
            <el-form-item label="头像">
              <el-upload
                class="avatar-uploader"
                action="/api/upload"
                :show-file-list="false"
                :on-success="handleAvatarSuccess"
              >
                <img v-if="editForm.avatar" :src="editForm.avatar" class="avatar" />
                <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
              </el-upload>
            </el-form-item>
            <el-form-item label="昵称" prop="name">
              <el-input v-model="editForm.name" />
            </el-form-item>
            <el-form-item label="邮箱" prop="email">
              <el-input v-model="editForm.email" />
            </el-form-item>
            <el-form-item label="手机" prop="phone">
              <el-input v-model="editForm.phone" />
            </el-form-item>
            <el-form-item label="性别">
              <el-radio-group v-model="editForm.gender">
                <el-radio label="male">男</el-radio>
                <el-radio label="female">女</el-radio>
                <el-radio label="other">其他</el-radio>
              </el-radio-group>
            </el-form-item>
            <el-form-item label="生日">
              <el-date-picker
                v-model="editForm.birthday"
                type="date"
                placeholder="选择生日"
              />
            </el-form-item>
            <el-form-item label="地区">
              <el-input v-model="editForm.location" />
            </el-form-item>
            <el-form-item label="个人简介">
              <el-input
                v-model="editForm.bio"
                type="textarea"
                :rows="3"
                placeholder="介绍一下自己吧"
              />
            </el-form-item>
          </el-form>
          <template #footer>
            <el-button @click="editDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="submitEdit">保存</el-button>
          </template>
        </el-dialog>
      </div>
    </template>
  </CommonLayout>
</template>

<script>
import { ref } from 'vue'
import CommonLayout from '@/layouts/CommonLayout.vue'
import { ElMessage } from 'element-plus'
import axios from 'axios'
import { Plus } from '@element-plus/icons-vue'

export default {
  name: 'Profile',
  components: {
    CommonLayout,
    Plus
  },
  data() {
    // 密码验证规则
    const validateConfirmPassword = (rule, value, callback) => {
      if (value !== this.passwordForm.newPassword) {
        callback(new Error('两次输入的密码不一致'))
      } else {
        callback()
      }
    }

    return {
      currentName: '个人中心',
      activeTab: 'info',
      editDialogVisible: false,
      userInfo: {},
      editForm: {
        avatar: '',
        name: '',
        email: '',
        phone: '',
        gender: '',
        birthday: '',
        location: '',
        bio: ''
      },
      passwordForm: {
        oldPassword: '',
        newPassword: '',
        confirmPassword: ''
      },
      privacySettings: {
        profileVisible: true,
        historyVisible: true
      },
      editRules: {
        name: [
          { required: true, message: '请输入昵称', trigger: 'blur' },
          { min: 2, max: 20, message: '长度在 2 到 20 个字符', trigger: 'blur' }
        ],
        email: [
          { required: true, message: '请输入邮箱地址', trigger: 'blur' },
          { type: 'email', message: '请输入正确的邮箱地址', trigger: 'blur' }
        ],
        phone: [
          { pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号码', trigger: 'blur' }
        ]
      },
      passwordRules: {
        oldPassword: [
          { required: true, message: '请输入原密码', trigger: 'blur' },
          { min: 6, message: '密码长度不能小于6位', trigger: 'blur' }
        ],
        newPassword: [
          { required: true, message: '请输入新密码', trigger: 'blur' },
          { min: 6, message: '密码长度不能小于6位', trigger: 'blur' }
        ],
        confirmPassword: [
          { required: true, message: '请再次输入新密码', trigger: 'blur' },
          { validator: validateConfirmPassword, trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.getUserInfo()
  },
  methods: {
    formatDate(date) {
      return new Date(date).toLocaleDateString()
    },

    // 获取用户信息
    async getUserInfo() {
      try {
        const response = await axios.get('/api/user/info')
        this.userInfo = response.data
        // 初始化编辑表单
        this.editForm = {
          avatar: this.userInfo.avatar,
          name: this.userInfo.name,
          email: this.userInfo.email,
          phone: this.userInfo.phone,
          gender: this.userInfo.gender,
          birthday: this.userInfo.birthday,
          location: this.userInfo.location,
          bio: this.userInfo.bio
        }
      } catch (error) {
        ElMessage.error('获取用户信息失败')
      }
    },

    showEditDialog() {
      this.editDialogVisible = true
    },

    handleAvatarSuccess(res) {
      this.editForm.avatar = res.url
    },

    // 提交编辑表单
    async submitEdit() {
      try {
        await this.$refs.editForm.validate()
        await axios.put('/api/user/info', this.editForm)
        ElMessage.success('个人资料已更新')
        this.editDialogVisible = false
        await this.getUserInfo()
      } catch (error) {
        if (error.response) {
          ElMessage.error(error.response.data.message || '更新失败')
        }
      }
    },

    // 更新密码
    async updatePassword() {
      try {
        await this.$refs.passwordForm.validate()
        await axios.put('/api/user/password', {
          oldPassword: this.passwordForm.oldPassword,
          newPassword: this.passwordForm.newPassword
        })
        ElMessage.success('密码修改成功')
        this.passwordForm = {
          oldPassword: '',
          newPassword: '',
          confirmPassword: ''
        }
      } catch (error) {
        if (error.response) {
          ElMessage.error(error.response.data.message || '修改密码失败')
        }
      }
    }
  }
}
</script>

<style scoped>
.profile-container {
  padding: 20px;
}

.profile-header {
  background: #fff;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
  margin-bottom: 20px;
}

.profile-cover {
  height: 200px;
  background: linear-gradient(135deg, #1890ff, #722ed1);
}

.profile-info {
  padding: 20px;
  margin-top: -50px;
  position: relative;
  display: flex;
  align-items: flex-end;
}

.profile-avatar {
  border: 4px solid #fff;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.profile-details {
  margin-left: 20px;
}

.profile-content {
  background: #fff;
  border-radius: 12px;
  padding: 20px;
  box-shadow: 0 2px 12px rgba(0, 0, 0, 0.1);
}

.user-form, .password-form {
  max-width: 500px;
  margin: 20px auto;
  padding: 20px;
}

.username {
  margin: 0;
  font-size: 24px;
  font-weight: 600;
}

.user-bio {
  margin: 8px 0;
  color: #666;
}

.info-list {
  margin-top: 20px;
}

.info-item {
  display: flex;
  justify-content: space-between;
  padding: 10px 0;
  border-bottom: 1px solid #eee;
}

.info-label {
  font-weight: 600;
}

.info-value {
  margin-left: 20px;
}

.music-stats {
  display: flex;
  justify-content: space-around;
  margin-top: 20px;
}

.stat-card {
  width: 30%;
  text-align: center;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}

.settings-section {
  margin-top: 20px;
}

.avatar-uploader {
  text-align: center;
}

.avatar-uploader .el-upload {
  border: 1px dashed #d9d9d9;
  border-radius: 6px;
  cursor: pointer;
  position: relative;
  overflow: hidden;
}

.avatar-uploader .el-upload:hover {
  border-color: #409eff;
}

.avatar-uploader-icon {
  font-size: 28px;
  color: #8c939d;
  width: 100%;
  height: 100%;
  text-align: center;
}

.avatar {
  width: 100px;
  height: 100px;
  display: block;
}
</style>
  