<script lang="ts">
    import Router from 'svelte-spa-router';
    import routes from './routes';
</script>

<main>
    <h1 class="text-center text-3xl bg-gray-300 font-bold pb-1 mb-2">Calculadora Básica</h1>
    <Router {routes} />
</main>
