<script>
    import axios from 'axios';
    import { resultadoAnterior } from '../shared/store';
	let display_number = "";
    let operador = "";
    let primerNum="";
    //Se toma el resultado del historial
    resultadoAnterior.subscribe(value => {
        primerNum = value;
        display_number += value;
    });

    let segundoNum = "";
    let resultado="";
	// @ts-ignore
	const AddNumber = (value) =>{
        if(operador==""){
            if (primerNum.length<2) {
                primerNum+=value;
                display_number += value;
            }
        }else if(segundoNum.length<2){
            segundoNum+=value;
        }
	}
	// @ts-ignore
	const addOperator = (value) => {
        if(operador=="" && primerNum!==""){
            operador = value;
            display_number += value;
        }
	}
	const clear_display = () => {
		display_number = "";
        operador ="";
        primerNum = "";
        segundoNum = "";
        resultado="";
	}
	const deleteLast = () => {
		display_number = display_number.slice(0,-1)
	}
    //funcion que resuelve operaciones llamando a la api
	function calculate() {
        if ((primerNum!=="" && segundoNum!=="" && operador!=="")) {
            axios(
                {
                    method: 'post',
                    url: 'http://localhost:8080/resolver',
                    data: {
                        primernumero: parseInt(primerNum),
                        segundonumero: parseInt(segundoNum),
                        operador:operador
                    }
                }).then(res=>
                {
                    let respuesta=res.data;
                    display_number+=segundoNum;
                    operador ="";
                    primerNum = respuesta.resultado;
                    segundoNum = "";
                    resultado=respuesta.resultado;

                }).catch(error => 
                {
                    display_number="Error al  realizar la operacion:", error;
                });;
        }
        
	}
	
</script>
<div class="bg-black py-7 text-white text-right m-2">
    <p class="text-2xl">{display_number}</p>
</div>
<div class="m-2">
    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-orange" on:click={() => clear_display()}>CE</button>
        <button class="btn btn-orange" on:click={() => deleteLast()}>C</button>
        <button class="btn btn-red">±</button>
        <button class="btn btn-red" on:click={() => addOperator("%")}>%</button>
    </div>
    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-blue"on:click={() => AddNumber(7)}>7</button>
        <button class="btn btn-blue"on:click={() => AddNumber(8)}>8</button>
        <button class="btn btn-blue"on:click={() => AddNumber(9)}>9</button>
        <button class="btn btn-red" on:click={() => addOperator("÷")}>÷</button>
    </div>
    
    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-blue"on:click={() => AddNumber(4)}>4</button>
        <button class="btn btn-blue"on:click={() => AddNumber(5)}>5</button>
        <button class="btn btn-blue"on:click={() => AddNumber(6)}>6</button>
        <button class="btn btn-red" on:click={() => addOperator("*")}>x</button>
    </div>
    
    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-blue"on:click={() => AddNumber(1)}>1</button>
        <button class="btn btn-blue"on:click={() => AddNumber(2)}>2</button>
        <button class="btn btn-blue"on:click={() => AddNumber(3)}>3</button>
        <button class="btn btn-red" on:click={() => addOperator("-")}>-</button>
    </div>
    <div class="flex w-full py-1 gap-1">
        <button class="btn btn-blue"on:click={() => AddNumber(0)}>0</button>
        <button class="btn btn-blue">.</button>
        <button class="btn btn-green" on:click={() => calculate()}>=</button>
        <button class="btn btn-red" on:click={() => addOperator("+")}>+</button>
    </div>
</div>
<style>
    .btn{
     @apply flex-grow text-center p-4 flex-1;
    } 
 .btn-blue{
     @apply bg-blue-700 text-white;
 }
 .btn-orange{
     @apply bg-orange-400 text-white;
 }
 .btn-green{
     @apply bg-green-600 text-white;
 }
 .btn-red{
     @apply bg-red-800 text-white;
 }
 </style>
