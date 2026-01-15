# SDK UI

A SvelteKit application rendering climate action widgets for checkout and post-purchase flows.

## Features

- **Checkout Widget**: Interactive widget with toggles for supporting climate action and environmental projects
- **Post-Purchase Widget**: Embedded SDK display with forest background showing purchase impact

## Getting Started

### Install dependencies

```bash
pnpm install
```

### Development

```bash
pnpm dev
```

The app will be available at `http://localhost:5173`

### Routes

- `/` - Home page with navigation
- `/checkout` - Checkout widget with climate action toggles
- `/post-purchase` - Post-purchase confirmation widget

### Build

```bash
pnpm build
```

### Preview production build

```bash
pnpm preview
```

## Tech Stack

- SvelteKit 2
- Svelte 5 (with runes)
- TypeScript
- Vite

## Components

- `CheckoutWidget` - Main checkout widget with toggle options
- `PostPurchaseWidget` - Post-purchase confirmation display
- `Toggle` - Reusable toggle switch component
- `EkkoLogo` - Ekko branding component
- `PartnerLogos` - Partner organization logos
- `GoldStandardLogo` - Gold Standard certification badge
