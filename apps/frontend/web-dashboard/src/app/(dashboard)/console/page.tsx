"use client";

import { Suspense } from "react";
import { useSearchParams } from "next/navigation";
import { createTranslator } from "@bilo/translations";

function ConsoleContent() {
  const searchParams = useSearchParams();
  const locale = searchParams.get("locale") || "en-GB";
  const t = createTranslator(locale as "en-GB" | "de-DE");

  return (
    <div className="p-8">
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-zinc-900 dark:text-white">
          {t("dashboard.pages.console.title")}
        </h1>
        <p className="mt-2 text-zinc-600 dark:text-zinc-400">
          {t("dashboard.pages.console.description")}
        </p>
      </div>

      <div className="space-y-6">
        <div className="rounded-lg border border-zinc-200 bg-white p-6 dark:border-zinc-800 dark:bg-zinc-900">
          <h2 className="mb-4 text-xl font-semibold text-zinc-900 dark:text-white">
            {t("dashboard.pages.console.apiKeys.title")}
          </h2>
          <p className="text-sm text-zinc-600 dark:text-zinc-400">
            {t("dashboard.pages.console.apiKeys.description")}
          </p>
        </div>

        <div className="rounded-lg border border-zinc-200 bg-white p-6 dark:border-zinc-800 dark:bg-zinc-900">
          <h2 className="mb-4 text-xl font-semibold text-zinc-900 dark:text-white">
            {t("dashboard.pages.console.webhooks.title")}
          </h2>
          <p className="text-sm text-zinc-600 dark:text-zinc-400">
            {t("dashboard.pages.console.webhooks.description")}
          </p>
        </div>

        <div className="rounded-lg border border-zinc-200 bg-white p-6 dark:border-zinc-800 dark:bg-zinc-900">
          <h2 className="mb-4 text-xl font-semibold text-zinc-900 dark:text-white">
            {t("dashboard.pages.console.logs.title")}
          </h2>
          <p className="text-sm text-zinc-600 dark:text-zinc-400">
            {t("dashboard.pages.console.logs.description")}
          </p>
        </div>
      </div>
    </div>
  );
}

export default function ConsolePage() {
  return (
    <Suspense fallback={<div className="p-8">Loading...</div>}>
      <ConsoleContent />
    </Suspense>
  );
}
