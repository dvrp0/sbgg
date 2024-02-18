<script lang="ts">
    import { onMount } from "svelte";
    import Icon from "$components/Icon.svelte";
    import Match from "$components/Match.svelte";
    import Skeleton from "$components/Skeleton.svelte";
    import { TAILWIND, user } from "$lib/store";
    import { getUserData } from "$lib/utils";

    let isMounted = false;

    $: won = $user.rankedWon;
    $: lose = $user.rankedPlayed - $user.rankedWon;
    $: winRate = ($user.rankedWon / $user.rankedPlayed * 100).toFixed(1);

    onMount(async () => {
        $user = await getUserData();
        document.documentElement.style.setProperty("--league", $TAILWIND.theme.colors[$user.userLeague]);

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
        <div class="flex items-center">
            <span class="font-black">{$user.username}</span><span class="text-gray-500">#{$user.userId}</span>
            <div class="mx-2xs">
                <Icon kind="sparkle" />
            </div>
            <Icon kind="base" /><span class="ml-3xs">{$user.userLevel}</span>
        </div>
        <div class="flex mt-xs items-center">
            <h1 class="text-league">{$user.userLeague} {$user.userDivision}</h1>
            <div class="ml-md grid grid-cols-4 gap-2xs">
                {#each Array($user.userStars) as _}
                    <Icon kind="star" big color="bg-league" />
                {/each}
            </div>
        </div>
        <div class="grid grid-cols-[repeat(4,auto)] grid-rows-2 w-fit mt-sm gap-y-2xs gap-x-4xl">
            <h2>{won}</h2>
            <h2>{lose}</h2>
            <h2>{winRate}%</h2>
            <h2>15</h2>
            <span>Wins</span>
            <span>Loses</span>
            <span>Win rate</span>
            <span>Avg. turns</span>
        </div>
    {/if}
    <div class="flex justify-between mt-4xl">
        <span>Matches</span>
        <span class="text-gray-500">Share</span>
    </div>
    <div class="grid grid-cols-[repeat(auto-fit,minmax(4rem,1fr))] mt-xs gap-xs mb-4xl">
        {#each Array(100) as _}
            <Match won={Math.random() > 0.5} />
        {/each}
    </div>
</main>
<div class="fixed bottom-0 w-full bg-gradient-to-t from-white from-10% h-4xl pointer-events-none" />