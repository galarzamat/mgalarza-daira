<script lang="ts">
    // @ts-nocheck
    import { http } from '@astric';
    import { onMount } from 'svelte';
    import { lastResult } from '../shared/store';
    import { push } from 'svelte-spa-router';
    let historial: object[] = [];
    const sendResult = async (result = '') => {
        lastResult.set(result);
        push('/calculadora');
    };
    const getHistorial = () => {
        http.get('operaciones/historial')
            .then(res => {
                historial = res.datos.reverse();
            })
            .catch(error => {
                error;
            })
            .finally(() => {});
    };
    onMount(getHistorial);
</script>

<h1 class="text-center font-bold text-3xl">Historial</h1>
<div class="min-h-screen m-4 bg-black p-4">
    <!-- Contenedor principal con margen, color de fondo y padding -->
    <table class="w-full text-white">
        <!-- Tabla de ancho completo con texto blanco -->
        <thead>
            <tr>
                <th class="table-cell">Fecha</th>
                <th class="table-cell">Operacion</th>
                <th class="table-cell">Resultado</th>
            </tr>
        </thead>
        <tbody>
            {#await historial}
                <p>Cargando...</p>
            {:then}
                {#if historial && historial.length != 0}
                    {#each historial as item}
                        <tr>
                            <td class="table-cell">{item.date}</td>
                            <td class="table-cell">{item.operation}</td>
                            <td class="table-cell">
                                <button on:click={() => sendResult(item.result)}>{item.result}</button>
                            </td>
                        </tr>
                    {/each}
                {/if}
            {:catch error}
                <p>No hay resultados todavia</p>
                <p>Error: {error}</p>
            {/await}
        </tbody>
    </table>
</div>

<style>
    .table-cell {
        @apply px-4 py-2 text-center;
    }
</style>
