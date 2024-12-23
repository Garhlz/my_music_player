import { createRouter, createWebHistory } from 'vue-router';

const Login = () => import('@/pages/Login.vue');
const Register = () => import('@/pages/Register.vue');
const PublicPlaylist = () => import('@/pages/PublicPlaylist.vue');
const Comments = () => import('@/pages/Comments.vue');
const ManageSongs = () => import('@/pages/ManageSongs.vue');
const Profile = () => import('@/pages/Profile.vue');
const Player = () => import('@/pages/Player.vue');
const ManageUsers = () => import('@/pages/ManageUsers.vue');
const MyPlaylist = () => import('@/pages/MyPlaylist.vue');
const MyLove = () => import('@/pages/MyLove.vue');
const MyAlbum = () => import('@/pages/MyAlbum.vue');
const Uploaded = () => import('@/pages/Uploaded.vue');
const routes = [
  {
    path: '/',
    name: 'Login1',
    component: Login
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/register',
    name: 'Register',
    component: Register
  },
  {
    path: '/public-playlist',
    name: 'PublicPlaylist',
    component: PublicPlaylist
  },
  {
    path: '/pub',
    name: 'PublicPlaylist1',
    component: PublicPlaylist
  },
  {
    path: '/comments',
    name: 'Comments',
    component: Comments
  },
  {
    path: '/manage-songs',
    name: 'ManageSongs',
    component: ManageSongs
  },
  {
    path: '/manage-users',
    name: 'ManageUsers',
    component: ManageUsers
  },
  {
    path: '/profile',
    name: 'Profile',
    component: Profile
  },
  {
    path: '/player',
    name: 'Player',
    component: Player
  },
  {
    path: '/my-playlist',
    name: 'MyPlaylist',
    component: MyPlaylist
  },
  {
    path: '/my-love',
    name: 'MyLove',
    component: MyLove
  },
  {
    path: '/my-album',
    name: 'MyAlbum',
    component: MyAlbum
  },
  {
    path: '/uploaded',
    name: 'Uploaded',
    component: Uploaded
  },
  {
    path: '/playlist/:id',
    name: 'PlaylistDetail',
    component: () => import('@/pages/MyNewPlaylist.vue')
  },
];

const router = createRouter({
  history: createWebHistory('/'),
  routes
});

export default router;
