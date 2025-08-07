import {defineConfig} from 'vite';
import vue from '@vitejs/plugin-vue';
import {fileURLToPath, URL} from 'node:url'; // 引入这两个模块

// https://vitejs.dev/config/
export default defineConfig({
	plugins: [
		vue(),
	],
	resolve: {
		alias: {
			// 配置 'src' 别名，使其指向项目根目录下的 'src' 文件夹
			'src': fileURLToPath(new URL('./src', import.meta.url)),

			// 另外，如果你还想保留或使用 '@' 别名 (Vue CLI 默认的别名)，也可以这样配置
			'@': fileURLToPath(new URL('./src', import.meta.url)),
		},
	},
});