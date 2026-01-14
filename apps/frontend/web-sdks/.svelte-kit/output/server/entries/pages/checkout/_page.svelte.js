import { x as attr_class, y as attr, z as bind_props, F as ensure_array_like, G as attr_style, J as stringify, w as head } from "../../../chunks/index.js";
import { E as EkkoLogo } from "../../../chunks/EkkoLogo.js";
import { e as escape_html } from "../../../chunks/context.js";
function Toggle($$renderer, $$props) {
  $$renderer.component(($$renderer2) => {
    let { checked = false, onchange } = $$props;
    $$renderer2.push(`<button${attr_class("toggle svelte-d39fdc", void 0, { "active": checked })} role="switch"${attr("aria-checked", checked)}><span class="toggle-thumb svelte-d39fdc"></span></button>`);
    bind_props($$props, { checked });
  });
}
function GoldStandardLogo($$renderer) {
  $$renderer.push(`<div class="gold-standard svelte-4m8jzk"><span class="gs-text svelte-4m8jzk">Gold Standard</span> <span class="gs-reg svelte-4m8jzk">®</span></div>`);
}
function PartnerLogos($$renderer) {
  const partners = [
    {
      name: "Conservation International",
      abbr: "CI",
      color: "#1a5632"
    },
    { name: "RSPB", abbr: "RSPB", color: "#2d5a27" },
    { name: "Tusk", abbr: "Tusk", color: "#d4432c" },
    { name: "1% for the Planet", abbr: "1%", color: "#006eb6" },
    {
      name: "Conservation Collective",
      abbr: "CC",
      color: "#2d7a3f"
    },
    { name: "Justdiggit", abbr: "JD", color: "#7cb342" }
  ];
  $$renderer.push(`<div class="partners svelte-1pgsogi"><!--[-->`);
  const each_array = ensure_array_like(partners);
  for (let $$index = 0, $$length = each_array.length; $$index < $$length; $$index++) {
    let partner = each_array[$$index];
    $$renderer.push(`<div class="partner-logo svelte-1pgsogi"${attr_style(`--partner-color: ${stringify(partner.color)}`)}${attr("title", partner.name)}><span class="partner-abbr svelte-1pgsogi">${escape_html(partner.abbr)}</span></div>`);
  }
  $$renderer.push(`<!--]--></div>`);
}
function CheckoutWidget($$renderer) {
  let climateActionEnabled = true;
  let roundUpEnabled = false;
  let isActive = climateActionEnabled || roundUpEnabled;
  const climateActionCost = 0.65;
  const roundUpCost = 0.85;
  let $$settled = true;
  let $$inner_renderer;
  function $$render_inner($$renderer2) {
    $$renderer2.push(`<article class="widget svelte-19uqmf4"><div class="widget-content svelte-19uqmf4"><div class="image-section svelte-19uqmf4"><img src="https://images.unsplash.com/photo-1529963183134-61a90db47eaf?w=400&amp;h=500&amp;fit=crop&amp;q=80" alt="Snow-capped mountain peak" loading="lazy" class="svelte-19uqmf4"/></div> <div class="content-section svelte-19uqmf4"><header class="header svelte-19uqmf4"><h2 class="title svelte-19uqmf4">Give a little. Change a lot.</h2> <p class="subtitle svelte-19uqmf4">Support <span class="highlight svelte-19uqmf4">environmental projects</span> and act on the ~21
					kgCO<sub class="svelte-19uqmf4">2</sub>e footprint of this purchase - about what <span class="highlight svelte-19uqmf4">1 tree</span> can capture in <span class="highlight svelte-19uqmf4">1 year</span>!</p></header> <div class="options svelte-19uqmf4"><div class="option svelte-19uqmf4"><div class="option-content svelte-19uqmf4"><div class="option-header svelte-19uqmf4"><span class="option-title svelte-19uqmf4">Support climate action</span> <span class="option-price svelte-19uqmf4">£${escape_html(climateActionCost.toFixed(2))}</span></div> <div class="option-partner svelte-19uqmf4"><span class="with-text svelte-19uqmf4">with</span> `);
    GoldStandardLogo($$renderer2);
    $$renderer2.push(`<!----></div></div> `);
    Toggle($$renderer2, {
      get checked() {
        return climateActionEnabled;
      },
      set checked($$value) {
        climateActionEnabled = $$value;
        $$settled = false;
      }
    });
    $$renderer2.push(`<!----></div> <div class="option svelte-19uqmf4"><div class="option-content svelte-19uqmf4"><div class="option-header svelte-19uqmf4"><span class="option-title svelte-19uqmf4">Round up to boost impact</span> <span class="option-price svelte-19uqmf4">£${escape_html(roundUpCost.toFixed(2))}</span></div> <div class="option-partner svelte-19uqmf4"><span class="with-text svelte-19uqmf4">with</span> `);
    PartnerLogos($$renderer2);
    $$renderer2.push(`<!----></div></div> `);
    Toggle($$renderer2, {
      get checked() {
        return roundUpEnabled;
      },
      set checked($$value) {
        roundUpEnabled = $$value;
        $$settled = false;
      }
    });
    $$renderer2.push(`<!----></div></div> <footer class="footer svelte-19uqmf4"><a href="#learn" class="learn-more svelte-19uqmf4">Learn more</a> <div class="powered-by svelte-19uqmf4"><span>Powered by</span> `);
    EkkoLogo($$renderer2, { size: "sm" });
    $$renderer2.push(`<!----></div></footer></div></div> <div${attr_class("thank-you svelte-19uqmf4", void 0, { "active": isActive })}><span>Thank you! Together, we're creating real change.</span> <span class="leaf svelte-19uqmf4">🌿</span></div></article>`);
  }
  do {
    $$settled = true;
    $$inner_renderer = $$renderer.copy();
    $$render_inner($$inner_renderer);
  } while (!$$settled);
  $$renderer.subsume($$inner_renderer);
}
function _page($$renderer) {
  head("jbcej5", $$renderer, ($$renderer2) => {
    $$renderer2.title(($$renderer3) => {
      $$renderer3.push(`<title>Checkout | Ekko SDK</title>`);
    });
  });
  $$renderer.push(`<main class="page svelte-jbcej5"><div class="container svelte-jbcej5">`);
  CheckoutWidget($$renderer);
  $$renderer.push(`<!----></div></main>`);
}
export {
  _page as default
};
