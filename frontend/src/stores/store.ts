import { writable, type Writable } from 'svelte/store';

export const sidebarState: Writable<string> = writable('init'); // TODO: adapt to more complex state behaviours