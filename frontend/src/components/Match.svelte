<script lang="ts">
    export let won: boolean | undefined = undefined;
</script>

<div class="w-full h-0 rounded pb-[100%] transition-colors duration-300
    {won === undefined ? "border-[3px] border-gray-100 border-solid" :
        won ? "bg-platinum" : "bg-gray-100"}">

</div>