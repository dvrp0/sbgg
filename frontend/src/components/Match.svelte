<script lang="ts">
    import { fade } from "svelte/transition";
    import Icon from "$components/Icon.svelte";

    export let date: string | undefined = undefined;
    export let turns: number | undefined = undefined;
    export let untracked: boolean | undefined = undefined;
    export let untrackedWins: number | undefined = undefined;
    export let untrackedLoses: number | undefined = undefined;
    export let won: boolean | undefined = undefined;
    export let streak: number | undefined = undefined;
    export let trophiesFrom: number | undefined = undefined;
    export let trophiesTo: number | undefined = undefined;

    const margin = { x: 3, y: 3 };

    let isHovered = false;
    let position = { x: 0, y : 0 };

    function handleMouseOver(event: MouseEvent)
    {
        isHovered = true;
        position = { x: event.pageX + margin.x, y: event.pageY + margin.y };
    }

    function handleMouseMove(event: MouseEvent)
    {
        position = { x: event.pageX + margin.x, y: event.pageY + margin.y };
    }

    function handleMouseLeave()
    {
        isHovered = false;
    }
</script>

<div class="w-full h-0 rounded pb-[100%] transition-all duration-300 hover:scale-[0.85]
    {date === undefined ? "border-2 border-gray-100 border-solid" :
        won ? "bg-league" : "bg-gray-100"}"
    on:focus={() => isHovered = true} on:mouseover={handleMouseOver} on:mousemove={handleMouseMove} on:mouseleave={handleMouseLeave}>
</div>

{#if isHovered && date}
    <div class="absolute z-50 bg-gray-900 flex flex-col rounded p-md" style="top: {position.y}px; left: {position.x}px;"
        transition:fade={{ duration: 100 }}>
        <div class="flex items-center">
            <h3 class="text-white">{untracked ? "Untracked Matches" : "Victory"}</h3>
            {#if !untracked}
                <div class="mx-2xs">
                    <Icon kind="sparkle" color="bg-white" />
                </div>
                <span class="text-white">{turns} turns</span>
            {/if}
        </div>
    </div>
{/if}