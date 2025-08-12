<template>
  <el-dialog
    :model-value="visible"
    :title="isEditMode ? '编辑歌单' : '创建新歌单'"
    width="450px"
    @update:modelValue="emit('update:visible', $event)"
    @close="handleClose"
    :close-on-click-modal="false"
  >
    <el-form :model="form" ref="formRef" label-position="top">
      <el-form-item label="歌单名称" prop="name"
                    :rules="[{ required: true, message: '歌单名称不能为空', trigger: 'blur' }]">
        <el-input v-model="form.name" placeholder="请输入歌单名称" maxlength="50" show-word-limit />
      </el-form-item>
      <el-form-item label="歌单描述" prop="description">
        <el-input v-model="form.description" type="textarea" :rows="4" placeholder="添加一些描述吧" maxlength="200"
                  show-word-limit />
      </el-form-item>
      <el-form-item label="是否公开" prop="is_public">
        <el-switch v-model="form.is_public" />
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="emit('update:visible', false)">取消</el-button>
        <el-button type="primary" :loading="isSubmitting" @click="handleSubmit">
          {{ isEditMode ? '保存' : '创建' }}
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>

<script setup lang="ts">
import { ref, reactive, watch, computed } from 'vue';
import type { PropType } from 'vue';
import { ElMessage, type FormInstance } from 'element-plus';
import { playlistApi } from '@/api';
import type {
  ModelsPlaylistInfoDTO,
  ModelsCreatePlaylistRequestDTO,
  ModelsUpdatePlaylistRequestDTO,
} from '@/api-client';

// --- Props ---
const props = defineProps({
  visible: {
    type: Boolean,
    required: true,
  },
  mode: {
    type: String as PropType<'create' | 'edit'>,
    required: true,
  },
  // 在编辑模式下，需要传入初始数据
  initialData: {
    type: Object as PropType<ModelsPlaylistInfoDTO | null>,
    default: null,
  },
});

// --- Emits ---
const emit = defineEmits<{
  (e: 'update:visible', value: boolean): void;
  (e: 'success'): void; // 操作成功后触发，通知父组件刷新列表
}>();

// --- State ---
const formRef = ref<FormInstance>();
const isSubmitting = ref(false);
const form = reactive<ModelsCreatePlaylistRequestDTO | ModelsUpdatePlaylistRequestDTO>({
  name: '',
  description: '',
  is_public: false,
});

// --- Computed ---
const isEditMode = computed(() => props.mode === 'edit');

// --- Watcher ---
// 监听弹窗的打开，如果是编辑模式，则填充表单
watch(() => props.visible, (newVal) => {
  if (newVal && isEditMode.value && props.initialData) {
    form.name = props.initialData.name;
    form.description = props.initialData.description || '';
    form.is_public = props.initialData.is_public || false;
  }
});

// --- Methods ---
const resetForm = () => {
  formRef.value?.resetFields();
  form.name = '';
  form.description = '';
  form.is_public = false;
};

const handleClose = () => {
  // 延迟重置表单，避免在关闭动画中看到内容变化
  setTimeout(resetForm, 200);
};

const handleSubmit = async () => {
  if (!formRef.value) return;
  await formRef.value.validate(async (valid) => {
    if (valid) {
      isSubmitting.value = true;
      try {
        if (isEditMode.value) {
          // 编辑模式
          if (!props.initialData?.id) throw new Error('缺少要编辑的歌单ID');
          await playlistApi.mePlaylistsPlaylistIdPut(props.initialData.id, form as ModelsUpdatePlaylistRequestDTO);
          ElMessage.success('歌单更新成功');
        } else {
          // 创建模式
          await playlistApi.mePlaylistsPost(form as ModelsCreatePlaylistRequestDTO);
          ElMessage.success('歌单创建成功');
        }
        emit('success'); // 通知父组件操作成功
        emit('update:visible', false); // 关闭弹窗
      } catch (error) {
        console.error('操作失败:', error);
        ElMessage.error('操作失败，请稍后重试');
      } finally {
        isSubmitting.value = false;
      }
    }
  });
};
</script>

<style scoped>
.dialog-footer {
  display: flex;
  justify-content: flex-end;
}
</style>