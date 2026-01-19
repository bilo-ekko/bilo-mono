"use client";

import * as React from "react";
import { useSearchParams, useRouter, usePathname } from "next/navigation";
import { createTranslator } from "@bilo/translations";

const tabs = [
  { id: "info", labelKey: "dashboard.pages.organisation.tabs.info" },
  { id: "hierarchy", labelKey: "dashboard.pages.organisation.tabs.hierarchy" },
  { id: "users", labelKey: "dashboard.pages.organisation.tabs.users" },
] as const;

type TabId = (typeof tabs)[number]["id"];

export default function OrganisationPage() {
  const router = useRouter();
  const pathname = usePathname();
  const searchParams = useSearchParams();
  const activeTab = (searchParams.get("tab") || "info") as TabId;
  
  const locale = searchParams.get("locale") || "en-GB";
  const t = createTranslator(locale as "en-GB" | "de-DE");

  const handleTabChange = (tabId: TabId) => {
    const params = new URLSearchParams(searchParams.toString());
    params.set("tab", tabId);
    router.push(`${pathname}?${params.toString()}`);
  };

  return (
    <div className="p-8">
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-zinc-900 dark:text-white">
          {t("dashboard.pages.organisation.title")}
        </h1>
        <p className="mt-2 text-zinc-600 dark:text-zinc-400">
          {t("dashboard.pages.organisation.description")}
        </p>
      </div>

      {/* Tabs */}
      <div className="mb-6 border-b border-zinc-200 dark:border-zinc-800">
        <nav className="-mb-px flex space-x-8">
          {tabs.map((tab) => (
            <button
              key={tab.id}
              onClick={() => handleTabChange(tab.id)}
              className={`whitespace-nowrap border-b-2 px-1 py-4 text-sm font-medium transition-colors ${
                activeTab === tab.id
                  ? "border-blue-600 text-blue-600 dark:border-blue-400 dark:text-blue-400"
                  : "border-transparent text-zinc-600 hover:border-zinc-300 hover:text-zinc-900 dark:text-zinc-400 dark:hover:border-zinc-600 dark:hover:text-white"
              }`}
            >
              {t(tab.labelKey)}
            </button>
          ))}
        </nav>
      </div>

      {/* Tab Content */}
      <div>
        {activeTab === "info" && (
          <div className="space-y-6">
            <div className="rounded-lg border border-zinc-200 bg-white p-6 dark:border-zinc-800 dark:bg-zinc-900">
              <h2 className="mb-4 text-xl font-semibold text-zinc-900 dark:text-white">
                {t("dashboard.pages.organisation.info.basicInformation")}
              </h2>
              <div className="space-y-4">
                <div>
                  <label className="block text-sm font-medium text-zinc-700 dark:text-zinc-300">
                    {t("dashboard.pages.organisation.info.organisationName")}
                  </label>
                  <input
                    type="text"
                    placeholder={t("dashboard.pages.organisation.info.organisationNamePlaceholder")}
                    className="mt-1 w-full rounded-lg border border-zinc-300 bg-white px-4 py-2 text-sm focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500 dark:border-zinc-700 dark:bg-zinc-800 dark:text-white"
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-zinc-700 dark:text-zinc-300">
                    {t("dashboard.pages.organisation.info.legalName")}
                  </label>
                  <input
                    type="text"
                    placeholder={t("dashboard.pages.organisation.info.legalNamePlaceholder")}
                    className="mt-1 w-full rounded-lg border border-zinc-300 bg-white px-4 py-2 text-sm focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500 dark:border-zinc-700 dark:bg-zinc-800 dark:text-white"
                  />
                </div>
              </div>
            </div>

            <div className="rounded-lg border border-zinc-200 bg-white p-6 dark:border-zinc-800 dark:bg-zinc-900">
              <h2 className="mb-4 text-xl font-semibold text-zinc-900 dark:text-white">
                {t("dashboard.pages.organisation.info.billingAddress")}
              </h2>
              <div className="space-y-4">
                <div>
                  <label className="block text-sm font-medium text-zinc-700 dark:text-zinc-300">
                    {t("dashboard.pages.organisation.info.streetAddress")}
                  </label>
                  <input
                    type="text"
                    placeholder={t("dashboard.pages.organisation.info.streetAddressPlaceholder")}
                    className="mt-1 w-full rounded-lg border border-zinc-300 bg-white px-4 py-2 text-sm focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500 dark:border-zinc-700 dark:bg-zinc-800 dark:text-white"
                  />
                </div>
                <div className="grid grid-cols-2 gap-4">
                  <div>
                    <label className="block text-sm font-medium text-zinc-700 dark:text-zinc-300">
                      {t("dashboard.pages.organisation.info.city")}
                    </label>
                    <input
                      type="text"
                      placeholder={t("dashboard.pages.organisation.info.cityPlaceholder")}
                      className="mt-1 w-full rounded-lg border border-zinc-300 bg-white px-4 py-2 text-sm focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500 dark:border-zinc-700 dark:bg-zinc-800 dark:text-white"
                    />
                  </div>
                  <div>
                    <label className="block text-sm font-medium text-zinc-700 dark:text-zinc-300">
                      {t("dashboard.pages.organisation.info.postalCode")}
                    </label>
                    <input
                      type="text"
                      placeholder={t("dashboard.pages.organisation.info.postalCodePlaceholder")}
                      className="mt-1 w-full rounded-lg border border-zinc-300 bg-white px-4 py-2 text-sm focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500 dark:border-zinc-700 dark:bg-zinc-800 dark:text-white"
                    />
                  </div>
                </div>
                <div>
                  <label className="block text-sm font-medium text-zinc-700 dark:text-zinc-300">
                    {t("dashboard.pages.organisation.info.country")}
                  </label>
                  <input
                    type="text"
                    placeholder={t("dashboard.pages.organisation.info.countryPlaceholder")}
                    className="mt-1 w-full rounded-lg border border-zinc-300 bg-white px-4 py-2 text-sm focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500 dark:border-zinc-700 dark:bg-zinc-800 dark:text-white"
                  />
                </div>
              </div>
            </div>

            <div className="rounded-lg border border-zinc-200 bg-white p-6 dark:border-zinc-800 dark:bg-zinc-900">
              <h2 className="mb-4 text-xl font-semibold text-zinc-900 dark:text-white">
                {t("dashboard.pages.organisation.info.contactInformation")}
              </h2>
              <div className="space-y-4">
                <div>
                  <label className="block text-sm font-medium text-zinc-700 dark:text-zinc-300">
                    {t("dashboard.pages.organisation.info.email")}
                  </label>
                  <input
                    type="email"
                    placeholder={t("dashboard.pages.organisation.info.emailPlaceholder")}
                    className="mt-1 w-full rounded-lg border border-zinc-300 bg-white px-4 py-2 text-sm focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500 dark:border-zinc-700 dark:bg-zinc-800 dark:text-white"
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium text-zinc-700 dark:text-zinc-300">
                    {t("dashboard.pages.organisation.info.phoneNumber")}
                  </label>
                  <input
                    type="tel"
                    placeholder={t("dashboard.pages.organisation.info.phoneNumberPlaceholder")}
                    className="mt-1 w-full rounded-lg border border-zinc-300 bg-white px-4 py-2 text-sm focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500 dark:border-zinc-700 dark:bg-zinc-800 dark:text-white"
                  />
                </div>
              </div>
            </div>

            <div className="flex justify-end">
              <button className="rounded-lg bg-blue-600 px-6 py-2 text-sm font-medium text-white hover:bg-blue-700">
                {t("dashboard.pages.organisation.info.saveChanges")}
              </button>
            </div>
          </div>
        )}

        {activeTab === "hierarchy" && (
          <div className="space-y-6">
            <div className="rounded-lg border border-zinc-200 bg-white p-6 dark:border-zinc-800 dark:bg-zinc-900">
              <h2 className="mb-4 text-xl font-semibold text-zinc-900 dark:text-white">
                {t("dashboard.pages.organisation.hierarchy.organisationStructure")}
              </h2>
              <p className="mb-4 text-sm text-zinc-600 dark:text-zinc-400">
                {t("dashboard.pages.organisation.hierarchy.description")}
              </p>
              <div className="mt-6 rounded-lg bg-zinc-50 p-8 text-center dark:bg-zinc-950">
                <p className="text-sm text-zinc-500 dark:text-zinc-400">
                  {t("dashboard.pages.organisation.hierarchy.chartPlaceholder")}
                </p>
              </div>
            </div>

            <div className="rounded-lg border border-zinc-200 bg-white p-6 dark:border-zinc-800 dark:bg-zinc-900">
              <h2 className="mb-4 text-xl font-semibold text-zinc-900 dark:text-white">
                {t("dashboard.pages.organisation.hierarchy.departments")}
              </h2>
              <p className="text-sm text-zinc-600 dark:text-zinc-400">
                {t("dashboard.pages.organisation.hierarchy.departmentsDescription")}
              </p>
            </div>
          </div>
        )}

        {activeTab === "users" && (
          <div className="space-y-6">
            <div className="rounded-lg border border-zinc-200 bg-white p-6 dark:border-zinc-800 dark:bg-zinc-900">
              <div className="mb-4 flex items-center justify-between">
                <h2 className="text-xl font-semibold text-zinc-900 dark:text-white">
                  {t("dashboard.pages.organisation.users.usersList")}
                </h2>
                <button className="rounded-lg bg-blue-600 px-4 py-2 text-sm font-medium text-white hover:bg-blue-700">
                  {t("dashboard.pages.organisation.users.addUser")}
                </button>
              </div>
              <p className="text-sm text-zinc-600 dark:text-zinc-400">
                {t("dashboard.pages.organisation.users.description")}
              </p>
            </div>

            <div className="rounded-lg border border-zinc-200 bg-white p-6 dark:border-zinc-800 dark:bg-zinc-900">
              <h2 className="mb-4 text-xl font-semibold text-zinc-900 dark:text-white">
                {t("dashboard.pages.organisation.users.rolesPermissions")}
              </h2>
              <p className="text-sm text-zinc-600 dark:text-zinc-400">
                {t("dashboard.pages.organisation.users.rolesDescription")}
              </p>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}
