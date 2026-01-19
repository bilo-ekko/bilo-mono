"use client";

import { useSearchParams } from "next/navigation";
import { createTranslator } from "@bilo/translations";

export default function MarketingToolkitPage() {
  const searchParams = useSearchParams();
  const locale = searchParams.get("locale") || "en-GB";
  const t = createTranslator(locale as "en-GB" | "de-DE");

  return (
    <div className="p-8">
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-zinc-900 dark:text-white">
          {t("dashboard.pages.marketingToolkit.title")}
        </h1>
        <p className="mt-2 text-zinc-600 dark:text-zinc-400">
          {t("dashboard.pages.marketingToolkit.description")}
        </p>
      </div>

      <div className="space-y-6">
        <div className="rounded-lg border border-zinc-200 bg-white p-6 dark:border-zinc-800 dark:bg-zinc-900">
          <h2 className="mb-4 text-xl font-semibold text-zinc-900 dark:text-white">
            {t("dashboard.pages.marketingToolkit.brandAssets.title")}
          </h2>
          <p className="mb-4 text-sm text-zinc-600 dark:text-zinc-400">
            {t("dashboard.pages.marketingToolkit.brandAssets.description")}
          </p>
          <div className="flex gap-3">
            <button className="rounded-lg border border-zinc-200 px-4 py-2 text-sm font-medium text-zinc-900 hover:bg-zinc-50 dark:border-zinc-700 dark:text-white dark:hover:bg-zinc-800">
              {t("dashboard.pages.marketingToolkit.brandAssets.downloadLogos")}
            </button>
            <button className="rounded-lg border border-zinc-200 px-4 py-2 text-sm font-medium text-zinc-900 hover:bg-zinc-50 dark:border-zinc-700 dark:text-white dark:hover:bg-zinc-800">
              {t("dashboard.pages.marketingToolkit.brandAssets.viewGuidelines")}
            </button>
          </div>
        </div>

        <div className="rounded-lg border border-zinc-200 bg-white p-6 dark:border-zinc-800 dark:bg-zinc-900">
          <h2 className="mb-4 text-xl font-semibold text-zinc-900 dark:text-white">
            {t("dashboard.pages.marketingToolkit.socialMediaTemplates.title")}
          </h2>
          <p className="mb-4 text-sm text-zinc-600 dark:text-zinc-400">
            {t("dashboard.pages.marketingToolkit.socialMediaTemplates.description")}
          </p>
          <button className="rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700">
            {t("dashboard.pages.marketingToolkit.socialMediaTemplates.browseTemplates")}
          </button>
        </div>

        <div className="rounded-lg border border-zinc-200 bg-white p-6 dark:border-zinc-800 dark:bg-zinc-900">
          <h2 className="mb-4 text-xl font-semibold text-zinc-900 dark:text-white">
            {t("dashboard.pages.marketingToolkit.emailTemplates.title")}
          </h2>
          <p className="mb-4 text-sm text-zinc-600 dark:text-zinc-400">
            {t("dashboard.pages.marketingToolkit.emailTemplates.description")}
          </p>
          <button className="rounded-lg border border-zinc-200 px-4 py-2 text-sm font-medium text-zinc-900 hover:bg-zinc-50 dark:border-zinc-700 dark:text-white dark:hover:bg-zinc-800">
            {t("dashboard.pages.marketingToolkit.emailTemplates.viewTemplates")}
          </button>
        </div>
      </div>
    </div>
  );
}
