# Playwright Tests

End-to-end testing for the bilo-mono project using [Playwright](https://playwright.dev/).

## Overview

Playwright is a modern end-to-end testing framework that enables reliable testing across Chromium, Firefox, and WebKit with a single API. It provides powerful automation capabilities with auto-wait, web-first assertions, and built-in test isolation.

## Getting Started

### Installation

Dependencies are already installed in this project. If you need to reinstall:

```bash
pnpm install
```

### Running Tests

```bash
# Run all tests
pnpm test

# Run tests in UI mode (interactive)
pnpm test:ui

# Run tests in debug mode
pnpm test:debug

# Run tests only on Desktop Chrome
pnpm test:chromium

# Auto-generate tests with Codegen
pnpm codegen
```

## Available Commands

| Command | Description |
|---------|-------------|
| `pnpm test` | Runs all end-to-end tests |
| `pnpm test:ui` | Starts the interactive UI mode for visual test debugging |
| `pnpm test:chromium` | Runs tests only on Desktop Chrome |
| `pnpm test:debug` | Runs tests in debug mode with step-by-step execution |
| `pnpm codegen` | Opens Codegen tool to auto-generate tests by recording interactions |

## Project Structure

```
tests/playwright/
├── tests/
│   └── example.spec.ts      # Example end-to-end test
├── playwright.config.ts      # Playwright configuration
├── package.json             # Dependencies and scripts
└── README.md               # This file
```

## Writing Tests

Tests are located in the `tests/` directory. Here's a basic example:

```typescript
import { test, expect } from '@playwright/test';

test('basic test', async ({ page }) => {
  await page.goto('https://playwright.dev/');
  await expect(page).toHaveTitle(/Playwright/);
});
```

### Running Specific Tests

```bash
# Run a specific test file
pnpm exec playwright test example

# Run tests matching a pattern
pnpm exec playwright test --grep "login"
```

## Configuration

The Playwright configuration is defined in `playwright.config.ts`. You can customize:

- Base URL
- Browser projects (Chromium, Firefox, WebKit)
- Test timeout settings
- Screenshot and video recording options
- Parallel execution settings

## Resources

- [Playwright Documentation](https://playwright.dev/docs/intro)
- [API Reference](https://playwright.dev/docs/api/class-playwright)
- [Best Practices](https://playwright.dev/docs/best-practices)
- [Test Generator (Codegen)](https://playwright.dev/docs/codegen)
- [Debugging Tests](https://playwright.dev/docs/debug)
- [CI/CD Integration](https://playwright.dev/docs/ci)

## Tips

- Use the UI mode (`pnpm test:ui`) for debugging and developing tests
- Use Codegen (`pnpm codegen`) to quickly generate test code by recording actions
- Tests run in parallel by default for faster execution
- Playwright automatically waits for elements to be actionable before performing actions
- Use `page.pause()` to pause test execution and inspect the page

## Support

For issues and questions:
- Check the [Playwright documentation](https://playwright.dev/docs/intro)
- Visit the [GitHub repository](https://github.com/microsoft/playwright)
- Join the [Discord community](https://aka.ms/playwright/discord)

---

Happy testing! 🎭
