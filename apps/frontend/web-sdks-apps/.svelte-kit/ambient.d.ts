
// this file is generated — do not edit it


/// <reference types="@sveltejs/kit" />

/**
 * Environment variables [loaded by Vite](https://vitejs.dev/guide/env-and-mode.html#env-files) from `.env` files and `process.env`. Like [`$env/dynamic/private`](https://svelte.dev/docs/kit/$env-dynamic-private), this module cannot be imported into client-side code. This module only includes variables that _do not_ begin with [`config.kit.env.publicPrefix`](https://svelte.dev/docs/kit/configuration#env) _and do_ start with [`config.kit.env.privatePrefix`](https://svelte.dev/docs/kit/configuration#env) (if configured).
 * 
 * _Unlike_ [`$env/dynamic/private`](https://svelte.dev/docs/kit/$env-dynamic-private), the values exported from this module are statically injected into your bundle at build time, enabling optimisations like dead code elimination.
 * 
 * ```ts
 * import { API_KEY } from '$env/static/private';
 * ```
 * 
 * Note that all environment variables referenced in your code should be declared (for example in an `.env` file), even if they don't have a value until the app is deployed:
 * 
 * ```
 * MY_FEATURE_FLAG=""
 * ```
 * 
 * You can override `.env` values from the command line like so:
 * 
 * ```sh
 * MY_FEATURE_FLAG="enabled" npm run dev
 * ```
 */
declare module '$env/static/private' {
	export const PROTO_APP_LOG: string;
	export const MOON_PROJECT_SOURCE: string;
	export const PROTO_MOON_VERSION: string;
	export const TERM_PROGRAM: string;
	export const VSCODE_GIT_IPC_AUTH_TOKEN: string;
	export const NODE: string;
	export const _P9K_TTY: string;
	export const INIT_CWD: string;
	export const SHELL: string;
	export const TERM: string;
	export const CLICOLOR: string;
	export const STARBASE_FORCE_TTY: string;
	export const PROTO_HOME: string;
	export const TMPDIR: string;
	export const PROTO_NO_PROGRESS: string;
	export const TERM_PROGRAM_VERSION: string;
	export const CURSOR_TRACE_ID: string;
	export const MOON_CACHE_DIR: string;
	export const MallocNanoZone: string;
	export const ORIGINAL_XDG_CURRENT_DESKTOP: string;
	export const PROTO_SHIM_NAME: string;
	export const STARBASE_LOG: string;
	export const MOON_TASK_RETRY_TOTAL: string;
	export const npm_config_registry: string;
	export const ZSH: string;
	export const PROTO_IGNORE_MIGRATE_WARNING: string;
	export const PROTO_NODE_VERSION: string;
	export const USER: string;
	export const LS_COLORS: string;
	export const COMMAND_MODE: string;
	export const MOON_PROJECT_ROOT: string;
	export const MOON_WORKSPACE_ROOT: string;
	export const PNPM_SCRIPT_SRC_DIR: string;
	export const CLAUDE_CODE_SSE_PORT: string;
	export const SSH_AUTH_SOCK: string;
	export const PROTO_LOOKUP_DIR: string;
	export const PROTO_VERSION: string;
	export const __CF_USER_TEXT_ENCODING: string;
	export const MOON_CACHE: string;
	export const npm_execpath: string;
	export const COLUMNS: string;
	export const PAGER: string;
	export const PROTO_MOON_DETECTED_FROM: string;
	export const LSCOLORS: string;
	export const npm_config_frozen_lockfile: string;
	export const npm_config_verify_deps_before_run: string;
	export const PATH: string;
	export const LaunchInstanceID: string;
	export const PROTO_PNPM_VERSION: string;
	export const npm_package_json: string;
	export const MOON_LOG: string;
	export const __CFBundleIdentifier: string;
	export const PWD: string;
	export const npm_command: string;
	export const PROTO_GO_VERSION: string;
	export const VSCODE_NONCE: string;
	export const P9K_SSH: string;
	export const npm_config__jsr_registry: string;
	export const npm_lifecycle_event: string;
	export const LANG: string;
	export const P9K_TTY: string;
	export const PROTO_CLI_VERSION: string;
	export const npm_package_name: string;
	export const NODE_PATH: string;
	export const VSCODE_GIT_ASKPASS_EXTRA_ARGS: string;
	export const XPC_FLAGS: string;
	export const MOON_TASK_HASH: string;
	export const PROTO_AUTO_INSTALL: string;
	export const FORCE_COLOR: string;
	export const LINES: string;
	export const MOON_TARGET: string;
	export const MOON_THEME: string;
	export const npm_config_node_gyp: string;
	export const XPC_SERVICE_NAME: string;
	export const npm_package_version: string;
	export const pnpm_config_verify_deps_before_run: string;
	export const HOME: string;
	export const SHLVL: string;
	export const STARBASE_THEME: string;
	export const VSCODE_GIT_ASKPASS_MAIN: string;
	export const PROTO_SHIM_PATH: string;
	export const MOON_VERSION: string;
	export const LESS: string;
	export const LOGNAME: string;
	export const MOON_WORKING_DIR: string;
	export const MOON_TASK_RETRY_ATTEMPT: string;
	export const npm_lifecycle_script: string;
	export const VSCODE_GIT_IPC_HANDLE: string;
	export const BUN_INSTALL: string;
	export const CLICOLOR_FORCE: string;
	export const MOON_TASK_ID: string;
	export const MOON_PROJECT_SNAPSHOT: string;
	export const npm_config_user_agent: string;
	export const GIT_ASKPASS: string;
	export const VSCODE_GIT_ASKPASS_NODE: string;
	export const PROTO_VERSION_CHECK: string;
	export const _P9K_SSH_TTY: string;
	export const MOON_PROJECT_ID: string;
	export const SECURITYSESSIONID: string;
	export const PROTO_OFFLINE_TIMEOUT: string;
	export const COLORTERM: string;
	export const npm_node_execpath: string;
	export const NODE_ENV: string;
}

