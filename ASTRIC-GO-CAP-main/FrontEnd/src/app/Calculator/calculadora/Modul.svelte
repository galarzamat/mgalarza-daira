<script lang="ts">
    import { http } from '@astric';
    import { lastResult } from '../shared/store';
    let displayNumber = '0'; // Inicializar con '0' para que el div sea visible desde el inicio
    let operator = '';
    let firstNumber = '';
    let secondNumber = '';
    let result = '';

    // Se toma el resultado del historial
    lastResult.subscribe(value => {
        firstNumber = value.toString();
        displayNumber = value.toString(); // Asegurar que el valor inicial se muestre
    });

    const AddNumber = (value: string) => {
        if (operator == '') {
            if (firstNumber.length < 2) {
                firstNumber += value;
                displayNumber = firstNumber; // Actualizar display_number directamente
            }
        } else {
            if (secondNumber.length < 2) {
                secondNumber += value;
                displayNumber += value; // Asegurar que el segundo número se añada al display
            }
        }
    };

    const addOperator = (value: string) => {
        if (operator == '' && firstNumber !== '') {
            operator = value;
            displayNumber += value;
        }
    };

    const clearDisplay = () => {
        displayNumber = '0'; // Resetea a '0' para mantener el div visible
        operator = '';
        firstNumber = '';
        secondNumber = '';
        result = '';
    };

    const deleteLast = () => {
        if (secondNumber !== '') {
            secondNumber = secondNumber.slice(0, -1);
        } else if (operator !== '') {
            operator = '';
        } else if (firstNumber !== '') {
            firstNumber = firstNumber.slice(0, -1);
        }
        updateDisplay();
    };

    // Función para actualizar el display basado en el estado actual
    function updateDisplay() {
        displayNumber = `${firstNumber}${operator}${secondNumber}`;
        if (displayNumber === '') displayNumber = '0';
    }

    //funcion que resuelve operaciones llamando a la api
    function calculate() {
        if (firstNumber !== '' && secondNumber !== '' && operator !== '') {
            http.post('operaciones/resolver', {
                firstnumber: parseInt(firstNumber),
                secondnumber: parseInt(secondNumber),
                operator: operator,
            })
                .then(res => {
                    let respuesta = res.datos;
                    let result = respuesta.result;
                    firstNumber = '';
                    secondNumber = '';
                    operator = '';
                    displayNumber = result;
                })
                .catch(err => {
                    err;
                    console.log(err);
                })
                .finally(() => {});
        }
    }
</script>

<div class="bg-black py-7 text-white text-right m-2">
    <p class="text-2xl">{displayNumber || secondNumber || '0'}</p>
    {#if result}
        <p class="text-2xl">{result}</p>
    {/if}
</div>
<div class="m-2">
    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-orange" on:click={() => clearDisplay()}>CE</button>
        <button class="btn btn-orange" on:click={() => deleteLast()}>C</button>
        <button class="btn btn-red">±</button>
        <button class="btn btn-red" on:click={() => addOperator('%')}>%</button>
    </div>
    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-blue" on:click={() => AddNumber('7')}>7</button>
        <button class="btn btn-blue" on:click={() => AddNumber('8')}>8</button>
        <button class="btn btn-blue" on:click={() => AddNumber('9')}>9</button>
        <button class="btn btn-red" on:click={() => addOperator('÷')}>÷</button>
    </div>

    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-blue" on:click={() => AddNumber('4')}>4</button>
        <button class="btn btn-blue" on:click={() => AddNumber('5')}>5</button>
        <button class="btn btn-blue" on:click={() => AddNumber('6')}>6</button>
        <button class="btn btn-red" on:click={() => addOperator('*')}>x</button>
    </div>

    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-blue" on:click={() => AddNumber('1')}>1</button>
        <button class="btn btn-blue" on:click={() => AddNumber('2')}>2</button>
        <button class="btn btn-blue" on:click={() => AddNumber('3')}>3</button>
        <button class="btn btn-red" on:click={() => addOperator('-')}>-</button>
    </div>
    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-blue" on:click={() => AddNumber('0')}>0</button>
        <button class="btn btn-blue">.</button>
        <button class="btn btn-green" on:click={() => calculate()}>=</button>
        <button class="btn btn-red" on:click={() => addOperator('+')}>+</button>
    </div>
</div>

<style>
    .btn {
        @apply flex-grow text-center p-4 flex-1 m-1 w-full;
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
