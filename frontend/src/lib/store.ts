import resolveConfig from "tailwindcss/resolveConfig";
import tailwindConfig from "../../tailwind.config";
import { readable, writable } from "svelte/store";
import type { main } from "$wails/go/models";

export const TAILWIND = readable(resolveConfig(tailwindConfig));
export const user = writable({} as main.RegistryData);
export const profile = writable({} as main.Profile);