/**
 * Similar to [`$env/static/private`](https://svelte.dev/docs/kit/$env-static-private), except that it only includes environment variables that begin with [`config.kit.env.publicPrefix`](https://svelte.dev/docs/kit/configuration#env) (which defaults to `PUBLIC_`), and can therefore safely be exposed to client-side code.
 * 
 * Values are replaced statically at build time.
 * 
 * ```ts
 * import { PUBLIC_BASE_URL } from '$env/static/public';
 * ```
 */
declare module '$env/static/public' {
	
}

/**
 * This module provides access to runtime environment variables, as defined by the platform you're running on. For example if you're using [`adapter-node`](https://github.com/sveltejs/kit/tree/main/packages/adapter-node) (or running [`vite preview`](https://svelte.dev/docs/kit/cli)), this is equivalent to `process.env`. This module only includes variables that _do not_ begin with [`config.kit.env.publicPrefix`](https://svelte.dev/docs/kit/configuration#env) _and do_ start with [`config.kit.env.privatePrefix`](https://svelte.dev/docs/kit/configuration#env) (if configured).
 * 
 * This module cannot be imported into client-side code.
 * 
 * ```ts
 * import { env } from '$env/dynamic/private';
 * console.log(env.DEPLOYMENT_SPECIFIC_VARIABLE);
 * ```
 * 
 * > [!NOTE] In `dev`, `$env/dynamic` always includes environment variables from `.env`. In `prod`, this behavior will depend on your adapter.
 */
