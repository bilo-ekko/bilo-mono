export const manifest = (() => {
function __memo(fn) {
	let value;
	return () => value ??= (value = fn());
}

return {
	appDir: "_app",
	appPath: "_app",
	assets: new Set(["favicon.png"]),
	mimeTypes: {".png":"image/png"},
	_: {
		client: {start:"_app/immutable/entry/start.BDRZr22f.js",app:"_app/immutable/entry/app.DXrtIdRR.js",imports:["_app/immutable/entry/start.BDRZr22f.js","_app/immutable/chunks/D6pdwf9A.js","_app/immutable/chunks/BjgrqnN-.js","_app/immutable/chunks/KthFeat2.js","_app/immutable/entry/app.DXrtIdRR.js","_app/immutable/chunks/BjgrqnN-.js","_app/immutable/chunks/KthFeat2.js","_app/immutable/chunks/CV4juGgb.js","_app/immutable/chunks/CEeEBKyQ.js","_app/immutable/chunks/BKYosCqX.js","_app/immutable/chunks/BqpafKht.js","_app/immutable/chunks/BukBgMWT.js","_app/immutable/chunks/Cc98ydwC.js"],stylesheets:[],fonts:[],uses_env_dynamic_public:false},
		nodes: [
			__memo(() => import('./nodes/0.js')),
			__memo(() => import('./nodes/1.js')),
			__memo(() => import('./nodes/2.js')),
			__memo(() => import('./nodes/3.js')),
			__memo(() => import('./nodes/4.js'))
		],
		remotes: {
			
		},
		routes: [
			{
				id: "/",
				pattern: /^\/$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 2 },
				endpoint: null
			},
			{
				id: "/checkout",
				pattern: /^\/checkout\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 3 },
				endpoint: null
			},
			{
				id: "/post-purchase",
				pattern: /^\/post-purchase\/?$/,
				params: [],
				page: { layouts: [0,], errors: [1,], leaf: 4 },
				endpoint: null
			}
		],
		prerendered_routes: new Set([]),
		matchers: async () => {
			
			return {  };
		},
		server_assets: {}
	}
}
})();
