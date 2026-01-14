import { w as head } from "../../chunks/index.js";
function _layout($$renderer, $$props) {
  let { children } = $$props;
  head("12qhfyh", $$renderer, ($$renderer2) => {
    $$renderer2.title(($$renderer3) => {
      $$renderer3.push(`<title>Ekko SDK</title>`);
    });
    $$renderer2.push(`<meta name="description" content="Support climate action with every purchase"/>`);
  });
  children($$renderer);
  $$renderer.push(`<!---->`);
}
export {
  _layout as default
};
