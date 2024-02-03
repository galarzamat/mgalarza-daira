<script>
    import axios from 'axios';
    import { onMount } from 'svelte';
    let historial=null;
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
    {#await historial}
        <p>Cargando...</p>
    {:then}
        {#if historial}
            {#if historial.length === 0}
                <p>Todavia no hay operaciones realizadas</p>
            {:else}
                {#each historial as item}
                <p>{item.fecha} {item.operacion} {item.resultado}  </p>
                {/each}
            {/if}
        {:else}
            <p>No hay datos disponibles.</p>
        {/if}
    {:catch error}
        <p>Error: {error}</p>
    {/await}