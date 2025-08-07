import {createRouter, createWebHistory} from 'vue-router';

const Login = () => import('src/pages/Login.vue');
const Register = () => import('src/pages/Register.vue');
const PublicPlaylist = () => import('src/pages/PublicPlaylist.vue');
// const Comments = () => import('src/pages/Comments.vue');
const ManageSongs = () => import('src/pages/ManageSongs.vue');

const Player = () => import('src/pages/Player.vue');
const ManageUsers = () => import('src/pages/ManageUsers.vue');
const MyPlaylist = () => import('src/pages/MyPlaylist.vue');
const MyLove = () => import('src/pages/MyLove.vue');
// const MyAlbum = () => import('src/pages/MyAlbum.vue');
const Uploaded = () => import('src/pages/Uploaded.vue');

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
	// {
	//   path: '/comments',
	//   name: 'Comments',
	//   component: Comments
	// },
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
		path: '/profile/:id',
		name: 'Profile',
		component: () => import('src/pages/UserProfile.vue')
		// component: () => import('src/pages/Profile.vue')
	},
	{
		path: '/player',
		name: 'Player',
		component: Player
	},
	{
		path: '/my-playlist/:id',
		name: 'MyPlaylist',
		component: MyPlaylist
	},
	{
		path: '/my-love/:id',
		name: 'MyLove',
		component: MyLove
	},
	// {
	//   path: '/my-album',
	//   name: 'MyAlbum',
	//   component: MyAlbum
	// },
	{
		path: '/uploaded',
		name: 'Uploaded',
		component: Uploaded
	},
	{
		path: '/playlist/:id',
		name: 'PlaylistDetail',
		component: () => import('src/pages/MyNewPlaylist.vue')
	},
	{
		path: '/album/:id',
		name: 'AlbumDetail',
		component: () => import('src/pages/Albums.vue')
	},
	{
		path: '/artist/:id',
		name: 'ArtistDetail',
		component: () => import('src/pages/Artists.vue')
	},
	{
		path: '/comment/:id',
		name: 'CommentDetail',
		component: () => import('src/pages/CommentOfSong.vue')
	},
	// {
	//   path: '/test/:id',
	//   name: 'Test',
	//   component: () => import('src/pages/test.vue')
	// }
];

const router = createRouter({
	history: createWebHistory('/'),
	routes
});

export default router;
