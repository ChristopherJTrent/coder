/**
 * @fileoverview TODO: centralize navigation code here! URL constants, URL formatting, all of it
 */

import { useEffectEvent } from "hooks/hookPolyfills";
import type { DashboardValue } from "./dashboard/DashboardProvider";
import { selectFeatureVisibility } from "./dashboard/entitlements";
import { useDashboard } from "./dashboard/useDashboard";

type LinkThunk = (state: DashboardValue) => string;

export function useLinks() {
  const dashboard = useDashboard();
  // Needs to be safe to call `get` from inside of a `useEffect` without causing
  // excess triggers from adding it as a dependency.
  const get = useEffectEvent((thunk: LinkThunk): string => thunk(dashboard));
  return get;
}

export function withFilter(path: string, filter: string) {
  return path + (filter ? `?filter=${encodeURIComponent(filter)}` : "");
}

export const linkToAuditing = "/audit";

export const linkToUsers = withFilter("/users", "status:active");

export const linkToTemplate =
  (organizationName: string, templateName: string): LinkThunk =>
  (dashboard) => {
    const hasMultipleOrganizations = dashboard.organizations.length > 1;
    const organizationsEnabled =
      dashboard.experiments.includes("multi-organization") &&
      selectFeatureVisibility(dashboard.entitlements).multiple_organizations;

    return hasMultipleOrganizations || organizationsEnabled
      ? `/templates/${organizationName}/${templateName}`
      : `/templates/${templateName}`;
  };
