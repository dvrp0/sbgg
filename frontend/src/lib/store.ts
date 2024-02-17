import { writable } from "svelte/store";
import type { UserData } from "$lib/types";

export const user = writable({} as UserData);