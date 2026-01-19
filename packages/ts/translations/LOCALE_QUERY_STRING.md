# Locale Query String Implementation

Both frontend apps now support locale selection via a query string parameter.

---

## Query String Parameter

**Parameter name:** `locale`

**Valid values:**
- `en-GB` (English - default)
- `de-DE` (German)

**Behavior:**
- If the `locale` parameter is present and valid, that locale is used
- If the parameter is missing or invalid, defaults to `en-GB`

---

## Usage Examples

### English (default)

```
http://localhost:4000/
http://localhost:4000/?locale=en-GB
```

### German

```
http://localhost:4000/?locale=de-DE
```

### Invalid locale (falls back to English)

```
http://localhost:4000/?locale=fr-FR
→ Renders in English (en-GB)
```

---

## Web Dashboard (Next.js)

The dashboard reads the locale from the `searchParams` in the server component:

**File:** `apps/frontend/web-dashboard/src/app/page.tsx`

```typescript
export default async function Home({
  searchParams,
}: {
  searchParams: Promise<{ locale?: string }>;
}) {
  const params = await searchParams;
  const localeParam = params.locale;
  const locale = localeParam && isValidLocale(localeParam) ? localeParam : "en-GB";
  const t = createTranslator(locale);
  // ...
}
```

### Testing

```bash
# Start the dashboard
cd apps/frontend/web-dashboard
pnpm dev

# Visit in browser:
# http://localhost:4000/
# http://localhost:4000/?locale=de-DE
```

---

## Web SDKs (SvelteKit)

The SDKs app reads the locale from the `$page` store in all components:

**Updated files:**
- `src/routes/+page.svelte`
- `src/routes/checkout/+page.svelte`
- `src/routes/post-purchase/+page.svelte`
- `src/lib/components/CheckoutWidget.svelte`
- `src/lib/components/PostPurchaseWidget.svelte`

**Pattern used:**

```typescript
import { page } from '$app/stores';
import { createTranslator, isValidLocale } from '@bilo/translations';

let localeParam = $derived($page.url.searchParams.get('locale'));
let locale = $derived(localeParam && isValidLocale(localeParam) ? localeParam : 'en-GB');
let t = $derived(createTranslator(locale));
```

### Testing

```bash
# Start the SDKs app
cd apps/frontend/web-sdks-apps
pnpm dev

# Visit in browser:
# http://localhost:5173/
# http://localhost:5173/?locale=de-DE
# http://localhost:5173/checkout?locale=de-DE
# http://localhost:5173/post-purchase?locale=de-DE
```

---

## Implementation Details

### Validation

The `isValidLocale()` function ensures only supported locales are used:

```typescript
import { isValidLocale } from '@bilo/translations';

const locale = localeParam && isValidLocale(localeParam) ? localeParam : 'en-GB';
```

This prevents errors from unsupported locale codes.

### Reactivity (SvelteKit)

In SvelteKit, the locale updates automatically when the URL changes:

```typescript
// Using $derived for reactivity
let locale = $derived(/* ... */);
let t = $derived(createTranslator(locale));
```

When the URL changes (e.g., via navigation or history), the translations update automatically.

### Server Rendering (Next.js)

In Next.js, the page is now dynamically rendered:

```
Route (app)
┌ ƒ /        ← Dynamic route (uses searchParams)
```

This ensures the locale is read server-side for each request.

---

## Adding Locale Switcher UI

To add a locale switcher dropdown:

### Next.js Example

```tsx
'use client';

import { useRouter, useSearchParams } from 'next/navigation';

export function LocaleSwitcher() {
  const router = useRouter();
  const searchParams = useSearchParams();
  const currentLocale = searchParams.get('locale') || 'en-GB';

  const handleChange = (newLocale: string) => {
    const params = new URLSearchParams(searchParams);
    params.set('locale', newLocale);
    router.push(`?${params.toString()}`);
  };

  return (
    <select value={currentLocale} onChange={(e) => handleChange(e.target.value)}>
      <option value="en-GB">English</option>
      <option value="de-DE">Deutsch</option>
    </select>
  );
}
```

### SvelteKit Example

```svelte
<script lang="ts">
  import { goto } from '$app/navigation';
  import { page } from '$app/stores';

  let currentLocale = $derived($page.url.searchParams.get('locale') || 'en-GB');

  function handleChange(event: Event) {
    const target = event.target as HTMLSelectElement;
    const newLocale = target.value;
    const newUrl = new URL($page.url);
    newUrl.searchParams.set('locale', newLocale);
    goto(newUrl.toString());
  }
</script>

<select value={currentLocale} onchange={handleChange}>
  <option value="en-GB">English</option>
  <option value="de-DE">Deutsch</option>
</select>
```

---

## Files Changed

### Translations Package
- ✅ `examples/nextjs-dashboard.tsx` → `examples/nextjs-dashboard.md`
- ✅ `examples/sveltekit-sdks.ts` → `examples/sveltekit-sdks.md`

### Web Dashboard
- ✅ `src/app/page.tsx` - Added `searchParams` handling

### Web SDKs
- ✅ `src/routes/+page.svelte` - Added `$page` store usage
- ✅ `src/routes/checkout/+page.svelte` - Added `$page` store usage
- ✅ `src/routes/post-purchase/+page.svelte` - Added `$page` store usage
- ✅ `src/lib/components/CheckoutWidget.svelte` - Added `$page` store usage
- ✅ `src/lib/components/PostPurchaseWidget.svelte` - Added `$page` store usage

---

## Verification

### Build Status
- ✅ Web Dashboard builds successfully
- ✅ Web SDKs builds successfully
- ✅ No linter errors
- ✅ No TypeScript errors

### Behavior
- ✅ Default locale is `en-GB`
- ✅ Valid `locale` parameter switches language
- ✅ Invalid `locale` parameter falls back to `en-GB`
- ✅ Translations update reactively in SvelteKit

---

## Next Steps

1. **Add locale switcher UI** - Add dropdown/buttons to switch languages
2. **Persist locale preference** - Store in cookies or localStorage
3. **Add more languages** - Extend to support additional locales
4. **SEO optimization** - Add `<link rel="alternate">` tags for different locales
5. **URL routing** - Consider moving locale to path (e.g., `/en-GB/`, `/de-DE/`)
