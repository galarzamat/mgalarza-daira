<script>
    import { http } from '@astric';
    import { resultadoAnterior } from '../shared/store';
    let display_number = '0'; // Inicializar con '0' para que el div sea visible desde el inicio
    let operador = '';
    let primerNum = '';
    let segundoNum = '';
    let resultado = '';

    // Se toma el resultado del historial
    resultadoAnterior.subscribe(value => {
        primerNum = value.toString();
        display_number = value.toString(); // Asegurar que el valor inicial se muestre
    });

    const AddNumber = value => {
        if (operador == '') {
            if (primerNum.length < 2) {
                primerNum += value;
                display_number = primerNum; // Actualizar display_number directamente
            }
        } else {
            if (segundoNum.length < 2) {
                segundoNum += value;
                display_number += value; // Asegurar que el segundo número se añada al display
            }
        }
    };

    const addOperator = value => {
        if (operador == '' && primerNum !== '') {
            operador = value;
            display_number += value;
        }
    };

    const clear_display = () => {
        display_number = '0'; // Resetea a '0' para mantener el div visible
        operador = '';
        primerNum = '';
        segundoNum = '';
        resultado = '';
    };

    const deleteLast = () => {
        if (segundoNum !== '') {
            segundoNum = segundoNum.slice(0, -1);
        } else if (operador !== '') {
            operador = '';
        } else if (primerNum !== '') {
            primerNum = primerNum.slice(0, -1);
        }
        updateDisplay();
    };

    // Función para actualizar el display basado en el estado actual
    function updateDisplay() {
        display_number = `${primerNum}${operador}${segundoNum}`;
        if (display_number === '') display_number = '0';
    }

    //funcion que resuelve operaciones llamando a la api
    function calculate() {
        if (primerNum !== '' && segundoNum !== '' && operador !== '') {
            http.post('operaciones/resolver', {
                primernumero: parseInt(primerNum),
                segundonumero: parseInt(segundoNum),
                operador: operador,
            })
                .then(res => {
                    let respuesta = res.data;
                    let resultado = respuesta.resultado;
                    primerNum = segundoNum = '';
                    operador = '';
                    display_number += resultado.toString();
                })
                .catch(err => {
                    err;
                })
                .finally(() => {});
        }
    }
</script>

<div class="bg-black py-7 text-white text-right m-2">
    <p class="text-2xl">{display_number || segundoNum || '0'}</p>
    {#if resultado}
        <p class="text-2xl">{resultado}</p>
    {/if}
</div>
<div class="m-2">
    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-orange" on:click={() => clear_display()}>CE</button>
        <button class="btn btn-orange" on:click={() => deleteLast()}>C</button>
        <button class="btn btn-red">±</button>
        <button class="btn btn-red" on:click={() => addOperator('%')}>%</button>
    </div>
    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-blue" on:click={() => AddNumber(7)}>7</button>
        <button class="btn btn-blue" on:click={() => AddNumber(8)}>8</button>
        <button class="btn btn-blue" on:click={() => AddNumber(9)}>9</button>
        <button class="btn btn-red" on:click={() => addOperator('÷')}>÷</button>
    </div>

    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-blue" on:click={() => AddNumber(4)}>4</button>
        <button class="btn btn-blue" on:click={() => AddNumber(5)}>5</button>
        <button class="btn btn-blue" on:click={() => AddNumber(6)}>6</button>
        <button class="btn btn-red" on:click={() => addOperator('*')}>x</button>
    </div>

    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-blue" on:click={() => AddNumber(1)}>1</button>
        <button class="btn btn-blue" on:click={() => AddNumber(2)}>2</button>
        <button class="btn btn-blue" on:click={() => AddNumber(3)}>3</button>
        <button class="btn btn-red" on:click={() => addOperator('-')}>-</button>
    </div>
    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-blue" on:click={() => AddNumber(0)}>0</button>
        <button class="btn btn-blue">.</button>
        <button class="btn btn-green" on:click={() => calculate()}>=</button>
        <button class="btn btn-red" on:click={() => addOperator('+')}>+</button>
    </div>
</div>

<style>
    .btn {
        @apply flex-grow text-center p-4 flex-1;
    }
    .btn-blue {
        @apply bg-blue-700 text-white;
    }
    .btn-orange {
        @apply bg-orange-400 text-white;
    }
    .btn-green {
        @apply bg-green-600 text-white;
    }
    .btn-red {
        @apply bg-red-800 text-white;
    }
</style>
