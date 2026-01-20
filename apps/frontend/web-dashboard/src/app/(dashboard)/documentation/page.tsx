"use client";

import { Suspense } from "react";
import { useSearchParams } from "next/navigation";
import { createTranslator } from "@bilo/translations";

function DocumentationContent() {
  const searchParams = useSearchParams();
  const locale = searchParams.get("locale") || "en-GB";
  const t = createTranslator(locale as "en-GB" | "de-DE");

  return (
    <div className="p-8">
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-zinc-900 dark:text-white">
          {t("dashboard.pages.documentation.title")}
        </h1>
        <p className="mt-2 text-zinc-600 dark:text-zinc-400">
          {t("dashboard.pages.documentation.description")}
        </p>
      </div>

      <div className="grid gap-6 md:grid-cols-2">
        <div className="rounded-lg border border-zinc-200 bg-white p-6 dark:border-zinc-800 dark:bg-zinc-900">
          <h2 className="mb-3 text-xl font-semibold text-zinc-900 dark:text-white">
            {t("dashboard.pages.documentation.gettingStarted.title")}
          </h2>
          <p className="mb-4 text-sm text-zinc-600 dark:text-zinc-400">
            {t("dashboard.pages.documentation.gettingStarted.description")}
          </p>
          <a
            href="#"
            className="text-sm font-medium text-blue-600 hover:text-blue-700 dark:text-blue-400"
          >
            {t("dashboard.pages.documentation.gettingStarted.readMore")}
          </a>
        </div>

        <div className="rounded-lg border border-zinc-200 bg-white p-6 dark:border-zinc-800 dark:bg-zinc-900">
          <h2 className="mb-3 text-xl font-semibold text-zinc-900 dark:text-white">
            {t("dashboard.pages.documentation.apiReference.title")}
          </h2>
          <p className="mb-4 text-sm text-zinc-600 dark:text-zinc-400">
            {t("dashboard.pages.documentation.apiReference.description")}
          </p>
          <a
            href="#"
            className="text-sm font-medium text-blue-600 hover:text-blue-700 dark:text-blue-400"
          >
            {t("dashboard.pages.documentation.apiReference.readMore")}
          </a>
        </div>

        <div className="rounded-lg border border-zinc-200 bg-white p-6 dark:border-zinc-800 dark:bg-zinc-900">
          <h2 className="mb-3 text-xl font-semibold text-zinc-900 dark:text-white">
            {t("dashboard.pages.documentation.sdksLibraries.title")}
          </h2>
          <p className="mb-4 text-sm text-zinc-600 dark:text-zinc-400">
            {t("dashboard.pages.documentation.sdksLibraries.description")}
          </p>
          <a
            href="#"
            className="text-sm font-medium text-blue-600 hover:text-blue-700 dark:text-blue-400"
          >
            {t("dashboard.pages.documentation.sdksLibraries.readMore")}
          </a>
        </div>

        <div className="rounded-lg border border-zinc-200 bg-white p-6 dark:border-zinc-800 dark:bg-zinc-900">
          <h2 className="mb-3 text-xl font-semibold text-zinc-900 dark:text-white">
            {t("dashboard.pages.documentation.tutorials.title")}
          </h2>
          <p className="mb-4 text-sm text-zinc-600 dark:text-zinc-400">
            {t("dashboard.pages.documentation.tutorials.description")}
          </p>
          <a
            href="#"
            className="text-sm font-medium text-blue-600 hover:text-blue-700 dark:text-blue-400"
          >
            {t("dashboard.pages.documentation.tutorials.readMore")}
          </a>
        </div>
      </div>
    </div>
  );
}

export default function DocumentationPage() {
  return (
    <Suspense fallback={<div className="p-8">Loading...</div>}>
      <DocumentationContent />
    </Suspense>
  );
}