declare module '$env/dynamic/private' {
	export const env: {
		PROTO_APP_LOG: string;
		MOON_PROJECT_SOURCE: string;
		PROTO_MOON_VERSION: string;
		TERM_PROGRAM: string;
		VSCODE_GIT_IPC_AUTH_TOKEN: string;
		NODE: string;
		_P9K_TTY: string;
		INIT_CWD: string;
		SHELL: string;
		TERM: string;
		CLICOLOR: string;
		STARBASE_FORCE_TTY: string;
		PROTO_HOME: string;
		TMPDIR: string;
		PROTO_NO_PROGRESS: string;
		TERM_PROGRAM_VERSION: string;
		CURSOR_TRACE_ID: string;
		MOON_CACHE_DIR: string;
		MallocNanoZone: string;
		ORIGINAL_XDG_CURRENT_DESKTOP: string;
		PROTO_SHIM_NAME: string;
		STARBASE_LOG: string;
		MOON_TASK_RETRY_TOTAL: string;
		npm_config_registry: string;
		ZSH: string;
		PROTO_IGNORE_MIGRATE_WARNING: string;
		PROTO_NODE_VERSION: string;
		USER: string;
		LS_COLORS: string;
		COMMAND_MODE: string;
		MOON_PROJECT_ROOT: string;
		MOON_WORKSPACE_ROOT: string;
		PNPM_SCRIPT_SRC_DIR: string;
		CLAUDE_CODE_SSE_PORT: string;
		SSH_AUTH_SOCK: string;
		PROTO_LOOKUP_DIR: string;
		PROTO_VERSION: string;
		__CF_USER_TEXT_ENCODING: string;
		MOON_CACHE: string;
		npm_execpath: string;
		COLUMNS: string;
		PAGER: string;
		PROTO_MOON_DETECTED_FROM: string;
		LSCOLORS: string;
		npm_config_frozen_lockfile: string;
		npm_config_verify_deps_before_run: string;
		PATH: string;
		LaunchInstanceID: string;
		PROTO_PNPM_VERSION: string;
		npm_package_json: string;
		MOON_LOG: string;
		__CFBundleIdentifier: string;
		PWD: string;
		npm_command: string;
		PROTO_GO_VERSION: string;
		VSCODE_NONCE: string;
		P9K_SSH: string;
		npm_config__jsr_registry: string;
		npm_lifecycle_event: string;
		LANG: string;
		P9K_TTY: string;
		PROTO_CLI_VERSION: string;
		npm_package_name: string;
		NODE_PATH: string;
		VSCODE_GIT_ASKPASS_EXTRA_ARGS: string;
		XPC_FLAGS: string;
		MOON_TASK_HASH: string;
		PROTO_AUTO_INSTALL: string;
		FORCE_COLOR: string;
		LINES: string;
		MOON_TARGET: string;
		MOON_THEME: string;
		npm_config_node_gyp: string;
		XPC_SERVICE_NAME: string;
		npm_package_version: string;
		pnpm_config_verify_deps_before_run: string;
		HOME: string;
		SHLVL: string;
		STARBASE_THEME: string;
		VSCODE_GIT_ASKPASS_MAIN: string;
		PROTO_SHIM_PATH: string;
		MOON_VERSION: string;
		LESS: string;
		LOGNAME: string;
		MOON_WORKING_DIR: string;
		MOON_TASK_RETRY_ATTEMPT: string;
		npm_lifecycle_script: string;
		VSCODE_GIT_IPC_HANDLE: string;
		BUN_INSTALL: string;
		CLICOLOR_FORCE: string;
		MOON_TASK_ID: string;
		MOON_PROJECT_SNAPSHOT: string;
		npm_config_user_agent: string;
		GIT_ASKPASS: string;
		VSCODE_GIT_ASKPASS_NODE: string;
		PROTO_VERSION_CHECK: string;
		_P9K_SSH_TTY: string;
		MOON_PROJECT_ID: string;
		SECURITYSESSIONID: string;
		PROTO_OFFLINE_TIMEOUT: string;
		COLORTERM: string;
		npm_node_execpath: string;
		NODE_ENV: string;
		[key: `PUBLIC_${string}`]: undefined;
		[key: `${string}`]: string | undefined;
	}
}

/**
 * Similar to [`$env/dynamic/private`](https://svelte.dev/docs/kit/$env-dynamic-private), but only includes variables that begin with [`config.kit.env.publicPrefix`](https://svelte.dev/docs/kit/configuration#env) (which defaults to `PUBLIC_`), and can therefore safely be exposed to client-side code.
 * 
 * Note that public dynamic environment variables must all be sent from the server to the client, causing larger network requests — when possible, use `$env/static/public` instead.
 * 
 * ```ts
 * import { env } from '$env/dynamic/public';
 * console.log(env.PUBLIC_DEPLOYMENT_SPECIFIC_VARIABLE);
 * ```
 */
declare module '$env/dynamic/public' {
	export const env: {
		[key: `PUBLIC_${string}`]: string | undefined;
	}
}
