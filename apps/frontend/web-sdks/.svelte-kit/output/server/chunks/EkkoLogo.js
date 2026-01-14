import { G as attr_style, J as stringify } from "./index.js";
function EkkoLogo($$renderer, $$props) {
  let { size = "md" } = $$props;
  const sizes = {
    sm: { logo: 24, text: 14 },
    md: { logo: 32, text: 18 },
    lg: { logo: 40, text: 22 }
  };
  $$renderer.push(`<div class="ekko-logo svelte-1mkmq2j"${attr_style(`--logo-size: ${stringify(sizes[size].logo)}px; --text-size: ${stringify(sizes[size].text)}px;`)}><div class="logo-circle svelte-1mkmq2j"><svg viewBox="0 0 32 32" fill="none" xmlns="http://www.w3.org/2000/svg" class="svelte-1mkmq2j"><circle cx="16" cy="16" r="16" fill="url(#gradient)"></circle><path d="M10 16C10 12.6863 12.6863 10 16 10C19.3137 10 22 12.6863 22 16" stroke="white" stroke-width="2.5" stroke-linecap="round"></path><path d="M13 16C13 14.3431 14.3431 13 16 13C17.6569 13 19 14.3431 19 16" stroke="white" stroke-width="2" stroke-linecap="round"></path><defs><linearGradient id="gradient" x1="0" y1="0" x2="32" y2="32"><stop offset="0%" stop-color="#22c55e"></stop><stop offset="100%" stop-color="#0ea5e9"></stop></linearGradient></defs></svg></div> <span class="logo-text svelte-1mkmq2j">e<span class="k-flip svelte-1mkmq2j">k</span>ko</span></div>`);
}
export {
  EkkoLogo as E
};
