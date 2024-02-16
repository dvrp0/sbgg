<script lang="ts">
    import Match from "$components/Match.svelte";
</script>

<main class="flex flex-col px-4xl pt-4xl">
    <div class="flex">
        <span class="font-black">dvrp</span><span class="text-gray-500">#1001413882</span>
    </div>
    <div class="flex mt-xs">
        <h1 class="text-platinum">Platinum 2</h1>
    </div>
    <div class="grid grid-cols-[repeat(4,auto)] grid-rows-2 w-fit mt-lg gap-y-2xs gap-x-4xl">
        <h2>116</h2>
        <h2>13</h2>
        <h2>89.4%</h2>
        <h2>15</h2>
        <span>Wins</span>
        <span>Loses</span>
        <span>Win rate</span>
        <span>Avg. turns</span>
    </div>
    <div class="flex justify-between mt-3xl">
        <span>Matches</span>
        <span class="text-gray-500">Share</span>
    </div>
    <div class="grid grid-cols-10 mt-xs gap-xs">
        {#each Array(27) as _}
            <Match won={Math.random() > 0.5} />
        {/each}
    </div>
</main>
<div class="fixed bottom-0 w-full bg-gradient-to-t from-white from-15% h-4xl" />