<script>
    import axios from 'axios';
    import { onMount } from 'svelte';
    import { resultadoAnterior } from './store';
    import {push} from 'svelte-spa-router';
    let historial=null;
    async function enviarResultado(resultado) {
        resultadoAnterior.set(resultado);
        push('/calculadora');
    }
    const getHistorial=()=>
    {
     axios.get('http://localhost:8080/historial').then(res=>{
        historial=res.data;
        console.log(res);
     }).catch(error => {
            console.error("Error al obtener el historial:", error);
        });
    }
    onMount(getHistorial);    
</script>
<h1 class="text-center font-bold text-3xl">Historial</h1>
<div class="min-h-screen m-4 bg-black p-4"> <!-- Contenedor principal con margen, color de fondo y padding -->
    <table class="w-full text-white"> <!-- Tabla de ancho completo con texto blanco -->
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
                <td class="table-cell">{item.fecha}</td>
                <td class="table-cell">{item.operacion}</td>
                <td class="table-cell">
                    <!-- Boton que tendria que actualizar el valor y cambiar de ruta -->
                    <button on:click={() => enviarResultado(item.resultado)}>{item.resultado}</button> 
                </td>
            </tr>
            {/each}
        {/if}        
    {:catch error}
        <p>Error: {error}</p>
    {/await}
        </tbody>
    </table>
</div>
<style>
   .table-cell{
    @apply px-4 py-2 text-center;
   }
</style>