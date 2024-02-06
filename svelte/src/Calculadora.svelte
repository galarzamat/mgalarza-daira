<script>
    import axios from 'axios';
<<<<<<< HEAD
=======
    import { onMount } from 'svelte';
>>>>>>> 545ab77e50e0072f2fafee4600052ea5ac71a6d5
	let display_number = "";
    let operador = "";
    let primero = "";
    let segundo = "";
    let operadores=[];
	const AddNumber = (value) =>{
<<<<<<< HEAD

        if(operador==""){
            if (primero.length<2) {
                primero+=value;
                display_number += primero; 
            }
        }else if(segundo.length<2){
=======
        if(primero.length<2 && (isFinite(value) || value === '.')){
            primero+=value;
            display_number += primero;
        }else if(segundo.length < 2 && (isFinite(value) || value === '.')){
>>>>>>> 545ab77e50e0072f2fafee4600052ea5ac71a6d5
            segundo+=value;
            display_number = primero + operador + segundo;
        }
	}
	const addOperator = (value) => {
<<<<<<< HEAD
        if(operador=="" && primero!==""){
=======
        if(!operador && primero!==""){
>>>>>>> 545ab77e50e0072f2fafee4600052ea5ac71a6d5
            operador = value;
            display_number += operador;
        }
            
	}
	const clear_display = () => {
		display_number = "";
        operador ="";
        primero = "";
        segundo = "";
	}
	const deleteLast = () => {
		display_number = display_number.slice(0,-1)
	}
	function calculate() {
<<<<<<< HEAD
        if (primero!=="" && segundo!=="" && operador!=="") {
            axios({
=======
        axios({
>>>>>>> 545ab77e50e0072f2fafee4600052ea5ac71a6d5
                method: 'post',
                url: 'http://localhost:8080/resolver',
                data: {
                    primernumero: parseInt(primero),
                    segundonumero: parseInt(segundo),
                    operador:operador
                }
                }).then(res=>{
                let respuesta=res.data;
                display_number+="="+respuesta.resultado;
<<<<<<< HEAD
                console.log(res.data);
        }).catch(error => {
            display_number="Error al  realizar la operacion:", error;
        });;
        }
        
=======
                console.log(respuesta.resultado);
        }).catch(error => {
            display_number="Error al obtener el realizar la operacion:", error;
            console.error("Error al obtener el realizar la operacion:", error);
        });;
>>>>>>> 545ab77e50e0072f2fafee4600052ea5ac71a6d5
	}
	
</script>
<div class="bg-black py-7 text-white text-right m-2">
    <p>{display_number}</p>
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
        <button class="btn btn-blue" on:click={() => AddNumber(".")}>.</button>
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
