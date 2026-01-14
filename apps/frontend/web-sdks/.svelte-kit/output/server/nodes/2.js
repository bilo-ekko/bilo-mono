

export const index = 2;
let component_cache;
export const component = async () => component_cache ??= (await import('../entries/pages/_page.svelte.js')).default;
export const imports = ["_app/immutable/nodes/2.C4jqDXYS.js","_app/immutable/chunks/BKYosCqX.js","_app/immutable/chunks/KthFeat2.js","_app/immutable/chunks/BPaJSgOP.js"];
export const stylesheets = ["_app/immutable/assets/2.DLihSrRy.css"];
export const fonts = [];
