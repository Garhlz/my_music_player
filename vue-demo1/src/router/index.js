import { createRouter, createWebHistory } from 'vue-router';

// Auth 组件现在是一个包含登录和注册逻辑的单页组件
const AuthPage = () => import('@/pages/Auth.vue');

// 其他页面组件统一使用 @/ 别名
const PublicPlaylist = () => import('@/pages/PublicPlaylist.vue');
const ManageSongs = () => import('@/pages/ManageSongs.vue');
const Player = () => import('@/pages/Player.vue');
const ManageUsers = () => import('@/pages/ManageUsers.vue');
const MyPlaylist = () => import('@/pages/MyPlaylist.vue');
const MyLove = () => import('@/pages/MyLove.vue');
const Uploaded = () => import('@/pages/Uploaded.vue');
const UserProfile = () => import('@/pages/UserProfile.vue');
const PlaylistDetail = () => import('@/pages/MyNewPlaylist.vue');
const AlbumDetail = () => import('@/pages/Albums.vue');
const ArtistDetail = () => import('@/pages/Artists.vue');
const CommentDetail = () => import('@/pages/CommentOfSong.vue');


const routes = [
  {
    path: '/',
    // 默认重定向到 /auth 即可
    redirect: '/auth',
  },
  {
    path: '/auth',
    name: 'Auth',
    component: AuthPage, // AuthPage (Auth.vue) 是一个包含所有逻辑的单页
    // 这里不再需要 children 数组，因为登录和注册逻辑在 Auth.vue 内部处理
  },
  {
    path: '/pub', // 建议统一别名，或者如果不再需要可以删除此路由
    name: 'PublicPlaylistAlias',
    component: PublicPlaylist,
  },

  {
    path: '/manage-songs',
    name: 'ManageSongs',
    component: ManageSongs,
  },
  {
    path: '/manage-users',
    name: 'ManageUsers',
    component: ManageUsers,
  },
  {
    path: '/profile/:id',
    name: 'Profile',
    component: UserProfile,
  },
  {
    path: '/player',
    name: 'Player',
    component: Player,
  },
  {
    path: '/my-playlist/:id',
    name: 'MyPlaylist',
    component: MyPlaylist,
  },

  {
    path: '/uploaded',
    name: 'Uploaded',
    component: Uploaded,
  },
  {
    path: '/playlist/:id',
    name: 'PlaylistDetail',
    component: PlaylistDetail,
  },
  {
    path: '/album/:id',
    name: 'AlbumDetail',
    component: AlbumDetail,
  },
  {
    path: '/artist/:id',
    name: 'ArtistDetail',
    component: ArtistDetail,
  },
  {
    path: '/comment/:id',
    name: 'CommentDetail',
    component: CommentDetail,
  },
];

const router = createRouter({
  history: createWebHistory('/'),
  routes,
});

export default router;