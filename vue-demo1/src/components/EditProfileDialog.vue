<template>
  <el-dialog
    :model-value="visible"
    title="编辑个人资料"
    width="500px"
    @update:modelValue="emit('update:visible', $event)"
    @close="handleClose"
    class="edit-profile-dialog"
    center
    append-to-body
  >
    <el-form :model="form" ref="formRef" label-position="top" class="profile-form">
      <el-form-item label="头像">
        <el-upload
          class="avatar-uploader"
          action="/api/upload"
          :show-file-list="false"
          :on-success="handleAvatarSuccess"
          :before-upload="beforeAvatarUpload"
        >
          <el-avatar :size="80" :src="form.avatar" class="uploaded-avatar" />
          <div class="upload-overlay">
            <el-icon>
              <EditPen />
            </el-icon>
          </div>
        </el-upload>
      </el-form-item>
      <el-form-item label="昵称" prop="name">
        <el-input v-model="form.name" placeholder="你的公开显示名称" />
      </el-form-item>
      <el-form-item label="个人简介" prop="bio">
        <el-input type="textarea" v-model="form.bio" :rows="3" placeholder="介绍一下自己..." />
      </el-form-item>

      <div class="form-grid">
        <el-form-item label="性别" prop="gender">
          <el-radio-group v-model="form.gender">
            <el-radio :label="1">男</el-radio>
            <el-radio :label="2">女</el-radio>
            <el-radio :label="0">保密</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="生日" prop="birthday">
          <el-date-picker v-model="form.birthday" type="date" placeholder="选择你的生日" style="width: 100%;"
                          format="YYYY-MM-DD" value-format="YYYY-MM-DDTHH:mm:ssZ" />
        </el-form-item>
      </div>
      <el-form-item label="所在地" prop="location">
        <el-input v-model="form.location" placeholder="例如：深圳" />
      </el-form-item>
      <el-form-item label="电子邮箱" prop="email">
        <el-input v-model="form.email" placeholder="你的电子邮箱地址" />
      </el-form-item>
      <el-form-item label="手机号" prop="phone">
        <el-input v-model="form.phone" placeholder="你的手机号码" />
      </el-form-item>

    </el-form>
    <template #footer>
      <div class="dialog-footer">
        <el-button @click="emit('update:visible', false)" class="cancel-btn">取消</el-button>
        <el-button type="primary" :loading="isSubmitting" @click="handleSubmit" class="submit-btn">
          保存
        </el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, watch, reactive } from 'vue';
import type { PropType } from 'vue';
import { ElMessage, type UploadProps, type FormInstance } from 'element-plus';
import type { ModelsUpdateUserRequest } from '@/api-client';
import { EditPen } from '@element-plus/icons-vue';

const props = defineProps({
  visible: { type: Boolean, required: true },
  initialData: { type: Object as PropType<ModelsUpdateUserRequest | null>, default: null },
});

const emit = defineEmits<{
  (e: 'update:visible', value: boolean): void;
  (e: 'success', data: ModelsUpdateUserRequest): void;
}>();

const formRef = ref<FormInstance>();
const isSubmitting = ref(false);

// [THE FIX] 扩展表单数据结构
const form = reactive<ModelsUpdateUserRequest>({
  name: '',
  bio: '',
  avatar: '',
  gender: 0,
  birthday: '',
  location: '',
  email: '',
  phone: '',
});

watch(() => props.visible, (newVal) => {
  if (newVal && props.initialData) {
    Object.assign(form, props.initialData);
  }
});

const handleClose = () => formRef.value?.resetFields();
const handleSubmit = async () => {
  isSubmitting.value = true;
  try {
    emit('success', { ...form });
  } finally {
    isSubmitting.value = false;
  }
};
const handleAvatarSuccess: UploadProps['onSuccess'] = (response) => {
  form.avatar = response.url;
  ElMessage.success('头像上传成功');
};
const beforeAvatarUpload: UploadProps['beforeUpload'] = (rawFile) => {
  if (!rawFile.type.startsWith('image/')) {
    ElMessage.error('头像只能是图片格式!');
    return false;
  }
  if (rawFile.size / 1024 / 1024 > 2) {
    ElMessage.error('头像图片大小不能超过 2MB!');
    return false;
  }
  return true;
};
</script>

<style scoped>
/* ... 之前的美化样式保持不变 ... */
.edit-profile-dialog :deep(.el-dialog) { /* ... */
}

/* [THE FIX] 新增 grid 布局用于排列性别和生日 */
.form-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px;
}
</style>