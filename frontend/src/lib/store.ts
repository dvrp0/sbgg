import resolveConfig from "tailwindcss/resolveConfig";
import tailwindConfig from "../../tailwind.config";
import { readable, writable } from "svelte/store";
import type { UserData } from "$lib/types";

export const TAILWIND = readable(resolveConfig(tailwindConfig));
export const user = writable({} as UserData);