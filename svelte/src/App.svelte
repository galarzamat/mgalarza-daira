<script >
    import Router from 'svelte-spa-router';
    import Calculadora from './routes/Calculadora.svelte';
    import Historial from './routes/Historial.svelte';
    const routes = {
    '/': Calculadora,
    '/calculadora': Calculadora,
    '/historial': Historial
    };
</script>
<h1 class="text-center text-3xl bg-gray-300 font-bold pb-1 mb-2">Calculadora Básica</h1>
<Router {routes} />

 