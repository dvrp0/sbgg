<script lang="ts">
    export let kind: "slot" | "h1" | "h2" | "body"= "slot";
    export let lines = 1;
</script>

<div class="w-full h-full flex flex-col items-start">
    {#each Array(lines) as _, i}
        <div class="skeleton {kind} {lines > 1 && i == lines - 1 ? "mb-xs" : ""}" class:partial={kind !== "slot"}>
            {#if kind === "slot"}
                <slot />
            {/if}
        </div>
    {/each}
</div>

<style lang="postcss">
    @keyframes flicker {
        0% {
            background-position: 100% 50%;
        }

        100% {
            background-position: 20% 50%;
        }
    }

    .skeleton {
        width: 100%;
        height: 100%;
        background-image: linear-gradient(to right, theme("colors.gray.100") 40%, theme("colors.gray.300") 60%, theme("colors.gray.300") 70%, theme("colors.gray.100") 70%);
        background-size: 500% auto;
        animation: flicker 2.5s infinite;
        border-radius: theme("borderRadius.DEFAULT");

        &.h1 {
            height: theme("spacing.2xl");

            &.partial {
                width: 35%;
            }
        }

        &.h2 {
            height: theme("spacing.lg");

            &.partial {
                width: 50%;
            }
        }

        &.body {
            height: theme("spacing.xs");

            &.partial {
                width: 65%;
            }
        }
    }
</style>