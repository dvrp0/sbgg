<script lang="ts">
    import { onMount } from "svelte";
    import Match from "$components/Match.svelte";
    import Skeleton from "$components/Skeleton.svelte";
    import { user } from "$lib/store";
    import { GetUserData } from "$wails/App";

    let isMounted = false;

    onMount(async () => {
        $user = await GetUserData();

        isMounted = true;
    });
</script>

<main class="flex flex-col px-4xl pt-4xl">
    {#if !isMounted}
        <div class="grid grid-rows-[repeat(4,auto)] gap-xs">
            <Skeleton kind="body" />
            <Skeleton kind="h1" />
            <Skeleton kind="h2" />
            <Skeleton kind="body" />
        </div>
    {:else}
        <div class="flex">
            <span class="font-black">dvrp</span><span class="text-gray-500">#1001413882</span>
        </div>
        <div class="flex mt-xs">
            <h1 class="text-platinum">Platinum 2</h1>
        </div>
        <div class="grid grid-cols-[repeat(4,auto)] grid-rows-2 w-fit mt-sm gap-y-2xs gap-x-4xl">
            <h2>116</h2>
            <h2>13</h2>
            <h2>89.4%</h2>
            <h2>15</h2>
            <span>Wins</span>
            <span>Loses</span>
            <span>Win rate</span>
            <span>Avg. turns</span>
        </div>
    {/if}
    <div class="border-solid border-[1px] border-gray-100 -mx-4xl my-md" />
    <div class="flex justify-between mt-xs">
        <span>Matches</span>
        <span class="text-gray-500">Share</span>
    </div>
    <div class="grid grid-cols-10 mt-xs gap-xs mb-4xl">
        {#each Array(100) as _}
            <Match won={Math.random() > 0.5} />
        {/each}
    </div>
</main>
<div class="fixed bottom-0 w-full bg-gradient-to-t from-white from-10% h-4xl pointer-events-none" />