<script lang="ts">
    export let won: boolean | undefined = undefined;
</script>

<div class="w-full h-0 rounded pb-[100%] transition-all duration-300 hover:scale-[0.85]
    {won === undefined ? "border-2 border-gray-100 border-solid" :
        won ? "bg-platinum" : "bg-gray-100"}">

</div>