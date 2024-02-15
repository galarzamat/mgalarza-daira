

// @ts-nocheck
import { wrap } from 'svelte-spa-router/wrap'

const routes = {'/': wrap({ asyncComponent: () => import('./calculadora/Modul.svelte') }) ,'/calculadora': wrap({ asyncComponent: () => import('./calculadora/Modul.svelte') }), '/historial': wrap({ asyncComponent: () => import('./historial/Modul.svelte') }),};

export default routes;
