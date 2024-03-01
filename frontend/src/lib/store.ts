import resolveConfig from "tailwindcss/resolveConfig";
import tailwindConfig from "../../tailwind.config";
import { readable, writable } from "svelte/store";
import type { RegistryData } from "$lib/types";

export const TAILWIND = readable(resolveConfig(tailwindConfig));
export const user = writable({} as RegistryData);