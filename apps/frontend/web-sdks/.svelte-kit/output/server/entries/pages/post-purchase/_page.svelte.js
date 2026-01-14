import { w as head } from "../../../chunks/index.js";
import { E as EkkoLogo } from "../../../chunks/EkkoLogo.js";
function PostPurchaseWidget($$renderer) {
  $$renderer.push(`<article class="widget svelte-zx0orb"><div class="background-image svelte-zx0orb"><img src="https://images.unsplash.com/photo-1448375240586-882707db888b?w=1200&amp;h=600&amp;fit=crop&amp;q=80" alt="Dense forest canopy from above" loading="lazy" class="svelte-zx0orb"/></div> <div class="widget-content svelte-zx0orb"><div class="card svelte-zx0orb"><div class="card-content svelte-zx0orb"><h2 class="title svelte-zx0orb">Give a little. Change a lot.</h2> <p class="description svelte-zx0orb">Support climate projects and act on the carbon
					footprint of this purchase (~21 kgCO2e).</p></div></div> <button class="info-button svelte-zx0orb" aria-label="More information"><svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg"><path d="M10 18C14.4183 18 18 14.4183 18 10C18 5.58172 14.4183 2 10 2C5.58172 2 2 5.58172 2 10C2 14.4183 5.58172 18 10 18Z" stroke="currentColor" stroke-width="1.5"></path><path d="M10 14V10M10 6H10.01" stroke="currentColor" stroke-width="2" stroke-linecap="round"></path></svg></button></div> <footer class="footer svelte-zx0orb"><div class="powered-by svelte-zx0orb"><span>Powered by</span> `);
  EkkoLogo($$renderer, { size: "md" });
  $$renderer.push(`<!----></div> <button class="cta-button svelte-zx0orb"><span>Find out more</span> <svg width="20" height="20" viewBox="0 0 20 20" fill="none" xmlns="http://www.w3.org/2000/svg" class="svelte-zx0orb"><path d="M4 10H16M16 10L11 5M16 10L11 15" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"></path></svg></button></footer> <div class="label svelte-zx0orb">Embedded SDK</div></article>`);
}
function _page($$renderer) {
  head("19ghvg1", $$renderer, ($$renderer2) => {
    $$renderer2.title(($$renderer3) => {
      $$renderer3.push(`<title>Post-Purchase | Ekko SDK</title>`);
    });
  });
  $$renderer.push(`<main class="page svelte-19ghvg1"><div class="container svelte-19ghvg1">`);
  PostPurchaseWidget($$renderer);
  $$renderer.push(`<!----></div></main>`);
}
export {
  _page as default
};
