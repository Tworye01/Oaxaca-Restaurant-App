export function buildFilterQuery(
  selectedFilters: string[],
  caloriesRange: number[]
): string {
  const queryParams = [];

  const allergens = selectedFilters.filter((filter) => filter !== "vegetarian");
  if (allergens.length > 0) {
    queryParams.push(`allergens=${allergens.join(",")}`);
  }

  if (selectedFilters.includes("vegetarian")) {
    queryParams.push("vegetarian=1");
  }

  if (caloriesRange[1] > 0) {
    queryParams.push(`calories=${caloriesRange[1]}`);
  }

  return queryParams.length > 0 ? `?${queryParams.join("&")}` : "";
}
